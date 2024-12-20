// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	parking "parking-lot-cli/parking"
	mock "github.com/stretchr/testify/mock"
)

// SortStyle is an autogenerated mock type for the SortStyle type
type SortStyle struct {
	mock.Mock
}

// AlgoStyle provides a mock function with given fields: _a0
func (_m *SortStyle) AlgoStyle(_a0 []*parking.ParkingLot) []*parking.ParkingLot {
	ret := _m.Called(_a0)

	var r0 []*parking.ParkingLot
	if rf, ok := ret.Get(0).(func([]*parking.ParkingLot) []*parking.ParkingLot); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*parking.ParkingLot)
		}
	}

	return r0
}

type mockConstructorTestingTNewSortStyle interface {
	mock.TestingT
	Cleanup(func())
}

// NewSortStyle creates a new instance of SortStyle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSortStyle(t mockConstructorTestingTNewSortStyle) *SortStyle {
	mock := &SortStyle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
