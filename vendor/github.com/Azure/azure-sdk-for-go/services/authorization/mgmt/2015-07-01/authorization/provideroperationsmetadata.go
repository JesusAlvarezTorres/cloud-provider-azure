package authorization

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProviderOperationsMetadataClient is the role based access control provides you a way to apply granular level policy
// administration down to individual resources or resource groups. These operations enable you to manage role
// definitions and role assignments. A role definition describes the set of actions that can be performed on resources.
// A role assignment grants access to Azure Active Directory users.
type ProviderOperationsMetadataClient struct {
	BaseClient
}

// NewProviderOperationsMetadataClient creates an instance of the ProviderOperationsMetadataClient client.
func NewProviderOperationsMetadataClient(subscriptionID string) ProviderOperationsMetadataClient {
	return NewProviderOperationsMetadataClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderOperationsMetadataClientWithBaseURI creates an instance of the ProviderOperationsMetadataClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewProviderOperationsMetadataClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationsMetadataClient {
	return ProviderOperationsMetadataClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets provider operations metadata for the specified resource provider.
// Parameters:
// resourceProviderNamespace - the namespace of the resource provider.
// APIVersion - the API version to use for the operation.
// expand - specifies whether to expand the values.
func (client ProviderOperationsMetadataClient) Get(ctx context.Context, resourceProviderNamespace string, APIVersion string, expand string) (result ProviderOperationsMetadata, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderOperationsMetadataClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceProviderNamespace, APIVersion, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProviderOperationsMetadataClient) GetPreparer(ctx context.Context, resourceProviderNamespace string, APIVersion string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	} else {
		queryParameters["$expand"] = autorest.Encode("query", "resourceTypes")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/providerOperations/{resourceProviderNamespace}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationsMetadataClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProviderOperationsMetadataClient) GetResponder(resp *http.Response) (result ProviderOperationsMetadata, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets provider operations metadata for all resource providers.
// Parameters:
// APIVersion - the API version to use for this operation.
// expand - specifies whether to expand the values.
func (client ProviderOperationsMetadataClient) List(ctx context.Context, APIVersion string, expand string) (result ProviderOperationsMetadataListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderOperationsMetadataClient.List")
		defer func() {
			sc := -1
			if result.pomlr.Response.Response != nil {
				sc = result.pomlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, APIVersion, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pomlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", resp, "Failure sending request")
		return
	}

	result.pomlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "List", resp, "Failure responding to request")
		return
	}
	if result.pomlr.hasNextLink() && result.pomlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client ProviderOperationsMetadataClient) ListPreparer(ctx context.Context, APIVersion string, expand string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	} else {
		queryParameters["$expand"] = autorest.Encode("query", "resourceTypes")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/providerOperations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationsMetadataClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProviderOperationsMetadataClient) ListResponder(resp *http.Response) (result ProviderOperationsMetadataListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProviderOperationsMetadataClient) listNextResults(ctx context.Context, lastResults ProviderOperationsMetadataListResult) (result ProviderOperationsMetadataListResult, err error) {
	req, err := lastResults.providerOperationsMetadataListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.ProviderOperationsMetadataClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderOperationsMetadataClient) ListComplete(ctx context.Context, APIVersion string, expand string) (result ProviderOperationsMetadataListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderOperationsMetadataClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, APIVersion, expand)
	return
}
