package parking

import (
	"fmt"

	"parking-lot-cli/entity"
	"parking-lot-cli/variable"
)

type ParkingLot struct {
	setCars        map[*entity.Ticket]*entity.Car
	MaxCapacityLot int
	IdLot          string
	observerList   []ObserverSubscriber
}

func NewLot(capacity int, IdLot string) *ParkingLot {
	return &ParkingLot{
		setCars:        make(map[*entity.Ticket]*entity.Car),
		MaxCapacityLot: capacity,
		IdLot:          IdLot,
		observerList:   []ObserverSubscriber{},
	}
}

func (pl *ParkingLot) Park(car *entity.Car) (*entity.Ticket, error) {
	if pl.IsFull() {
		fmt.Println(variable.ErrorCapacityPark)
		return nil, variable.ErrorCapacityPark
	}
	if pl.HasTheSameCar(car) {
		fmt.Println(variable.ErrorSameCar)
		return nil, variable.ErrorSameCar
	}
	generateTicket := entity.NewTicket()
	pl.setCars[&generateTicket] = car
	fmt.Printf("[PARK %v], Car : %v, Ticket : %v\n", pl.IdLot, car.PlateNumber, generateTicket.ID)
	if pl.IsFull() {
		pl.NotifyAll()
	}
	return &generateTicket, nil
}

func (pl *ParkingLot) UnPark(ticket *entity.Ticket) (*entity.Car, error) {
	if _, ok := pl.setCars[ticket]; !ok {
		fmt.Println(variable.ErrorWrongTicket)
		return nil, variable.ErrorWrongTicket
	}
	getCar := pl.setCars[ticket]
	fmt.Printf("[UNPARK %v], Car : %v, Unpark sucessfully\n", pl.IdLot, getCar.PlateNumber)
	if pl.IsFull() {
		pl.NotifyForLotIsAvailable()
	}
	delete(pl.setCars, ticket)
	return getCar, nil
}

func (pl *ParkingLot) IsFull() bool {
	return len(pl.setCars) == pl.MaxCapacityLot
}

func (pl *ParkingLot) HasTheSameCar(car *entity.Car) bool {
	for _, carInPark := range pl.setCars {
		if carInPark == car {
			return true
		}
	}
	return false
}

func (pl *ParkingLot) NotifyAll() {
	for _, observer := range pl.observerList {
		observer.Update(pl)
	}
}

func (pl *ParkingLot) NotifyForLotIsAvailable() {
	for _, observer := range pl.observerList {
		observer.UpdateParkNowAvailable(pl)
	}
}

func (pl *ParkingLot) RegisterObserver(a ObserverSubscriber) {
	pl.observerList = append(pl.observerList, a)
}
