// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/storage_drivers/gcp/gcnvapi (interfaces: GCNV)
//
// Generated by this command:
//
//	mockgen -package mock_api -destination=../../../mocks/mock_storage_drivers/mock_gcp/mock_gcnvapi.go github.com/netapp/trident/storage_drivers/gcp/gcnvapi GCNV
//

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	reflect "reflect"
	time "time"

	storage "github.com/netapp/trident/storage"
	gcnvapi "github.com/netapp/trident/storage_drivers/gcp/gcnvapi"
	gomock "go.uber.org/mock/gomock"
)

// MockGCNV is a mock of GCNV interface.
type MockGCNV struct {
	ctrl     *gomock.Controller
	recorder *MockGCNVMockRecorder
	isgomock struct{}
}

// MockGCNVMockRecorder is the mock recorder for MockGCNV.
type MockGCNVMockRecorder struct {
	mock *MockGCNV
}

// NewMockGCNV creates a new mock instance.
func NewMockGCNV(ctrl *gomock.Controller) *MockGCNV {
	mock := &MockGCNV{ctrl: ctrl}
	mock.recorder = &MockGCNVMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGCNV) EXPECT() *MockGCNVMockRecorder {
	return m.recorder
}

// CapacityPools mocks base method.
func (m *MockGCNV) CapacityPools() *[]*gcnvapi.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapacityPools")
	ret0, _ := ret[0].(*[]*gcnvapi.CapacityPool)
	return ret0
}

// CapacityPools indicates an expected call of CapacityPools.
func (mr *MockGCNVMockRecorder) CapacityPools() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityPools", reflect.TypeOf((*MockGCNV)(nil).CapacityPools))
}

// CapacityPoolsForStoragePool mocks base method.
func (m *MockGCNV) CapacityPoolsForStoragePool(arg0 context.Context, arg1 storage.Pool, arg2 string) []*gcnvapi.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapacityPoolsForStoragePool", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*gcnvapi.CapacityPool)
	return ret0
}

// CapacityPoolsForStoragePool indicates an expected call of CapacityPoolsForStoragePool.
func (mr *MockGCNVMockRecorder) CapacityPoolsForStoragePool(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityPoolsForStoragePool", reflect.TypeOf((*MockGCNV)(nil).CapacityPoolsForStoragePool), arg0, arg1, arg2)
}

// CapacityPoolsForStoragePools mocks base method.
func (m *MockGCNV) CapacityPoolsForStoragePools(arg0 context.Context) []*gcnvapi.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapacityPoolsForStoragePools", arg0)
	ret0, _ := ret[0].([]*gcnvapi.CapacityPool)
	return ret0
}

// CapacityPoolsForStoragePools indicates an expected call of CapacityPoolsForStoragePools.
func (mr *MockGCNVMockRecorder) CapacityPoolsForStoragePools(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityPoolsForStoragePools", reflect.TypeOf((*MockGCNV)(nil).CapacityPoolsForStoragePools), arg0)
}

// CreateSnapshot mocks base method.
func (m *MockGCNV) CreateSnapshot(arg0 context.Context, arg1 *gcnvapi.Volume, arg2 string) (*gcnvapi.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*gcnvapi.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockGCNVMockRecorder) CreateSnapshot(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockGCNV)(nil).CreateSnapshot), arg0, arg1, arg2)
}

// CreateVolume mocks base method.
func (m *MockGCNV) CreateVolume(arg0 context.Context, arg1 *gcnvapi.VolumeCreateRequest) (*gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1)
	ret0, _ := ret[0].(*gcnvapi.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockGCNVMockRecorder) CreateVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockGCNV)(nil).CreateVolume), arg0, arg1)
}

// DeleteSnapshot mocks base method.
func (m *MockGCNV) DeleteSnapshot(arg0 context.Context, arg1 *gcnvapi.Volume, arg2 *gcnvapi.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockGCNVMockRecorder) DeleteSnapshot(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockGCNV)(nil).DeleteSnapshot), arg0, arg1, arg2)
}

// DeleteVolume mocks base method.
func (m *MockGCNV) DeleteVolume(arg0 context.Context, arg1 *gcnvapi.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockGCNVMockRecorder) DeleteVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockGCNV)(nil).DeleteVolume), arg0, arg1)
}

