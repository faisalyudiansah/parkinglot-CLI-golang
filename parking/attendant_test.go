package parking_test

import (
	"testing"

	"parking-lot-cli/entity"
	"parking-lot-cli/parking"
	"parking-lot-cli/parking/mocks"
	"parking-lot-cli/variable"

	"github.com/stretchr/testify/assert"
)

func TestAttendant(t *testing.T) {
	t.Run("should be successful to park and unpark the car when the attendant gets the car", func(t *testing.T) {
		slotWithOneCapacity := parking.NewLot(1, "idLot-1")
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")
		newAttendant.AddLotToTheList(slotWithOneCapacity)
		getTicket, _ := newAttendant.Park(entity.NewCarToPark("BN-9304-II"))
		getCar, _ := newAttendant.UnPark(getTicket)

		assert.NotNil(t, getTicket)
		assert.NotNil(t, getCar)
	})

	t.Run("should be an error when the parking lot is full when the attendant wants to park the car", func(t *testing.T) {
		slotWithOneCapacity := parking.NewLot(1, "idLot-1")
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")
		newAttendant.AddLotToTheList(slotWithOneCapacity)

		newAttendant.Park(entity.NewCarToPark("BN-3027-CX"))
		_, err := newAttendant.Park(entity.NewCarToPark("BN-1022-DD"))

		expectedResult := variable.ErrorCapacityPark

		assert.NotNil(t, err)
		assert.Equal(t, expectedResult, err)
	})

	t.Run("should success when attendant register to observerlist park a car", func(t *testing.T) {
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")
		observer := newAttendant.AddRegister(newAttendant)
		assert.NotNil(t, observer)
	})

	t.Run("should return an error if the ticket provided is wrong or invalid when you want to unpark the car", func(t *testing.T) {
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.Park(entity.NewCarToPark("BN-3027-CX"))
		wrongTicket := entity.NewTicket()
		_, err := newAttendant.UnPark(&wrongTicket)
		expectedResult := variable.ErrorWrongTicket

		assert.NotNil(t, wrongTicket)
		assert.NotNil(t, err)
		assert.Equal(t, expectedResult, err)
	})

	t.Run("should be error when car parked twice even in different parking lot", func(t *testing.T) {
		NewLot := parking.NewLot(2, "idLot-1")
		anotherNewLot := parking.NewLot(2, "idLot-2")

		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.AddLotToTheList(NewLot)
		newAttendant.AddLotToTheList(anotherNewLot)

		car := entity.NewCarToPark("BN-2029-CN")

		// lot 1
		newAttendant.Park(car)
		newAttendant.Park(entity.NewCarToPark("BN-1230-LP"))

		// lot 2
		_, err := newAttendant.Park(car)
		newAttendant.Park(entity.NewCarToPark("BN-900-OP"))

		expectedResult := variable.ErrorSameCar

		assert.Equal(t, expectedResult, err)
	})

	t.Run("should get notification when park a lot back to available (after full) for subscription", func(t *testing.T) {
		newLot := parking.NewLot(3, "idLot-1")
		anotherNewLot := parking.NewLot(3, "idLot-2")
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.AddLotToTheList(newLot)
		newAttendant.AddLotToTheList(anotherNewLot)

		newAttendant.AddRegister(newAttendant)

		subscriberMock := &mocks.ObserverSubscriber{}
		subscriberMock.On("Update", newAttendant.Lot[0])
		subscriberMock.On("Update", newAttendant.Lot[1])
		subscriberMock.On("UpdateParkNowAvailable", newAttendant.Lot[0])
		subscriberMock.On("UpdateParkNowAvailable", newAttendant.Lot[1])
		newAttendant.Lot[0].RegisterObserver(subscriberMock)
		newAttendant.Lot[1].RegisterObserver(subscriberMock)

		newAttendant.Park(entity.NewCarToPark("BN-4044-BB"))
		ticket1, _ := newAttendant.Park(entity.NewCarToPark("BN-4899-IR"))
		ticket2, _ := newAttendant.Park(entity.NewCarToPark("BN-3838-LI"))
		newAttendant.UnPark(ticket2)
		newAttendant.UnPark(ticket1)
		newAttendant.Park(entity.NewCarToPark("BN-4209-IR"))
		newAttendant.Park(entity.NewCarToPark("BN-1838-WI"))

		assert.NotNil(t, ticket1)
		assert.NotNil(t, ticket2)
		subscriberMock.AssertNumberOfCalls(t, "Update", 2)
		subscriberMock.AssertNumberOfCalls(t, "UpdateParkNowAvailable", 1)
	})

	t.Run("should trigger the update method and the UpdateParkNowAvailable method on the attendant section when park a lot back to available (after full) for subscription", func(t *testing.T) {
		newLot := parking.NewLot(3, "idLot-1")
		anotherNewLot := parking.NewLot(3, "idLot-2")
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.AddLotToTheList(newLot)
		newAttendant.AddLotToTheList(anotherNewLot)

		newAttendant.AddRegister(newAttendant)

		subscriberMock := &mocks.ObserverSubscriber{}
		subscriberMock.On("Update", newAttendant.Lot[0])
		subscriberMock.On("UpdateParkNowAvailable", newAttendant.Lot[0])
		newAttendant.AddRegister(newAttendant)

		// SLOT 1
		ticket, _ := newAttendant.Park(entity.NewCarToPark("BN-3432-OR"))
		newAttendant.Park(entity.NewCarToPark("BN-3044-SD"))
		newAttendant.Park(entity.NewCarToPark("BN-4938-L3"))

		// SLOT 2
		newAttendant.Park(entity.NewCarToPark("BN-2003-LO"))
		newAttendant.Park(entity.NewCarToPark("BN-4920-CN"))

		getCar, _ := newAttendant.UnPark(ticket)

		newAttendant.Park(entity.NewCarToPark("BN-2343-HH"))
		newAttendant.Park(entity.NewCarToPark("BN-4938-EE"))

		newAttendant.Update(newAttendant.Lot[0])
		newAttendant.UpdateParkNowAvailable(newAttendant.Lot[0])

		assert.NotNil(t, ticket)
		assert.NotNil(t, getCar.PlateNumber)
		expectedResult := "BN-3432-OR"
		assert.Equal(t, expectedResult, getCar.PlateNumber)
	})

	t.Run("should error if park the same car when attendant park a car in the available lot", func(t *testing.T) {
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")
		newLot := parking.NewLot(3, "idLot-1")
		newAttendant.AddLotToTheList(newLot)

		car := entity.NewCarToPark("BN-3432-OR")
		ticket, _ := newAttendant.Park(car)
		_, err := newAttendant.Park(car)

		assert.NotNil(t, ticket)
		assert.NotNil(t, err)

		expectedResult := variable.ErrorSameCar
		assert.Equal(t, expectedResult, err)
	})
}
