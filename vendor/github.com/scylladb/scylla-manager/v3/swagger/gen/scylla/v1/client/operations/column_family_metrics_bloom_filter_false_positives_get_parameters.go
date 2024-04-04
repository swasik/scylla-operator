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

// NewColumnFamilyMetricsBloomFilterFalsePositivesGetParams creates a new ColumnFamilyMetricsBloomFilterFalsePositivesGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsBloomFilterFalsePositivesGetParams() *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams {

	return &ColumnFamilyMetricsBloomFilterFalsePositivesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsBloomFilterFalsePositivesGetParamsWithTimeout creates a new ColumnFamilyMetricsBloomFilterFalsePositivesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsBloomFilterFalsePositivesGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams {

	return &ColumnFamilyMetricsBloomFilterFalsePositivesGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsBloomFilterFalsePositivesGetParamsWithContext creates a new ColumnFamilyMetricsBloomFilterFalsePositivesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsBloomFilterFalsePositivesGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams {

	return &ColumnFamilyMetricsBloomFilterFalsePositivesGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsBloomFilterFalsePositivesGetParamsWithHTTPClient creates a new ColumnFamilyMetricsBloomFilterFalsePositivesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsBloomFilterFalsePositivesGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams {

	return &ColumnFamilyMetricsBloomFilterFalsePositivesGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsBloomFilterFalsePositivesGetParams contains all the parameters to send to the API endpoint
for the column family metrics bloom filter false positives get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsBloomFilterFalsePositivesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics bloom filter false positives get params
func (o *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics bloom filter false positives get params
func (o *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics bloom filter false positives get params
func (o *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics bloom filter false positives get params
func (o *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics bloom filter false positives get params
func (o *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics bloom filter false positives get params
func (o *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsBloomFilterFalsePositivesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
