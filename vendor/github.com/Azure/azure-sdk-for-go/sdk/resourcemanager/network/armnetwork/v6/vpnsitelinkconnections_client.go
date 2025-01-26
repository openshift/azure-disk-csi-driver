//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VPNSiteLinkConnectionsClient contains the methods for the VPNSiteLinkConnections group.
// Don't use this type directly, use NewVPNSiteLinkConnectionsClient() instead.
type VPNSiteLinkConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVPNSiteLinkConnectionsClient creates a new instance of VPNSiteLinkConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVPNSiteLinkConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNSiteLinkConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VPNSiteLinkConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieves the details of a vpn site link connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the vpn connection.
//   - linkConnectionName - The name of the vpn connection.
//   - options - VPNSiteLinkConnectionsClientGetOptions contains the optional parameters for the VPNSiteLinkConnectionsClient.Get
//     method.
func (client *VPNSiteLinkConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNSiteLinkConnectionsClientGetOptions) (VPNSiteLinkConnectionsClientGetResponse, error) {
	var err error
	const operationName = "VPNSiteLinkConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName, options)
	if err != nil {
		return VPNSiteLinkConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VPNSiteLinkConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VPNSiteLinkConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VPNSiteLinkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNSiteLinkConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNSiteLinkConnectionsClient) getHandleResponse(resp *http.Response) (VPNSiteLinkConnectionsClientGetResponse, error) {
	result := VPNSiteLinkConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNSiteLinkConnection); err != nil {
		return VPNSiteLinkConnectionsClientGetResponse{}, err
	}
	return result, nil
}
