// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TaskService is an autogenerated mock type for the TaskService type
type TaskService struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: input
func (_m *TaskService) AddTask(input []string) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeTaskStatus provides a mock function with given fields: i, active
func (_m *TaskService) ChangeTaskStatus(i int, active bool) error {
	ret := _m.Called(i, active)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, bool) error); ok {
		r0 = rf(i, active)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CleanDoneTasks provides a mock function with given fields:
func (_m *TaskService) CleanDoneTasks() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTasks provides a mock function with given fields:
func (_m *TaskService) GetTasks() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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
