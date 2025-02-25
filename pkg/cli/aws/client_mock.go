// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/radius-project/radius/pkg/cli/aws (interfaces: Client)
//
// Generated by this command:
//
//	mockgen -typed -destination=./client_mock.go -package=aws -self_package github.com/radius-project/radius/pkg/cli/aws github.com/radius-project/radius/pkg/cli/aws Client
//

// Package aws is a generated GoMock package.
package aws

import (
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	sts "github.com/aws/aws-sdk-go-v2/service/sts"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetCallerIdentity mocks base method.
func (m *MockClient) GetCallerIdentity(arg0 context.Context) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentity", arg0)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity.
func (mr *MockClientMockRecorder) GetCallerIdentity(arg0 any) *MockClientGetCallerIdentityCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockClient)(nil).GetCallerIdentity), arg0)
	return &MockClientGetCallerIdentityCall{Call: call}
}

// MockClientGetCallerIdentityCall wrap *gomock.Call
type MockClientGetCallerIdentityCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientGetCallerIdentityCall) Return(arg0 *sts.GetCallerIdentityOutput, arg1 error) *MockClientGetCallerIdentityCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientGetCallerIdentityCall) Do(f func(context.Context) (*sts.GetCallerIdentityOutput, error)) *MockClientGetCallerIdentityCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientGetCallerIdentityCall) DoAndReturn(f func(context.Context) (*sts.GetCallerIdentityOutput, error)) *MockClientGetCallerIdentityCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListRegions mocks base method.
func (m *MockClient) ListRegions(arg0 context.Context) (*ec2.DescribeRegionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRegions", arg0)
	ret0, _ := ret[0].(*ec2.DescribeRegionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegions indicates an expected call of ListRegions.
func (mr *MockClientMockRecorder) ListRegions(arg0 any) *MockClientListRegionsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegions", reflect.TypeOf((*MockClient)(nil).ListRegions), arg0)
	return &MockClientListRegionsCall{Call: call}
}

// MockClientListRegionsCall wrap *gomock.Call
type MockClientListRegionsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClientListRegionsCall) Return(arg0 *ec2.DescribeRegionsOutput, arg1 error) *MockClientListRegionsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClientListRegionsCall) Do(f func(context.Context) (*ec2.DescribeRegionsOutput, error)) *MockClientListRegionsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClientListRegionsCall) DoAndReturn(f func(context.Context) (*ec2.DescribeRegionsOutput, error)) *MockClientListRegionsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
