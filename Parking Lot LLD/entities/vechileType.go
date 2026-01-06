package entities

type VechileType int

const (
	MOTORCYCLE VechileType = iota
	CAR
	TRUCK
)

type Vechile interface {
	Type() VechileType
	GetNumberPlate() string
}

type MotorCycle struct{
	numberPlate string
}

func (mc MotorCycle) Type() VechileType{
	return MOTORCYCLE
}

func (mc MotorCycle) getNumberPlate() string{
	return mc.numberPlate
}

type Car struct{
	numberPlate string
}

func (c Car) Type() VechileType{
	return CAR
}

func (c Car) getNumberPlate() string{
	return c.numberPlate
}

type Truck struct{
	numberPlate string
}

func (t Truck) Type() VechileType{
	return TRUCK
}

func (t Truck) getNumberPlate() string{
	return t.numberPlate
}