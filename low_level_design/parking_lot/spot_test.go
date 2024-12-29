package parking_lot

import "testing"

func TestSpotIn(t *testing.T) {
	spot := NewSpot(CAR)

	vehicle := NewCar("123ABC")

	err := spot.In(vehicle)

	if spot.GetVehicle() != vehicle {
		t.Errorf("expected vehicle %v, got %v", vehicle, spot.GetVehicle())
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestSpotInWithUnsupportedVehicle(t *testing.T) {
	spot := NewSpot(CAR)

	vehicle := NewMotorcycle("123ABC")

	err := spot.In(vehicle)

	if spot.GetVehicle() != nil {
		t.Errorf("expected no vehicle, got %v", spot.GetVehicle())
	}

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
