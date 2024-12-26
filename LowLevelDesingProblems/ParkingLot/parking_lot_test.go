package parking_lot

import (
	"sync"
	"testing"
)

func TestParkingLot_Join(t *testing.T) {
	parkingLot := newParkingLot(
		[]*ParkingLevel{newParkingLevelWithSpots(1, 1, 0, 0)},
		make(map[string]uint),
	)

	car1 := NewCar("CAR-001")
	car2 := NewCar("CAR-002")
	car3 := NewCar("CAR-003")

	if !parkingLot.Join(car1) {
		t.Errorf("expected to park vehicle %s", car1.Plate)
	}

	if parkingLot.Join(car2) {
		t.Errorf("expected to not park vehicle %s", car2.Plate)
	}

	if parkingLot.Join(car3) {
		t.Errorf("expected to not park vehicle %s", car3.Plate)
	}

	_, ok := parkingLot.db[car1.Plate]
	if !ok {
		t.Errorf("expected to find vehicle %s in the database", car1.Plate)
	}
}

func TestParkingLevel_RaceConditionTest(t *testing.T) {
	parkingLot := newParkingLot(
		[]*ParkingLevel{newParkingLevelWithSpots(1, 1, 0, 0)},
		make(map[string]uint),
	)

	car1 := &Vehicle{Plate: "CAR-001", Type: CAR}
	car2 := &Vehicle{Plate: "CAR-002", Type: CAR}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() { defer wg.Done(); parkingLot.Join(car1) }()
	go func() { defer wg.Done(); parkingLot.Join(car2) }()
	wg.Wait()

	if len(parkingLot.db) != 1 {
		t.Errorf("expected to have only one vehicle parked")
	}
}

func TestParkingLot_JoinAndLeave(t *testing.T) {
	parkingLot := newParkingLot(
		[]*ParkingLevel{newParkingLevelWithSpots(1, 1, 0, 0)},
		make(map[string]uint),
	)

	car1 := NewCar("CAR-001")
	car2 := NewCar("CAR-002")

	if !parkingLot.Join(car1) {
		t.Errorf("expected to park vehicle %s", car1.Plate)
	}

	if parkingLot.Join(car2) {
		t.Errorf("expected to not park vehicle %s", car2.Plate)
	}

	if !parkingLot.Leave(car1) {
		t.Errorf("expected to unpark vehicle %s", car1.Plate)
	}

	if !parkingLot.Join(car2) {
		t.Errorf("expected to park vehicle %s", car2.Plate)
	}

	if !parkingLot.Leave(car2) {
		t.Errorf("expected to unpark vehicle %s", car2.Plate)
	}
}
