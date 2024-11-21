package parking

import (
	"sort"
)

type SortStyle interface {
	AlgoStyle([]*ParkingLot) []*ParkingLot
}

type Default struct{}

func (d *Default) AlgoStyle(lots []*ParkingLot) []*ParkingLot {
	return lots
}

type MostCapacity struct{}

func (mc *MostCapacity) AlgoStyle(lots []*ParkingLot) []*ParkingLot {
	sort.SliceStable(lots, func(i, j int) bool {
		return lots[i].MaxCapacityLot > lots[j].MaxCapacityLot
	})
	return lots
}

type HighestNumber struct{}

func (hn *HighestNumber) AlgoStyle(lots []*ParkingLot) []*ParkingLot {
	sort.SliceStable(lots, func(i, j int) bool {
		return lots[i].MaxCapacityLot-len(lots[i].setCars) > lots[j].MaxCapacityLot-len(lots[j].setCars)
	})
	return lots
}
