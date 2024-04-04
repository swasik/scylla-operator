// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewColumnFamilyMetricsReadLatencyByNameGetParams creates a new ColumnFamilyMetricsReadLatencyByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsReadLatencyByNameGetParams() *ColumnFamilyMetricsReadLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsReadLatencyByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsReadLatencyByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsReadLatencyByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsReadLatencyByNameGetParamsWithContext creates a new ColumnFamilyMetricsReadLatencyByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsReadLatencyByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsReadLatencyByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsReadLatencyByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsReadLatencyByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyByNameGetParams {
	var ()
	return &ColumnFamilyMetricsReadLatencyByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsReadLatencyByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics read latency by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsReadLatencyByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsReadLatencyByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsReadLatencyByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsReadLatencyByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) WithName(name string) *ColumnFamilyMetricsReadLatencyByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics read latency by name get params
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsReadLatencyByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
