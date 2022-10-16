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

package extras

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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewExtrasCustomLinksUpdateParams creates a new ExtrasCustomLinksUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomLinksUpdateParams() *ExtrasCustomLinksUpdateParams {
	return &ExtrasCustomLinksUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomLinksUpdateParamsWithTimeout creates a new ExtrasCustomLinksUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomLinksUpdateParamsWithTimeout(timeout time.Duration) *ExtrasCustomLinksUpdateParams {
	return &ExtrasCustomLinksUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomLinksUpdateParamsWithContext creates a new ExtrasCustomLinksUpdateParams object
// with the ability to set a context for a request.
func NewExtrasCustomLinksUpdateParamsWithContext(ctx context.Context) *ExtrasCustomLinksUpdateParams {
	return &ExtrasCustomLinksUpdateParams{
		Context: ctx,
	}
}

// NewExtrasCustomLinksUpdateParamsWithHTTPClient creates a new ExtrasCustomLinksUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomLinksUpdateParamsWithHTTPClient(client *http.Client) *ExtrasCustomLinksUpdateParams {
	return &ExtrasCustomLinksUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasCustomLinksUpdateParams contains all the parameters to send to the API endpoint

	for the extras custom links update operation.

	Typically these are written to a http.Request.
*/
type ExtrasCustomLinksUpdateParams struct {

	// Data.
	Data *models.CustomLink

	/* ID.

	   A unique integer value identifying this custom link.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom links update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksUpdateParams) WithDefaults() *ExtrasCustomLinksUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom links update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) WithTimeout(timeout time.Duration) *ExtrasCustomLinksUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) WithContext(ctx context.Context) *ExtrasCustomLinksUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) WithHTTPClient(client *http.Client) *ExtrasCustomLinksUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) WithData(data *models.CustomLink) *ExtrasCustomLinksUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) SetData(data *models.CustomLink) {
	o.Data = data
}

// WithID adds the id to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) WithID(id int64) *ExtrasCustomLinksUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom links update params
func (o *ExtrasCustomLinksUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomLinksUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
