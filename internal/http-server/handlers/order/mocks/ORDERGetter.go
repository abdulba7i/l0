// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	postgres "l0wb/internal/storage/postgres"
)

// ORDERGetter is an autogenerated mock type for the ORDERGetter type
type ORDERGetter struct {
	mock.Mock
}

// GetOrderById provides a mock function with given fields: id
func (_m *ORDERGetter) GetOrderById(id string) (postgres.Order, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderById")
	}

	var r0 postgres.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (postgres.Order, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) postgres.Order); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(postgres.Order)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewORDERGetter creates a new instance of ORDERGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORDERGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORDERGetter {
	mock := &ORDERGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
