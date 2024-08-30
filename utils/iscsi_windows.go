// Copyright 2022 NetApp, Inc. All Rights Reserved.

// This file should only contain functions for handling the filesystem for Windows flavor

package utils

import (
	"context"

	. "github.com/netapp/trident/logging"
	"github.com/netapp/trident/utils/errors"
	"github.com/netapp/trident/utils/models"
)

// ISCSIActiveOnHost unused stub function
func ISCSIActiveOnHost(ctx context.Context, host models.HostSystem) (bool, error) {
	Logc(ctx).Debug(">>>> iscsi_windows.ISCSIActiveOnHost")
	defer Logc(ctx).Debug("<<<< iscsi_windows.ISCSIActiveOnHost")
	return false, errors.UnsupportedError("ISCSIActiveOnHost is not supported for windows")
}