// DiscoverGCNVResources mocks base method.
func (m *MockGCNV) DiscoverGCNVResources(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverGCNVResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscoverGCNVResources indicates an expected call of DiscoverGCNVResources.
func (mr *MockGCNVMockRecorder) DiscoverGCNVResources(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverGCNVResources", reflect.TypeOf((*MockGCNV)(nil).DiscoverGCNVResources), arg0)
}

// EnsureVolumeInValidCapacityPool mocks base method.
func (m *MockGCNV) EnsureVolumeInValidCapacityPool(arg0 context.Context, arg1 *gcnvapi.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVolumeInValidCapacityPool", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVolumeInValidCapacityPool indicates an expected call of EnsureVolumeInValidCapacityPool.
func (mr *MockGCNVMockRecorder) EnsureVolumeInValidCapacityPool(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVolumeInValidCapacityPool", reflect.TypeOf((*MockGCNV)(nil).EnsureVolumeInValidCapacityPool), arg0, arg1)
}

// FilterCapacityPoolsOnTopology mocks base method.
func (m *MockGCNV) FilterCapacityPoolsOnTopology(arg0 context.Context, arg1 []*gcnvapi.CapacityPool, arg2, arg3 []map[string]string) []*gcnvapi.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterCapacityPoolsOnTopology", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*gcnvapi.CapacityPool)
	return ret0
}

// FilterCapacityPoolsOnTopology indicates an expected call of FilterCapacityPoolsOnTopology.
func (mr *MockGCNVMockRecorder) FilterCapacityPoolsOnTopology(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterCapacityPoolsOnTopology", reflect.TypeOf((*MockGCNV)(nil).FilterCapacityPoolsOnTopology), arg0, arg1, arg2, arg3)
}

// Init mocks base method.
func (m *MockGCNV) Init(arg0 context.Context, arg1 map[string]storage.Pool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockGCNVMockRecorder) Init(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockGCNV)(nil).Init), arg0, arg1)
}

