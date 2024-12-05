// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: virtualmachineclient/interface.go
//
// Generated by this command:
//
//	mockgen -package mock_virtualmachineclient -source virtualmachineclient/interface.go -typed -write_generate_directive -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt
//

// Package mock_virtualmachineclient is a generated GoMock package.
package mock_virtualmachineclient

import (
	context "context"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -package mock_virtualmachineclient -source virtualmachineclient/interface.go -typed -write_generate_directive -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
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

// BeginAttachDetachDataDisks mocks base method.
func (m *MockInterface) BeginAttachDetachDataDisks(ctx context.Context, resourceGroupName, vmName string, parameters armcompute.AttachDetachDataDisksRequest, options *armcompute.VirtualMachinesClientBeginAttachDetachDataDisksOptions) (*runtime.Poller[armcompute.VirtualMachinesClientAttachDetachDataDisksResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginAttachDetachDataDisks", ctx, resourceGroupName, vmName, parameters, options)
	ret0, _ := ret[0].(*runtime.Poller[armcompute.VirtualMachinesClientAttachDetachDataDisksResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginAttachDetachDataDisks indicates an expected call of BeginAttachDetachDataDisks.
func (mr *MockInterfaceMockRecorder) BeginAttachDetachDataDisks(ctx, resourceGroupName, vmName, parameters, options any) *MockInterfaceBeginAttachDetachDataDisksCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginAttachDetachDataDisks", reflect.TypeOf((*MockInterface)(nil).BeginAttachDetachDataDisks), ctx, resourceGroupName, vmName, parameters, options)
	return &MockInterfaceBeginAttachDetachDataDisksCall{Call: call}
}

// MockInterfaceBeginAttachDetachDataDisksCall wrap *gomock.Call
type MockInterfaceBeginAttachDetachDataDisksCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceBeginAttachDetachDataDisksCall) Return(arg0 *runtime.Poller[armcompute.VirtualMachinesClientAttachDetachDataDisksResponse], arg1 error) *MockInterfaceBeginAttachDetachDataDisksCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceBeginAttachDetachDataDisksCall) Do(f func(context.Context, string, string, armcompute.AttachDetachDataDisksRequest, *armcompute.VirtualMachinesClientBeginAttachDetachDataDisksOptions) (*runtime.Poller[armcompute.VirtualMachinesClientAttachDetachDataDisksResponse], error)) *MockInterfaceBeginAttachDetachDataDisksCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceBeginAttachDetachDataDisksCall) DoAndReturn(f func(context.Context, string, string, armcompute.AttachDetachDataDisksRequest, *armcompute.VirtualMachinesClientBeginAttachDetachDataDisksOptions) (*runtime.Poller[armcompute.VirtualMachinesClientAttachDetachDataDisksResponse], error)) *MockInterfaceBeginAttachDetachDataDisksCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BeginUpdate mocks base method.
func (m *MockInterface) BeginUpdate(ctx context.Context, resourceGroupName, vmName string, parameters armcompute.VirtualMachineUpdate, options *armcompute.VirtualMachinesClientBeginUpdateOptions) (*runtime.Poller[armcompute.VirtualMachinesClientUpdateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginUpdate", ctx, resourceGroupName, vmName, parameters, options)
	ret0, _ := ret[0].(*runtime.Poller[armcompute.VirtualMachinesClientUpdateResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginUpdate indicates an expected call of BeginUpdate.
func (mr *MockInterfaceMockRecorder) BeginUpdate(ctx, resourceGroupName, vmName, parameters, options any) *MockInterfaceBeginUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginUpdate", reflect.TypeOf((*MockInterface)(nil).BeginUpdate), ctx, resourceGroupName, vmName, parameters, options)
	return &MockInterfaceBeginUpdateCall{Call: call}
}

// MockInterfaceBeginUpdateCall wrap *gomock.Call
type MockInterfaceBeginUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceBeginUpdateCall) Return(arg0 *runtime.Poller[armcompute.VirtualMachinesClientUpdateResponse], arg1 error) *MockInterfaceBeginUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceBeginUpdateCall) Do(f func(context.Context, string, string, armcompute.VirtualMachineUpdate, *armcompute.VirtualMachinesClientBeginUpdateOptions) (*runtime.Poller[armcompute.VirtualMachinesClientUpdateResponse], error)) *MockInterfaceBeginUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceBeginUpdateCall) DoAndReturn(f func(context.Context, string, string, armcompute.VirtualMachineUpdate, *armcompute.VirtualMachinesClientBeginUpdateOptions) (*runtime.Poller[armcompute.VirtualMachinesClientUpdateResponse], error)) *MockInterfaceBeginUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CreateOrUpdate mocks base method.
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, resourceName string, resourceParam armcompute.VirtualMachine) (*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceName, resourceParam)
	ret0, _ := ret[0].(*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceName, resourceParam any) *MockInterfaceCreateOrUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceName, resourceParam)
	return &MockInterfaceCreateOrUpdateCall{Call: call}
}

