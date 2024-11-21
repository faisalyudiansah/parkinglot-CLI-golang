package variable

import "errors"

var (
	ErrorWrongTicket    = errors.New("unrecognized parking ticket")
	ErrorCapacityPark   = errors.New("no available position")
	ErrorSameCar        = errors.New("should not be able to park twice with a same car")
	ErrorLotIsEmpty     = errors.New("the parking lot is currently empty")
	ErrorInputNotNumber = errors.New("input must be numbers")
	ErrorInputInvalid   = errors.New("input invalid")
	ErrorThereIsNoCar   = errors.New("not a single car was parked")
)
