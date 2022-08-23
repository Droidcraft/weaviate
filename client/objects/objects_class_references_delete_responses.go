//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ObjectsClassReferencesDeleteReader is a Reader for the ObjectsClassReferencesDelete structure.
type ObjectsClassReferencesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsClassReferencesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewObjectsClassReferencesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewObjectsClassReferencesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsClassReferencesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewObjectsClassReferencesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewObjectsClassReferencesDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsClassReferencesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewObjectsClassReferencesDeleteNoContent creates a ObjectsClassReferencesDeleteNoContent with default headers values
func NewObjectsClassReferencesDeleteNoContent() *ObjectsClassReferencesDeleteNoContent {
	return &ObjectsClassReferencesDeleteNoContent{}
}

/*
ObjectsClassReferencesDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type ObjectsClassReferencesDeleteNoContent struct {
}

func (o *ObjectsClassReferencesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}/references/{propertyName}][%d] objectsClassReferencesDeleteNoContent ", 204)
}

func (o *ObjectsClassReferencesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassReferencesDeleteUnauthorized creates a ObjectsClassReferencesDeleteUnauthorized with default headers values
func NewObjectsClassReferencesDeleteUnauthorized() *ObjectsClassReferencesDeleteUnauthorized {
	return &ObjectsClassReferencesDeleteUnauthorized{}
}

/*
ObjectsClassReferencesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsClassReferencesDeleteUnauthorized struct {
}

func (o *ObjectsClassReferencesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}/references/{propertyName}][%d] objectsClassReferencesDeleteUnauthorized ", 401)
}

func (o *ObjectsClassReferencesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsClassReferencesDeleteForbidden creates a ObjectsClassReferencesDeleteForbidden with default headers values
func NewObjectsClassReferencesDeleteForbidden() *ObjectsClassReferencesDeleteForbidden {
	return &ObjectsClassReferencesDeleteForbidden{}
}

/*
ObjectsClassReferencesDeleteForbidden handles this case with default header values.

Forbidden
*/
type ObjectsClassReferencesDeleteForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassReferencesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}/references/{propertyName}][%d] objectsClassReferencesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsClassReferencesDeleteForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassReferencesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassReferencesDeleteNotFound creates a ObjectsClassReferencesDeleteNotFound with default headers values
func NewObjectsClassReferencesDeleteNotFound() *ObjectsClassReferencesDeleteNotFound {
	return &ObjectsClassReferencesDeleteNotFound{}
}

/*
ObjectsClassReferencesDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type ObjectsClassReferencesDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassReferencesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}/references/{propertyName}][%d] objectsClassReferencesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ObjectsClassReferencesDeleteNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassReferencesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassReferencesDeleteUnprocessableEntity creates a ObjectsClassReferencesDeleteUnprocessableEntity with default headers values
func NewObjectsClassReferencesDeleteUnprocessableEntity() *ObjectsClassReferencesDeleteUnprocessableEntity {
	return &ObjectsClassReferencesDeleteUnprocessableEntity{}
}

/*
ObjectsClassReferencesDeleteUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type ObjectsClassReferencesDeleteUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassReferencesDeleteUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}/references/{propertyName}][%d] objectsClassReferencesDeleteUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsClassReferencesDeleteUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassReferencesDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsClassReferencesDeleteInternalServerError creates a ObjectsClassReferencesDeleteInternalServerError with default headers values
func NewObjectsClassReferencesDeleteInternalServerError() *ObjectsClassReferencesDeleteInternalServerError {
	return &ObjectsClassReferencesDeleteInternalServerError{}
}

/*
ObjectsClassReferencesDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsClassReferencesDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsClassReferencesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /objects/{className}/{id}/references/{propertyName}][%d] objectsClassReferencesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsClassReferencesDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsClassReferencesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}