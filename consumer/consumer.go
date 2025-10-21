package consumer

import (
	apiv1 "github.com/bennettzhu1/bazel_go_demo_dep/pkg/api/v1"
)

// GetAPIVersionFromExternal calls pkg/api from external repo
func GetAPIVersionFromExternal() string {
	return "external-consumer says: " + apiv1.GetAPIVersion()
}
