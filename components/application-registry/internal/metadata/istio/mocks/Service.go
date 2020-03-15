// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	apperrors "github.com/kyma-project/kyma/components/application-registry/internal/apperrors"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: application, appUID, serviceId, resourceName
func (_m *Service) Create(application string, appUID types.UID, serviceId string, resourceName string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, resourceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, resourceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: resourceName
func (_m *Service) Delete(resourceName string) apperrors.AppError {
	ret := _m.Called(resourceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(resourceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Upsert provides a mock function with given fields: application, appUID, serviceId, resourceName
func (_m *Service) Upsert(application string, appUID types.UID, serviceId string, resourceName string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, resourceName)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, resourceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
