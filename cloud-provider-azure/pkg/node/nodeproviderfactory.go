/*
Copyright 2021 The Kubernetes Authors.

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

package node

import (
	nodemanager "sigs.k8s.io/cloud-provider-azure/pkg/nodemanager"
)

// NewNodeProvider returns a node provider depending on the use case
func NewNodeProvider(useMetadata bool, cloudConfigFilePath string) nodemanager.NodeProvider {
	var nodeProvider nodemanager.NodeProvider

	if useMetadata {
		nodeProvider = NewIMDSNodeProvider()
	} else {
		nodeProvider = NewARMNodeProvider(cloudConfigFilePath)
	}

	return nodeProvider
}
