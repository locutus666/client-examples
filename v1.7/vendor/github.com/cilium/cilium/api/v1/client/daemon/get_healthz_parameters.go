// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetHealthzParams creates a new GetHealthzParams object
// with the default values initialized.
func NewGetHealthzParams() *GetHealthzParams {
	var ()
	return &GetHealthzParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthzParamsWithTimeout creates a new GetHealthzParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHealthzParamsWithTimeout(timeout time.Duration) *GetHealthzParams {
	var ()
	return &GetHealthzParams{

		timeout: timeout,
	}
}

// NewGetHealthzParamsWithContext creates a new GetHealthzParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHealthzParamsWithContext(ctx context.Context) *GetHealthzParams {
	var ()
	return &GetHealthzParams{

		Context: ctx,
	}
}

// NewGetHealthzParamsWithHTTPClient creates a new GetHealthzParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHealthzParamsWithHTTPClient(client *http.Client) *GetHealthzParams {
	var ()
	return &GetHealthzParams{
		HTTPClient: client,
	}
}

/*GetHealthzParams contains all the parameters to send to the API endpoint
for the get healthz operation typically these are written to a http.Request
*/
type GetHealthzParams struct {

	/*Brief
	  Brief will return a brief representation of the Cilium status.


	*/
	Brief *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get healthz params
func (o *GetHealthzParams) WithTimeout(timeout time.Duration) *GetHealthzParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get healthz params
func (o *GetHealthzParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get healthz params
func (o *GetHealthzParams) WithContext(ctx context.Context) *GetHealthzParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get healthz params
func (o *GetHealthzParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get healthz params
func (o *GetHealthzParams) WithHTTPClient(client *http.Client) *GetHealthzParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get healthz params
func (o *GetHealthzParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrief adds the brief to the get healthz params
func (o *GetHealthzParams) WithBrief(brief *bool) *GetHealthzParams {
	o.SetBrief(brief)
	return o
}

// SetBrief adds the brief to the get healthz params
func (o *GetHealthzParams) SetBrief(brief *bool) {
	o.Brief = brief
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthzParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Brief != nil {

		// header param brief
		if err := r.SetHeaderParam("brief", swag.FormatBool(*o.Brief)); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
