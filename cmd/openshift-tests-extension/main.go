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
		Name:       "openshift/csi/azure-disk",
		Parents:    []string{"openshift/conformance/parallel"},
		Qualifiers: []string{`!labels.exists(l, l=="azure-disk-upstream")`},
	})

	ext.AddSuite(e.Suite{
		Name:       "openshift/csi/azure-disk/upstream",
		Qualifiers: []string{`labels.exists(l, l=="azure-disk-upstream")`},
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

	// Label tests that require Azure API credentials (not available in generic conformance jobs)
	specs.Select(et.NameContains("azuredisk with tag")).AddLabel("azure-disk-upstream")
	specs.Select(et.NameContains("detach disk")).AddLabel("azure-disk-upstream")
	specs.Select(et.NameContains("separate resource group")).AddLabel("azure-disk-upstream")
	specs.Select(et.NameContains("volume snapshot")).AddLabel("azure-disk-upstream")
	specs.Select(et.NameContains("resize")).AddLabel("azure-disk-upstream")

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
