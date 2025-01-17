// SPDX-License-Identifier: Apache-2.0
// Copyright 2016-2017 Authors of Cilium

package defaults

import (
	daemon "github.com/cilium/cilium/pkg/defaults"
)

const (
	// SockPath is the path to the UNIX domain socket exposing the API to clients locally
	SockPath = daemon.RuntimePath + "/health.sock"

	// SockPathEnv is the environment variable to overwrite SockPath
	SockPathEnv = "CILIUM_HEALTH_SOCK"

	// HTTPPathPort is used for probing base HTTP path connectivity
	HTTPPathPort = daemon.ClusterHealthPort
)
