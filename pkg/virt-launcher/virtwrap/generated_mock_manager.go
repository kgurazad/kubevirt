// Automatically generated by MockGen. DO NOT EDIT!
// Source: manager.go

package virtwrap

import (
	gomock "github.com/golang/mock/gomock"

	v1 "kubevirt.io/kubevirt/pkg/api/v1"
	api "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
)

// Mock of DomainManager interface
type MockDomainManager struct {
	ctrl     *gomock.Controller
	recorder *_MockDomainManagerRecorder
}

// Recorder for MockDomainManager (not exported)
type _MockDomainManagerRecorder struct {
	mock *MockDomainManager
}

func NewMockDomainManager(ctrl *gomock.Controller) *MockDomainManager {
	mock := &MockDomainManager{ctrl: ctrl}
	mock.recorder = &_MockDomainManagerRecorder{mock}
	return mock
}

func (_m *MockDomainManager) EXPECT() *_MockDomainManagerRecorder {
	return _m.recorder
}

func (_m *MockDomainManager) SyncVMI(_param0 *v1.VirtualMachineInstance, _param1 bool) (*api.DomainSpec, error) {
	ret := _m.ctrl.Call(_m, "SyncVMI", _param0, _param1)
	ret0, _ := ret[0].(*api.DomainSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDomainManagerRecorder) SyncVMI(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SyncVMI", arg0, arg1)
}

func (_m *MockDomainManager) KillVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "KillVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) KillVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KillVMI", arg0)
}

func (_m *MockDomainManager) DeleteVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "DeleteVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) DeleteVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteVMI", arg0)
}

func (_m *MockDomainManager) SignalShutdownVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "SignalShutdownVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) SignalShutdownVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SignalShutdownVMI", arg0)
}

func (_m *MockDomainManager) ListAllDomains() ([]*api.Domain, error) {
	ret := _m.ctrl.Call(_m, "ListAllDomains")
	ret0, _ := ret[0].([]*api.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDomainManagerRecorder) ListAllDomains() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAllDomains")
}

func (_m *MockDomainManager) MigrateVMI(_param0 *v1.VirtualMachineInstance, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "MigrateVMI", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) MigrateVMI(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateVMI", arg0, arg1)
}

func (_m *MockDomainManager) PrepareMigrationTarget(_param0 *v1.VirtualMachineInstance, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "PrepareMigrationTarget", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) PrepareMigrationTarget(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PrepareMigrationTarget", arg0, arg1)
}
