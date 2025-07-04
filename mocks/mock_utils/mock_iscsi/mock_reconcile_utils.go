// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils/iscsi (interfaces: IscsiReconcileUtils)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/mock_utils/mock_iscsi/mock_reconcile_utils.go github.com/netapp/trident/utils/iscsi IscsiReconcileUtils
//

// Package mock_iscsi is a generated GoMock package.
package mock_iscsi

import (
	context "context"
	reflect "reflect"

	models "github.com/netapp/trident/utils/models"
	gomock "go.uber.org/mock/gomock"
)

// MockIscsiReconcileUtils is a mock of IscsiReconcileUtils interface.
type MockIscsiReconcileUtils struct {
	ctrl     *gomock.Controller
	recorder *MockIscsiReconcileUtilsMockRecorder
	isgomock struct{}
}

// MockIscsiReconcileUtilsMockRecorder is the mock recorder for MockIscsiReconcileUtils.
type MockIscsiReconcileUtilsMockRecorder struct {
	mock *MockIscsiReconcileUtils
}

// NewMockIscsiReconcileUtils creates a new mock instance.
func NewMockIscsiReconcileUtils(ctrl *gomock.Controller) *MockIscsiReconcileUtils {
	mock := &MockIscsiReconcileUtils{ctrl: ctrl}
	mock.recorder = &MockIscsiReconcileUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIscsiReconcileUtils) EXPECT() *MockIscsiReconcileUtilsMockRecorder {
	return m.recorder
}

// DiscoverSCSIAddressMapForTarget mocks base method.
func (m *MockIscsiReconcileUtils) DiscoverSCSIAddressMapForTarget(ctx context.Context, targetIQN string) (map[string]models.ScsiDeviceAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverSCSIAddressMapForTarget", ctx, targetIQN)
	ret0, _ := ret[0].(map[string]models.ScsiDeviceAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverSCSIAddressMapForTarget indicates an expected call of DiscoverSCSIAddressMapForTarget.
func (mr *MockIscsiReconcileUtilsMockRecorder) DiscoverSCSIAddressMapForTarget(ctx, targetIQN any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverSCSIAddressMapForTarget", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).DiscoverSCSIAddressMapForTarget), ctx, targetIQN)
}

// GetDevicesForLUN mocks base method.
func (m *MockIscsiReconcileUtils) GetDevicesForLUN(paths []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesForLUN", paths)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesForLUN indicates an expected call of GetDevicesForLUN.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetDevicesForLUN(paths any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesForLUN", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetDevicesForLUN), paths)
}

// GetISCSIHostSessionMapForTarget mocks base method.
func (m *MockIscsiReconcileUtils) GetISCSIHostSessionMapForTarget(arg0 context.Context, arg1 string) map[int]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetISCSIHostSessionMapForTarget", arg0, arg1)
	ret0, _ := ret[0].(map[int]int)
	return ret0
}

// GetISCSIHostSessionMapForTarget indicates an expected call of GetISCSIHostSessionMapForTarget.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetISCSIHostSessionMapForTarget(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetISCSIHostSessionMapForTarget", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetISCSIHostSessionMapForTarget), arg0, arg1)
}

// GetSysfsBlockDirsForLUN mocks base method.
func (m *MockIscsiReconcileUtils) GetSysfsBlockDirsForLUN(arg0 int, arg1 map[int]int) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSysfsBlockDirsForLUN", arg0, arg1)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetSysfsBlockDirsForLUN indicates an expected call of GetSysfsBlockDirsForLUN.
func (mr *MockIscsiReconcileUtilsMockRecorder) GetSysfsBlockDirsForLUN(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSysfsBlockDirsForLUN", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).GetSysfsBlockDirsForLUN), arg0, arg1)
}

// ReconcileISCSIVolumeInfo mocks base method.
func (m *MockIscsiReconcileUtils) ReconcileISCSIVolumeInfo(ctx context.Context, trackingInfo *models.VolumeTrackingInfo) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileISCSIVolumeInfo", ctx, trackingInfo)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileISCSIVolumeInfo indicates an expected call of ReconcileISCSIVolumeInfo.
func (mr *MockIscsiReconcileUtilsMockRecorder) ReconcileISCSIVolumeInfo(ctx, trackingInfo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileISCSIVolumeInfo", reflect.TypeOf((*MockIscsiReconcileUtils)(nil).ReconcileISCSIVolumeInfo), ctx, trackingInfo)
}
