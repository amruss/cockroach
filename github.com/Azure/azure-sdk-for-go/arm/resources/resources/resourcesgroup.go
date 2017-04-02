package resources

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// GroupClient is the provides operations for working with resources and
// resource groups.
type GroupClient struct {
	ManagementClient
}

// NewGroupClient creates an instance of the GroupClient client.
func NewGroupClient(subscriptionID string) GroupClient {
	return NewGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupClientWithBaseURI creates an instance of the GroupClient client.
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return GroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckExistence checks whether a resource exists.
//
// resourceGroupName is the name of the resource group containing the resource
// to check. The name is case insensitive. resourceProviderNamespace is the
// resource provider of the resource to check. parentResourcePath is the parent
// resource identity. resourceType is the resource type. resourceName is the
// name of the resource to check whether it exists.
func (client GroupClient) CheckExistence(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resources.GroupClient", "CheckExistence")
	}

	req, err := client.CheckExistencePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CheckExistence", nil, "Failure preparing request")
	}

	resp, err := client.CheckExistenceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CheckExistence", resp, "Failure sending request")
	}

	result, err = client.CheckExistenceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "CheckExistence", resp, "Failure responding to request")
	}

	return
}

// CheckExistencePreparer prepares the CheckExistence request.
func (client GroupClient) CheckExistencePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckExistenceSender sends the CheckExistence request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CheckExistenceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckExistenceResponder handles the response to the CheckExistence request. The method always
// closes the http.Response Body.
func (client GroupClient) CheckExistenceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CheckExistenceByID checks by ID whether a resource exists.
//
// resourceID is the fully qualified ID of the resource, including the resource
// name and resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/{resource-provider-namespace}/{resource-type}/{resource-name}
func (client GroupClient) CheckExistenceByID(resourceID string) (result autorest.Response, err error) {
	req, err := client.CheckExistenceByIDPreparer(resourceID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CheckExistenceByID", nil, "Failure preparing request")
	}

	resp, err := client.CheckExistenceByIDSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CheckExistenceByID", resp, "Failure sending request")
	}

	result, err = client.CheckExistenceByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "CheckExistenceByID", resp, "Failure responding to request")
	}

	return
}

