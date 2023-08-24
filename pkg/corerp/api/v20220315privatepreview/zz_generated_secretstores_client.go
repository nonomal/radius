//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

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

// SecretStoresClient contains the methods for the SecretStores group.
// Don't use this type directly, use NewSecretStoresClient() instead.
type SecretStoresClient struct {
	internal *arm.Client
	rootScope string
}

// NewSecretStoresClient creates a new instance of SecretStoresClient with the specified values.
//   - rootScope - The scope in which the resource is present. UCP Scope is /planes/{planeType}/{planeName}/resourceGroup/{resourcegroupID}
//     and Azure resource scope is
//     /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecretStoresClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecretStoresClient, error) {
	cl, err := arm.NewClient(moduleName+".SecretStoresClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecretStoresClient{
		rootScope: rootScope,
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a SecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - secretStoreName - SecretStore name
//   - resource - Resource create parameters.
//   - options - SecretStoresClientBeginCreateOrUpdateOptions contains the optional parameters for the SecretStoresClient.BeginCreateOrUpdate
//     method.
func (client *SecretStoresClient) BeginCreateOrUpdate(ctx context.Context, secretStoreName string, resource SecretStoreResource, options *SecretStoresClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecretStoresClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, secretStoreName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecretStoresClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SecretStoresClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a SecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
func (client *SecretStoresClient) createOrUpdate(ctx context.Context, secretStoreName string, resource SecretStoreResource, options *SecretStoresClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, secretStoreName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecretStoresClient) createOrUpdateCreateRequest(ctx context.Context, secretStoreName string, resource SecretStoreResource, options *SecretStoresClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/secretStores/{secretStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if secretStoreName == "" {
		return nil, errors.New("parameter secretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretStoreName}", url.PathEscape(secretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
	return nil, err
}
	return req, nil
}

// BeginDelete - Delete a SecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - secretStoreName - SecretStore name
//   - options - SecretStoresClientBeginDeleteOptions contains the optional parameters for the SecretStoresClient.BeginDelete
//     method.
func (client *SecretStoresClient) BeginDelete(ctx context.Context, secretStoreName string, options *SecretStoresClientBeginDeleteOptions) (*runtime.Poller[SecretStoresClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, secretStoreName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecretStoresClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SecretStoresClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a SecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
func (client *SecretStoresClient) deleteOperation(ctx context.Context, secretStoreName string, options *SecretStoresClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, secretStoreName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecretStoresClient) deleteCreateRequest(ctx context.Context, secretStoreName string, options *SecretStoresClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/secretStores/{secretStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if secretStoreName == "" {
		return nil, errors.New("parameter secretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretStoreName}", url.PathEscape(secretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a SecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - secretStoreName - SecretStore name
//   - options - SecretStoresClientGetOptions contains the optional parameters for the SecretStoresClient.Get method.
func (client *SecretStoresClient) Get(ctx context.Context, secretStoreName string, options *SecretStoresClientGetOptions) (SecretStoresClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, secretStoreName, options)
	if err != nil {
		return SecretStoresClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretStoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecretStoresClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecretStoresClient) getCreateRequest(ctx context.Context, secretStoreName string, options *SecretStoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/secretStores/{secretStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if secretStoreName == "" {
		return nil, errors.New("parameter secretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretStoreName}", url.PathEscape(secretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecretStoresClient) getHandleResponse(resp *http.Response) (SecretStoresClientGetResponse, error) {
	result := SecretStoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretStoreResource); err != nil {
		return SecretStoresClientGetResponse{}, err
	}
	return result, nil
}

// NewListByScopePager - List SecretStoreResource resources by Scope
//
// Generated from API version 2022-03-15-privatepreview
//   - options - SecretStoresClientListByScopeOptions contains the optional parameters for the SecretStoresClient.NewListByScopePager
//     method.
func (client *SecretStoresClient) NewListByScopePager(options *SecretStoresClientListByScopeOptions) (*runtime.Pager[SecretStoresClientListByScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[SecretStoresClientListByScopeResponse]{
		More: func(page SecretStoresClientListByScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecretStoresClientListByScopeResponse) (SecretStoresClientListByScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecretStoresClientListByScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SecretStoresClientListByScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecretStoresClientListByScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByScopeHandleResponse(resp)
		},
	})
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *SecretStoresClient) listByScopeCreateRequest(ctx context.Context, options *SecretStoresClientListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/secretStores"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByScopeHandleResponse handles the ListByScope response.
func (client *SecretStoresClient) listByScopeHandleResponse(resp *http.Response) (SecretStoresClientListByScopeResponse, error) {
	result := SecretStoresClientListByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretStoreResourceListResult); err != nil {
		return SecretStoresClientListByScopeResponse{}, err
	}
	return result, nil
}

// ListSecrets - List the secrets of a secret stores.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - secretStoreName - SecretStore name
//   - body - The content of the action request
//   - options - SecretStoresClientListSecretsOptions contains the optional parameters for the SecretStoresClient.ListSecrets
//     method.
func (client *SecretStoresClient) ListSecrets(ctx context.Context, secretStoreName string, body map[string]any, options *SecretStoresClientListSecretsOptions) (SecretStoresClientListSecretsResponse, error) {
	var err error
	req, err := client.listSecretsCreateRequest(ctx, secretStoreName, body, options)
	if err != nil {
		return SecretStoresClientListSecretsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretStoresClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecretStoresClientListSecretsResponse{}, err
	}
	resp, err := client.listSecretsHandleResponse(httpResp)
	return resp, err
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *SecretStoresClient) listSecretsCreateRequest(ctx context.Context, secretStoreName string, body map[string]any, options *SecretStoresClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/secretStores/{secretStoreName}/listSecrets"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if secretStoreName == "" {
		return nil, errors.New("parameter secretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretStoreName}", url.PathEscape(secretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *SecretStoresClient) listSecretsHandleResponse(resp *http.Response) (SecretStoresClientListSecretsResponse, error) {
	result := SecretStoresClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretStoreListSecretsResult); err != nil {
		return SecretStoresClientListSecretsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a SecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
//   - secretStoreName - SecretStore name
//   - properties - The resource properties to be updated.
//   - options - SecretStoresClientBeginUpdateOptions contains the optional parameters for the SecretStoresClient.BeginUpdate
//     method.
func (client *SecretStoresClient) BeginUpdate(ctx context.Context, secretStoreName string, properties SecretStoreResourceUpdate, options *SecretStoresClientBeginUpdateOptions) (*runtime.Poller[SecretStoresClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, secretStoreName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecretStoresClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SecretStoresClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a SecretStoreResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-15-privatepreview
func (client *SecretStoresClient) update(ctx context.Context, secretStoreName string, properties SecretStoreResourceUpdate, options *SecretStoresClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, secretStoreName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *SecretStoresClient) updateCreateRequest(ctx context.Context, secretStoreName string, properties SecretStoreResourceUpdate, options *SecretStoresClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/secretStores/{secretStoreName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if secretStoreName == "" {
		return nil, errors.New("parameter secretStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secretStoreName}", url.PathEscape(secretStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
	return nil, err
}
	return req, nil
}

