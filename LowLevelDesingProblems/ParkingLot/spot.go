package parking_lot

import (
	"errors"
	"log/slog"
)

const VehicleNotSupportedErrorMsg = "the vehicle is not supported in the selected spot"

type Spot struct {
	vType   VehicleType
	vehicle *Vehicle
}

func NewSpot(vType VehicleType) *Spot {
	return &Spot{
		vType:   vType,
		vehicle: nil,
	}
}

func (s *Spot) GetType() VehicleType {
	return s.vType
}

func (s *Spot) GetVehicle() *Vehicle {
	return s.vehicle
}

func (s *Spot) IsFree() bool {
	return s.vehicle == nil
}

func (s *Spot) In(vehicle *Vehicle) error {
	if s.vType != vehicle.Type {
		slog.Error(VehicleNotSupportedErrorMsg)
		return errors.New(VehicleNotSupportedErrorMsg)
	}

	s.vehicle = vehicle
	return nil
}

func (s *Spot) Out() {
	s.vehicle = nil
}
