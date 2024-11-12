// Code generated by go-swagger; DO NOT EDIT.

package service_order

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

// NewCreateServiceOrderParams creates a new CreateServiceOrderParams object
//
// There are no default values defined in the spec.
func NewCreateServiceOrderParams() CreateServiceOrderParams {

	return CreateServiceOrderParams{}
}

// CreateServiceOrderParams contains all the bound params for the create service order operation
// typically these are obtained from a http.Request
//
// swagger:parameters createServiceOrder
type CreateServiceOrderParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ServiceOrder to be created
	  Required: true
	  In: body
	*/
	ServiceOrder *models.ServiceOrderCreate
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateServiceOrderParams() beforehand.
func (o *CreateServiceOrderParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ServiceOrderCreate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("serviceOrder", "body", ""))
			} else {
				res = append(res, errors.NewParseError("serviceOrder", "body", "", err))
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
				o.ServiceOrder = &body
			}
		}
	} else {
		res = append(res, errors.Required("serviceOrder", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
