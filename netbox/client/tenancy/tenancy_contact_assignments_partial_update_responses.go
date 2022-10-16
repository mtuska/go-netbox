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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// TenancyContactAssignmentsPartialUpdateReader is a Reader for the TenancyContactAssignmentsPartialUpdate structure.
type TenancyContactAssignmentsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactAssignmentsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactAssignmentsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactAssignmentsPartialUpdateOK creates a TenancyContactAssignmentsPartialUpdateOK with default headers values
func NewTenancyContactAssignmentsPartialUpdateOK() *TenancyContactAssignmentsPartialUpdateOK {
	return &TenancyContactAssignmentsPartialUpdateOK{}
}

/*
TenancyContactAssignmentsPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactAssignmentsPartialUpdateOK tenancy contact assignments partial update o k
*/
type TenancyContactAssignmentsPartialUpdateOK struct {
	Payload *models.ContactAssignment
}

// IsSuccess returns true when this tenancy contact assignments partial update o k response has a 2xx status code
func (o *TenancyContactAssignmentsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact assignments partial update o k response has a 3xx status code
func (o *TenancyContactAssignmentsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact assignments partial update o k response has a 4xx status code
func (o *TenancyContactAssignmentsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact assignments partial update o k response has a 5xx status code
func (o *TenancyContactAssignmentsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact assignments partial update o k response a status code equal to that given
func (o *TenancyContactAssignmentsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactAssignmentsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactAssignmentsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-assignments/{id}/][%d] tenancyContactAssignmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactAssignmentsPartialUpdateOK) GetPayload() *models.ContactAssignment {
	return o.Payload
}

func (o *TenancyContactAssignmentsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactAssignmentsPartialUpdateDefault creates a TenancyContactAssignmentsPartialUpdateDefault with default headers values
func NewTenancyContactAssignmentsPartialUpdateDefault(code int) *TenancyContactAssignmentsPartialUpdateDefault {
	return &TenancyContactAssignmentsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactAssignmentsPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyContactAssignmentsPartialUpdateDefault tenancy contact assignments partial update default
*/
type TenancyContactAssignmentsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact assignments partial update default response
func (o *TenancyContactAssignmentsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact assignments partial update default response has a 2xx status code
func (o *TenancyContactAssignmentsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact assignments partial update default response has a 3xx status code
func (o *TenancyContactAssignmentsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact assignments partial update default response has a 4xx status code
func (o *TenancyContactAssignmentsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact assignments partial update default response has a 5xx status code
func (o *TenancyContactAssignmentsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact assignments partial update default response a status code equal to that given
func (o *TenancyContactAssignmentsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactAssignmentsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactAssignmentsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-assignments/{id}/][%d] tenancy_contact-assignments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactAssignmentsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactAssignmentsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
