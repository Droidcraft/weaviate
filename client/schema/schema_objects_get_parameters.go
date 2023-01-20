//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewSchemaObjectsGetParams creates a new SchemaObjectsGetParams object
// with the default values initialized.
func NewSchemaObjectsGetParams() *SchemaObjectsGetParams {
	var ()
	return &SchemaObjectsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaObjectsGetParamsWithTimeout creates a new SchemaObjectsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaObjectsGetParamsWithTimeout(timeout time.Duration) *SchemaObjectsGetParams {
	var ()
	return &SchemaObjectsGetParams{

		timeout: timeout,
	}
}

// NewSchemaObjectsGetParamsWithContext creates a new SchemaObjectsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaObjectsGetParamsWithContext(ctx context.Context) *SchemaObjectsGetParams {
	var ()
	return &SchemaObjectsGetParams{

		Context: ctx,
	}
}

// NewSchemaObjectsGetParamsWithHTTPClient creates a new SchemaObjectsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaObjectsGetParamsWithHTTPClient(client *http.Client) *SchemaObjectsGetParams {
	var ()
	return &SchemaObjectsGetParams{
		HTTPClient: client,
	}
}

/*
SchemaObjectsGetParams contains all the parameters to send to the API endpoint
for the schema objects get operation typically these are written to a http.Request
*/
type SchemaObjectsGetParams struct {

	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema objects get params
func (o *SchemaObjectsGetParams) WithTimeout(timeout time.Duration) *SchemaObjectsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema objects get params
func (o *SchemaObjectsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema objects get params
func (o *SchemaObjectsGetParams) WithContext(ctx context.Context) *SchemaObjectsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema objects get params
func (o *SchemaObjectsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema objects get params
func (o *SchemaObjectsGetParams) WithHTTPClient(client *http.Client) *SchemaObjectsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema objects get params
func (o *SchemaObjectsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the schema objects get params
func (o *SchemaObjectsGetParams) WithClassName(className string) *SchemaObjectsGetParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema objects get params
func (o *SchemaObjectsGetParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaObjectsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
