package service

import (
	"errors"
	"fmt"
	"sync"

	"../entities"
)

var (
	ErrParkingLotFull      = errors.New("parking lot is full")
	ErrInvalidVehicle      = errors.New("invalid vehicle")
	ErrTicketNotFound      = errors.New("ticket not found")
	ErrVehicleNotParked    = errors.New("vehicle is not currently parked")
	ErrInvalidFloor         = errors.New("invalid floor number")
)

// ParkingLotService manages the entire parking lot operations
// This is the main service layer that coordinates between floors, spots, and tickets
type ParkingLotService struct {
	floors          []*entities.ParkingSpace
	tickets         map[string]*entities.Ticket // ticketID -> ticket
	activeTickets   map[string]*entities.Ticket // vehicle number plate -> ticket (for quick lookup)
	pricing         map[entities.VehicleType]int // price per hour for each vehicle type
	mu              sync.RWMutex
}

// NewParkingLotService creates a new parking lot service
// floorsConfig: array of floor configurations, each containing [carCapacity, motorcycleCapacity, truckCapacity]
func NewParkingLotService(floorsConfig [][3]int, pricing map[entities.VehicleType]int) *ParkingLotService {
	floors := make([]*entities.ParkingSpace, len(floorsConfig))
	for i, config := range floorsConfig {
		floors[i] = entities.NewParkingSpace(i+1, config[0], config[1], config[2])
	}

	// Default pricing if not provided
	if pricing == nil {
		pricing = map[entities.VehicleType]int{
			entities.MOTORCYCLE: 10, // $0.10 per hour
			entities.CAR:        20, // $0.20 per hour
			entities.TRUCK:      50, // $0.50 per hour
		}
	}

	return &ParkingLotService{
		floors:        floors,
		tickets:       make(map[string]*entities.Ticket),
		activeTickets: make(map[string]*entities.Ticket),
		pricing:       pricing,
	}
}

// ParkVehicle parks a vehicle and returns a ticket
// Strategy: Find nearest available spot (check floors from bottom to top)
func (pls *ParkingLotService) ParkVehicle(vehicle entities.Vehicle) (*entities.Ticket, error) {
	if vehicle == nil {
		return nil, ErrInvalidVehicle
	}

	pls.mu.Lock()
	defer pls.mu.Unlock()

	// Check if vehicle is already parked
	if _, exists := pls.activeTickets[vehicle.GetNumberPlate()]; exists {
		return nil, fmt.Errorf("vehicle %s is already parked", vehicle.GetNumberPlate())
	}

	vehicleType := vehicle.Type()

	// Find first available spot across all floors
	for _, floor := range pls.floors {
		spotCollection := floor.GetSpotByVehicleType(vehicleType)
		if spotCollection == nil {
			continue
		}

		spotID, err := spotCollection.FindVacantSpot()
		if err != nil {
			continue // No spot on this floor, try next
		}

		// Occupy the spot
		if err := spotCollection.OccupySpot(spotID, vehicle); err != nil {
			continue // Failed to occupy, try next floor
		}

		// Create ticket
		pricePerHour := pls.pricing[vehicleType]
		ticket := entities.NewTicket(vehicle, floor.ID, spotID, pricePerHour)

		// Store ticket
		pls.tickets[ticket.ID] = ticket
		pls.activeTickets[vehicle.GetNumberPlate()] = ticket

		return ticket, nil
	}

	return nil, ErrParkingLotFull
}

// UnparkVehicle releases a vehicle and calculates the final price
func (pls *ParkingLotService) UnparkVehicle(ticketID string) (*entities.Ticket, int, error) {
	pls.mu.Lock()
	defer pls.mu.Unlock()

	ticket, exists := pls.tickets[ticketID]
	if !exists {
		return nil, 0, ErrTicketNotFound
	}

	if !ticket.IsActive() {
		return ticket, ticket.CalculatePrice(), nil // Already unparked, return existing price
	}

	// Validate floor and spot
	if ticket.FloorID < 1 || ticket.FloorID > len(pls.floors) {
		return nil, 0, ErrInvalidFloor
	}

	floor := pls.floors[ticket.FloorID-1]
	spotCollection := floor.GetSpotByVehicleType(ticket.VehicleType)

	if spotCollection == nil {
		return nil, 0, errors.New("invalid spot collection")
	}

	// Release the spot
	if err := spotCollection.ReleaseSpot(ticket.SpotID); err != nil {
		return nil, 0, fmt.Errorf("failed to release spot: %w", err)
	}

	// Mark ticket as exited
	ticket.MarkExit()

	// Remove from active tickets
	delete(pls.activeTickets, ticket.Vehicle.GetNumberPlate())

	// Calculate final price
	price := ticket.CalculatePrice()

	return ticket, price, nil
}

// GetTicket retrieves a ticket by ID
func (pls *ParkingLotService) GetTicket(ticketID string) (*entities.Ticket, error) {
	pls.mu.RLock()
	defer pls.mu.RUnlock()

	ticket, exists := pls.tickets[ticketID]
	if !exists {
		return nil, ErrTicketNotFound
	}

	return ticket, nil
}

// GetActiveTicketByVehicle retrieves active ticket for a vehicle
func (pls *ParkingLotService) GetActiveTicketByVehicle(numberPlate string) (*entities.Ticket, error) {
	pls.mu.RLock()
	defer pls.mu.RUnlock()

	ticket, exists := pls.activeTickets[numberPlate]
	if !exists {
		return nil, ErrVehicleNotParked
	}

	return ticket, nil
}

// GetParkingLotStatus returns the current status of the parking lot
func (pls *ParkingLotService) GetParkingLotStatus() *ParkingLotStatus {
	pls.mu.RLock()
	defer pls.mu.RUnlock()

	status := &ParkingLotStatus{
		Floors: make([]FloorStatus, len(pls.floors)),
		TotalActiveTickets: len(pls.activeTickets),
	}

	for i, floor := range pls.floors {
		status.Floors[i] = FloorStatus{
			FloorID: floor.ID,
			CarSpots: SpotStatus{
				Total:    floor.CarSpots.GetTotalSpots(),
				Occupied: floor.CarSpots.GetOccupiedCount(),
				Vacant:   floor.CarSpots.GetVacantCount(),
			},
			MotorcycleSpots: SpotStatus{
				Total:    floor.MotorBikeSpots.GetTotalSpots(),
				Occupied: floor.MotorBikeSpots.GetOccupiedCount(),
				Vacant:   floor.MotorBikeSpots.GetVacantCount(),
			},
			TruckSpots: SpotStatus{
				Total:    floor.TruckSpots.GetTotalSpots(),
				Occupied: floor.TruckSpots.GetOccupiedCount(),
				Vacant:   floor.TruckSpots.GetVacantCount(),
			},
		}
	}

	return status
}

// ParkingLotStatus represents the current status of the parking lot
type ParkingLotStatus struct {
	Floors            []FloorStatus
	TotalActiveTickets int
}

// FloorStatus represents the status of a single floor
type FloorStatus struct {
	FloorID          int
	CarSpots         SpotStatus
	MotorcycleSpots  SpotStatus
	TruckSpots       SpotStatus
}

// SpotStatus represents the status of spots of a particular type
type SpotStatus struct {
	Total    int
	Occupied int
	Vacant   int
}

