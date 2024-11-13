//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleAssignmentScheduleInstancesClient contains the methods for the RoleAssignmentScheduleInstances group.
// Don't use this type directly, use NewRoleAssignmentScheduleInstancesClient() instead.
type RoleAssignmentScheduleInstancesClient struct {
	internal *arm.Client
}

// NewRoleAssignmentScheduleInstancesClient creates a new instance of RoleAssignmentScheduleInstancesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoleAssignmentScheduleInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleAssignmentScheduleInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoleAssignmentScheduleInstancesClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets the specified role assignment schedule instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role assignments schedules.
//   - roleAssignmentScheduleInstanceName - The name (hash of schedule name + time) of the role assignment schedule to get.
//   - options - RoleAssignmentScheduleInstancesClientGetOptions contains the optional parameters for the RoleAssignmentScheduleInstancesClient.Get
//     method.
func (client *RoleAssignmentScheduleInstancesClient) Get(ctx context.Context, scope string, roleAssignmentScheduleInstanceName string, options *RoleAssignmentScheduleInstancesClientGetOptions) (RoleAssignmentScheduleInstancesClientGetResponse, error) {
	var err error
	const operationName = "RoleAssignmentScheduleInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentScheduleInstanceName, options)
	if err != nil {
		return RoleAssignmentScheduleInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentScheduleInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleAssignmentScheduleInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentScheduleInstancesClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleInstanceName string, options *RoleAssignmentScheduleInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleInstances/{roleAssignmentScheduleInstanceName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleInstanceName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleInstanceName}", url.PathEscape(roleAssignmentScheduleInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentScheduleInstancesClient) getHandleResponse(resp *http.Response) (RoleAssignmentScheduleInstancesClientGetResponse, error) {
	result := RoleAssignmentScheduleInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleInstance); err != nil {
		return RoleAssignmentScheduleInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role assignment schedule instances of a role assignment schedule.
//
// Generated from API version 2020-10-01
//   - scope - The scope of the role assignment schedule.
//   - options - RoleAssignmentScheduleInstancesClientListForScopeOptions contains the optional parameters for the RoleAssignmentScheduleInstancesClient.NewListForScopePager
//     method.
func (client *RoleAssignmentScheduleInstancesClient) NewListForScopePager(scope string, options *RoleAssignmentScheduleInstancesClientListForScopeOptions) *runtime.Pager[RoleAssignmentScheduleInstancesClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentScheduleInstancesClientListForScopeResponse]{
		More: func(page RoleAssignmentScheduleInstancesClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentScheduleInstancesClientListForScopeResponse) (RoleAssignmentScheduleInstancesClientListForScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoleAssignmentScheduleInstancesClient.NewListForScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForScopeCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return RoleAssignmentScheduleInstancesClientListForScopeResponse{}, err
			}
			return client.listForScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentScheduleInstancesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentScheduleInstancesClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentScheduleInstances"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentScheduleInstancesClient) listForScopeHandleResponse(resp *http.Response) (RoleAssignmentScheduleInstancesClientListForScopeResponse, error) {
	result := RoleAssignmentScheduleInstancesClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleInstanceListResult); err != nil {
		return RoleAssignmentScheduleInstancesClientListForScopeResponse{}, err
	}
	return result, nil
}
