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

// NewFailureDetectorEndpointsGetParams creates a new FailureDetectorEndpointsGetParams object
// with the default values initialized.
func NewFailureDetectorEndpointsGetParams() *FailureDetectorEndpointsGetParams {

	return &FailureDetectorEndpointsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFailureDetectorEndpointsGetParamsWithTimeout creates a new FailureDetectorEndpointsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFailureDetectorEndpointsGetParamsWithTimeout(timeout time.Duration) *FailureDetectorEndpointsGetParams {

	return &FailureDetectorEndpointsGetParams{

		timeout: timeout,
	}
}

// NewFailureDetectorEndpointsGetParamsWithContext creates a new FailureDetectorEndpointsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFailureDetectorEndpointsGetParamsWithContext(ctx context.Context) *FailureDetectorEndpointsGetParams {

	return &FailureDetectorEndpointsGetParams{

		Context: ctx,
	}
}

// NewFailureDetectorEndpointsGetParamsWithHTTPClient creates a new FailureDetectorEndpointsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFailureDetectorEndpointsGetParamsWithHTTPClient(client *http.Client) *FailureDetectorEndpointsGetParams {

	return &FailureDetectorEndpointsGetParams{
		HTTPClient: client,
	}
}

/*
FailureDetectorEndpointsGetParams contains all the parameters to send to the API endpoint
for the failure detector endpoints get operation typically these are written to a http.Request
*/
type FailureDetectorEndpointsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the failure detector endpoints get params
func (o *FailureDetectorEndpointsGetParams) WithTimeout(timeout time.Duration) *FailureDetectorEndpointsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the failure detector endpoints get params
func (o *FailureDetectorEndpointsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the failure detector endpoints get params
func (o *FailureDetectorEndpointsGetParams) WithContext(ctx context.Context) *FailureDetectorEndpointsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the failure detector endpoints get params
func (o *FailureDetectorEndpointsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the failure detector endpoints get params
func (o *FailureDetectorEndpointsGetParams) WithHTTPClient(client *http.Client) *FailureDetectorEndpointsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the failure detector endpoints get params
func (o *FailureDetectorEndpointsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FailureDetectorEndpointsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
