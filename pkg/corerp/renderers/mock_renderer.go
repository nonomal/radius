// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/corerp/renderers (interfaces: Renderer)

// Package renderers is a generated GoMock package.
package renderers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	resources "github.com/radius-project/radius/pkg/ucp/resources"
)

// MockRenderer is a mock of Renderer interface.
type MockRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockRendererMockRecorder
}

// MockRendererMockRecorder is the mock recorder for MockRenderer.
type MockRendererMockRecorder struct {
	mock *MockRenderer
}

// NewMockRenderer creates a new mock instance.
func NewMockRenderer(ctrl *gomock.Controller) *MockRenderer {
	mock := &MockRenderer{ctrl: ctrl}
	mock.recorder = &MockRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderer) EXPECT() *MockRendererMockRecorder {
	return m.recorder
}

// GetDependencyIDs mocks base method.
func (m *MockRenderer) GetDependencyIDs(arg0 context.Context, arg1 v1.DataModelInterface) ([]*resources.ID, []*resources.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDependencyIDs", arg0, arg1)
	ret0, _ := ret[0].([]*resources.ID)
	ret1, _ := ret[1].([]*resources.ID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDependencyIDs indicates an expected call of GetDependencyIDs.
func (mr *MockRendererMockRecorder) GetDependencyIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDependencyIDs", reflect.TypeOf((*MockRenderer)(nil).GetDependencyIDs), arg0, arg1)
}

// Render mocks base method.
func (m *MockRenderer) Render(arg0 context.Context, arg1 v1.DataModelInterface, arg2 RenderOptions) (RendererOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1, arg2)
	ret0, _ := ret[0].(RendererOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render.
func (mr *MockRendererMockRecorder) Render(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockRenderer)(nil).Render), arg0, arg1, arg2)
}
