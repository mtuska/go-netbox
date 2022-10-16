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

package dcim

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
	"github.com/go-openapi/swag"
)

// NewDcimModuleBaysReadParams creates a new DcimModuleBaysReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleBaysReadParams() *DcimModuleBaysReadParams {
	return &DcimModuleBaysReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleBaysReadParamsWithTimeout creates a new DcimModuleBaysReadParams object
// with the ability to set a timeout on a request.
func NewDcimModuleBaysReadParamsWithTimeout(timeout time.Duration) *DcimModuleBaysReadParams {
	return &DcimModuleBaysReadParams{
		timeout: timeout,
	}
}

// NewDcimModuleBaysReadParamsWithContext creates a new DcimModuleBaysReadParams object
// with the ability to set a context for a request.
func NewDcimModuleBaysReadParamsWithContext(ctx context.Context) *DcimModuleBaysReadParams {
	return &DcimModuleBaysReadParams{
		Context: ctx,
	}
}

// NewDcimModuleBaysReadParamsWithHTTPClient creates a new DcimModuleBaysReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleBaysReadParamsWithHTTPClient(client *http.Client) *DcimModuleBaysReadParams {
	return &DcimModuleBaysReadParams{
		HTTPClient: client,
	}
}

/*
DcimModuleBaysReadParams contains all the parameters to send to the API endpoint

	for the dcim module bays read operation.

	Typically these are written to a http.Request.
*/
type DcimModuleBaysReadParams struct {

	/* ID.

	   A unique integer value identifying this module bay.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module bays read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBaysReadParams) WithDefaults() *DcimModuleBaysReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module bays read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBaysReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module bays read params
func (o *DcimModuleBaysReadParams) WithTimeout(timeout time.Duration) *DcimModuleBaysReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module bays read params
func (o *DcimModuleBaysReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module bays read params
func (o *DcimModuleBaysReadParams) WithContext(ctx context.Context) *DcimModuleBaysReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module bays read params
func (o *DcimModuleBaysReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module bays read params
func (o *DcimModuleBaysReadParams) WithHTTPClient(client *http.Client) *DcimModuleBaysReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module bays read params
func (o *DcimModuleBaysReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim module bays read params
func (o *DcimModuleBaysReadParams) WithID(id int64) *DcimModuleBaysReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim module bays read params
func (o *DcimModuleBaysReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleBaysReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
