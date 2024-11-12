// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
)

// NewListenToCancelServiceOrderInformationRequiredEventParams creates a new ListenToCancelServiceOrderInformationRequiredEventParams object
//
// There are no default values defined in the spec.
func NewListenToCancelServiceOrderInformationRequiredEventParams() ListenToCancelServiceOrderInformationRequiredEventParams {

	return ListenToCancelServiceOrderInformationRequiredEventParams{}
}

// ListenToCancelServiceOrderInformationRequiredEventParams contains all the bound params for the listen to cancel service order information required event operation
// typically these are obtained from a http.Request
//
// swagger:parameters listenToCancelServiceOrderInformationRequiredEvent
type ListenToCancelServiceOrderInformationRequiredEventParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The event data
	  Required: true
	  In: body
	*/
	Data *models.CancelServiceOrderInformationRequiredEvent
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListenToCancelServiceOrderInformationRequiredEventParams() beforehand.
func (o *ListenToCancelServiceOrderInformationRequiredEventParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CancelServiceOrderInformationRequiredEvent
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body", ""))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
