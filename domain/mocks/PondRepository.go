// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "go-template/domain"

	mock "github.com/stretchr/testify/mock"
)

// PondRepository is an autogenerated mock type for the PondRepository type
type PondRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: pond
func (_m *PondRepository) Create(pond domain.Pond) (uint, error) {
	ret := _m.Called(pond)

	var r0 uint
	if rf, ok := ret.Get(0).(func(domain.Pond) uint); ok {
		r0 = rf(pond)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Pond) error); ok {
		r1 = rf(pond)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: pond
func (_m *PondRepository) Delete(pond domain.Pond) error {
	ret := _m.Called(pond)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Pond) error); ok {
		r0 = rf(pond)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *PondRepository) GetAll() ([]domain.Pond, error) {
	ret := _m.Called()

	var r0 []domain.Pond
	if rf, ok := ret.Get(0).(func() []domain.Pond); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Pond)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByFarmId provides a mock function with given fields: id
func (_m *PondRepository) GetByFarmId(id uint) ([]domain.Pond, error) {
	ret := _m.Called(id)

	var r0 []domain.Pond
	if rf, ok := ret.Get(0).(func(uint) []domain.Pond); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Pond)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *PondRepository) GetById(id uint) (domain.Pond, error) {
	ret := _m.Called(id)

	var r0 domain.Pond
	if rf, ok := ret.Get(0).(func(uint) domain.Pond); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Pond)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: pond
func (_m *PondRepository) Update(pond domain.Pond) error {
	ret := _m.Called(pond)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Pond) error); ok {
		r0 = rf(pond)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}