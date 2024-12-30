package parking_lot

import (
	"math/rand"
	"sync"
)

type ParkingLevel struct {
	Floor        uint
	MaxSize      int
	Size         int
	spots        []*Spot
	spotsPerType map[VehicleType]int

	l sync.RWMutex
}

// NewParkingLevel 50% for CARS, 25% for MOTORCYCLES, 25% for TRUCKS
func NewParkingLevel(floor uint, size int) *ParkingLevel {
	spots := make([]*Spot, size)
	spotsPerType := make(map[VehicleType]int)

	for i := 0; i < size; i++ {
		vehicleType := VehicleType(rand.Intn(3))
		spots[i] = NewSpot(vehicleType)
		spotsPerType[vehicleType]++
	}

	return &ParkingLevel{
		Floor:        floor,
		MaxSize:      size,
		Size:         0,
		spots:        spots,
		spotsPerType: spotsPerType,
	}
}

func (pl *ParkingLevel) IsFull() bool {
	defer pl.l.RUnlock()
	pl.l.RLock()
	return pl.Size == pl.MaxSize
}

func (pl *ParkingLevel) CanPark(vehicle *Vehicle) bool {
	defer pl.l.RUnlock()
	pl.l.RLock()
	return pl.spotsPerType[vehicle.Type] > 0
}

func (pl *ParkingLevel) Park(vehicle *Vehicle) bool {
	defer pl.l.Unlock()
	pl.l.Lock()

	for _, spot := range pl.spots {
		if spot.IsFree() && spot.GetType() == vehicle.Type {
			if err := spot.In(vehicle); err != nil {
				return false
			}

			pl.Size++
			pl.spotsPerType[vehicle.Type]--
			return true
		}
	}

	return false
}

func (pl *ParkingLevel) Unpark(vehicle *Vehicle) bool {
	defer pl.l.Unlock()
	pl.l.Lock()

	for _, spot := range pl.spots {
		if spot.GetVehicle() == vehicle {
			spot.Out()
			pl.Size--
			pl.spotsPerType[vehicle.Type]++
			return true
		}
	}

	return false
}
