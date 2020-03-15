// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/apps/v1"

// gqlVersionInfoConverter is an autogenerated mock type for the gqlVersionInfoConverter type
type gqlVersionInfoConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlVersionInfoConverter) ToGQL(in *v1.Deployment) gqlschema.VersionInfo {
	ret := _m.Called(in)

	var r0 gqlschema.VersionInfo
	if rf, ok := ret.Get(0).(func(*v1.Deployment) gqlschema.VersionInfo); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(gqlschema.VersionInfo)
	}

	return r0
}
