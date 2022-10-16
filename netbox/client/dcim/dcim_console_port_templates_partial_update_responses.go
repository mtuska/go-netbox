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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimConsolePortTemplatesPartialUpdateReader is a Reader for the DcimConsolePortTemplatesPartialUpdate structure.
type DcimConsolePortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesPartialUpdateOK creates a DcimConsolePortTemplatesPartialUpdateOK with default headers values
func NewDcimConsolePortTemplatesPartialUpdateOK() *DcimConsolePortTemplatesPartialUpdateOK {
	return &DcimConsolePortTemplatesPartialUpdateOK{}
}

/*
DcimConsolePortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesPartialUpdateOK dcim console port templates partial update o k
*/
type DcimConsolePortTemplatesPartialUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

// IsSuccess returns true when this dcim console port templates partial update o k response has a 2xx status code
func (o *DcimConsolePortTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates partial update o k response has a 3xx status code
func (o *DcimConsolePortTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates partial update o k response has a 4xx status code
func (o *DcimConsolePortTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates partial update o k response has a 5xx status code
func (o *DcimConsolePortTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates partial update o k response a status code equal to that given
func (o *DcimConsolePortTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsolePortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesPartialUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesPartialUpdateDefault creates a DcimConsolePortTemplatesPartialUpdateDefault with default headers values
func NewDcimConsolePortTemplatesPartialUpdateDefault(code int) *DcimConsolePortTemplatesPartialUpdateDefault {
	return &DcimConsolePortTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimConsolePortTemplatesPartialUpdateDefault dcim console port templates partial update default
*/
type DcimConsolePortTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates partial update default response
func (o *DcimConsolePortTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console port templates partial update default response has a 2xx status code
func (o *DcimConsolePortTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console port templates partial update default response has a 3xx status code
func (o *DcimConsolePortTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console port templates partial update default response has a 4xx status code
func (o *DcimConsolePortTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console port templates partial update default response has a 5xx status code
func (o *DcimConsolePortTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console port templates partial update default response a status code equal to that given
func (o *DcimConsolePortTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
