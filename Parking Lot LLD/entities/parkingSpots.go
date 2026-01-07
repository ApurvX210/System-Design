package entities

import "sync"

type ParkingSpot interface {
	SetOcc(bool)
	Status() bool
}

type MotorCycleSpot struct {
	SpotId   int
	mu       sync.RWMutex
	occupied bool
}

func (m *MotorCycleSpot) SetOcc(flag bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.occupied = flag
}

func (m *MotorCycleSpot) Status() bool{
	m.mu.RLock()
	defer m.mu.RLocker().Unlock()

	return m.occupied
}

type CarSpot struct {
	SpotId   int
	mu       sync.RWMutex
	occupied bool
}

func (c *CarSpot) SetOcc(flag bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.occupied = flag
}

func (c *CarSpot) Status() bool{
	c.mu.RLock()
	defer c.mu.RLocker().Unlock()

	return c.occupied
}

type TruckSpot struct {
	SpotId   int
	mu       sync.RWMutex
	occupied bool
}

func (t *TruckSpot) SetOcc(flag bool) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.occupied = flag
}

func (t *TruckSpot) Status() bool{
	t.mu.RLock()
	defer t.mu.RLocker().Unlock()

	return t.occupied
}