// ModifyVolume mocks base method.
func (m *MockGCNV) ModifyVolume(arg0 context.Context, arg1 *gcnvapi.Volume, arg2 map[string]string, arg3 *string, arg4 *bool, arg5 *gcnvapi.ExportRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyVolume", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyVolume indicates an expected call of ModifyVolume.
func (mr *MockGCNVMockRecorder) ModifyVolume(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyVolume", reflect.TypeOf((*MockGCNV)(nil).ModifyVolume), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RefreshGCNVResources mocks base method.
func (m *MockGCNV) RefreshGCNVResources(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshGCNVResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshGCNVResources indicates an expected call of RefreshGCNVResources.
func (mr *MockGCNVMockRecorder) RefreshGCNVResources(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshGCNVResources", reflect.TypeOf((*MockGCNV)(nil).RefreshGCNVResources), arg0)
}

// ResizeVolume mocks base method.
func (m *MockGCNV) ResizeVolume(arg0 context.Context, arg1 *gcnvapi.Volume, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResizeVolume indicates an expected call of ResizeVolume.
func (mr *MockGCNVMockRecorder) ResizeVolume(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeVolume", reflect.TypeOf((*MockGCNV)(nil).ResizeVolume), arg0, arg1, arg2)
}

// RestoreSnapshot mocks base method.
func (m *MockGCNV) RestoreSnapshot(arg0 context.Context, arg1 *gcnvapi.Volume, arg2 *gcnvapi.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreSnapshot indicates an expected call of RestoreSnapshot.
func (mr *MockGCNVMockRecorder) RestoreSnapshot(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSnapshot", reflect.TypeOf((*MockGCNV)(nil).RestoreSnapshot), arg0, arg1, arg2)
}

// SnapshotForVolume mocks base method.
func (m *MockGCNV) SnapshotForVolume(arg0 context.Context, arg1 *gcnvapi.Volume, arg2 string) (*gcnvapi.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotForVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(*gcnvapi.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotForVolume indicates an expected call of SnapshotForVolume.
func (mr *MockGCNVMockRecorder) SnapshotForVolume(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotForVolume", reflect.TypeOf((*MockGCNV)(nil).SnapshotForVolume), arg0, arg1, arg2)
}

// SnapshotsForVolume mocks base method.
func (m *MockGCNV) SnapshotsForVolume(arg0 context.Context, arg1 *gcnvapi.Volume) (*[]*gcnvapi.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotsForVolume", arg0, arg1)
	ret0, _ := ret[0].(*[]*gcnvapi.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotsForVolume indicates an expected call of SnapshotsForVolume.
func (mr *MockGCNVMockRecorder) SnapshotsForVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotsForVolume", reflect.TypeOf((*MockGCNV)(nil).SnapshotsForVolume), arg0, arg1)
}

// Volume mocks base method.
func (m *MockGCNV) Volume(arg0 context.Context, arg1 *storage.VolumeConfig) (*gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Volume", arg0, arg1)
	ret0, _ := ret[0].(*gcnvapi.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Volume indicates an expected call of Volume.
func (mr *MockGCNVMockRecorder) Volume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volume", reflect.TypeOf((*MockGCNV)(nil).Volume), arg0, arg1)
}

// VolumeByID mocks base method.
func (m *MockGCNV) VolumeByID(arg0 context.Context, arg1 string) (*gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeByID", arg0, arg1)
	ret0, _ := ret[0].(*gcnvapi.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeByID indicates an expected call of VolumeByID.
func (mr *MockGCNVMockRecorder) VolumeByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeByID", reflect.TypeOf((*MockGCNV)(nil).VolumeByID), arg0, arg1)
}

// VolumeByName mocks base method.
func (m *MockGCNV) VolumeByName(arg0 context.Context, arg1 string) (*gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeByName", arg0, arg1)
	ret0, _ := ret[0].(*gcnvapi.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeByName indicates an expected call of VolumeByName.
func (mr *MockGCNVMockRecorder) VolumeByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeByName", reflect.TypeOf((*MockGCNV)(nil).VolumeByName), arg0, arg1)
}

// VolumeByNameOrID mocks base method.
func (m *MockGCNV) VolumeByNameOrID(arg0 context.Context, arg1 string) (*gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeByNameOrID", arg0, arg1)
	ret0, _ := ret[0].(*gcnvapi.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeByNameOrID indicates an expected call of VolumeByNameOrID.
func (mr *MockGCNVMockRecorder) VolumeByNameOrID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeByNameOrID", reflect.TypeOf((*MockGCNV)(nil).VolumeByNameOrID), arg0, arg1)
}

// VolumeExists mocks base method.
func (m *MockGCNV) VolumeExists(arg0 context.Context, arg1 *storage.VolumeConfig) (bool, *gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*gcnvapi.Volume)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumeExists indicates an expected call of VolumeExists.
func (mr *MockGCNVMockRecorder) VolumeExists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExists", reflect.TypeOf((*MockGCNV)(nil).VolumeExists), arg0, arg1)
}

// VolumeExistsByID mocks base method.
func (m *MockGCNV) VolumeExistsByID(arg0 context.Context, arg1 string) (bool, *gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExistsByID", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*gcnvapi.Volume)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumeExistsByID indicates an expected call of VolumeExistsByID.
func (mr *MockGCNVMockRecorder) VolumeExistsByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExistsByID", reflect.TypeOf((*MockGCNV)(nil).VolumeExistsByID), arg0, arg1)
}

// VolumeExistsByName mocks base method.
func (m *MockGCNV) VolumeExistsByName(arg0 context.Context, arg1 string) (bool, *gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExistsByName", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*gcnvapi.Volume)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumeExistsByName indicates an expected call of VolumeExistsByName.
func (mr *MockGCNVMockRecorder) VolumeExistsByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExistsByName", reflect.TypeOf((*MockGCNV)(nil).VolumeExistsByName), arg0, arg1)
}

// Volumes mocks base method.
func (m *MockGCNV) Volumes(arg0 context.Context) (*[]*gcnvapi.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Volumes", arg0)
	ret0, _ := ret[0].(*[]*gcnvapi.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Volumes indicates an expected call of Volumes.
func (mr *MockGCNVMockRecorder) Volumes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volumes", reflect.TypeOf((*MockGCNV)(nil).Volumes), arg0)
}

// WaitForSnapshotState mocks base method.
func (m *MockGCNV) WaitForSnapshotState(arg0 context.Context, arg1 *gcnvapi.Snapshot, arg2 *gcnvapi.Volume, arg3 string, arg4 []string, arg5 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForSnapshotState", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForSnapshotState indicates an expected call of WaitForSnapshotState.
func (mr *MockGCNVMockRecorder) WaitForSnapshotState(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForSnapshotState", reflect.TypeOf((*MockGCNV)(nil).WaitForSnapshotState), arg0, arg1, arg2, arg3, arg4, arg5)
}

// WaitForVolumeState mocks base method.
func (m *MockGCNV) WaitForVolumeState(arg0 context.Context, arg1 *gcnvapi.Volume, arg2 string, arg3 []string, arg4 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForVolumeState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForVolumeState indicates an expected call of WaitForVolumeState.
func (mr *MockGCNVMockRecorder) WaitForVolumeState(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForVolumeState", reflect.TypeOf((*MockGCNV)(nil).WaitForVolumeState), arg0, arg1, arg2, arg3, arg4)
}
