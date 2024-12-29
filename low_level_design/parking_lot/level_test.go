package parking_lot

import (
	"sync"
	"testing"
)

func TestNewParkingLevel(t *testing.T) {
	size := 8
	parkingLevel := newParkingLevelWithSpots(1, 4, 2, 2)

	if len(parkingLevel.spots) != size {
		t.Errorf("expected %d spots, got %d", size, len(parkingLevel.spots))
	}

	if parkingLevel.spotsPerType[CAR] != 4 || parkingLevel.spotsPerType[MOTORCYCLE] != 2 || parkingLevel.spotsPerType[TRUCK] != 2 {
		t.Errorf("unexpected spots distribution: %+v", parkingLevel.spotsPerType)
	}
}

func TestParkingLevel_ParkAndUnpark(t *testing.T) {
	parkingLevel := newParkingLevelWithSpots(1, 2, 1, 1)

	car := NewCar("CAR-001")
	motorcycle := NewMotorcycle("MC-001")

	if !parkingLevel.Park(car) {
		t.Errorf("expected car to be parked")
	}

	if parkingLevel.Size != 1 {
		t.Errorf("expected size to be 1, got %d", parkingLevel.Size)
	}

	if parkingLevel.spotsPerType[CAR] != 1 {
		t.Errorf("expected 1 car spots left, got %d", parkingLevel.spotsPerType[CAR])
	}

	if !parkingLevel.Unpark(car) {
		t.Errorf("expected car to be unparked")
	}

	if parkingLevel.Size != 0 {
		t.Errorf("expected size to be 0, got %d", parkingLevel.Size)
	}

	if parkingLevel.spotsPerType[CAR] != 2 {
		t.Errorf("expected 2 car spots left, got %d", parkingLevel.spotsPerType[CAR])
	}

	if parkingLevel.Unpark(motorcycle) {
		t.Errorf("unparking a non-parked vehicle should return false")
	}
}

func TestParkingLevel_IsFull(t *testing.T) {
	parkingLevel := newParkingLevelWithSpots(1, 2, 0, 0)

	car1 := NewCar("CAR-001")
	car2 := NewCar("CAR-002")

	if parkingLevel.IsFull() {
		t.Errorf("expected IsFull to return false, got true")
	}

	parkingLevel.Park(car1)
	parkingLevel.Park(car2)

	if !parkingLevel.IsFull() {
		t.Errorf("expected IsFull to return true, got false")
	}
}

func TestParkingLevel_CanPark(t *testing.T) {
	parkingLevel := newParkingLevelWithSpots(1, 2, 1, 1)

	car1 := NewCar("CAR-001")
	car2 := NewCar("CAR-002")
	motorcycle := NewMotorcycle("MC-001")
	truck := NewTruck("TRK-001")

	if !parkingLevel.Park(car1) {
		t.Errorf("expected to park car1")
	}

	if !parkingLevel.CanPark(car2) {
		t.Errorf("expected to be able to park car2")
	}
	if !parkingLevel.Park(car2) {
		t.Errorf("expected to park car2")
	}

	if !parkingLevel.Park(motorcycle) {
		t.Errorf("expected to park motorcycle")
	}

	if !parkingLevel.Park(truck) {
		t.Errorf("expected to park truck")
	}

	car3 := NewCar("CAR-003")
	if parkingLevel.CanPark(car3) {
		t.Errorf("expected not to park car3")
	}
}

func TestParkingLevel_Concurrency(t *testing.T) {
	parkingLevel := newParkingLevelWithSpots(1, 5, 3, 2)

	car := NewCar("CAR-001")
	motorcycle := NewMotorcycle("MC-001")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		parkingLevel.Park(car)
	}()

	go func() {
		defer wg.Done()
		parkingLevel.Park(motorcycle)
	}()

	wg.Wait()

	if parkingLevel.Size != 2 {
		t.Errorf("expected parking level size to be 2, got %d", parkingLevel.Size)
	}

	if parkingLevel.spotsPerType[CAR] != 4 {
		t.Errorf("expected 4 car spots left, got %d", parkingLevel.spotsPerType[CAR])
	}

	if parkingLevel.spotsPerType[MOTORCYCLE] != 2 {
		t.Errorf("expected 2 motorcycle spots left, got %d", parkingLevel.spotsPerType[MOTORCYCLE])
	}

}
