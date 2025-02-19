// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
)

// AuthSchemaClient is an autogenerated mock type for the AuthSchemaClient type
type AuthSchemaClient struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *AuthSchemaClient) Login(ctx context.Context, in *v1.LoginRequest, opts ...grpc.CallOption) (*v1.AuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *v1.AuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LoginRequest, ...grpc.CallOption) (*v1.AuthResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.LoginRequest, ...grpc.CallOption) *v1.AuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.AuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, in, opts
func (_m *AuthSchemaClient) Register(ctx context.Context, in *v1.RegisterRequest, opts ...grpc.CallOption) (*v1.AuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *v1.AuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RegisterRequest, ...grpc.CallOption) (*v1.AuthResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RegisterRequest, ...grpc.CallOption) *v1.AuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.AuthResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.RegisterRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthSchemaClient creates a new instance of AuthSchemaClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthSchemaClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthSchemaClient {
	mock := &AuthSchemaClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
