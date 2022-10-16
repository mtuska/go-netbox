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
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimLocationsListReader is a Reader for the DcimLocationsList structure.
type DcimLocationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsListOK creates a DcimLocationsListOK with default headers values
func NewDcimLocationsListOK() *DcimLocationsListOK {
	return &DcimLocationsListOK{}
}

/*
DcimLocationsListOK describes a response with status code 200, with default header values.

DcimLocationsListOK dcim locations list o k
*/
type DcimLocationsListOK struct {
	Payload *DcimLocationsListOKBody
}

// IsSuccess returns true when this dcim locations list o k response has a 2xx status code
func (o *DcimLocationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations list o k response has a 3xx status code
func (o *DcimLocationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations list o k response has a 4xx status code
func (o *DcimLocationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations list o k response has a 5xx status code
func (o *DcimLocationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations list o k response a status code equal to that given
func (o *DcimLocationsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimLocationsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/locations/][%d] dcimLocationsListOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsListOK) String() string {
	return fmt.Sprintf("[GET /dcim/locations/][%d] dcimLocationsListOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsListOK) GetPayload() *DcimLocationsListOKBody {
	return o.Payload
}

func (o *DcimLocationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimLocationsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsListDefault creates a DcimLocationsListDefault with default headers values
func NewDcimLocationsListDefault(code int) *DcimLocationsListDefault {
	return &DcimLocationsListDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsListDefault describes a response with status code -1, with default header values.

DcimLocationsListDefault dcim locations list default
*/
type DcimLocationsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim locations list default response
func (o *DcimLocationsListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim locations list default response has a 2xx status code
func (o *DcimLocationsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations list default response has a 3xx status code
func (o *DcimLocationsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations list default response has a 4xx status code
func (o *DcimLocationsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations list default response has a 5xx status code
func (o *DcimLocationsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations list default response a status code equal to that given
func (o *DcimLocationsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimLocationsListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/locations/][%d] dcim_locations_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/locations/][%d] dcim_locations_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimLocationsListOKBody dcim locations list o k body
swagger:model DcimLocationsListOKBody
*/
type DcimLocationsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Location `json:"results"`
}

// Validate validates this dcim locations list o k body
func (o *DcimLocationsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimLocationsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimLocationsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimLocationsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimLocationsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimLocationsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimLocationsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimLocationsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimLocationsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimLocationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimLocationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim locations list o k body based on the context it is used
func (o *DcimLocationsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimLocationsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimLocationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimLocationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimLocationsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimLocationsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimLocationsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
