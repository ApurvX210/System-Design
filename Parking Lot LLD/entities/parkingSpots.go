package entities

import (
	"errors"
	"sync"
)

var (
	ErrNoVacantSpot        = errors.New("no vacant spot available")
	ErrSpotNotFound        = errors.New("spot not found")
	ErrSpotAlreadyOccupied = errors.New("spot is already occupied")
	ErrSpotAlreadyVacant   = errors.New("spot is already vacant")
)

// ParkingSpot interface defines methods for managing parking spots
type ParkingSpot interface {
	FindVacantSpot() (int, error)
	OccupySpot(spotID int, vehicle Vehicle) error
	ReleaseSpot(spotID int) error
	IsOccupied(spotID int) bool
	GetVehicle(spotID int) (Vehicle, error)
	GetTotalSpots() int
	GetOccupiedCount() int
	GetVacantCount() int
}

// Spot represents a single parking spot
type Spot struct {
	ID       int
	Occupied bool
	Vehicle  Vehicle
}

// MotorCycleSpot manages multiple motorcycle parking spots
type MotorCycleSpot struct {
	spots []*Spot
	mu    sync.RWMutex
}

// NewMotorCycleSpot creates a new MotorCycleSpot with specified capacity
func NewMotorCycleSpot(capacity int) *MotorCycleSpot {
	spots := make([]*Spot, capacity)
	for i := 0; i < capacity; i++ {
		spots[i] = &Spot{ID: i + 1, Occupied: false}
	}
	return &MotorCycleSpot{spots: spots}
}

func (m *MotorCycleSpot) FindVacantSpot() (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, spot := range m.spots {
		if !spot.Occupied {
			return spot.ID, nil
		}
	}
	return 0, ErrNoVacantSpot
}

func (m *MotorCycleSpot) OccupySpot(spotID int, vehicle Vehicle) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if spotID < 1 || spotID > len(m.spots) {
		return ErrSpotNotFound
	}

	spot := m.spots[spotID-1]
	if spot.Occupied {
		return ErrSpotAlreadyOccupied
	}

	spot.Occupied = true
	spot.Vehicle = vehicle
	return nil
}

func (m *MotorCycleSpot) ReleaseSpot(spotID int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if spotID < 1 || spotID > len(m.spots) {
		return ErrSpotNotFound
	}

	spot := m.spots[spotID-1]
	if !spot.Occupied {
		return ErrSpotAlreadyVacant
	}

	spot.Occupied = false
	spot.Vehicle = nil
	return nil
}

func (m *MotorCycleSpot) IsOccupied(spotID int) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if spotID < 1 || spotID > len(m.spots) {
		return false
	}
	return m.spots[spotID-1].Occupied
}

func (m *MotorCycleSpot) GetVehicle(spotID int) (Vehicle, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if spotID < 1 || spotID > len(m.spots) {
		return nil, ErrSpotNotFound
	}

	spot := m.spots[spotID-1]
	if !spot.Occupied {
		return nil, ErrSpotAlreadyVacant
	}

	return spot.Vehicle, nil
}

func (m *MotorCycleSpot) GetTotalSpots() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.spots)
}

func (m *MotorCycleSpot) GetOccupiedCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, spot := range m.spots {
		if spot.Occupied {
			count++
		}
	}
	return count
}

func (m *MotorCycleSpot) GetVacantCount() int {
	return m.GetTotalSpots() - m.GetOccupiedCount()
}

// CarSpot manages multiple car parking spots
type CarSpot struct {
	spots []*Spot
	mu    sync.RWMutex
}

// NewCarSpot creates a new CarSpot with specified capacity
func NewCarSpot(capacity int) *CarSpot {
	spots := make([]*Spot, capacity)
	for i := 0; i < capacity; i++ {
		spots[i] = &Spot{ID: i + 1, Occupied: false}
	}
	return &CarSpot{spots: spots}
}

func (c *CarSpot) FindVacantSpot() (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, spot := range c.spots {
		if !spot.Occupied {
			return spot.ID, nil
		}
	}
	return 0, ErrNoVacantSpot
}

