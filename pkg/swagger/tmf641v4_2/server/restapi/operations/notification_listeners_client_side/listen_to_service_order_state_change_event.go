// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListenToServiceOrderStateChangeEventHandlerFunc turns a function with the right signature into a listen to service order state change event handler
type ListenToServiceOrderStateChangeEventHandlerFunc func(ListenToServiceOrderStateChangeEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListenToServiceOrderStateChangeEventHandlerFunc) Handle(params ListenToServiceOrderStateChangeEventParams) middleware.Responder {
	return fn(params)
}

// ListenToServiceOrderStateChangeEventHandler interface for that can handle valid listen to service order state change event params
type ListenToServiceOrderStateChangeEventHandler interface {
	Handle(ListenToServiceOrderStateChangeEventParams) middleware.Responder
}

// NewListenToServiceOrderStateChangeEvent creates a new http.Handler for the listen to service order state change event operation
func NewListenToServiceOrderStateChangeEvent(ctx *middleware.Context, handler ListenToServiceOrderStateChangeEventHandler) *ListenToServiceOrderStateChangeEvent {
	return &ListenToServiceOrderStateChangeEvent{Context: ctx, Handler: handler}
}

/*
	ListenToServiceOrderStateChangeEvent swagger:route POST /listener/serviceOrderStateChangeEvent notification listeners (client side) listenToServiceOrderStateChangeEvent

# Client listener for entity ServiceOrderStateChangeEvent

Example of a client listener for receiving the notification ServiceOrderStateChangeEvent
*/
type ListenToServiceOrderStateChangeEvent struct {
	Context *middleware.Context
	Handler ListenToServiceOrderStateChangeEventHandler
}

func (o *ListenToServiceOrderStateChangeEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListenToServiceOrderStateChangeEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}