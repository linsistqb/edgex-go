// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// IntervalDeleter is an autogenerated mock type for the IntervalDeleter type
type IntervalDeleter struct {
	mock.Mock
}

// DeleteIntervalById provides a mock function with given fields: id
func (_m *IntervalDeleter) DeleteIntervalById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IntervalActionsByIntervalName provides a mock function with given fields: name
func (_m *IntervalDeleter) IntervalActionsByIntervalName(name string) ([]models.IntervalAction, error) {
	ret := _m.Called(name)

	var r0 []models.IntervalAction
	if rf, ok := ret.Get(0).(func(string) []models.IntervalAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.IntervalAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalById provides a mock function with given fields: id
func (_m *IntervalDeleter) IntervalById(id string) (models.Interval, error) {
	ret := _m.Called(id)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IntervalByName provides a mock function with given fields: name
func (_m *IntervalDeleter) IntervalByName(name string) (models.Interval, error) {
	ret := _m.Called(name)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryIntervalByID provides a mock function with given fields: intervalId
func (_m *IntervalDeleter) QueryIntervalByID(intervalId string) (models.Interval, error) {
	ret := _m.Called(intervalId)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(intervalId)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(intervalId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryIntervalByName provides a mock function with given fields: intervalName
func (_m *IntervalDeleter) QueryIntervalByName(intervalName string) (models.Interval, error) {
	ret := _m.Called(intervalName)

	var r0 models.Interval
	if rf, ok := ret.Get(0).(func(string) models.Interval); ok {
		r0 = rf(intervalName)
	} else {
		r0 = ret.Get(0).(models.Interval)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(intervalName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveIntervalInQueue provides a mock function with given fields: intervalId
func (_m *IntervalDeleter) RemoveIntervalInQueue(intervalId string) error {
	ret := _m.Called(intervalId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(intervalId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
