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

// NewStorageServiceTombstoneFailureThresholdGetParams creates a new StorageServiceTombstoneFailureThresholdGetParams object
// with the default values initialized.
func NewStorageServiceTombstoneFailureThresholdGetParams() *StorageServiceTombstoneFailureThresholdGetParams {

	return &StorageServiceTombstoneFailureThresholdGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceTombstoneFailureThresholdGetParamsWithTimeout creates a new StorageServiceTombstoneFailureThresholdGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceTombstoneFailureThresholdGetParamsWithTimeout(timeout time.Duration) *StorageServiceTombstoneFailureThresholdGetParams {

	return &StorageServiceTombstoneFailureThresholdGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceTombstoneFailureThresholdGetParamsWithContext creates a new StorageServiceTombstoneFailureThresholdGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceTombstoneFailureThresholdGetParamsWithContext(ctx context.Context) *StorageServiceTombstoneFailureThresholdGetParams {

	return &StorageServiceTombstoneFailureThresholdGetParams{

		Context: ctx,
	}
}

// NewStorageServiceTombstoneFailureThresholdGetParamsWithHTTPClient creates a new StorageServiceTombstoneFailureThresholdGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceTombstoneFailureThresholdGetParamsWithHTTPClient(client *http.Client) *StorageServiceTombstoneFailureThresholdGetParams {

	return &StorageServiceTombstoneFailureThresholdGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceTombstoneFailureThresholdGetParams contains all the parameters to send to the API endpoint
for the storage service tombstone failure threshold get operation typically these are written to a http.Request
*/
type StorageServiceTombstoneFailureThresholdGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service tombstone failure threshold get params
func (o *StorageServiceTombstoneFailureThresholdGetParams) WithTimeout(timeout time.Duration) *StorageServiceTombstoneFailureThresholdGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service tombstone failure threshold get params
func (o *StorageServiceTombstoneFailureThresholdGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service tombstone failure threshold get params
func (o *StorageServiceTombstoneFailureThresholdGetParams) WithContext(ctx context.Context) *StorageServiceTombstoneFailureThresholdGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service tombstone failure threshold get params
func (o *StorageServiceTombstoneFailureThresholdGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service tombstone failure threshold get params
func (o *StorageServiceTombstoneFailureThresholdGetParams) WithHTTPClient(client *http.Client) *StorageServiceTombstoneFailureThresholdGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service tombstone failure threshold get params
func (o *StorageServiceTombstoneFailureThresholdGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceTombstoneFailureThresholdGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
