package entity

type Car struct {
	PlateNumber string
}

func NewCarToPark(PlateNumber string) *Car {
	return &Car{PlateNumber: PlateNumber}
}
