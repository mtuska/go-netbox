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
)

// DcimCablesDeleteReader is a Reader for the DcimCablesDelete structure.
type DcimCablesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimCablesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesDeleteNoContent creates a DcimCablesDeleteNoContent with default headers values
func NewDcimCablesDeleteNoContent() *DcimCablesDeleteNoContent {
	return &DcimCablesDeleteNoContent{}
}

/*
DcimCablesDeleteNoContent describes a response with status code 204, with default header values.

DcimCablesDeleteNoContent dcim cables delete no content
*/
type DcimCablesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim cables delete no content response has a 2xx status code
func (o *DcimCablesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables delete no content response has a 3xx status code
func (o *DcimCablesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables delete no content response has a 4xx status code
func (o *DcimCablesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables delete no content response has a 5xx status code
func (o *DcimCablesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables delete no content response a status code equal to that given
func (o *DcimCablesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimCablesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cables/{id}/][%d] dcimCablesDeleteNoContent ", 204)
}

func (o *DcimCablesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/cables/{id}/][%d] dcimCablesDeleteNoContent ", 204)
}

func (o *DcimCablesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimCablesDeleteDefault creates a DcimCablesDeleteDefault with default headers values
func NewDcimCablesDeleteDefault(code int) *DcimCablesDeleteDefault {
	return &DcimCablesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimCablesDeleteDefault describes a response with status code -1, with default header values.

DcimCablesDeleteDefault dcim cables delete default
*/
type DcimCablesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables delete default response
func (o *DcimCablesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cables delete default response has a 2xx status code
func (o *DcimCablesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cables delete default response has a 3xx status code
func (o *DcimCablesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cables delete default response has a 4xx status code
func (o *DcimCablesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cables delete default response has a 5xx status code
func (o *DcimCablesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cables delete default response a status code equal to that given
func (o *DcimCablesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCablesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cables/{id}/][%d] dcim_cables_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/cables/{id}/][%d] dcim_cables_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
