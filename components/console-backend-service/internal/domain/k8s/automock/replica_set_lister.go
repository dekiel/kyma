// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/apps/v1"

// replicaSetLister is an autogenerated mock type for the replicaSetLister type
type replicaSetLister struct {
	mock.Mock
}

// ListReplicaSets provides a mock function with given fields: namespace
func (_m *replicaSetLister) ListReplicaSets(namespace string) ([]*v1.ReplicaSet, error) {
	ret := _m.Called(namespace)

	var r0 []*v1.ReplicaSet
	if rf, ok := ret.Get(0).(func(string) []*v1.ReplicaSet); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.ReplicaSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
