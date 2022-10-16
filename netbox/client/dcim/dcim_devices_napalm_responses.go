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

// DcimDevicesNapalmReader is a Reader for the DcimDevicesNapalm structure.
type DcimDevicesNapalmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesNapalmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesNapalmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesNapalmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesNapalmOK creates a DcimDevicesNapalmOK with default headers values
func NewDcimDevicesNapalmOK() *DcimDevicesNapalmOK {
	return &DcimDevicesNapalmOK{}
}

/*
DcimDevicesNapalmOK describes a response with status code 200, with default header values.

DcimDevicesNapalmOK dcim devices napalm o k
*/
type DcimDevicesNapalmOK struct {
	Payload *models.DeviceNAPALM
}

// IsSuccess returns true when this dcim devices napalm o k response has a 2xx status code
func (o *DcimDevicesNapalmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices napalm o k response has a 3xx status code
func (o *DcimDevicesNapalmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices napalm o k response has a 4xx status code
func (o *DcimDevicesNapalmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices napalm o k response has a 5xx status code
func (o *DcimDevicesNapalmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices napalm o k response a status code equal to that given
func (o *DcimDevicesNapalmOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDevicesNapalmOK) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/napalm/][%d] dcimDevicesNapalmOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesNapalmOK) String() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/napalm/][%d] dcimDevicesNapalmOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesNapalmOK) GetPayload() *models.DeviceNAPALM {
	return o.Payload
}

func (o *DcimDevicesNapalmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceNAPALM)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesNapalmDefault creates a DcimDevicesNapalmDefault with default headers values
func NewDcimDevicesNapalmDefault(code int) *DcimDevicesNapalmDefault {
	return &DcimDevicesNapalmDefault{
		_statusCode: code,
	}
}

/*
DcimDevicesNapalmDefault describes a response with status code -1, with default header values.

DcimDevicesNapalmDefault dcim devices napalm default
*/
type DcimDevicesNapalmDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices napalm default response
func (o *DcimDevicesNapalmDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim devices napalm default response has a 2xx status code
func (o *DcimDevicesNapalmDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim devices napalm default response has a 3xx status code
func (o *DcimDevicesNapalmDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim devices napalm default response has a 4xx status code
func (o *DcimDevicesNapalmDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim devices napalm default response has a 5xx status code
func (o *DcimDevicesNapalmDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim devices napalm default response a status code equal to that given
func (o *DcimDevicesNapalmDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDevicesNapalmDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/napalm/][%d] dcim_devices_napalm default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesNapalmDefault) String() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/napalm/][%d] dcim_devices_napalm default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesNapalmDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesNapalmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
