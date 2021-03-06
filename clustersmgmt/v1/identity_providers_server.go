/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// IdentityProvidersServer represents the interface the manages the 'identity_providers' resource.
type IdentityProvidersServer interface {

	// Add handles a request for the 'add' method.
	//
	// Adds a new identity provider to the cluster.
	Add(ctx context.Context, request *IdentityProvidersAddServerRequest, response *IdentityProvidersAddServerResponse) error

	// List handles a request for the 'list' method.
	//
	// Retrieves the list of identity providers.
	List(ctx context.Context, request *IdentityProvidersListServerRequest, response *IdentityProvidersListServerResponse) error

	// IdentityProvider returns the target 'identity_provider' server for the given identifier.
	//
	// Reference to the service that manages an specific identity provider.
	IdentityProvider(id string) IdentityProviderServer
}

// IdentityProvidersAddServerRequest is the request for the 'add' method.
type IdentityProvidersAddServerRequest struct {
	body *IdentityProvider
}

// Body returns the value of the 'body' parameter.
//
// Description of the cluster.
func (r *IdentityProvidersAddServerRequest) Body() *IdentityProvider {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Description of the cluster.
func (r *IdentityProvidersAddServerRequest) GetBody() (value *IdentityProvider, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// unmarshal is the method used internally to unmarshal request to the
// 'add' method.
func (r *IdentityProvidersAddServerRequest) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(identityProviderData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// IdentityProvidersAddServerResponse is the response for the 'add' method.
type IdentityProvidersAddServerResponse struct {
	status int
	err    *errors.Error
	body   *IdentityProvider
}

// Body sets the value of the 'body' parameter.
//
// Description of the cluster.
func (r *IdentityProvidersAddServerResponse) Body(value *IdentityProvider) *IdentityProvidersAddServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *IdentityProvidersAddServerResponse) SetStatusCode(status int) *IdentityProvidersAddServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'add' method.
func (r *IdentityProvidersAddServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// IdentityProvidersListServerRequest is the request for the 'list' method.
type IdentityProvidersListServerRequest struct {
}

// IdentityProvidersListServerResponse is the response for the 'list' method.
type IdentityProvidersListServerResponse struct {
	status int
	err    *errors.Error
	items  *IdentityProviderList
	page   *int
	size   *int
	total  *int
}

// Items sets the value of the 'items' parameter.
//
// Retrieved list of identity providers.
func (r *IdentityProvidersListServerResponse) Items(value *IdentityProviderList) *IdentityProvidersListServerResponse {
	r.items = value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *IdentityProvidersListServerResponse) Page(value int) *IdentityProvidersListServerResponse {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *IdentityProvidersListServerResponse) Size(value int) *IdentityProvidersListServerResponse {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *IdentityProvidersListServerResponse) Total(value int) *IdentityProvidersListServerResponse {
	r.total = &value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *IdentityProvidersListServerResponse) SetStatusCode(status int) *IdentityProvidersListServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'list' method.
func (r *IdentityProvidersListServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data := new(identityProvidersListServerResponseData)
	data.Items, err = r.items.wrap()
	if err != nil {
		return err
	}
	data.Page = r.page
	data.Size = r.size
	data.Total = r.total
	err = encoder.Encode(data)
	return err
}

// identityProvidersListServerResponseData is the structure used internally to write the request of the
// 'list' method.
type identityProvidersListServerResponseData struct {
	Items identityProviderListData "json:\"items,omitempty\""
	Page  *int                     "json:\"page,omitempty\""
	Size  *int                     "json:\"size,omitempty\""
	Total *int                     "json:\"total,omitempty\""
}

// IdentityProvidersServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type IdentityProvidersServerAdapter struct {
	server IdentityProvidersServer
	router *mux.Router
}

func NewIdentityProvidersServerAdapter(server IdentityProvidersServer, router *mux.Router) *IdentityProvidersServerAdapter {
	adapter := new(IdentityProvidersServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.PathPrefix("/{id}").HandlerFunc(adapter.identityProviderHandler)
	adapter.router.Methods("POST").Path("").HandlerFunc(adapter.addHandler)
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.listHandler)
	return adapter
}
func (a *IdentityProvidersServerAdapter) identityProviderHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	target := a.server.IdentityProvider(id)
	targetAdapter := NewIdentityProviderServerAdapter(target, a.router.PathPrefix("/{id}").Subrouter())
	targetAdapter.ServeHTTP(w, r)
	return
}
func (a *IdentityProvidersServerAdapter) readIdentityProvidersAddServerRequest(r *http.Request) (*IdentityProvidersAddServerRequest, error) {
	var err error
	result := new(IdentityProvidersAddServerRequest)
	err = result.unmarshal(r.Body)
	if err != nil {
		return nil, err
	}
	return result, err
}
func (a *IdentityProvidersServerAdapter) writeIdentityProvidersAddServerResponse(w http.ResponseWriter, r *IdentityProvidersAddServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *IdentityProvidersServerAdapter) addHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readIdentityProvidersAddServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(IdentityProvidersAddServerResponse)
	err = a.server.Add(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Add: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeIdentityProvidersAddServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *IdentityProvidersServerAdapter) readIdentityProvidersListServerRequest(r *http.Request) (*IdentityProvidersListServerRequest, error) {
	var err error
	result := new(IdentityProvidersListServerRequest)
	return result, err
}
func (a *IdentityProvidersServerAdapter) writeIdentityProvidersListServerResponse(w http.ResponseWriter, r *IdentityProvidersListServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *IdentityProvidersServerAdapter) listHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readIdentityProvidersListServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(IdentityProvidersListServerResponse)
	err = a.server.List(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method List: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeIdentityProvidersListServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *IdentityProvidersServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
