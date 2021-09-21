// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclientv3

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DaprIoStateStoreComponentClient contains the methods for the DaprIoStateStoreComponent group.
// Don't use this type directly, use NewDaprIoStateStoreComponentClient() instead.
type DaprIoStateStoreComponentClient struct {
	con *armcore.Connection
	subscriptionID string
}

// NewDaprIoStateStoreComponentClient creates a new instance of DaprIoStateStoreComponentClient with the specified values.
func NewDaprIoStateStoreComponentClient(con *armcore.Connection, subscriptionID string) *DaprIoStateStoreComponentClient {
	return &DaprIoStateStoreComponentClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a dapr.io.StateStoreComponent resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoStateStoreComponentClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, daprStateStoreComponentName string, parameters DaprStateStoreComponentResource, options *DaprIoStateStoreComponentCreateOrUpdateOptions) (DaprStateStoreComponentResourceResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationName, daprStateStoreComponentName, parameters, options)
	if err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return DaprStateStoreComponentResourceResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprIoStateStoreComponentClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, daprStateStoreComponentName string, parameters DaprStateStoreComponentResource, options *DaprIoStateStoreComponentCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.StateStoreComponent/{daprStateStoreComponentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if daprStateStoreComponentName == "" {
		return nil, errors.New("parameter daprStateStoreComponentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprStateStoreComponentName}", url.PathEscape(daprStateStoreComponentName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprIoStateStoreComponentClient) createOrUpdateHandleResponse(resp *azcore.Response) (DaprStateStoreComponentResourceResponse, error) {
	var val *DaprStateStoreComponentResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
return DaprStateStoreComponentResourceResponse{RawResponse: resp.Response, DaprStateStoreComponentResource: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DaprIoStateStoreComponentClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes a dapr.io.StateStoreComponent resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoStateStoreComponentClient) Delete(ctx context.Context, resourceGroupName string, applicationName string, daprStateStoreComponentName string, options *DaprIoStateStoreComponentDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationName, daprStateStoreComponentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprIoStateStoreComponentClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, daprStateStoreComponentName string, options *DaprIoStateStoreComponentDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.StateStoreComponent/{daprStateStoreComponentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if daprStateStoreComponentName == "" {
		return nil, errors.New("parameter daprStateStoreComponentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprStateStoreComponentName}", url.PathEscape(daprStateStoreComponentName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DaprIoStateStoreComponentClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets a dapr.io.StateStoreComponent resource by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoStateStoreComponentClient) Get(ctx context.Context, resourceGroupName string, applicationName string, daprStateStoreComponentName string, options *DaprIoStateStoreComponentGetOptions) (DaprStateStoreComponentResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationName, daprStateStoreComponentName, options)
	if err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DaprStateStoreComponentResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprIoStateStoreComponentClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, daprStateStoreComponentName string, options *DaprIoStateStoreComponentGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.StateStoreComponent/{daprStateStoreComponentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if daprStateStoreComponentName == "" {
		return nil, errors.New("parameter daprStateStoreComponentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprStateStoreComponentName}", url.PathEscape(daprStateStoreComponentName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprIoStateStoreComponentClient) getHandleResponse(resp *azcore.Response) (DaprStateStoreComponentResourceResponse, error) {
	var val *DaprStateStoreComponentResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
return DaprStateStoreComponentResourceResponse{RawResponse: resp.Response, DaprStateStoreComponentResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *DaprIoStateStoreComponentClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - List the dapr.io.StateStoreComponent resources deployed in the application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DaprIoStateStoreComponentClient) List(ctx context.Context, resourceGroupName string, applicationName string, options *DaprIoStateStoreComponentListOptions) (DaprStateStoreComponentListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return DaprStateStoreComponentListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DaprStateStoreComponentListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DaprStateStoreComponentListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DaprIoStateStoreComponentClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *DaprIoStateStoreComponentListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/dapr.io.StateStoreComponent"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DaprIoStateStoreComponentClient) listHandleResponse(resp *azcore.Response) (DaprStateStoreComponentListResponse, error) {
	var val *DaprStateStoreComponentList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DaprStateStoreComponentListResponse{}, err
	}
return DaprStateStoreComponentListResponse{RawResponse: resp.Response, DaprStateStoreComponentList: val}, nil
}

// listHandleError handles the List error response.
func (client *DaprIoStateStoreComponentClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

