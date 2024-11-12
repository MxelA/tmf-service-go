// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListenToCancelServiceOrderCreateEventHandlerFunc turns a function with the right signature into a listen to cancel service order create event handler
type ListenToCancelServiceOrderCreateEventHandlerFunc func(ListenToCancelServiceOrderCreateEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListenToCancelServiceOrderCreateEventHandlerFunc) Handle(params ListenToCancelServiceOrderCreateEventParams) middleware.Responder {
	return fn(params)
}

// ListenToCancelServiceOrderCreateEventHandler interface for that can handle valid listen to cancel service order create event params
type ListenToCancelServiceOrderCreateEventHandler interface {
	Handle(ListenToCancelServiceOrderCreateEventParams) middleware.Responder
}

// NewListenToCancelServiceOrderCreateEvent creates a new http.Handler for the listen to cancel service order create event operation
func NewListenToCancelServiceOrderCreateEvent(ctx *middleware.Context, handler ListenToCancelServiceOrderCreateEventHandler) *ListenToCancelServiceOrderCreateEvent {
	return &ListenToCancelServiceOrderCreateEvent{Context: ctx, Handler: handler}
}

/*
	ListenToCancelServiceOrderCreateEvent swagger:route POST /listener/cancelServiceOrderCreateEvent notification listeners (client side) listenToCancelServiceOrderCreateEvent

# Client listener for entity CancelServiceOrderCreateEvent

Example of a client listener for receiving the notification CancelServiceOrderCreateEvent
*/
type ListenToCancelServiceOrderCreateEvent struct {
	Context *middleware.Context
	Handler ListenToCancelServiceOrderCreateEventHandler
}

func (o *ListenToCancelServiceOrderCreateEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListenToCancelServiceOrderCreateEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
