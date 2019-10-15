// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import store "github.com/eko/gache/store"

// StoreInterface is an autogenerated mock type for the StoreInterface type
type StoreInterface struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *StoreInterface) Get(key interface{}) (interface{}, error) {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetType provides a mock function with given fields:
func (_m *StoreInterface) GetType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Set provides a mock function with given fields: key, value, options
func (_m *StoreInterface) Set(key interface{}, value interface{}, options *store.Options) error {
	ret := _m.Called(key, value, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, *store.Options) error); ok {
		r0 = rf(key, value, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
