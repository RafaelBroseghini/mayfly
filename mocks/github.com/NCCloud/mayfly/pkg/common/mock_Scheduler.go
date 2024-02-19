// Code generated by mockery v2.42.0. DO NOT EDIT.

package common

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockScheduler is an autogenerated mock type for the Scheduler type
type MockScheduler struct {
	mock.Mock
}

type MockScheduler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockScheduler) EXPECT() *MockScheduler_Expecter {
	return &MockScheduler_Expecter{mock: &_m.Mock}
}

// CreateOrUpdateTask provides a mock function with given fields: tag, date, task
func (_m *MockScheduler) CreateOrUpdateTask(tag string, date time.Time, task func() error) error {
	ret := _m.Called(tag, date, task)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrUpdateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Time, func() error) error); ok {
		r0 = rf(tag, date, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockScheduler_CreateOrUpdateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrUpdateTask'
type MockScheduler_CreateOrUpdateTask_Call struct {
	*mock.Call
}

// CreateOrUpdateTask is a helper method to define mock.On call
//   - tag string
//   - date time.Time
//   - task func() error
func (_e *MockScheduler_Expecter) CreateOrUpdateTask(tag interface{}, date interface{}, task interface{}) *MockScheduler_CreateOrUpdateTask_Call {
	return &MockScheduler_CreateOrUpdateTask_Call{Call: _e.mock.On("CreateOrUpdateTask", tag, date, task)}
}

func (_c *MockScheduler_CreateOrUpdateTask_Call) Run(run func(tag string, date time.Time, task func() error)) *MockScheduler_CreateOrUpdateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(time.Time), args[2].(func() error))
	})
	return _c
}

func (_c *MockScheduler_CreateOrUpdateTask_Call) Return(_a0 error) *MockScheduler_CreateOrUpdateTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScheduler_CreateOrUpdateTask_Call) RunAndReturn(run func(string, time.Time, func() error) error) *MockScheduler_CreateOrUpdateTask_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTask provides a mock function with given fields: tag
func (_m *MockScheduler) DeleteTask(tag string) error {
	ret := _m.Called(tag)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockScheduler_DeleteTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTask'
type MockScheduler_DeleteTask_Call struct {
	*mock.Call
}

// DeleteTask is a helper method to define mock.On call
//   - tag string
func (_e *MockScheduler_Expecter) DeleteTask(tag interface{}) *MockScheduler_DeleteTask_Call {
	return &MockScheduler_DeleteTask_Call{Call: _e.mock.On("DeleteTask", tag)}
}

func (_c *MockScheduler_DeleteTask_Call) Run(run func(tag string)) *MockScheduler_DeleteTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockScheduler_DeleteTask_Call) Return(_a0 error) *MockScheduler_DeleteTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockScheduler_DeleteTask_Call) RunAndReturn(run func(string) error) *MockScheduler_DeleteTask_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockScheduler creates a new instance of MockScheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockScheduler {
	mock := &MockScheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
