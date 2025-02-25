// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/recipes/terraform/config/providers (interfaces: Provider)
//
// Generated by this command:
//
//	mockgen -typed -destination=./mock_provider.go -package=providers -self_package github.com/radius-project/radius/pkg/recipes/terraform/config/providers github.com/radius-project/radius/pkg/recipes/terraform/config/providers Provider
//

// Package providers is a generated GoMock package.
package providers

import (
	context "context"
	reflect "reflect"

	recipes "github.com/radius-project/radius/pkg/recipes"
	gomock "go.uber.org/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// BuildConfig mocks base method.
func (m *MockProvider) BuildConfig(arg0 context.Context, arg1 *recipes.Configuration) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildConfig", arg0, arg1)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildConfig indicates an expected call of BuildConfig.
func (mr *MockProviderMockRecorder) BuildConfig(arg0, arg1 any) *MockProviderBuildConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildConfig", reflect.TypeOf((*MockProvider)(nil).BuildConfig), arg0, arg1)
	return &MockProviderBuildConfigCall{Call: call}
}

// MockProviderBuildConfigCall wrap *gomock.Call
type MockProviderBuildConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProviderBuildConfigCall) Return(arg0 map[string]any, arg1 error) *MockProviderBuildConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProviderBuildConfigCall) Do(f func(context.Context, *recipes.Configuration) (map[string]any, error)) *MockProviderBuildConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProviderBuildConfigCall) DoAndReturn(f func(context.Context, *recipes.Configuration) (map[string]any, error)) *MockProviderBuildConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
