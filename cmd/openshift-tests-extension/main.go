package main

import (
	"context"
	"encoding/json"
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

	_ "sigs.k8s.io/azuredisk-csi-driver/test/e2e"
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

	// Exclude retain reclaim policy test (uses in-process CSI driver for cleanup)
	specs.Select(et.NameContains(`reclaimPolicy "Retain"`)).Exclude("true")

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
// azure-cloud-provider secret in kube-system and sets the environment
// variables that CreateAzureCredentialFile() expects. This allows tests
// that make direct Azure API calls to work when running via OTE, where
// BeforeSuite is not executed.
func populateAzureCredentialsFromCluster() {
	if os.Getenv("AZURE_TENANT_ID") != "" {
		return
	}

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
		context.Background(), "azure-cloud-provider", metav1.GetOptions{})
	if err != nil {
		log.Printf("Warning: could not read azure-cloud-provider secret: %v", err)
		return
	}

	cloudConfigData := secret.Data["cloud-config"]
	if len(cloudConfigData) == 0 {
		log.Printf("Warning: azure-cloud-provider secret has no cloud-config data")
		return
	}

	var cloudConfig struct {
		TenantID              string `json:"tenantId"`
		SubscriptionID        string `json:"subscriptionId"`
		AADClientID           string `json:"aadClientId"`
		AADClientSecret       string `json:"aadClientSecret"`
		ResourceGroup         string `json:"resourceGroup"`
		Location              string `json:"location"`
		AADFederatedTokenFile string `json:"aadFederatedTokenFile"`
	}
	if err := json.Unmarshal(cloudConfigData, &cloudConfig); err != nil {
		log.Printf("Warning: could not parse azure-cloud-provider cloud-config: %v", err)
		return
	}

	setEnvIfEmpty("AZURE_TENANT_ID", cloudConfig.TenantID)
	setEnvIfEmpty("AZURE_SUBSCRIPTION_ID", cloudConfig.SubscriptionID)
	setEnvIfEmpty("AZURE_CLIENT_ID", cloudConfig.AADClientID)
	setEnvIfEmpty("AZURE_CLIENT_SECRET", cloudConfig.AADClientSecret)
	setEnvIfEmpty("AZURE_RESOURCE_GROUP", cloudConfig.ResourceGroup)
	setEnvIfEmpty("AZURE_LOCATION", cloudConfig.Location)
	setEnvIfEmpty("AZURE_FEDERATED_TOKEN_FILE", cloudConfig.AADFederatedTokenFile)

	log.Printf("Azure credentials populated from cluster secret azure-cloud-provider in kube-system")
}

func setEnvIfEmpty(key, value string) {
	if os.Getenv(key) == "" && value != "" {
		os.Setenv(key, value)
	}
}
