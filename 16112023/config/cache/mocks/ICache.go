// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// ICache is an autogenerated mock type for the ICache type
type ICache struct {
	mock.Mock
}

type ICache_Expecter struct {
	mock *mock.Mock
}

func (_m *ICache) EXPECT() *ICache_Expecter {
	return &ICache_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, key
func (_m *ICache) Get(ctx context.Context, key string) ([]byte, error) {
	ret := _m.Called(ctx, key)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICache_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ICache_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *ICache_Expecter) Get(ctx interface{}, key interface{}) *ICache_Get_Call {
	return &ICache_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *ICache_Get_Call) Run(run func(ctx context.Context, key string)) *ICache_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ICache_Get_Call) Return(_a0 []byte, _a1 error) *ICache_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ICache_Get_Call) RunAndReturn(run func(context.Context, string) ([]byte, error)) *ICache_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields: ctx
func (_m *ICache) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICache_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type ICache_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ICache_Expecter) Ping(ctx interface{}) *ICache_Ping_Call {
	return &ICache_Ping_Call{Call: _e.mock.On("Ping", ctx)}
}

func (_c *ICache_Ping_Call) Run(run func(ctx context.Context)) *ICache_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ICache_Ping_Call) Return(_a0 error) *ICache_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ICache_Ping_Call) RunAndReturn(run func(context.Context) error) *ICache_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value, expiration
func (_m *ICache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	ret := _m.Called(ctx, key, value, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICache_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type ICache_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
//   - expiration time.Duration
func (_e *ICache_Expecter) Set(ctx interface{}, key interface{}, value interface{}, expiration interface{}) *ICache_Set_Call {
	return &ICache_Set_Call{Call: _e.mock.On("Set", ctx, key, value, expiration)}
}

func (_c *ICache_Set_Call) Run(run func(ctx context.Context, key string, value interface{}, expiration time.Duration)) *ICache_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *ICache_Set_Call) Return(_a0 error) *ICache_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ICache_Set_Call) RunAndReturn(run func(context.Context, string, interface{}, time.Duration) error) *ICache_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewICache creates a new instance of ICache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICache(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICache {
	mock := &ICache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
