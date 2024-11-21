package parking_test

import (
	"testing"

	"parking-lot-cli/entity"
	"parking-lot-cli/parking"
	"parking-lot-cli/parking/mocks"

	"github.com/stretchr/testify/assert"
)

func TestSortStyle(t *testing.T) {
	t.Run("should successfully park the car using the default sort style when without specifying which style to use", func(t *testing.T) {
		slotWithOneCapacity := parking.NewLot(1, "idLot-1")
		slotWithFiveCapacity := parking.NewLot(5, "idLot-2")
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.AddLotToTheList(slotWithOneCapacity)
		newAttendant.AddLotToTheList(slotWithFiveCapacity)

		newAttendant.AddRegister(newAttendant)

		getTicket, _ := newAttendant.Park(entity.NewCarToPark("BN-3421-EE"))
		getCar, _ := newAttendant.UnPark(getTicket)

		assert.NotNil(t, getTicket)
		assert.Equal(t, "BN-3421-EE", getCar.PlateNumber)
	})

	t.Run("should success to park cars when use default sort style", func(t *testing.T) {
		slotWithOneCapacity := parking.NewLot(1, "idLot-1")
		slotWithFiveCapacity := parking.NewLot(5, "idLot-2")
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.AddLotToTheList(slotWithOneCapacity)
		newAttendant.AddLotToTheList(slotWithFiveCapacity)

		newAttendant.ChangeStylePark(newAttendant.OptionStyle("default"))

		newAttendant.AddRegister(newAttendant)

		subscriberMock := &mocks.ObserverSubscriber{}
		subscriberMock.On("Update", newAttendant.Lot[0])
		subscriberMock.On("Update", newAttendant.Lot[1])
		subscriberMock.On("UpdateParkNowAvailable", newAttendant.Lot[0])
		subscriberMock.On("UpdateParkNowAvailable", newAttendant.Lot[1])
		newAttendant.Lot[0].RegisterObserver(subscriberMock)
		newAttendant.Lot[1].RegisterObserver(subscriberMock)

		getTicket, _ := newAttendant.Park(entity.NewCarToPark("BN-3421-EE"))
		getCar, _ := newAttendant.UnPark(getTicket)

		assert.NotNil(t, getTicket)
		assert.Equal(t, "BN-3421-EE", getCar.PlateNumber)
		subscriberMock.AssertNumberOfCalls(t, "Update", 1)
		subscriberMock.AssertNumberOfCalls(t, "UpdateParkNowAvailable", 1)
	})

	t.Run("should success to park cars when use most capacity sort style", func(t *testing.T) {
		slotWithOneCapacity := parking.NewLot(1, "idLot-1")
		slotWithFiveCapacity := parking.NewLot(5, "idLot-2")
		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.AddLotToTheList(slotWithOneCapacity)
		newAttendant.AddLotToTheList(slotWithFiveCapacity)

		newAttendant.AddRegister(newAttendant)

		subscriberMock := &mocks.ObserverSubscriber{}
		subscriberMock.On("Update", newAttendant.Lot[0])
		subscriberMock.On("Update", newAttendant.Lot[1])
		subscriberMock.On("UpdateParkNowAvailable", newAttendant.Lot[0])
		subscriberMock.On("UpdateParkNowAvailable", newAttendant.Lot[1])
		newAttendant.Lot[0].RegisterObserver(subscriberMock)
		newAttendant.Lot[1].RegisterObserver(subscriberMock)

		sortStyleMock := &mocks.SortStyle{}
		sortStyleMock.On("AlgoStyle", newAttendant.Lot).Return(newAttendant.Lot)
		newAttendant.ChangeStylePark(newAttendant.OptionStyle("mostCapacity"))

		_, err1 := newAttendant.Park(entity.NewCarToPark("BN-1134-CN"))
		_, err2 := newAttendant.Park(entity.NewCarToPark("BN-2344-CN"))
		_, err3 := newAttendant.Park(entity.NewCarToPark("BN-3314-CN"))
		_, err4 := newAttendant.Park(entity.NewCarToPark("BN-4302-CN"))
		newAttendant.ChangeStylePark(sortStyleMock)
		getTicket, err5 := newAttendant.Park(entity.NewCarToPark("BN-3344-CN"))

		getCar, errUnpark := newAttendant.UnPark(getTicket)

		assert.Nil(t, err1)
		assert.Nil(t, err2)
		assert.Nil(t, err3)
		assert.Nil(t, err4)
		assert.Nil(t, err5)
		assert.Nil(t, errUnpark)
		assert.NotNil(t, getTicket)
		subscriberMock.AssertNumberOfCalls(t, "Update", 1)
		subscriberMock.AssertNumberOfCalls(t, "UpdateParkNowAvailable", 1)
		sortStyleMock.AssertNumberOfCalls(t, "AlgoStyle", 1)
		assert.Equal(t, "BN-3344-CN", getCar.PlateNumber)
	})

	t.Run("should success to park cars when use highest number free capacity sort style", func(t *testing.T) {
		slotWithTwoCapacity := parking.NewLot(2, "idLot-1")
		slotWithFiveCapacity := parking.NewLot(5, "idLot-2")
		slotWithSevenCapacity := parking.NewLot(7, "idLot-3")

		newAttendant := parking.NewAttendant([]*parking.ParkingLot{}, "att-01")

		newAttendant.AddLotToTheList(slotWithTwoCapacity)
		newAttendant.AddLotToTheList(slotWithFiveCapacity)
		newAttendant.AddLotToTheList(slotWithSevenCapacity)

		newAttendant.ChangeStylePark(newAttendant.OptionStyle("highestNumber"))

		newAttendant.AddRegister(newAttendant)

		newAttendant.Park(entity.NewCarToPark("BN-1134-CN"))
		newAttendant.Park(entity.NewCarToPark("BN-2131-CN"))
		newAttendant.Park(entity.NewCarToPark("BN-9888-CN"))

		carToIdLot2 := entity.NewCarToPark("BN-1234-CN")
		ticketOnIdLot2, _ := newAttendant.Park(carToIdLot2)
		getCar, _ := slotWithFiveCapacity.UnPark(ticketOnIdLot2)

		newAttendant.Park(entity.NewCarToPark("BN-1114-CN"))
		newAttendant.Park(entity.NewCarToPark("BN-5131-CN"))
		newAttendant.Park(entity.NewCarToPark("BN-3132-CN"))

		assert.NotNil(t, ticketOnIdLot2)
		assert.Same(t, newAttendant.Lot[0], slotWithSevenCapacity)
		assert.Same(t, newAttendant.Lot[1], slotWithFiveCapacity)
		assert.Same(t, newAttendant.Lot[2], slotWithTwoCapacity)
		assert.Same(t, carToIdLot2, getCar)
	})
}