// MockInterfaceCreateOrUpdateCall wrap *gomock.Call
type MockInterfaceCreateOrUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceCreateOrUpdateCall) Return(arg0 *armcompute.VirtualMachine, arg1 error) *MockInterfaceCreateOrUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceCreateOrUpdateCall) Do(f func(context.Context, string, string, armcompute.VirtualMachine) (*armcompute.VirtualMachine, error)) *MockInterfaceCreateOrUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceCreateOrUpdateCall) DoAndReturn(f func(context.Context, string, string, armcompute.VirtualMachine) (*armcompute.VirtualMachine, error)) *MockInterfaceCreateOrUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockInterface) Delete(ctx context.Context, resourceGroupName, resourceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, resourceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(ctx, resourceGroupName, resourceName any) *MockInterfaceDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), ctx, resourceGroupName, resourceName)
	return &MockInterfaceDeleteCall{Call: call}
}

// MockInterfaceDeleteCall wrap *gomock.Call
type MockInterfaceDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceDeleteCall) Return(arg0 error) *MockInterfaceDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceDeleteCall) Do(f func(context.Context, string, string) error) *MockInterfaceDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceDeleteCall) DoAndReturn(f func(context.Context, string, string) error) *MockInterfaceDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, resourceName string, expand *string) (*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, expand)
	ret0, _ := ret[0].(*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, resourceName, expand any) *MockInterfaceGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, resourceName, expand)
	return &MockInterfaceGetCall{Call: call}
}

// MockInterfaceGetCall wrap *gomock.Call
type MockInterfaceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceGetCall) Return(result *armcompute.VirtualMachine, rerr error) *MockInterfaceGetCall {
	c.Call = c.Call.Return(result, rerr)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceGetCall) Do(f func(context.Context, string, string, *string) (*armcompute.VirtualMachine, error)) *MockInterfaceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceGetCall) DoAndReturn(f func(context.Context, string, string, *string) (*armcompute.VirtualMachine, error)) *MockInterfaceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InstanceView mocks base method.
func (m *MockInterface) InstanceView(ctx context.Context, resourceGroupName, vmName string) (*armcompute.VirtualMachineInstanceView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceView", ctx, resourceGroupName, vmName)
	ret0, _ := ret[0].(*armcompute.VirtualMachineInstanceView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceView indicates an expected call of InstanceView.
func (mr *MockInterfaceMockRecorder) InstanceView(ctx, resourceGroupName, vmName any) *MockInterfaceInstanceViewCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceView", reflect.TypeOf((*MockInterface)(nil).InstanceView), ctx, resourceGroupName, vmName)
	return &MockInterfaceInstanceViewCall{Call: call}
}

// MockInterfaceInstanceViewCall wrap *gomock.Call
type MockInterfaceInstanceViewCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceInstanceViewCall) Return(arg0 *armcompute.VirtualMachineInstanceView, arg1 error) *MockInterfaceInstanceViewCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceInstanceViewCall) Do(f func(context.Context, string, string) (*armcompute.VirtualMachineInstanceView, error)) *MockInterfaceInstanceViewCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceInstanceViewCall) DoAndReturn(f func(context.Context, string, string) (*armcompute.VirtualMachineInstanceView, error)) *MockInterfaceInstanceViewCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockInterface) List(ctx context.Context, resourceGroupName string) ([]*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName)
	ret0, _ := ret[0].([]*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName any) *MockInterfaceListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName)
	return &MockInterfaceListCall{Call: call}
}

// MockInterfaceListCall wrap *gomock.Call
type MockInterfaceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceListCall) Return(result []*armcompute.VirtualMachine, rerr error) *MockInterfaceListCall {
	c.Call = c.Call.Return(result, rerr)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceListCall) Do(f func(context.Context, string) ([]*armcompute.VirtualMachine, error)) *MockInterfaceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceListCall) DoAndReturn(f func(context.Context, string) ([]*armcompute.VirtualMachine, error)) *MockInterfaceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListVMInstanceView mocks base method.
func (m *MockInterface) ListVMInstanceView(ctx context.Context, resourceGroupName string) ([]*armcompute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVMInstanceView", ctx, resourceGroupName)
	ret0, _ := ret[0].([]*armcompute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVMInstanceView indicates an expected call of ListVMInstanceView.
func (mr *MockInterfaceMockRecorder) ListVMInstanceView(ctx, resourceGroupName any) *MockInterfaceListVMInstanceViewCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVMInstanceView", reflect.TypeOf((*MockInterface)(nil).ListVMInstanceView), ctx, resourceGroupName)
	return &MockInterfaceListVMInstanceViewCall{Call: call}
}

// MockInterfaceListVMInstanceViewCall wrap *gomock.Call
type MockInterfaceListVMInstanceViewCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceListVMInstanceViewCall) Return(result []*armcompute.VirtualMachine, rerr error) *MockInterfaceListVMInstanceViewCall {
	c.Call = c.Call.Return(result, rerr)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceListVMInstanceViewCall) Do(f func(context.Context, string) ([]*armcompute.VirtualMachine, error)) *MockInterfaceListVMInstanceViewCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceListVMInstanceViewCall) DoAndReturn(f func(context.Context, string) ([]*armcompute.VirtualMachine, error)) *MockInterfaceListVMInstanceViewCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
