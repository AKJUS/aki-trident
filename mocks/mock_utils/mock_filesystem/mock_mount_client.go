// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils/filesystem (interfaces: Mount)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/mock_utils/mock_filesystem/mock_mount_client.go github.com/netapp/trident/utils/filesystem Mount
//

// Package mock_filesystem is a generated GoMock package.
package mock_filesystem

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMount is a mock of Mount interface.
type MockMount struct {
	ctrl     *gomock.Controller
	recorder *MockMountMockRecorder
	isgomock struct{}
}

// MockMountMockRecorder is the mock recorder for MockMount.
type MockMountMockRecorder struct {
	mock *MockMount
}

// NewMockMount creates a new mock instance.
func NewMockMount(ctrl *gomock.Controller) *MockMount {
	mock := &MockMount{ctrl: ctrl}
	mock.recorder = &MockMountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMount) EXPECT() *MockMountMockRecorder {
	return m.recorder
}

// MountFilesystemForResize mocks base method.
func (m *MockMount) MountFilesystemForResize(ctx context.Context, devicePath, stagedTargetPath, mountOptions string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountFilesystemForResize", ctx, devicePath, stagedTargetPath, mountOptions)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MountFilesystemForResize indicates an expected call of MountFilesystemForResize.
func (mr *MockMountMockRecorder) MountFilesystemForResize(ctx, devicePath, stagedTargetPath, mountOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountFilesystemForResize", reflect.TypeOf((*MockMount)(nil).MountFilesystemForResize), ctx, devicePath, stagedTargetPath, mountOptions)
}

// RemoveMountPoint mocks base method.
func (m *MockMount) RemoveMountPoint(ctx context.Context, mountPointPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMountPoint", ctx, mountPointPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMountPoint indicates an expected call of RemoveMountPoint.
func (mr *MockMountMockRecorder) RemoveMountPoint(ctx, mountPointPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMountPoint", reflect.TypeOf((*MockMount)(nil).RemoveMountPoint), ctx, mountPointPath)
}
