// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new webhooks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for webhooks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteV1WebhooksWebhookID deletes a specific webhook

Delete a specific webhook
*/
func (a *Client) DeleteV1WebhooksWebhookID(params *DeleteV1WebhooksWebhookIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteV1WebhooksWebhookIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1WebhooksWebhookIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteV1WebhooksWebhookId",
		Method:             "DELETE",
		PathPattern:        "/v1/webhooks/{webhook_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteV1WebhooksWebhookIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteV1WebhooksWebhookIDNoContent), nil

}

/*
GetV1Webhooks lists webhooks

Lists webhooks
*/
func (a *Client) GetV1Webhooks(params *GetV1WebhooksParams, authInfo runtime.ClientAuthInfoWriter) (*GetV1WebhooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WebhooksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getV1Webhooks",
		Method:             "GET",
		PathPattern:        "/v1/webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1WebhooksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetV1WebhooksOK), nil

}

/*
GetV1WebhooksWebhookID retrieves a specific webhook

Retrieve a specific webhook
*/
func (a *Client) GetV1WebhooksWebhookID(params *GetV1WebhooksWebhookIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetV1WebhooksWebhookIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1WebhooksWebhookIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getV1WebhooksWebhookId",
		Method:             "GET",
		PathPattern:        "/v1/webhooks/{webhook_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1WebhooksWebhookIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetV1WebhooksWebhookIDOK), nil

}

/*
PatchV1WebhooksWebhookID updates a specific webhook

Update a specific webhook
*/
func (a *Client) PatchV1WebhooksWebhookID(params *PatchV1WebhooksWebhookIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchV1WebhooksWebhookIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1WebhooksWebhookIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchV1WebhooksWebhookId",
		Method:             "PATCH",
		PathPattern:        "/v1/webhooks/{webhook_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchV1WebhooksWebhookIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchV1WebhooksWebhookIDOK), nil

}

/*
PostV1Webhooks creates webhook

Create a new webhook
*/
func (a *Client) PostV1Webhooks(params *PostV1WebhooksParams, authInfo runtime.ClientAuthInfoWriter) (*PostV1WebhooksCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1WebhooksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postV1Webhooks",
		Method:             "POST",
		PathPattern:        "/v1/webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostV1WebhooksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostV1WebhooksCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}