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

// The external node manager is responsible for running node reconciler loops that
// are cloud provider dependent. It uses the API to listen to new events on resources.

package main

import (
	"math/rand"
	"os"
	"time"

	"k8s.io/component-base/logs"
	"sigs.k8s.io/cloud-provider-azure/cmd/cloud-node-manager/app"

	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load all the prometheus client-go plugins

	_ "sigs.k8s.io/cloud-provider-azure/pkg/provider"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewCloudNodeManagerCommand()

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
