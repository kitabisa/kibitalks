// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IRabbitPublisher is an autogenerated mock type for the IRabbitPublisher type
type IRabbitPublisher struct {
	mock.Mock
}

type IRabbitPublisher_Expecter struct {
	mock *mock.Mock
}

func (_m *IRabbitPublisher) EXPECT() *IRabbitPublisher_Expecter {
	return &IRabbitPublisher_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: queueName, message
func (_m *IRabbitPublisher) Publish(queueName string, message string) error {
	ret := _m.Called(queueName, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(queueName, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRabbitPublisher_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type IRabbitPublisher_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - queueName string
//   - message string
func (_e *IRabbitPublisher_Expecter) Publish(queueName interface{}, message interface{}) *IRabbitPublisher_Publish_Call {
	return &IRabbitPublisher_Publish_Call{Call: _e.mock.On("Publish", queueName, message)}
}

func (_c *IRabbitPublisher_Publish_Call) Run(run func(queueName string, message string)) *IRabbitPublisher_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *IRabbitPublisher_Publish_Call) Return(_a0 error) *IRabbitPublisher_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRabbitPublisher_Publish_Call) RunAndReturn(run func(string, string) error) *IRabbitPublisher_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// NewIRabbitPublisher creates a new instance of IRabbitPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRabbitPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRabbitPublisher {
	mock := &IRabbitPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
