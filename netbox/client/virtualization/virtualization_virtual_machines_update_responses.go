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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// VirtualizationVirtualMachinesUpdateReader is a Reader for the VirtualizationVirtualMachinesUpdate structure.
type VirtualizationVirtualMachinesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesUpdateOK creates a VirtualizationVirtualMachinesUpdateOK with default headers values
func NewVirtualizationVirtualMachinesUpdateOK() *VirtualizationVirtualMachinesUpdateOK {
	return &VirtualizationVirtualMachinesUpdateOK{}
}

/*VirtualizationVirtualMachinesUpdateOK handles this case with default header values.

VirtualizationVirtualMachinesUpdateOK virtualization virtual machines update o k
*/
type VirtualizationVirtualMachinesUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesUpdateDefault creates a VirtualizationVirtualMachinesUpdateDefault with default headers values
func NewVirtualizationVirtualMachinesUpdateDefault(code int) *VirtualizationVirtualMachinesUpdateDefault {
	return &VirtualizationVirtualMachinesUpdateDefault{
		_statusCode: code,
	}
}

/*VirtualizationVirtualMachinesUpdateDefault handles this case with default header values.

VirtualizationVirtualMachinesUpdateDefault virtualization virtual machines update default
*/
type VirtualizationVirtualMachinesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines update default response
func (o *VirtualizationVirtualMachinesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualMachinesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