// CheckExistenceByIDPreparer prepares the CheckExistenceByID request.
func (client GroupClient) CheckExistenceByIDPreparer(resourceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId": resourceID,
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckExistenceByIDSender sends the CheckExistenceByID request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CheckExistenceByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckExistenceByIDResponder handles the response to the CheckExistenceByID request. The method always
// closes the http.Response Body.
func (client GroupClient) CheckExistenceByIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate creates a resource. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group for the resource. The
// name is case insensitive. resourceProviderNamespace is the namespace of the
// resource provider. parentResourcePath is the parent resource identity.
// resourceType is the resource type of the resource to create. resourceName is
// the name of the resource to create. parameters is parameters for creating or
// updating the resource.
func (client GroupClient) CreateOrUpdate(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, parameters GenericResource, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Kind", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Kind", Name: validation.Pattern, Rule: `^[-\w\._,\(\)]+$`, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resources.GroupClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GroupClient) CreateOrUpdatePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, parameters GenericResource, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GroupClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdateByID create a resource by ID. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceID is the fully qualified ID of the resource, including the resource
// name and resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/{resource-provider-namespace}/{resource-type}/{resource-name}
// parameters is create or update resource parameters.
func (client GroupClient) CreateOrUpdateByID(resourceID string, parameters GenericResource, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Kind", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Kind", Name: validation.Pattern, Rule: `^[-\w\._,\(\)]+$`, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resources.GroupClient", "CreateOrUpdateByID")
	}

	req, err := client.CreateOrUpdateByIDPreparer(resourceID, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CreateOrUpdateByID", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateByIDSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "CreateOrUpdateByID", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "CreateOrUpdateByID", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateByIDPreparer prepares the CreateOrUpdateByID request.
func (client GroupClient) CreateOrUpdateByIDPreparer(resourceID string, parameters GenericResource, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId": resourceID,
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateByIDSender sends the CreateOrUpdateByID request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) CreateOrUpdateByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateByIDResponder handles the response to the CreateOrUpdateByID request. The method always
// closes the http.Response Body.
func (client GroupClient) CreateOrUpdateByIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes a resource. This method may poll for completion. Polling can
// be canceled by passing the cancel channel argument. The channel will be used
// to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group that contains the
// resource to delete. The name is case insensitive. resourceProviderNamespace
// is the namespace of the resource provider. parentResourcePath is the parent
// resource identity. resourceType is the resource type. resourceName is the
// name of the resource to delete.
func (client GroupClient) Delete(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resources.GroupClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GroupClient) DeletePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByID deletes a resource by ID. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceID is the fully qualified ID of the resource, including the resource
// name and resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/{resource-provider-namespace}/{resource-type}/{resource-name}
func (client GroupClient) DeleteByID(resourceID string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeleteByIDPreparer(resourceID, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "DeleteByID", nil, "Failure preparing request")
	}

	resp, err := client.DeleteByIDSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "DeleteByID", resp, "Failure sending request")
	}

	result, err = client.DeleteByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "DeleteByID", resp, "Failure responding to request")
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client GroupClient) DeleteByIDPreparer(resourceID string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId": resourceID,
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) DeleteByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client GroupClient) DeleteByIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a resource.
//
// resourceGroupName is the name of the resource group containing the resource
// to get. The name is case insensitive. resourceProviderNamespace is the
// namespace of the resource provider. parentResourcePath is the parent
// resource identity. resourceType is the resource type of the resource.
// resourceName is the name of the resource to get.
func (client GroupClient) Get(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result GenericResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resources.GroupClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GroupClient) GetPreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupClient) GetResponder(resp *http.Response) (result GenericResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets a resource by ID.
//
// resourceID is the fully qualified ID of the resource, including the resource
// name and resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/{resource-provider-namespace}/{resource-type}/{resource-name}
func (client GroupClient) GetByID(resourceID string) (result GenericResource, err error) {
	req, err := client.GetByIDPreparer(resourceID)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "GetByID", nil, "Failure preparing request")
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "GetByID", resp, "Failure sending request")
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client GroupClient) GetByIDPreparer(resourceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId": resourceID,
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client GroupClient) GetByIDResponder(resp *http.Response) (result GenericResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all the resources in a subscription.
//
// filter is the filter to apply on the operation. expand is the $expand query
// parameter. top is the number of results to return. If null is passed,
// returns all resource groups.
func (client GroupClient) List(filter string, expand string, top *int32) (result ListResult, err error) {
	req, err := client.ListPreparer(filter, expand, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupClient) ListPreparer(filter string, expand string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupClient) ListResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client GroupClient) ListNextResults(lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.ListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// MoveResources the resources to move must be in the same source resource
// group. The target resource group may be in a different subscription. When
// moving resources, both the source group and the target group are locked for
// the duration of the operation. Write and delete operations are blocked on
// the groups until the move completes.  This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// sourceResourceGroupName is the name of the resource group containing the
// rsources to move. parameters is parameters for moving resources.
func (client GroupClient) MoveResources(sourceResourceGroupName string, parameters MoveInfo, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: sourceResourceGroupName,
			Constraints: []validation.Constraint{{Target: "sourceResourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "sourceResourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "sourceResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resources.GroupClient", "MoveResources")
	}

	req, err := client.MoveResourcesPreparer(sourceResourceGroupName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "MoveResources", nil, "Failure preparing request")
	}

	resp, err := client.MoveResourcesSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources.GroupClient", "MoveResources", resp, "Failure sending request")
	}

	result, err = client.MoveResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.GroupClient", "MoveResources", resp, "Failure responding to request")
	}

	return
}

// MoveResourcesPreparer prepares the MoveResources request.
func (client GroupClient) MoveResourcesPreparer(sourceResourceGroupName string, parameters MoveInfo, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceResourceGroupName": autorest.Encode("path", sourceResourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{sourceResourceGroupName}/moveResources", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// MoveResourcesSender sends the MoveResources request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) MoveResourcesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// MoveResourcesResponder handles the response to the MoveResources request. The method always
// closes the http.Response Body.
func (client GroupClient) MoveResourcesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
