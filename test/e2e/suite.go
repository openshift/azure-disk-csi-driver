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
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"sigs.k8s.io/azuredisk-csi-driver/pkg/azuredisk"
	"sigs.k8s.io/azuredisk-csi-driver/test/e2e/driver"
)

const (
	kubeconfigEnvVar       = "KUBECONFIG"
	reportDirEnvVar        = "ARTIFACTS"
	testMigrationEnvVar    = "TEST_MIGRATION"
	testWindowsEnvVar      = "TEST_WINDOWS"
	testWinServerVerEnvVar = "WINDOWS_SERVER_VERSION"
	cloudNameEnvVar        = "AZURE_CLOUD_NAME"
	defaultReportDir       = "/workspace/_artifacts"
	inTreeStorageClass     = "kubernetes.io/azure-disk"
)

var (
	azurediskDriver           azuredisk.CSIDriver
	isUsingInTreeVolumePlugin = os.Getenv(driver.AzureDriverNameVar) == inTreeStorageClass
	isTestingMigration        = os.Getenv(testMigrationEnvVar) != ""
	isWindowsCluster          = os.Getenv(testWindowsEnvVar) != ""
	winServerVer              = os.Getenv(testWinServerVerEnvVar)
	isAzureStackCloud         = strings.EqualFold(os.Getenv(cloudNameEnvVar), "AZURESTACKCLOUD")
	isWindowsHPCDeployment    = strings.EqualFold(os.Getenv("WINDOWS_USE_HOST_PROCESS_CONTAINERS"), "true")
	isCapzTest                = os.Getenv("NODE_MACHINE_TYPE") != ""
	location                  string
	supportsZRS               bool
)

type testCmd struct {
	command  string
	args     []string
	startLog string
	endLog   string
}

func execTestCmd(cmds []testCmd) {
	err := os.Chdir("../..")
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	defer func() {
		err := os.Chdir("test/e2e")
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	}()

	projectRoot, err := os.Getwd()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	gomega.Expect(strings.HasSuffix(projectRoot, "azuredisk-csi-driver") || strings.HasSuffix(projectRoot, "azure-disk-csi-driver")).To(gomega.Equal(true))

	for _, cmd := range cmds {
		log.Println(cmd.startLog)
		cmdSh := exec.Command(cmd.command, cmd.args...)
		cmdSh.Dir = projectRoot
		cmdSh.Stdout = os.Stdout
		cmdSh.Stderr = os.Stderr
		err = cmdSh.Run()
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		log.Println(cmd.endLog)
	}
}

func skipIfTestingInWindowsCluster() {
	if isWindowsCluster {
		ginkgo.Skip("test case not supported by Windows clusters")
	}
}

func skipIfUsingInTreeVolumePlugin() {
	if isUsingInTreeVolumePlugin {
		ginkgo.Skip("test case is only available for CSI drivers")
	}
}

func skipIfOnAzureStackCloud() {
	if isAzureStackCloud {
		ginkgo.Skip("test case not supported on Azure Stack Cloud")
	}
}

func skipIfNotZRSSupported() {
	if !supportsZRS {
		ginkgo.Skip("test case not supported on regions without ZRS")
	}
}

func convertToPowershellorCmdCommandIfNecessary(command string) string {
	if !isWindowsCluster {
		return command
	}

	switch command {
	case "echo 'hello world' > /mnt/test-1/data && grep 'hello world' /mnt/test-1/data":
		return "echo 'hello world' | Out-File -FilePath C:\\mnt\\test-1\\data.txt; Get-Content C:\\mnt\\test-1\\data.txt | findstr 'hello world'"
	case "touch /mnt/test-1/data":
		return "echo $null >> C:\\mnt\\test-1\\data"
	case "while true; do echo $(date -u) >> /mnt/test-1/data; sleep 3600; done":
		return "while (1) { Add-Content -Encoding Ascii C:\\mnt\\test-1\\data.txt $(Get-Date -Format u); sleep 3600 }"
	case "echo 'hello world' >> /mnt/test-1/data && while true; do sleep 3600; done":
		return "echo 'hello world' | Out-File -Append -FilePath C:\\mnt\\test-1\\data.txt; Get-Content C:\\mnt\\test-1\\data.txt | findstr 'hello world'; Start-Sleep 3600"
	case "echo 'hello world' > /mnt/test-1/data && echo 'hello world' > /mnt/test-2/data && echo 'hello world' > /mnt/test-3/data && grep 'hello world' /mnt/test-1/data && grep 'hello world' /mnt/test-2/data && grep 'hello world' /mnt/test-3/data":
		return "echo 'hello world' | Out-File -FilePath C:\\mnt\\test-1\\data.txt; Get-Content C:\\mnt\\test-1\\data.txt | findstr 'hello world'; echo 'hello world' | Out-File -FilePath C:\\mnt\\test-2\\data.txt; Get-Content C:\\mnt\\test-2\\data.txt | findstr 'hello world'; echo 'hello world' | Out-File -FilePath C:\\mnt\\test-3\\data.txt; Get-Content C:\\mnt\\test-3\\data.txt | findstr 'hello world'"
	case "while true; do sleep 5; done":
		return "while ($true) { Start-Sleep 5 }"
	case "echo 'hello world' > /mnt/test-1/data":
		return "echo 'hello world' | Out-File -FilePath C:\\mnt\\test-1\\data.txt"
	case "echo 'overwrite' > /mnt/test-1/data; sleep 3600":
		return "echo 'overwrite' | Out-File -FilePath C:\\mnt\\test-1\\data.txt; Start-Sleep 3600"
	case "grep 'hello world' /mnt/test-1/data":
		return "Get-Content C:\\mnt\\test-1\\data.txt | findstr 'hello world'"
	case "df -h | grep /mnt/test- | awk '{print $2}' | grep -E '19|20'":
		return "fsutil volume diskfree C:\\mnt\\ | Select-String 'Total bytes' | Select-String -Pattern '19|20'"
	}

	return command
}

func getFSType(IsWindowsCluster bool) string {
	if IsWindowsCluster {
		return "ntfs"
	}
	return "ext4"
}
