// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ticketing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ticketing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1TicketingTicketsTicketID(params *DeleteV1TicketingTicketsTicketIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1TicketingTicketsTicketIDNoContent, error)

	GetV1TicketingProjects(params *GetV1TicketingProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TicketingProjectsOK, error)

	GetV1TicketingTickets(params *GetV1TicketingTicketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TicketingTicketsOK, error)

	GetV1TicketingTicketsTicketID(params *GetV1TicketingTicketsTicketIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TicketingTicketsTicketIDOK, error)

	PatchV1TicketingTicketsTicketID(params *PatchV1TicketingTicketsTicketIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1TicketingTicketsTicketIDOK, error)

	PostV1TicketingTickets(params *PostV1TicketingTicketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1TicketingTicketsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteV1TicketingTicketsTicketID delete v1 ticketing tickets ticket Id API
*/
func (a *Client) DeleteV1TicketingTicketsTicketID(params *DeleteV1TicketingTicketsTicketIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1TicketingTicketsTicketIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1TicketingTicketsTicketIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1TicketingTicketsTicketId",
		Method:             "DELETE",
		PathPattern:        "/v1/ticketing/tickets/{ticket_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1TicketingTicketsTicketIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1TicketingTicketsTicketIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1TicketingTicketsTicketId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1TicketingProjects lists all ticketing projects

  List all ticketing projects available to the organization
*/
func (a *Client) GetV1TicketingProjects(params *GetV1TicketingProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TicketingProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TicketingProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1TicketingProjects",
		Method:             "GET",
		PathPattern:        "/v1/ticketing/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TicketingProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TicketingProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1TicketingProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1TicketingTickets lists all functionalities

  List all of the functionalities that have been added to the organiation
*/
func (a *Client) GetV1TicketingTickets(params *GetV1TicketingTicketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TicketingTicketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TicketingTicketsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1TicketingTickets",
		Method:             "GET",
		PathPattern:        "/v1/ticketing/tickets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TicketingTicketsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TicketingTicketsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1TicketingTickets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1TicketingTicketsTicketID retrieves a single ticket

  Retrieves a single ticket by ID
*/
func (a *Client) GetV1TicketingTicketsTicketID(params *GetV1TicketingTicketsTicketIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1TicketingTicketsTicketIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TicketingTicketsTicketIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1TicketingTicketsTicketId",
		Method:             "GET",
		PathPattern:        "/v1/ticketing/tickets/{ticket_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TicketingTicketsTicketIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TicketingTicketsTicketIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1TicketingTicketsTicketId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1TicketingTicketsTicketID updates a ticket

  Update a ticket's attributes
*/
func (a *Client) PatchV1TicketingTicketsTicketID(params *PatchV1TicketingTicketsTicketIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1TicketingTicketsTicketIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1TicketingTicketsTicketIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1TicketingTicketsTicketId",
		Method:             "PATCH",
		PathPattern:        "/v1/ticketing/tickets/{ticket_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1TicketingTicketsTicketIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchV1TicketingTicketsTicketIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1TicketingTicketsTicketId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1TicketingTickets creates a ticket

  Creates a ticket for a project
*/
func (a *Client) PostV1TicketingTickets(params *PostV1TicketingTicketsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1TicketingTicketsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1TicketingTicketsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1TicketingTickets",
		Method:             "POST",
		PathPattern:        "/v1/ticketing/tickets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1TicketingTicketsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1TicketingTicketsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1TicketingTickets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
