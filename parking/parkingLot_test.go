package parking

import (
	"testing"

	"parking-lot-cli/entity"
	"parking-lot-cli/variable"

	"github.com/stretchr/testify/assert"
)

func TestNewLot(t *testing.T) {
	t.Run("should successfully to return a ticket when customer successfully park the car", func(t *testing.T) {
		car := entity.NewCarToPark("BN-3740-OU")
		newParkCar := NewLot(2, "idLot-1")
		getTicket, _ := newParkCar.Park(car)
		expectedResult := getTicket

		assert.Equal(t, expectedResult, getTicket)
	})
	t.Run("should successfully to return a car when customer successfully unpark the car", func(t *testing.T) {
		expectedCar := entity.NewCarToPark("BN-3740-OU")
		newParkCar := NewLot(2, "idLot-1")
		getTicket, _ := newParkCar.Park(expectedCar)
		getCar, _ := newParkCar.UnPark(getTicket)

		assert.Equal(t, expectedCar, getCar)
	})
	t.Run("should to return an error when customer use the wrong ticket", func(t *testing.T) {
		getTicket := entity.NewTicket()
		car := entity.NewCarToPark("BN-3740-OU")
		newParkCar := NewLot(2, "idLot-1")
		newParkCar.Park(car)
		_, err := newParkCar.UnPark(&getTicket)
		expectedResult := variable.ErrorWrongTicket

		assert.Equal(t, expectedResult, err)
	})
	t.Run("should return an error when parking lot is full", func(t *testing.T) {
		car := entity.NewCarToPark("BN-0433-OP")
		newParkCar := NewLot(1, "idLot-1")
		newParkCar.Park(car)

		car4 := entity.NewCarToPark("BN-2233-RR")
		_, err := newParkCar.Park(car4)

		expectedResult := variable.ErrorCapacityPark

		assert.Equal(t, expectedResult, err)
	})
	t.Run("should not be able when to park a car or park twice", func(t *testing.T) {
		car := entity.NewCarToPark("BN-1200-FF")
		newParkCar := NewLot(5, "idLot-1")
		newParkCar.Park(car)
		_, err := newParkCar.Park(car)

		expectedResult := variable.ErrorSameCar

		assert.Equal(t, expectedResult, err, "exp : %v, actuall : %v", expectedResult, err)
	})
}
