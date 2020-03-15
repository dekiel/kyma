// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	apperrors "github.com/kyma-project/kyma/components/application-gateway/pkg/apperrors"
	authorization "github.com/kyma-project/kyma/components/application-gateway/pkg/authorization"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Strategy is an autogenerated mock type for the Strategy type
type Strategy struct {
	mock.Mock
}

// AddAuthorization provides a mock function with given fields: r, setter
func (_m *Strategy) AddAuthorization(r *http.Request, setter authorization.TransportSetter) apperrors.AppError {
	ret := _m.Called(r, setter)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(*http.Request, authorization.TransportSetter) apperrors.AppError); ok {
		r0 = rf(r, setter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Invalidate provides a mock function with given fields:
func (_m *Strategy) Invalidate() {
	_m.Called()
}
