// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewIpamFhrpGroupsBulkUpdateParams creates a new IpamFhrpGroupsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamFhrpGroupsBulkUpdateParams() *IpamFhrpGroupsBulkUpdateParams {
	return &IpamFhrpGroupsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamFhrpGroupsBulkUpdateParamsWithTimeout creates a new IpamFhrpGroupsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamFhrpGroupsBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamFhrpGroupsBulkUpdateParams {
	return &IpamFhrpGroupsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamFhrpGroupsBulkUpdateParamsWithContext creates a new IpamFhrpGroupsBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamFhrpGroupsBulkUpdateParamsWithContext(ctx context.Context) *IpamFhrpGroupsBulkUpdateParams {
	return &IpamFhrpGroupsBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamFhrpGroupsBulkUpdateParamsWithHTTPClient creates a new IpamFhrpGroupsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamFhrpGroupsBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamFhrpGroupsBulkUpdateParams {
	return &IpamFhrpGroupsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamFhrpGroupsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the ipam fhrp groups bulk update operation.

	Typically these are written to a http.Request.
*/
type IpamFhrpGroupsBulkUpdateParams struct {

	// Data.
	Data *models.FHRPGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam fhrp groups bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupsBulkUpdateParams) WithDefaults() *IpamFhrpGroupsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam fhrp groups bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamFhrpGroupsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamFhrpGroupsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) WithContext(ctx context.Context) *IpamFhrpGroupsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamFhrpGroupsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) WithData(data *models.FHRPGroup) *IpamFhrpGroupsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam fhrp groups bulk update params
func (o *IpamFhrpGroupsBulkUpdateParams) SetData(data *models.FHRPGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamFhrpGroupsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
