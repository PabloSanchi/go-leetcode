package parking_lot

import "testing"

func TestCreateVehicleTypeFromNumber(t *testing.T) {
	vehicleType := VehicleType(0)

	if vehicleType != CAR {
		t.Errorf("expected CAR(0), got %v", vehicleType)
	}
}
