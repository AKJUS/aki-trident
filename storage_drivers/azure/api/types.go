// Copyright 2021 NetApp, Inc. All Rights Reserved.

package api

import (
	"context"
	"time"

	"github.com/netapp/trident/storage"
)

//go:generate mockgen -destination=../../../mocks/mock_storage_drivers/mock_azure/mock_api.go github.com/netapp/trident/storage_drivers/azure/api Azure

type Azure interface {
	Init(context.Context, map[string]storage.Pool) error

	RefreshAzureResources(context.Context) error
	DiscoverAzureResources(context.Context) error
	EnableAzureFeatures(context.Context, ...string) error
	Features() map[string]bool
	HasFeature(string) bool
	CapacityPools() *[]*CapacityPool
	CapacityPoolsForStoragePools(context.Context) []*CapacityPool
	CapacityPoolsForStoragePool(context.Context, storage.Pool, string, string) []*CapacityPool
	EnsureVolumeInValidCapacityPool(context.Context, *FileSystem) error
	SubnetsForStoragePool(context.Context, storage.Pool) []*Subnet
	RandomSubnetForStoragePool(context.Context, storage.Pool) *Subnet
	FilteredCapacityPoolMap(ctx context.Context, rgFilter, naFilter, cpFilter []string) map[string]*CapacityPool
	FilteredSubnetMap(ctx context.Context, rgFilter []string, vnFilter, snFilter string) map[string]*Subnet

	Volumes(context.Context) (*[]*FileSystem, error)
	Volume(context.Context, *storage.VolumeConfig) (*FileSystem, error)
	VolumeExists(context.Context, *storage.VolumeConfig) (bool, *FileSystem, error)
	VolumeByCreationToken(context.Context, string) (*FileSystem, error)
	VolumeExistsByCreationToken(context.Context, string) (bool, *FileSystem, error)
	VolumeByID(context.Context, string) (*FileSystem, error)
	VolumeExistsByID(context.Context, string) (bool, *FileSystem, error)
	WaitForVolumeState(context.Context, *FileSystem, string, []string, time.Duration, Operation) (string, error)
	CreateVolume(context.Context, *FilesystemCreateRequest) (*FileSystem, error)
	ModifyVolume(context.Context, *FileSystem, map[string]string, *string, *bool, *ExportRule) error
	ResizeVolume(context.Context, *FileSystem, int64) error
	DeleteVolume(context.Context, *FileSystem) error

	SnapshotsForVolume(context.Context, *FileSystem) (*[]*Snapshot, error)
	SnapshotForVolume(context.Context, *FileSystem, string) (*Snapshot, error)
	WaitForSnapshotState(context.Context, *Snapshot, *FileSystem, string, []string, time.Duration) error
	CreateSnapshot(context.Context, *FileSystem, string) (*Snapshot, error)
	RestoreSnapshot(context.Context, *FileSystem, *Snapshot) error
	DeleteSnapshot(context.Context, *FileSystem, *Snapshot) error

	AvailabilityZones(ctx context.Context) ([]string, error)
}
