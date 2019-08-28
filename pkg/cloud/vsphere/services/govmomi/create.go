/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govmomi

import (
	"sigs.k8s.io/cluster-api-provider-vsphere/pkg/cloud/vsphere/context"
	"sigs.k8s.io/cluster-api-provider-vsphere/pkg/cloud/vsphere/services/govmomi/esxi"
	"sigs.k8s.io/cluster-api-provider-vsphere/pkg/cloud/vsphere/services/govmomi/vcenter"
)

const (
	// nodeRole is the label assigned to every node in the cluster.
	nodeRole = "node-role.kubernetes.io/node="

	// the Kubernetes cloud provider to use
	cloudProvider = "vsphere"

	// the cloud config path read by the cloud provider
	cloudConfigPath = "/etc/kubernetes/vsphere.conf"
)

func createVM(ctx *context.MachineContext, bootstrapData []byte) error {
	if ctx.Session.IsVC() {
		return vcenter.Clone(ctx, bootstrapData)
	}
	return esxi.Clone(ctx, bootstrapData)
}
