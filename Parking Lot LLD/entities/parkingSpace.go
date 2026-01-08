package entities

// ParkingSpace represents a single floor in the parking lot
// Each floor can have multiple spots of different vehicle types
type ParkingSpace struct {
	ID             int
	CarSpots       *CarSpot
	MotorBikeSpots *MotorCycleSpot
	TruckSpots     *TruckSpot
}

// NewParkingSpace creates a new parking space (floor) with specified capacities
func NewParkingSpace(floorID, carCapacity, motorcycleCapacity, truckCapacity int) *ParkingSpace {
	return &ParkingSpace{
		ID:             floorID,
		CarSpots:       NewCarSpot(carCapacity),
		MotorBikeSpots: NewMotorCycleSpot(motorcycleCapacity),
		TruckSpots:     NewTruckSpot(truckCapacity),
	}
}

// GetSpotByVehicleType returns the appropriate spot collection based on vehicle type
func (ps *ParkingSpace) GetSpotByVehicleType(vehicleType VehicleType) ParkingSpot {
	switch vehicleType {
	case MOTORCYCLE:
		return ps.MotorBikeSpots
	case CAR:
		return ps.CarSpots
	case TRUCK:
		return ps.TruckSpots
	default:
		return nil
	}
}