func (c *CarSpot) OccupySpot(spotID int, vehicle Vehicle) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if spotID < 1 || spotID > len(c.spots) {
		return ErrSpotNotFound
	}

	spot := c.spots[spotID-1]
	if spot.Occupied {
		return ErrSpotAlreadyOccupied
	}

	spot.Occupied = true
	spot.Vehicle = vehicle
	return nil
}

func (c *CarSpot) ReleaseSpot(spotID int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if spotID < 1 || spotID > len(c.spots) {
		return ErrSpotNotFound
	}

	spot := c.spots[spotID-1]
	if !spot.Occupied {
		return ErrSpotAlreadyVacant
	}

	spot.Occupied = false
	spot.Vehicle = nil
	return nil
}

func (c *CarSpot) IsOccupied(spotID int) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if spotID < 1 || spotID > len(c.spots) {
		return false
	}
	return c.spots[spotID-1].Occupied
}

func (c *CarSpot) GetVehicle(spotID int) (Vehicle, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if spotID < 1 || spotID > len(c.spots) {
		return nil, ErrSpotNotFound
	}

	spot := c.spots[spotID-1]
	if !spot.Occupied {
		return nil, ErrSpotAlreadyVacant
	}

	return spot.Vehicle, nil
}

func (c *CarSpot) GetTotalSpots() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.spots)
}

func (c *CarSpot) GetOccupiedCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	count := 0
	for _, spot := range c.spots {
		if spot.Occupied {
			count++
		}
	}
	return count
}

func (c *CarSpot) GetVacantCount() int {
	return c.GetTotalSpots() - c.GetOccupiedCount()
}

// TruckSpot manages multiple truck parking spots
type TruckSpot struct {
	spots []*Spot
	mu    sync.RWMutex
}

// NewTruckSpot creates a new TruckSpot with specified capacity
func NewTruckSpot(capacity int) *TruckSpot {
	spots := make([]*Spot, capacity)
	for i := 0; i < capacity; i++ {
		spots[i] = &Spot{ID: i + 1, Occupied: false}
	}
	return &TruckSpot{spots: spots}
}

func (t *TruckSpot) FindVacantSpot() (int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, spot := range t.spots {
		if !spot.Occupied {
			return spot.ID, nil
		}
	}
	return 0, ErrNoVacantSpot
}

func (t *TruckSpot) OccupySpot(spotID int, vehicle Vehicle) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if spotID < 1 || spotID > len(t.spots) {
		return ErrSpotNotFound
	}

	spot := t.spots[spotID-1]
	if spot.Occupied {
		return ErrSpotAlreadyOccupied
	}

	spot.Occupied = true
	spot.Vehicle = vehicle
	return nil
}

func (t *TruckSpot) ReleaseSpot(spotID int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if spotID < 1 || spotID > len(t.spots) {
		return ErrSpotNotFound
	}

	spot := t.spots[spotID-1]
	if !spot.Occupied {
		return ErrSpotAlreadyVacant
	}

	spot.Occupied = false
	spot.Vehicle = nil
	return nil
}

func (t *TruckSpot) IsOccupied(spotID int) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if spotID < 1 || spotID > len(t.spots) {
		return false
	}
	return t.spots[spotID-1].Occupied
}

func (t *TruckSpot) GetVehicle(spotID int) (Vehicle, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if spotID < 1 || spotID > len(t.spots) {
		return nil, ErrSpotNotFound
	}

	spot := t.spots[spotID-1]
	if !spot.Occupied {
		return nil, ErrSpotAlreadyVacant
	}

	return spot.Vehicle, nil
}

func (t *TruckSpot) GetTotalSpots() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.spots)
}

func (t *TruckSpot) GetOccupiedCount() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	count := 0
	for _, spot := range t.spots {
		if spot.Occupied {
			count++
		}
	}
	return count
}

func (t *TruckSpot) GetVacantCount() int {
	return t.GetTotalSpots() - t.GetOccupiedCount()
}
