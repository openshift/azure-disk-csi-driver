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

package credentials

import (
	"fmt"
	"html/template"
	"os"

	"k8s.io/apimachinery/pkg/util/uuid"
)

const (
	AzurePublicCloud            = "AzurePublicCloud"
	ResourceGroupPrefix         = "azuredisk-csi-driver-test-"
	TempAzureCredentialFilePath = "/tmp/azure.json"

	azureCredentialFileTemplate = `{
    "cloud": "{{.Cloud}}",
    "tenantId": "{{.TenantID}}",
    "subscriptionId": "{{.SubscriptionID}}",
    "aadClientId": "{{.AADClientID}}",
    "aadClientSecret": "{{.AADClientSecret}}",
    "resourceGroup": "{{.ResourceGroup}}",
    "location": "{{.Location}}",
    "vmType": "{{.VMType}}",
    "aadFederatedTokenFile": "{{.AADFederatedTokenFile}}"
}`
	defaultAzurePublicCloudLocation = "eastus2"
	defaultAzurePublicCloudVMType   = "vmss"

	// Env vars
	cloudNameEnvVar       = "AZURE_CLOUD_NAME"
	tenantIDEnvVar        = "AZURE_TENANT_ID"
	subscriptionIDEnvVar  = "AZURE_SUBSCRIPTION_ID"
	aadClientIDEnvVar     = "AZURE_CLIENT_ID"
	aadClientSecretEnvVar = "AZURE_CLIENT_SECRET"
	resourceGroupEnvVar   = "AZURE_RESOURCE_GROUP"
	locationEnvVar        = "AZURE_LOCATION"
	vmTypeEnvVar          = "AZURE_VM_TYPE"
	federatedTokenFileVar = "AZURE_FEDERATED_TOKEN_FILE"
)

// Config is used in Prow to store Azure credentials
// https://github.com/kubernetes/test-infra/blob/master/kubetest/azure.go#L116-L118
type Config struct {
	Creds FromProw
}

// FromProw is used in Prow to store Azure credentials
// https://github.com/kubernetes/test-infra/blob/master/kubetest/azure.go#L107-L114
type FromProw struct {
	ClientID           string
	ClientSecret       string
	TenantID           string
	SubscriptionID     string
	StorageAccountName string
	StorageAccountKey  string
}

// Credentials is used in Azure Disk CSI driver to store Azure credentials
type Credentials struct {
	Cloud                 string
	TenantID              string
	SubscriptionID        string
	AADClientID           string
	AADClientSecret       string
	ResourceGroup         string
	Location              string
	VMType                string
	AADFederatedTokenFile string
}

// CreateAzureCredentialFile creates a temporary Azure credential file for
// Azure Disk CSI driver tests and returns the credentials
func CreateAzureCredentialFile() (*Credentials, error) {
	// Search credentials through env vars first
	var cloud, tenantID, subscriptionID, aadClientID, aadClientSecret, resourceGroup, location, vmType, aadFederatedTokenFile string
	cloud = os.Getenv(cloudNameEnvVar)
	if cloud == "" {
		cloud = AzurePublicCloud
	}
	tenantID = os.Getenv(tenantIDEnvVar)
	subscriptionID = os.Getenv(subscriptionIDEnvVar)
	aadClientID = os.Getenv(aadClientIDEnvVar)
	aadClientSecret = os.Getenv(aadClientSecretEnvVar)
	resourceGroup = os.Getenv(resourceGroupEnvVar)
	location = os.Getenv(locationEnvVar)
	vmType = os.Getenv(vmTypeEnvVar)
	aadFederatedTokenFile = os.Getenv(federatedTokenFileVar)

	if resourceGroup == "" {
		resourceGroup = ResourceGroupPrefix + string(uuid.NewUUID())
	}

	if location == "" {
		location = defaultAzurePublicCloudLocation
	}

	if vmType == "" {
		vmType = defaultAzurePublicCloudVMType
	}

	if tenantID != "" && subscriptionID != "" && aadClientID != "" && (aadClientSecret != "" || aadFederatedTokenFile != "") {
		return parseAndExecuteTemplate(cloud, tenantID, subscriptionID, aadClientID, aadClientSecret, aadFederatedTokenFile, resourceGroup, location, vmType)
	}

	return nil, fmt.Errorf("If you are running tests locally, you will need to set the following env vars: $%s, $%s, $%s, $%s, $%s, $%s",
		tenantIDEnvVar, subscriptionIDEnvVar, aadClientIDEnvVar, aadClientSecretEnvVar, resourceGroupEnvVar, locationEnvVar)
}

// DeleteAzureCredentialFile deletes the temporary Azure credential file
func DeleteAzureCredentialFile() error {
	if err := os.Remove(TempAzureCredentialFilePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error removing %s %v", TempAzureCredentialFilePath, err)
	}

	return nil
}

// parseAndExecuteTemplate replaces credential placeholders in azureCredentialFileTemplate with actual credentials
func parseAndExecuteTemplate(cloud, tenantID, subscriptionID, aadClientID, aadClientSecret, aadFederatedTokenFile, resourceGroup, location, vmType string) (*Credentials, error) {
	t := template.New("AzureCredentialFileTemplate")
	t, err := t.Parse(azureCredentialFileTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing azureCredentialFileTemplate %v", err)
	}

	f, err := os.Create(TempAzureCredentialFilePath)
	if err != nil {
		return nil, fmt.Errorf("error creating %s %v", TempAzureCredentialFilePath, err)
	}
	defer f.Close()

	c := Credentials{
		cloud,
		tenantID,
		subscriptionID,
		aadClientID,
		aadClientSecret,
		resourceGroup,
		location,
		vmType,
		aadFederatedTokenFile,
	}
	err = t.Execute(f, c)
	if err != nil {
		return nil, fmt.Errorf("error executing parsed azure credential file template %v", err)
	}

	return &c, nil
}
