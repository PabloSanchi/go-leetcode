package parking_lot

type ParkingLot struct {
	levels []*ParkingLevel
	db     map[string]uint
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		levels: make([]*ParkingLevel, 0),
		db:     make(map[string]uint),
	}
}

func (pl *ParkingLot) AddLevel(size int) {
	floor := len(pl.levels)
	pl.levels = append(pl.levels, NewParkingLevel(uint(floor), size))
}

func (pl *ParkingLot) Join(vehicle *Vehicle) bool {
	for floor := range pl.levels {
		level := pl.levels[floor]

		if !level.IsFull() && level.CanPark(vehicle) {
			ok := level.Park(vehicle)
			if ok {
				pl.db[vehicle.Plate] = uint(floor)
			}

			return ok
		}
	}

	return false
}

func (pl *ParkingLot) Leave(vehicle *Vehicle) bool {

	floor, ok := pl.db[vehicle.Plate]
	if !ok {
		return false
	}

	level := pl.levels[floor]

	if ok := level.Unpark(vehicle); ok {
		delete(pl.db, vehicle.Plate)
		return true
	}

	return false
}
