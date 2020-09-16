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

// DcimConsolePortTemplatesCreateReader is a Reader for the DcimConsolePortTemplatesCreate structure.
type DcimConsolePortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsolePortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesCreateCreated creates a DcimConsolePortTemplatesCreateCreated with default headers values
func NewDcimConsolePortTemplatesCreateCreated() *DcimConsolePortTemplatesCreateCreated {
	return &DcimConsolePortTemplatesCreateCreated{}
}

/*DcimConsolePortTemplatesCreateCreated handles this case with default header values.

DcimConsolePortTemplatesCreateCreated dcim console port templates create created
*/
type DcimConsolePortTemplatesCreateCreated struct {
	Payload *models.ConsolePortTemplate
}

func (o *DcimConsolePortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-port-templates/][%d] dcimConsolePortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimConsolePortTemplatesCreateCreated) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesCreateDefault creates a DcimConsolePortTemplatesCreateDefault with default headers values
func NewDcimConsolePortTemplatesCreateDefault(code int) *DcimConsolePortTemplatesCreateDefault {
	return &DcimConsolePortTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*DcimConsolePortTemplatesCreateDefault handles this case with default header values.

DcimConsolePortTemplatesCreateDefault dcim console port templates create default
*/
type DcimConsolePortTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates create default response
func (o *DcimConsolePortTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/console-port-templates/][%d] dcim_console-port-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
