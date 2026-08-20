package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	et "github.com/openshift-eng/openshift-tests-extension/pkg/extension/extensiontests"
	"github.com/spf13/cobra"

	"github.com/openshift-eng/openshift-tests-extension/pkg/cmd"
	e "github.com/openshift-eng/openshift-tests-extension/pkg/extension"
	g "github.com/openshift-eng/openshift-tests-extension/pkg/ginkgo"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/config"

	consts "sigs.k8s.io/azuredisk-csi-driver/pkg/azureconstants"
	"sigs.k8s.io/azuredisk-csi-driver/pkg/azuredisk"
	e2e "sigs.k8s.io/azuredisk-csi-driver/test/e2e"
	"sigs.k8s.io/azuredisk-csi-driver/test/utils/credentials"
)

func main() {
	os.Setenv("SKIP_DRIVER_INSTALL", "true")

	if os.Getenv("KUBECONFIG") == "" {
		os.Setenv("KUBECONFIG", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	}
	config.CopyFlags(config.Flags, flag.CommandLine)
	framework.RegisterCommonFlags(flag.CommandLine)
	framework.RegisterClusterFlags(flag.CommandLine)
	flag.Parse()
	framework.AfterReadingAllFlags(&framework.TestContext)

	registry := e.NewRegistry()
	ext := e.NewExtension("openshift", "payload", "azure-disk-csi-driver-test")

	ext.AddSuite(e.Suite{
		Name: "openshift/csi/azure-disk/upstream",
	})

	specs, err := g.BuildExtensionTestSpecsFromOpenShiftGinkgoSuite()
	if err != nil {
		panic(fmt.Sprintf("failed to build extension test specs from ginkgo: %s", err.Error()))
	}

	// Only run on Azure platform
	specs.Include(et.PlatformEquals("azure"))

	// Exclude pre-provisioned tests (require in-process CSI driver instance)
	specs.Select(et.NameContains("Pre-Provisioned")).Exclude("true")

	// Exclude cross-region snapshot tests (require location from BeforeSuite)
	specs.Select(et.NameContains("snapshot cross region")).Exclude("true")

	// Label all tests as upstream azure-disk tests
	specs.AddLabel("azure-disk-upstream")

	specs.AddBeforeAll(populateAzureCredentialsFromCluster)

	ext.AddSpecs(specs)
	registry.Register(ext)

	root := &cobra.Command{
		Long: "Azure Disk CSI Driver E2E Tests for OpenShift",
	}

	root.AddCommand(cmd.DefaultExtensionCommands(registry)...)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

// populateAzureCredentialsFromCluster reads Azure credentials from the
// azure-credentials secret in kube-system and sets the environment
// variables that CreateAzureCredentialFile() expects. It also initializes
// the in-process CSI driver so that retain-reclaim tests can call
// CreateVolume/DeleteVolume directly via the Azure SDK.
func populateAzureCredentialsFromCluster() {
	if os.Getenv("AZURE_TENANT_ID") == "" {
		kubeconfigPath := os.Getenv("KUBECONFIG")
		restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			log.Printf("Warning: could not build kubeconfig for Azure credential extraction: %v", err)
			return
		}

		clientset, err := kubernetes.NewForConfig(restConfig)
		if err != nil {
			log.Printf("Warning: could not create Kubernetes client for Azure credential extraction: %v", err)
			return
		}

		secret, err := clientset.CoreV1().Secrets("kube-system").Get(
			context.Background(), "azure-credentials", metav1.GetOptions{})
		if err != nil {
			log.Printf("Warning: could not read azure-credentials secret: %v", err)
			return
		}

		keyToEnv := map[string]string{
			"azure_tenant_id":       "AZURE_TENANT_ID",
			"azure_subscription_id": "AZURE_SUBSCRIPTION_ID",
			"azure_client_id":       "AZURE_CLIENT_ID",
			"azure_client_secret":   "AZURE_CLIENT_SECRET",
			"azure_resourcegroup":   "AZURE_RESOURCE_GROUP",
			"azure_region":          "AZURE_LOCATION",
		}

		for secretKey, envVar := range keyToEnv {
			if val, ok := secret.Data[secretKey]; ok {
				setEnvIfEmpty(envVar, string(val))
			}
		}

		log.Printf("Azure credentials populated from cluster secret azure-credentials in kube-system")
	}

	initAzureDiskDriver()
}

// initAzureDiskDriver creates the in-process CSI driver instance so that
// tests using CreateVolume/DeleteVolume (retain-reclaim, pre-provisioned)
// can talk to the Azure API without a running gRPC server.
func initAzureDiskDriver() {
	creds, err := credentials.CreateAzureCredentialFile()
	if err != nil {
		log.Printf("Warning: could not create Azure credential file: %v", err)
		return
	}
	os.Setenv("AZURE_CREDENTIAL_FILE", credentials.TempAzureCredentialFilePath)
	log.Printf("Azure credential file created at %s (location: %s, rg: %s)", credentials.TempAzureCredentialFilePath, creds.Location, creds.ResourceGroup)

	driverOptions := azuredisk.DriverOptions{
		DriverName:              consts.DefaultDriverName,
		Kubeconfig:              os.Getenv("KUBECONFIG"),
		Endpoint:                fmt.Sprintf("unix:///tmp/csi-ote-%d.sock", os.Getpid()),
		GetDiskTimeoutInSeconds: 15,
	}
	driver := azuredisk.NewDriver(&driverOptions)
	e2e.SetAzureDiskDriver(driver)
	log.Printf("In-process Azure Disk CSI driver initialized for API calls")
}

func setEnvIfEmpty(key, value string) {
	if os.Getenv(key) == "" && value != "" {
		os.Setenv(key, value)
	}
}
