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
)

// VirtualizationClustersDeleteReader is a Reader for the VirtualizationClustersDelete structure.
type VirtualizationClustersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationClustersDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClustersDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersDeleteNoContent creates a VirtualizationClustersDeleteNoContent with default headers values
func NewVirtualizationClustersDeleteNoContent() *VirtualizationClustersDeleteNoContent {
	return &VirtualizationClustersDeleteNoContent{}
}

/*
VirtualizationClustersDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationClustersDeleteNoContent virtualization clusters delete no content
*/
type VirtualizationClustersDeleteNoContent struct {
}

// IsSuccess returns true when this virtualization clusters delete no content response has a 2xx status code
func (o *VirtualizationClustersDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters delete no content response has a 3xx status code
func (o *VirtualizationClustersDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters delete no content response has a 4xx status code
func (o *VirtualizationClustersDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters delete no content response has a 5xx status code
func (o *VirtualizationClustersDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters delete no content response a status code equal to that given
func (o *VirtualizationClustersDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *VirtualizationClustersDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualizationClustersDeleteNoContent ", 204)
}

func (o *VirtualizationClustersDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualizationClustersDeleteNoContent ", 204)
}

func (o *VirtualizationClustersDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationClustersDeleteDefault creates a VirtualizationClustersDeleteDefault with default headers values
func NewVirtualizationClustersDeleteDefault(code int) *VirtualizationClustersDeleteDefault {
	return &VirtualizationClustersDeleteDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClustersDeleteDefault describes a response with status code -1, with default header values.

VirtualizationClustersDeleteDefault virtualization clusters delete default
*/
type VirtualizationClustersDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization clusters delete default response
func (o *VirtualizationClustersDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization clusters delete default response has a 2xx status code
func (o *VirtualizationClustersDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization clusters delete default response has a 3xx status code
func (o *VirtualizationClustersDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization clusters delete default response has a 4xx status code
func (o *VirtualizationClustersDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization clusters delete default response has a 5xx status code
func (o *VirtualizationClustersDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization clusters delete default response a status code equal to that given
func (o *VirtualizationClustersDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClustersDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualization_clusters_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /virtualization/clusters/{id}/][%d] virtualization_clusters_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
