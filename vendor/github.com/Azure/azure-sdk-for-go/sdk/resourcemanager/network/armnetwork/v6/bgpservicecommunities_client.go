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

// BgpServiceCommunitiesClient contains the methods for the BgpServiceCommunities group.
// Don't use this type directly, use NewBgpServiceCommunitiesClient() instead.
type BgpServiceCommunitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBgpServiceCommunitiesClient creates a new instance of BgpServiceCommunitiesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBgpServiceCommunitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BgpServiceCommunitiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BgpServiceCommunitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets all the available bgp service communities.
//
// Generated from API version 2024-05-01
//   - options - BgpServiceCommunitiesClientListOptions contains the optional parameters for the BgpServiceCommunitiesClient.NewListPager
//     method.
func (client *BgpServiceCommunitiesClient) NewListPager(options *BgpServiceCommunitiesClientListOptions) *runtime.Pager[BgpServiceCommunitiesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BgpServiceCommunitiesClientListResponse]{
		More: func(page BgpServiceCommunitiesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BgpServiceCommunitiesClientListResponse) (BgpServiceCommunitiesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BgpServiceCommunitiesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BgpServiceCommunitiesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BgpServiceCommunitiesClient) listCreateRequest(ctx context.Context, options *BgpServiceCommunitiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listHandleResponse handles the List response.
func (client *BgpServiceCommunitiesClient) listHandleResponse(resp *http.Response) (BgpServiceCommunitiesClientListResponse, error) {
	result := BgpServiceCommunitiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BgpServiceCommunityListResult); err != nil {
		return BgpServiceCommunitiesClientListResponse{}, err
	}
	return result, nil
}
