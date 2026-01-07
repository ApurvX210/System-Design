package entities

import "sync"

type ParkingSpot interface {
	ValidSpot() int
	// Capacity() int
}

type MotorCycleSpot struct {
	SpotsCount	int
	mu       sync.RWMutex
	occupied	[]int
	vacant		[]int
}

func (m *MotorCycleSpot) ValidSpot() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Logic to find vacant spot
	return 0
}

// func (m *MotorCycleSpot) Status() bool{
// 	m.mu.RLock()
// 	defer m.mu.RLocker().Unlock()

// 	return m.occupied
// }

type CarSpot struct {
	SpotId   int
	mu       sync.RWMutex
	occupied bool
}

func (c *CarSpot) ValidSpot() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Logic to find vacant spot
	return 0
}

// func (c *CarSpot) Status() bool{
// 	c.mu.RLock()
// 	defer c.mu.RLocker().Unlock()

// 	return c.occupied
// }

type TruckSpot struct {
	SpotId   int
	mu       sync.RWMutex
	occupied bool
}

func (t *TruckSpot) ValidSpot() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	// Logic to find vacant spot
	return 0
}

// func (t *TruckSpot) Status() bool{
// 	t.mu.RLock()
// 	defer t.mu.RLocker().Unlock()

// 	return t.occupied
// }
