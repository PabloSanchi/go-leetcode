package parking_lot

// newParkingLevelWithSpots for testing only
// no fancy code here, just to make tests easier to write
func newParkingLevelWithSpots(floor, carSpots, motorcycleSpots, truckSpots uint) *ParkingLevel {
	size := carSpots + motorcycleSpots + truckSpots
	spots := make([]*Spot, size)
	spotsPerType := make(map[VehicleType]int)

	for i := uint(0); i < carSpots; i++ {
		spots[i] = NewSpot(CAR)
		spotsPerType[CAR]++
	}

	for i := carSpots; i < carSpots+motorcycleSpots; i++ {
		spots[i] = NewSpot(MOTORCYCLE)
		spotsPerType[MOTORCYCLE]++
	}

	for i := carSpots + motorcycleSpots; i < size; i++ {
		spots[i] = NewSpot(TRUCK)
		spotsPerType[TRUCK]++
	}

	return &ParkingLevel{
		Floor:        floor,
		MaxSize:      int(size),
		Size:         0,
		spots:        spots,
		spotsPerType: spotsPerType,
	}

}

// newParkingLot only for testing purposes
func newParkingLot(levels []*ParkingLevel, db map[string]uint) *ParkingLot {
	return &ParkingLot{
		levels: levels,
		db:     db,
	}
}
