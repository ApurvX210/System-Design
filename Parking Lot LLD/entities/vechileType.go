package entities

// VehicleType represents the type of vehicle
type VehicleType int

const (
	MOTORCYCLE VehicleType = iota
	CAR
	TRUCK
)

// Vehicle interface defines methods that all vehicles must implement
type Vehicle interface {
	Type() VehicleType
	GetNumberPlate() string
}

// MotorCycle represents a motorcycle vehicle
type MotorCycle struct {
	numberPlate string
}

func (mc MotorCycle) Type() VehicleType {
	return MOTORCYCLE
}

// GetNumberPlate returns the license plate number (exported method)
func (mc MotorCycle) GetNumberPlate() string {
	return mc.numberPlate
}

// NewMotorCycle creates a new MotorCycle instance
func NewMotorCycle(numberPlate string) *MotorCycle {
	return &MotorCycle{numberPlate: numberPlate}
}

// Car represents a car vehicle
type Car struct {
	numberPlate string
}

func (c Car) Type() VehicleType {
	return CAR
}

// GetNumberPlate returns the license plate number (exported method)
func (c Car) GetNumberPlate() string {
	return c.numberPlate
}

// NewCar creates a new Car instance
func NewCar(numberPlate string) *Car {
	return &Car{numberPlate: numberPlate}
}

// Truck represents a truck vehicle
type Truck struct {
	numberPlate string
}

func (t Truck) Type() VehicleType {
	return TRUCK
}

// GetNumberPlate returns the license plate number (exported method)
func (t Truck) GetNumberPlate() string {
	return t.numberPlate
}

// NewTruck creates a new Truck instance
func NewTruck(numberPlate string) *Truck {
	return &Truck{numberPlate: numberPlate}
}
