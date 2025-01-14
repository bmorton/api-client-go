// Code generated by go-swagger; DO NOT EDIT.

package attachments

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

// NewPostV1AttachmentsParams creates a new PostV1AttachmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1AttachmentsParams() *PostV1AttachmentsParams {
	return &PostV1AttachmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1AttachmentsParamsWithTimeout creates a new PostV1AttachmentsParams object
// with the ability to set a timeout on a request.
func NewPostV1AttachmentsParamsWithTimeout(timeout time.Duration) *PostV1AttachmentsParams {
	return &PostV1AttachmentsParams{
		timeout: timeout,
	}
}

// NewPostV1AttachmentsParamsWithContext creates a new PostV1AttachmentsParams object
// with the ability to set a context for a request.
func NewPostV1AttachmentsParamsWithContext(ctx context.Context) *PostV1AttachmentsParams {
	return &PostV1AttachmentsParams{
		Context: ctx,
	}
}

// NewPostV1AttachmentsParamsWithHTTPClient creates a new PostV1AttachmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1AttachmentsParamsWithHTTPClient(client *http.Client) *PostV1AttachmentsParams {
	return &PostV1AttachmentsParams{
		HTTPClient: client,
	}
}

/* PostV1AttachmentsParams contains all the parameters to send to the API endpoint
   for the post v1 attachments operation.

   Typically these are written to a http.Request.
*/
type PostV1AttachmentsParams struct {

	// Description.
	Description *string

	// File.
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 attachments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AttachmentsParams) WithDefaults() *PostV1AttachmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 attachments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AttachmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 attachments params
func (o *PostV1AttachmentsParams) WithTimeout(timeout time.Duration) *PostV1AttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 attachments params
func (o *PostV1AttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 attachments params
func (o *PostV1AttachmentsParams) WithContext(ctx context.Context) *PostV1AttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 attachments params
func (o *PostV1AttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 attachments params
func (o *PostV1AttachmentsParams) WithHTTPClient(client *http.Client) *PostV1AttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 attachments params
func (o *PostV1AttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the post v1 attachments params
func (o *PostV1AttachmentsParams) WithDescription(description *string) *PostV1AttachmentsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post v1 attachments params
func (o *PostV1AttachmentsParams) SetDescription(description *string) {
	o.Description = description
}

// WithFile adds the file to the post v1 attachments params
func (o *PostV1AttachmentsParams) WithFile(file runtime.NamedReadCloser) *PostV1AttachmentsParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the post v1 attachments params
func (o *PostV1AttachmentsParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1AttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
