// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	apperrors "github.com/kyma-project/kyma/components/application-registry/internal/apperrors"

	clusterassetgroup "github.com/kyma-project/kyma/components/application-registry/internal/metadata/specification/rafter/clusterassetgroup"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *Service) Get(id string) ([]byte, []byte, []byte, apperrors.AppError) {
	ret := _m.Called(id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string) []byte); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 []byte
	if rf, ok := ret.Get(2).(func(string) []byte); ok {
		r2 = rf(id)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]byte)
		}
	}

	var r3 apperrors.AppError
	if rf, ok := ret.Get(3).(func(string) apperrors.AppError); ok {
		r3 = rf(id)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(apperrors.AppError)
		}
	}

	return r0, r1, r2, r3
}

// Put provides a mock function with given fields: id, apiType, documentation, apiSpec, eventsSpec
func (_m *Service) Put(id string, apiType clusterassetgroup.ApiType, documentation []byte, apiSpec []byte, eventsSpec []byte) apperrors.AppError {
	ret := _m.Called(id, apiType, documentation, apiSpec, eventsSpec)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, clusterassetgroup.ApiType, []byte, []byte, []byte) apperrors.AppError); ok {
		r0 = rf(id, apiType, documentation, apiSpec, eventsSpec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Remove provides a mock function with given fields: id
func (_m *Service) Remove(id string) apperrors.AppError {
	ret := _m.Called(id)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
