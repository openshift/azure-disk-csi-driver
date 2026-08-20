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

package e2e

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/reporters"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/config"
	consts "sigs.k8s.io/azuredisk-csi-driver/pkg/azureconstants"
	"sigs.k8s.io/azuredisk-csi-driver/pkg/azuredisk"
	"sigs.k8s.io/azuredisk-csi-driver/test/e2e/driver"
	"sigs.k8s.io/azuredisk-csi-driver/test/utils/azure"
	"sigs.k8s.io/azuredisk-csi-driver/test/utils/credentials"
)

var _ = ginkgo.BeforeSuite(func(ctx ginkgo.SpecContext) {
	log.Println(driver.AzureDriverNameVar, os.Getenv(driver.AzureDriverNameVar), fmt.Sprintf("%v", isUsingInTreeVolumePlugin))
	log.Println(testMigrationEnvVar, os.Getenv(testMigrationEnvVar), fmt.Sprintf("%v", isTestingMigration))
	log.Println(testWindowsEnvVar, os.Getenv(testWindowsEnvVar), fmt.Sprintf("%v", isWindowsCluster))
	log.Println(testWinServerVerEnvVar, os.Getenv(testWinServerVerEnvVar), fmt.Sprintf("%v", winServerVer))

	if os.Getenv(kubeconfigEnvVar) == "" {
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		os.Setenv(kubeconfigEnvVar, kubeconfig)
	}

	if isTestingMigration || !isUsingInTreeVolumePlugin {
		creds, err := credentials.CreateAzureCredentialFile()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		azureClient, err := azure.GetAzureClient(creds.Cloud, creds.SubscriptionID, creds.AADClientID, creds.TenantID, creds.AADClientSecret, creds.AADFederatedTokenFile)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		_, err = azureClient.EnsureResourceGroup(ctx, creds.ResourceGroup, creds.Location, nil)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		location = creds.Location
		supportZRSRegions := []string{
			"southafricanorth", "eastasia", "southeastasia", "australiaeast",
			"brazilsouth", "westeurope", "northeurope", "francecentral",
			"centralindia", "italynorth", "japaneast", "koreacentral",
			"norwayeast", "polandcentral", "qatarcentral", "swedencentral",
			"switzerlandnorth", "uaenorth", "uksouth",
			"eastus", "eastus2", "southcentralus", "westus2", "westus3",
		}
		supportsZRS = false
		for _, region := range supportZRSRegions {
			if location == region {
				supportsZRS = true
				break
			}
		}

		if os.Getenv("SKIP_DRIVER_INSTALL") != "true" {
			e2eBootstrap := testCmd{
				command:  "make",
				args:     []string{"e2e-bootstrap"},
				startLog: "Installing Azure Disk CSI Driver...",
				endLog:   "Azure Disk CSI Driver installed",
			}
			createMetricsSVC := testCmd{
				command:  "make",
				args:     []string{"create-metrics-svc"},
				startLog: "create metrics service ...",
				endLog:   "metrics service created",
			}
			execTestCmd([]testCmd{e2eBootstrap, createMetricsSVC})
		}

		driverOptions := azuredisk.DriverOptions{
			NodeID:                  os.Getenv("nodeid"),
			DriverName:              consts.DefaultDriverName,
			EnablePerfOptimization:  false,
			Kubeconfig:              os.Getenv(kubeconfigEnvVar),
			Endpoint:                fmt.Sprintf("unix:///tmp/csi-%s.sock", string(uuid.NewUUID())),
			GetDiskTimeoutInSeconds: 15,
		}
		os.Setenv("AZURE_CREDENTIAL_FILE", credentials.TempAzureCredentialFilePath)
		azurediskDriver = azuredisk.NewDriver(&driverOptions)

		go func() {
			err := azurediskDriver.Run(context.Background())
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		}()
	}
})

var _ = ginkgo.AfterSuite(func(_ ginkgo.SpecContext) {
	if isTestingMigration || isUsingInTreeVolumePlugin {
		cmLog := testCmd{
			command:  "bash",
			args:     []string{"test/utils/controller-manager-log.sh"},
			startLog: "===================controller-manager log=======",
			endLog:   "===================================================",
		}
		execTestCmd([]testCmd{cmLog})
	}

	if isTestingMigration || !isUsingInTreeVolumePlugin {
		checkPodsRestart := testCmd{
			command:  "bash",
			args:     []string{"test/utils/check_driver_pods_restart.sh", "log"},
			startLog: "Check driver pods if restarts ...",
			endLog:   "Check successfully",
		}
		execTestCmd([]testCmd{checkPodsRestart})

		testOS := "linux"
		cloud := "azurepubliccloud"
		if isWindowsCluster {
			testOS = "windows"
			if winServerVer == "windows-2022" {
				testOS = winServerVer
			}
		}
		if isAzureStackCloud {
			cloud = "azurestackcloud"
		}

		azurediskLog := testCmd{
			command:  "bash",
			args:     []string{"test/utils/azuredisk_log.sh"},
			startLog: "===================azuredisk log===================",
			endLog:   "===================================================",
		}
		execTestCmd([]testCmd{azurediskLog})

		if os.Getenv("SKIP_DRIVER_INSTALL") != "true" {
			createExampleDeployment := testCmd{
				command:  "bash",
				args:     []string{"hack/verify-examples.sh", testOS, cloud},
				startLog: "create example deployments",
				endLog:   "example deployments created",
			}
			execTestCmd([]testCmd{createExampleDeployment})

			deleteMetricsSVC := testCmd{
				command:  "make",
				args:     []string{"delete-metrics-svc"},
				startLog: "delete metrics service...",
				endLog:   "metrics service deleted",
			}
			e2eTeardown := testCmd{
				command:  "make",
				args:     []string{"e2e-teardown"},
				startLog: "Uninstalling Azure Disk CSI Driver...",
				endLog:   "Azure Disk CSI Driver uninstalled",
			}
			execTestCmd([]testCmd{deleteMetricsSVC, e2eTeardown})

			if !isTestingMigration {
				installDriver := testCmd{
					command:  "bash",
					args:     []string{"deploy/install-driver.sh", "master", "windows,snapshot,local"},
					startLog: "===================install Azure Disk CSI Driver deployment scripts test===================",
					endLog:   "===================================================",
				}
				execTestCmd([]testCmd{installDriver})

				createExampleDeployment := testCmd{
					command:  "bash",
					args:     []string{"hack/verify-examples.sh", testOS, cloud},
					startLog: "create example deployments#2",
					endLog:   "example deployments#2 created",
				}
				execTestCmd([]testCmd{createExampleDeployment})

				uninstallDriver := testCmd{
					command:  "bash",
					args:     []string{"deploy/uninstall-driver.sh", "master", "windows,snapshot,local"},
					startLog: "===================uninstall Azure Disk CSI Driver deployment scripts test===================",
					endLog:   "===================================================",
				}
				execTestCmd([]testCmd{uninstallDriver})
			}
		}
		err := credentials.DeleteAzureCredentialFile()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	}
})

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	reportDir := os.Getenv(reportDirEnvVar)
	if reportDir == "" {
		reportDir = defaultReportDir
	}
	r := []ginkgo.Reporter{reporters.NewJUnitReporter(path.Join(reportDir, "junit_01.xml"))}
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "AzureDisk CSI Driver End-to-End Tests", r)
}

func TestMain(m *testing.M) {
	config.CopyFlags(config.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	framework.AfterReadingAllFlags(&framework.TestContext)
	flag.Parse()
	os.Exit(m.Run())
}
