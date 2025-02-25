// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/cli/kubernetes/logstream (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -typed -destination=./mock_logstream.go -package=logstream -self_package github.com/radius-project/radius/pkg/cli/kubernetes/logstream github.com/radius-project/radius/pkg/cli/kubernetes/logstream Interface
//

// Package logstream is a generated GoMock package.
package logstream

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Stream mocks base method.
func (m *MockInterface) Stream(arg0 context.Context, arg1 Options) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockInterfaceMockRecorder) Stream(arg0, arg1 any) *MockInterfaceStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockInterface)(nil).Stream), arg0, arg1)
	return &MockInterfaceStreamCall{Call: call}
}

// MockInterfaceStreamCall wrap *gomock.Call
type MockInterfaceStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceStreamCall) Return(arg0 error) *MockInterfaceStreamCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceStreamCall) Do(f func(context.Context, Options) error) *MockInterfaceStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceStreamCall) DoAndReturn(f func(context.Context, Options) error) *MockInterfaceStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
