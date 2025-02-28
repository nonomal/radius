/*
Copyright 2023 The Radius Authors.

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

package clients

import (
	"context"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"golang.org/x/sync/errgroup"

	"github.com/radius-project/radius/pkg/azure/clientv2"
	aztoken "github.com/radius-project/radius/pkg/azure/tokencredentials"
	"github.com/radius-project/radius/pkg/cli/clients_new/generated"
	corerpv20231001 "github.com/radius-project/radius/pkg/corerp/api/v20231001preview"
	cntr_ctrl "github.com/radius-project/radius/pkg/corerp/frontend/controller/containers"
	ext_ctrl "github.com/radius-project/radius/pkg/corerp/frontend/controller/extenders"
	gtwy_ctrl "github.com/radius-project/radius/pkg/corerp/frontend/controller/gateways"
	sstr_ctrl "github.com/radius-project/radius/pkg/corerp/frontend/controller/secretstores"
	dapr_ctrl "github.com/radius-project/radius/pkg/daprrp/frontend/controller"
	ds_ctrl "github.com/radius-project/radius/pkg/datastoresrp/frontend/controller"
	msg_ctrl "github.com/radius-project/radius/pkg/messagingrp/frontend/controller"
	ucpv20231001 "github.com/radius-project/radius/pkg/ucp/api/v20231001preview"
	"github.com/radius-project/radius/pkg/ucp/resources"
	resources_radius "github.com/radius-project/radius/pkg/ucp/resources/radius"
)

type UCPApplicationsManagementClient struct {
	RootScope     string
	ClientOptions *arm.ClientOptions
}

var _ ApplicationsManagementClient = (*UCPApplicationsManagementClient)(nil)

var (
	ResourceTypesList = []string{
		ds_ctrl.MongoDatabasesResourceType,
		msg_ctrl.RabbitMQQueuesResourceType,
		ds_ctrl.RedisCachesResourceType,
		ds_ctrl.SqlDatabasesResourceType,
		dapr_ctrl.DaprStateStoresResourceType,
		dapr_ctrl.DaprSecretStoresResourceType,
		dapr_ctrl.DaprPubSubBrokersResourceType,
		ext_ctrl.ResourceTypeName,
		gtwy_ctrl.ResourceTypeName,
		cntr_ctrl.ResourceTypeName,
		sstr_ctrl.ResourceTypeName,
	}
)

// ListAllResourcesByType lists the all the resources within a scope
//

// ListAllResourcesByType retrieves a list of all resources of a given type from the root
// scope, and returns them in a slice of GenericResource objects, or an error if one occurs.
func (amc *UCPApplicationsManagementClient) ListAllResourcesByType(ctx context.Context, resourceType string) ([]generated.GenericResource, error) {
	results := []generated.GenericResource{}

	client, err := generated.NewGenericResourcesClient(amc.RootScope, resourceType, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return results, err
	}

	pager := client.NewListByRootScopePager(&generated.GenericResourcesClientListByRootScopeOptions{})
	for pager.More() {
		nextPage, err := pager.NextPage(ctx)
		if err != nil {
			return results, err
		}
		applicationList := nextPage.GenericResourcesList.Value
		for _, application := range applicationList {
			results = append(results, *application)
		}
	}

	return results, nil
}

// ListAllResourceOfTypeInApplication lists the resources of a particular type in an application
//

// ListAllResourcesOfTypeInApplication takes in a context, an application name and a
// resource type and returns a slice of GenericResources and an error if one occurs.
func (amc *UCPApplicationsManagementClient) ListAllResourcesOfTypeInApplication(ctx context.Context, applicationName string, resourceType string) ([]generated.GenericResource, error) {
	results := []generated.GenericResource{}
	resourceList, err := amc.ListAllResourcesByType(ctx, resourceType)
	if err != nil {
		return nil, err
	}
	for _, resource := range resourceList {
		isResourceWithApplication := isResourceInApplication(resource, applicationName)
		if isResourceWithApplication {
			results = append(results, resource)
		}
	}
	return results, nil
}

// ListAllResourcesByApplication lists the resources of a particular application
//

// ListAllResourcesByApplication takes in a context and an application name and returns
// a slice of GenericResources and an error if one occurs.
func (amc *UCPApplicationsManagementClient) ListAllResourcesByApplication(ctx context.Context, applicationName string) ([]generated.GenericResource, error) {
	results := []generated.GenericResource{}
	for _, resourceType := range ResourceTypesList {
		resourceList, err := amc.ListAllResourcesOfTypeInApplication(ctx, applicationName, resourceType)
		if err != nil {
			return nil, err
		}
		results = append(results, resourceList...)
	}

	return results, nil
}

// ListAllResourcesByEnvironment lists the all the resources of a particular environment
//

// ListAllResourcesByEnvironment iterates through a list of resource types and calls ListAllResourcesOfTypeInEnvironment
// for each one, appending the results to a slice of GenericResources and returning it. If an error is encountered, it is returned.
func (amc *UCPApplicationsManagementClient) ListAllResourcesByEnvironment(ctx context.Context, environmentName string) ([]generated.GenericResource, error) {
	results := []generated.GenericResource{}
	for _, resourceType := range ResourceTypesList {
		resourceList, err := amc.ListAllResourcesOfTypeInEnvironment(ctx, environmentName, resourceType)
		if err != nil {
			return nil, err
		}
		results = append(results, resourceList...)
	}

	return results, nil
}

// ListAllResourcesByTypeInEnvironment lists the all the resources of a particular type in an environment
//

// ListAllResourcesOfTypeInEnvironment takes in a context, an environment name and a
// resource type and returns a slice of GenericResources and an error if one occurs.
func (amc *UCPApplicationsManagementClient) ListAllResourcesOfTypeInEnvironment(ctx context.Context, environmentName string, resourceType string) ([]generated.GenericResource, error) {
	results := []generated.GenericResource{}
	resourceList, err := amc.ListAllResourcesByType(ctx, resourceType)
	if err != nil {
		return nil, err
	}
	for _, resource := range resourceList {
		isResourceWithApplication := isResourceInEnvironment(resource, environmentName)
		if isResourceWithApplication {
			results = append(results, resource)
		}
	}
	return results, nil
}

// ShowResource creates a new client for a given resource type and attempts to retrieve the resource with the given name,
// returning the resource or an error if one occurs.
func (amc *UCPApplicationsManagementClient) ShowResource(ctx context.Context, resourceType string, resourceName string) (generated.GenericResource, error) {
	client, err := generated.NewGenericResourcesClient(amc.RootScope, resourceType, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return generated.GenericResource{}, err
	}

	getResponse, err := client.Get(ctx, resourceName, &generated.GenericResourcesClientGetOptions{})
	if err != nil {
		return generated.GenericResource{}, err
	}

	return getResponse.GenericResource, nil
}

// DeleteResource creates a new client, sends a delete request to the resource, polls until the request is completed,
// and returns a boolean indicating whether the resource was successfully deleted or not, and an error if one occurred.
func (amc *UCPApplicationsManagementClient) DeleteResource(ctx context.Context, resourceType string, resourceName string) (bool, error) {
	client, err := generated.NewGenericResourcesClient(amc.RootScope, resourceType, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return false, err
	}

	var respFromCtx *http.Response
	ctxWithResp := runtime.WithCaptureResponse(ctx, &respFromCtx)

	poller, err := client.BeginDelete(ctxWithResp, resourceName, nil)
	if err != nil {
		return false, err
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		return false, err
	}

	return respFromCtx.StatusCode != 204, nil
}

// ListApplications() retrieves a list of ApplicationResource objects from the Azure API
// and returns them in a slice, or an error if one occurs.
func (amc *UCPApplicationsManagementClient) ListApplications(ctx context.Context) ([]corerpv20231001.ApplicationResource, error) {
	results := []corerpv20231001.ApplicationResource{}

	client, err := corerpv20231001.NewApplicationsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return results, err
	}

	pager := client.NewListByScopePager(&corerpv20231001.ApplicationsClientListByScopeOptions{})
	for pager.More() {
		nextPage, err := pager.NextPage(ctx)
		if err != nil {
			return results, err
		}
		applicationList := nextPage.ApplicationResourceListResult.Value
		for _, application := range applicationList {
			results = append(results, *application)
		}
	}

	return results, nil
}

// ListApplicationsByEnv takes in a context and an environment name and returns a slice of ApplicationResource objects
// and an error if one occurs.
func (amc *UCPApplicationsManagementClient) ListApplicationsByEnv(ctx context.Context, envName string) ([]corerpv20231001.ApplicationResource, error) {
	results := []corerpv20231001.ApplicationResource{}
	applicationsList, err := amc.ListApplications(ctx)
	if err != nil {
		return nil, err
	}
	envID := "/" + amc.RootScope + "/providers/applications.core/environments/" + envName
	for _, application := range applicationsList {
		if strings.EqualFold(envID, *application.Properties.Environment) {
			results = append(results, application)
		}
	}
	return results, nil
}

// ShowApplication creates a new ApplicationsClient, attempts to get an application
// resource from the Azure Cognitive Search service, and returns the resource or an error if one occurs.
func (amc *UCPApplicationsManagementClient) ShowApplication(ctx context.Context, applicationName string) (corerpv20231001.ApplicationResource, error) {
	client, err := corerpv20231001.NewApplicationsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return corerpv20231001.ApplicationResource{}, err
	}

	getResponse, err := client.Get(ctx, applicationName, &corerpv20231001.ApplicationsClientGetOptions{})
	var result corerpv20231001.ApplicationResource
	if err != nil {
		return result, err
	}
	result = getResponse.ApplicationResource
	return result, nil
}

// GetGraph creates a new ApplicationsClient, returns the application graph or an error if one occurs.
func (amc *UCPApplicationsManagementClient) GetGraph(ctx context.Context, applicationName string) (corerpv20231001.ApplicationGraphResponse, error) {
	client, err := corerpv20231001.NewApplicationsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return corerpv20231001.ApplicationGraphResponse{}, err
	}

	getResponse, err := client.GetGraph(ctx, applicationName, map[string]any{}, &corerpv20231001.ApplicationsClientGetGraphOptions{})
	if err != nil {
		return corerpv20231001.ApplicationGraphResponse{}, err
	}

	return getResponse.ApplicationGraphResponse, nil
}

// DeleteApplication deletes an application and all its associated resources, and returns an error if any of the operations fail.
func (amc *UCPApplicationsManagementClient) DeleteApplication(ctx context.Context, applicationName string) (bool, error) {
	// This handles the case where the application doesn't exist.
	resourcesWithApplication, err := amc.ListAllResourcesByApplication(ctx, applicationName)
	if err != nil && !clientv2.Is404Error(err) {
		return false, err
	}

	g, groupCtx := errgroup.WithContext(ctx)
	for _, resource := range resourcesWithApplication {
		resource := resource
		g.Go(func() error {
			_, err := amc.DeleteResource(groupCtx, *resource.Type, *resource.Name)
			if err != nil {
				return err
			}
			return nil
		})
	}

	err = g.Wait()
	if err != nil {
		return false, err
	}

	client, err := corerpv20231001.NewApplicationsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return false, err
	}

	var respFromCtx *http.Response
	ctxWithResp := runtime.WithCaptureResponse(ctx, &respFromCtx)

	_, err = client.Delete(ctxWithResp, applicationName, nil)
	if err != nil {
		return false, err
	}

	return respFromCtx.StatusCode != 204, nil
}

// CreateOrUpdateApplication creates or updates an application.
//

// CreateOrUpdateApplication creates or updates an application resource in Azure using the
// given application name and resource. It returns an error if the creation or update fails.
func (amc *UCPApplicationsManagementClient) CreateOrUpdateApplication(ctx context.Context, applicationName string, resource corerpv20231001.ApplicationResource) error {
	client, err := corerpv20231001.NewApplicationsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return err
	}

	_, err = client.CreateOrUpdate(ctx, applicationName, resource, nil)
	if err != nil {
		return err
	}

	return nil
}

// CreateApplicationIfNotFound creates an application if it does not exist.
//

// CreateApplicationIfNotFound checks if an application exists and creates it if it does
// not exist, returning an error if any occurs.
func (amc *UCPApplicationsManagementClient) CreateApplicationIfNotFound(ctx context.Context, applicationName string, resource corerpv20231001.ApplicationResource) error {
	client, err := corerpv20231001.NewApplicationsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return err
	}

	_, err = client.Get(ctx, applicationName, nil)
	if Is404Error(err) {
		// continue
	} else if err != nil {
		return err
	} else {
		// Application already exists, nothing to do.
		return nil
	}

	_, err = client.CreateOrUpdate(ctx, applicationName, resource, nil)
	if err != nil {
		return err
	}

	return nil
}

// Creates a Radius Environment resource
//

// CreateEnvironment creates or updates an environment with the given name, location and
// properties, and returns an error if one occurs.
func (amc *UCPApplicationsManagementClient) CreateEnvironment(ctx context.Context, envName string, location string, envProperties *corerpv20231001.EnvironmentProperties) error {
	client, err := corerpv20231001.NewEnvironmentsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return err
	}

	_, err = client.CreateOrUpdate(ctx, envName, corerpv20231001.EnvironmentResource{Location: &location, Properties: envProperties}, &corerpv20231001.EnvironmentsClientCreateOrUpdateOptions{})
	if err != nil {
		return err
	}

	return nil

}

func isResourceInApplication(resource generated.GenericResource, applicationName string) bool {
	obj, found := resource.Properties["application"]
	// A resource may not have an application associated with it.
	if !found {
		return false
	}

	associatedAppId, ok := obj.(string)
	if !ok || associatedAppId == "" {
		return false
	}

	idParsed, err := resources.ParseResource(associatedAppId)
	if err != nil {
		return false
	}

	if strings.EqualFold(idParsed.Name(), applicationName) {
		return true
	}

	return false
}

func isResourceInEnvironment(resource generated.GenericResource, environmentName string) bool {
	obj, found := resource.Properties["environment"]
	// A resource may not have an environment associated with it.
	if !found {
		return false
	}

	associatedEnvId, ok := obj.(string)
	if !ok || associatedEnvId == "" {
		return false
	}

	idParsed, err := resources.ParseResource(associatedEnvId)
	if err != nil {
		return false
	}

	if strings.EqualFold(idParsed.Name(), environmentName) {
		return true
	}

	return false
}

// ListEnvironmentsInResourceGroup creates a list of environment resources by paging through the list of environments in
// the resource group and appending each environment to the list. It returns the list of environment resources or an error
// if one occurs.
func (amc *UCPApplicationsManagementClient) ListEnvironmentsInResourceGroup(ctx context.Context) ([]corerpv20231001.EnvironmentResource, error) {
	envResourceList := []corerpv20231001.EnvironmentResource{}

	envClient, err := corerpv20231001.NewEnvironmentsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return envResourceList, err
	}

	pager := envClient.NewListByScopePager(&corerpv20231001.EnvironmentsClientListByScopeOptions{})
	for pager.More() {
		nextPage, err := pager.NextPage(ctx)
		if err != nil {
			return envResourceList, err
		}
		applicationList := nextPage.EnvironmentResourceListResult.Value
		for _, application := range applicationList {
			envResourceList = append(envResourceList, *application)
		}
	}

	return envResourceList, nil
}

// ListEnvironmentsAll queries the scope for all environment resources and returns a slice of environment resources or an error if one occurs.
func (amc *UCPApplicationsManagementClient) ListEnvironmentsAll(ctx context.Context) ([]corerpv20231001.EnvironmentResource, error) {
	scope, err := resources.ParseScope("/" + amc.RootScope)
	if err != nil {
		return []corerpv20231001.EnvironmentResource{}, err
	}

	// Query at plane scope, not resource group scope. We don't enforce the exact structure of the scope, so handle both cases.
	//
	// - /planes/radius/local
	// - /planes/radius/local/resourceGroups/my-group
	if scope.FindScope(resources_radius.ScopeResourceGroups) != "" {
		scope = scope.Truncate()
	}

	environments := []corerpv20231001.EnvironmentResource{}
	client, err := corerpv20231001.NewEnvironmentsClient(scope.String(), &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return []corerpv20231001.EnvironmentResource{}, err
	}

	pager := client.NewListByScopePager(&corerpv20231001.EnvironmentsClientListByScopeOptions{})
	for pager.More() {
		nextPage, err := pager.NextPage(ctx)
		if err != nil {
			return []corerpv20231001.EnvironmentResource{}, err
		}

		for _, environment := range nextPage.EnvironmentResourceListResult.Value {
			environments = append(environments, *environment)
		}
	}

	return environments, nil
}

// GetEnvDetails attempts to retrieve an environment resource from an environment client, and returns the environment
// resource or an error if unsuccessful.
func (amc *UCPApplicationsManagementClient) GetEnvDetails(ctx context.Context, envName string) (corerpv20231001.EnvironmentResource, error) {
	envClient, err := corerpv20231001.NewEnvironmentsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return corerpv20231001.EnvironmentResource{}, err
	}

	envGetResp, err := envClient.Get(ctx, envName, &corerpv20231001.EnvironmentsClientGetOptions{})
	if err == nil {
		return envGetResp.EnvironmentResource, nil
	}

	return corerpv20231001.EnvironmentResource{}, err

}

// DeleteEnv function checks if there are any applications associated with the given environment, deletes them if found,
// and then deletes the environment itself. It returns a boolean and an error if one occurs.
func (amc *UCPApplicationsManagementClient) DeleteEnv(ctx context.Context, envName string) (bool, error) {
	applicationsWithEnv, err := amc.ListApplicationsByEnv(ctx, envName)
	if err != nil {
		return false, err
	}

	for _, application := range applicationsWithEnv {
		_, err := amc.DeleteApplication(ctx, *application.Name)
		if err != nil {
			return false, err
		}
	}

	envClient, err := corerpv20231001.NewEnvironmentsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return false, err
	}

	var respFromCtx *http.Response
	ctxWithResp := runtime.WithCaptureResponse(ctx, &respFromCtx)

	_, err = envClient.Delete(ctxWithResp, envName, nil)
	if err != nil {
		return false, err
	}

	return respFromCtx.StatusCode != 204, nil
}

// CreateUCPGroup creates a new resource group in the specified plane type and plane name using the provided resource
// group resource and returns an error if one occurs.
func (amc *UCPApplicationsManagementClient) CreateUCPGroup(ctx context.Context, planeName string, resourceGroupName string, resourceGroup ucpv20231001.ResourceGroupResource) error {
	var resourceGroupOptions *ucpv20231001.ResourceGroupsClientCreateOrUpdateOptions
	resourcegroupClient, err := ucpv20231001.NewResourceGroupsClient(&aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return err
	}

	_, err = resourcegroupClient.CreateOrUpdate(ctx, planeName, resourceGroupName, resourceGroup, resourceGroupOptions)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUCPGroup attempts to delete a UCP resource group using the provided plane type, plane name and resource group
// name, and returns a boolean indicating success or failure and an error if one occurs.
func (amc *UCPApplicationsManagementClient) DeleteUCPGroup(ctx context.Context, planeName string, resourceGroupName string) (bool, error) {
	var resourceGroupOptions *ucpv20231001.ResourceGroupsClientDeleteOptions
	resourcegroupClient, err := ucpv20231001.NewResourceGroupsClient(&aztoken.AnonymousCredential{}, amc.ClientOptions)

	var respFromCtx *http.Response
	ctxWithResp := runtime.WithCaptureResponse(ctx, &respFromCtx)
	if err != nil {
		return false, err
	}

	_, err = resourcegroupClient.Delete(ctxWithResp, planeName, resourceGroupName, resourceGroupOptions)
	if err != nil {
		return false, err
	}

	return respFromCtx.StatusCode == 204, nil

}

// ShowUCPGroup is a function that retrieves a resource group from the UCP API using the given plane type,
// plane name and resource group name, and returns the resource group resource or an error if one occurs.
func (amc *UCPApplicationsManagementClient) ShowUCPGroup(ctx context.Context, planeName string, resourceGroupName string) (ucpv20231001.ResourceGroupResource, error) {
	var resourceGroupOptions *ucpv20231001.ResourceGroupsClientGetOptions
	resourcegroupClient, err := ucpv20231001.NewResourceGroupsClient(&aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return ucpv20231001.ResourceGroupResource{}, err
	}

	resp, err := resourcegroupClient.Get(ctx, planeName, resourceGroupName, resourceGroupOptions)
	if err != nil {
		return ucpv20231001.ResourceGroupResource{}, err
	}

	return resp.ResourceGroupResource, nil
}

// ListUCPGroup is a function that retrieves a list of resource groups from the UCP API and returns them as a slice of
// ResourceGroupResource objects. It may return an error if there is an issue with the API request.
func (amc *UCPApplicationsManagementClient) ListUCPGroup(ctx context.Context, planeName string) ([]ucpv20231001.ResourceGroupResource, error) {
	var resourceGroupOptions *ucpv20231001.ResourceGroupsClientListOptions
	resourceGroupResources := []ucpv20231001.ResourceGroupResource{}
	resourcegroupClient, err := ucpv20231001.NewResourceGroupsClient(&aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return resourceGroupResources, err
	}

	pager := resourcegroupClient.NewListPager(planeName, resourceGroupOptions)

	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			return resourceGroupResources, err
		}

		resourceGroupList := resp.Value
		for _, resourceGroup := range resourceGroupList {
			resourceGroupResources = append(resourceGroupResources, *resourceGroup)

		}
	}

	return resourceGroupResources, nil
}

// ShowRecipe creates a new EnvironmentsClient, gets the recipe metadata from the
// environment, and returns the EnvironmentRecipeProperties or an error if one occurs.
func (amc *UCPApplicationsManagementClient) ShowRecipe(ctx context.Context, environmentName string, recipeName corerpv20231001.RecipeGetMetadata) (corerpv20231001.RecipeGetMetadataResponse, error) {
	client, err := corerpv20231001.NewEnvironmentsClient(amc.RootScope, &aztoken.AnonymousCredential{}, amc.ClientOptions)
	if err != nil {
		return corerpv20231001.RecipeGetMetadataResponse{}, err
	}

	resp, err := client.GetMetadata(ctx, environmentName, recipeName, &corerpv20231001.EnvironmentsClientGetMetadataOptions{})
	if err != nil {
		return corerpv20231001.RecipeGetMetadataResponse{}, err
	}

	return corerpv20231001.RecipeGetMetadataResponse(resp.RecipeGetMetadataResponse), nil
}
