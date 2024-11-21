package parking

import (
	"fmt"

	"parking-lot-cli/entity"
	"parking-lot-cli/variable"
)

type ObserverSubscriber interface {
	Update(*ParkingLot)
	UpdateParkNowAvailable(*ParkingLot)
}

type Attendant struct {
	Lot            []*ParkingLot
	LotUnavailable []*ParkingLot
	IdAttendant    string
	TypeStylePark  SortStyle
}

func NewAttendant(listOfLot []*ParkingLot, IdAttendant string) *Attendant {
	return &Attendant{
		Lot:            listOfLot,
		LotUnavailable: []*ParkingLot{},
		IdAttendant:    IdAttendant,
		TypeStylePark:  &Default{},
	}
}

func (a *Attendant) Park(car *entity.Car) (*entity.Ticket, error) {
	if a.HasTheSameCar(car) {
		fmt.Println(variable.ErrorSameCar)
		return nil, variable.ErrorSameCar
	}
	a.Lot = a.TypeStylePark.AlgoStyle(a.Lot)
	for _, v := range a.Lot {
		if !v.IsFull() {
			getTicket, _ := v.Park(car)
			if v.IsFull() {
				a.MoveToLotUnavailable(v)
			}
			return getTicket, nil
		}
	}
	fmt.Println(variable.ErrorCapacityPark)
	return nil, variable.ErrorCapacityPark
}

func (a *Attendant) HasTheSameCar(car *entity.Car) bool {
	for _, v := range a.Lot {
		if v.HasTheSameCar(car) {
			return true
		}
	}
	for _, v := range a.LotUnavailable {
		if v.HasTheSameCar(car) {
			return true
		}
	}
	return false
}

func (a *Attendant) UnPark(ticket *entity.Ticket) (*entity.Car, error) {
	for _, lotCar := range a.Lot {
		if lotCar.setCars[ticket] != nil {
			getCar, _ := lotCar.UnPark(ticket)
			return getCar, nil
		}
	}
	for _, lotCar := range a.LotUnavailable {
		if lotCar.setCars[ticket] != nil {
			getCar, _ := lotCar.UnPark(ticket)
			return getCar, nil
		}
	}
	fmt.Println(variable.ErrorWrongTicket)
	return nil, variable.ErrorWrongTicket
}

func (a *Attendant) Update(pl *ParkingLot) {
	fmt.Printf("INFORMATION!!! = Attendant: %v, Parking lot %v is full\n", a.IdAttendant, pl.IdLot)
}

func (a *Attendant) UpdateParkNowAvailable(lot *ParkingLot) {
	var tempLot []*ParkingLot
	tempLot = append(tempLot, a.Lot...)
	a.Lot = append(a.Lot[:0], lot)
	a.Lot = append(a.Lot, tempLot...)
	for idx, v := range a.LotUnavailable {
		if v == lot {
			a.LotUnavailable = append(a.LotUnavailable[:idx], a.LotUnavailable[idx+1:]...)
		}
	}
	fmt.Printf("Park lot with ID : %v is now available\n", lot.IdLot)
}

func (a *Attendant) AddRegister(attendant *Attendant) *Attendant {
	for _, lot := range attendant.Lot {
		lot.RegisterObserver(attendant)
	}
	return a
}

func (a *Attendant) AddLotToTheList(lot *ParkingLot) {
	a.Lot = append(a.Lot, lot)
}

func (a *Attendant) ChangeStylePark(style SortStyle) {
	a.TypeStylePark = style
}

func (a *Attendant) OptionStyle(option string) SortStyle {
	var style SortStyle
	switch option {
	case "default":
		style = &Default{}
	case "mostCapacity":
		style = &MostCapacity{}
	case "highestNumber":
		style = &HighestNumber{}
	}
	return style
}

func (a *Attendant) MoveToLotUnavailable(lot *ParkingLot) {
	a.LotUnavailable = append(a.LotUnavailable, lot)
	for idx, v := range a.Lot {
		if v == lot {
			a.Lot = append(a.Lot[:idx], a.Lot[idx+1:]...)
		}
	}
}
