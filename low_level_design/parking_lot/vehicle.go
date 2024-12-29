package parking_lot

type VehicleType int

const (
	CAR VehicleType = iota
	MOTORCYCLE
	TRUCK
)

type Vehicle struct {
	Plate string
	Type  VehicleType
}

func NewCar(plate string) *Vehicle {
	return &Vehicle{
		Plate: plate,
		Type:  CAR,
	}
}

func NewMotorcycle(plate string) *Vehicle {
	return &Vehicle{
		Plate: plate,
		Type:  MOTORCYCLE,
	}
}

func NewTruck(plate string) *Vehicle {
	return &Vehicle{
		Plate: plate,
		Type:  TRUCK,
	}
}
