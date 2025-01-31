// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewVerifyLdapUserSearchParams creates a new VerifyLdapUserSearchParams object
// with the default values initialized.
func NewVerifyLdapUserSearchParams() *VerifyLdapUserSearchParams {

	return &VerifyLdapUserSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVerifyLdapUserSearchParamsWithTimeout creates a new VerifyLdapUserSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVerifyLdapUserSearchParamsWithTimeout(timeout time.Duration) *VerifyLdapUserSearchParams {

	return &VerifyLdapUserSearchParams{

		timeout: timeout,
	}
}

// NewVerifyLdapUserSearchParamsWithContext creates a new VerifyLdapUserSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewVerifyLdapUserSearchParamsWithContext(ctx context.Context) *VerifyLdapUserSearchParams {

	return &VerifyLdapUserSearchParams{

		Context: ctx,
	}
}

// NewVerifyLdapUserSearchParamsWithHTTPClient creates a new VerifyLdapUserSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVerifyLdapUserSearchParamsWithHTTPClient(client *http.Client) *VerifyLdapUserSearchParams {

	return &VerifyLdapUserSearchParams{
		HTTPClient: client,
	}
}

/*
VerifyLdapUserSearchParams contains all the parameters to send to the API endpoint
for the verify ldap user search operation typically these are written to a http.Request
*/
type VerifyLdapUserSearchParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the verify ldap user search params
func (o *VerifyLdapUserSearchParams) WithTimeout(timeout time.Duration) *VerifyLdapUserSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the verify ldap user search params
func (o *VerifyLdapUserSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the verify ldap user search params
func (o *VerifyLdapUserSearchParams) WithContext(ctx context.Context) *VerifyLdapUserSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the verify ldap user search params
func (o *VerifyLdapUserSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the verify ldap user search params
func (o *VerifyLdapUserSearchParams) WithHTTPClient(client *http.Client) *VerifyLdapUserSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the verify ldap user search params
func (o *VerifyLdapUserSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VerifyLdapUserSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
