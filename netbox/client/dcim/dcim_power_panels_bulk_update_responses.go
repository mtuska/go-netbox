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

// DcimPowerPanelsBulkUpdateReader is a Reader for the DcimPowerPanelsBulkUpdate structure.
type DcimPowerPanelsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPanelsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsBulkUpdateOK creates a DcimPowerPanelsBulkUpdateOK with default headers values
func NewDcimPowerPanelsBulkUpdateOK() *DcimPowerPanelsBulkUpdateOK {
	return &DcimPowerPanelsBulkUpdateOK{}
}

/*
DcimPowerPanelsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerPanelsBulkUpdateOK dcim power panels bulk update o k
*/
type DcimPowerPanelsBulkUpdateOK struct {
	Payload *models.PowerPanel
}

// IsSuccess returns true when this dcim power panels bulk update o k response has a 2xx status code
func (o *DcimPowerPanelsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels bulk update o k response has a 3xx status code
func (o *DcimPowerPanelsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels bulk update o k response has a 4xx status code
func (o *DcimPowerPanelsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels bulk update o k response has a 5xx status code
func (o *DcimPowerPanelsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels bulk update o k response a status code equal to that given
func (o *DcimPowerPanelsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerPanelsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-panels/][%d] dcimPowerPanelsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-panels/][%d] dcimPowerPanelsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsBulkUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsBulkUpdateDefault creates a DcimPowerPanelsBulkUpdateDefault with default headers values
func NewDcimPowerPanelsBulkUpdateDefault(code int) *DcimPowerPanelsBulkUpdateDefault {
	return &DcimPowerPanelsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPanelsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPanelsBulkUpdateDefault dcim power panels bulk update default
*/
type DcimPowerPanelsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power panels bulk update default response
func (o *DcimPowerPanelsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power panels bulk update default response has a 2xx status code
func (o *DcimPowerPanelsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power panels bulk update default response has a 3xx status code
func (o *DcimPowerPanelsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power panels bulk update default response has a 4xx status code
func (o *DcimPowerPanelsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power panels bulk update default response has a 5xx status code
func (o *DcimPowerPanelsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power panels bulk update default response a status code equal to that given
func (o *DcimPowerPanelsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPanelsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-panels/][%d] dcim_power-panels_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-panels/][%d] dcim_power-panels_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
