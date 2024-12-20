// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListenToCancelServiceOrderStateChangeEventHandlerFunc turns a function with the right signature into a listen to cancel service order state change event handler
type ListenToCancelServiceOrderStateChangeEventHandlerFunc func(ListenToCancelServiceOrderStateChangeEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListenToCancelServiceOrderStateChangeEventHandlerFunc) Handle(params ListenToCancelServiceOrderStateChangeEventParams) middleware.Responder {
	return fn(params)
}

// ListenToCancelServiceOrderStateChangeEventHandler interface for that can handle valid listen to cancel service order state change event params
type ListenToCancelServiceOrderStateChangeEventHandler interface {
	Handle(ListenToCancelServiceOrderStateChangeEventParams) middleware.Responder
}

// NewListenToCancelServiceOrderStateChangeEvent creates a new http.Handler for the listen to cancel service order state change event operation
func NewListenToCancelServiceOrderStateChangeEvent(ctx *middleware.Context, handler ListenToCancelServiceOrderStateChangeEventHandler) *ListenToCancelServiceOrderStateChangeEvent {
	return &ListenToCancelServiceOrderStateChangeEvent{Context: ctx, Handler: handler}
}

/*
	ListenToCancelServiceOrderStateChangeEvent swagger:route POST /listener/cancelServiceOrderStateChangeEvent notification listeners (client side) listenToCancelServiceOrderStateChangeEvent

# Client listener for entity CancelServiceOrderStateChangeEvent

Example of a client listener for receiving the notification CancelServiceOrderStateChangeEvent
*/
type ListenToCancelServiceOrderStateChangeEvent struct {
	Context *middleware.Context
	Handler ListenToCancelServiceOrderStateChangeEventHandler
}

func (o *ListenToCancelServiceOrderStateChangeEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListenToCancelServiceOrderStateChangeEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
