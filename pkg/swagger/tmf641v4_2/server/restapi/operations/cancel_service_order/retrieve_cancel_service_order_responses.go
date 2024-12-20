// Code generated by go-swagger; DO NOT EDIT.

package cancel_service_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
)

// RetrieveCancelServiceOrderOKCode is the HTTP code returned for type RetrieveCancelServiceOrderOK
const RetrieveCancelServiceOrderOKCode int = 200

/*
RetrieveCancelServiceOrderOK Success

swagger:response retrieveCancelServiceOrderOK
*/
type RetrieveCancelServiceOrderOK struct {

	/*
	  In: Body
	*/
	Payload *models.CancelServiceOrder `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderOK creates RetrieveCancelServiceOrderOK with default headers values
func NewRetrieveCancelServiceOrderOK() *RetrieveCancelServiceOrderOK {

	return &RetrieveCancelServiceOrderOK{}
}

// WithPayload adds the payload to the retrieve cancel service order o k response
func (o *RetrieveCancelServiceOrderOK) WithPayload(payload *models.CancelServiceOrder) *RetrieveCancelServiceOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order o k response
func (o *RetrieveCancelServiceOrderOK) SetPayload(payload *models.CancelServiceOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveCancelServiceOrderBadRequestCode is the HTTP code returned for type RetrieveCancelServiceOrderBadRequest
const RetrieveCancelServiceOrderBadRequestCode int = 400

/*
RetrieveCancelServiceOrderBadRequest Bad Request

swagger:response retrieveCancelServiceOrderBadRequest
*/
type RetrieveCancelServiceOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderBadRequest creates RetrieveCancelServiceOrderBadRequest with default headers values
func NewRetrieveCancelServiceOrderBadRequest() *RetrieveCancelServiceOrderBadRequest {

	return &RetrieveCancelServiceOrderBadRequest{}
}

// WithPayload adds the payload to the retrieve cancel service order bad request response
func (o *RetrieveCancelServiceOrderBadRequest) WithPayload(payload *models.Error) *RetrieveCancelServiceOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order bad request response
func (o *RetrieveCancelServiceOrderBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveCancelServiceOrderUnauthorizedCode is the HTTP code returned for type RetrieveCancelServiceOrderUnauthorized
const RetrieveCancelServiceOrderUnauthorizedCode int = 401

/*
RetrieveCancelServiceOrderUnauthorized Unauthorized

swagger:response retrieveCancelServiceOrderUnauthorized
*/
type RetrieveCancelServiceOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderUnauthorized creates RetrieveCancelServiceOrderUnauthorized with default headers values
func NewRetrieveCancelServiceOrderUnauthorized() *RetrieveCancelServiceOrderUnauthorized {

	return &RetrieveCancelServiceOrderUnauthorized{}
}

// WithPayload adds the payload to the retrieve cancel service order unauthorized response
func (o *RetrieveCancelServiceOrderUnauthorized) WithPayload(payload *models.Error) *RetrieveCancelServiceOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order unauthorized response
func (o *RetrieveCancelServiceOrderUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveCancelServiceOrderForbiddenCode is the HTTP code returned for type RetrieveCancelServiceOrderForbidden
const RetrieveCancelServiceOrderForbiddenCode int = 403

/*
RetrieveCancelServiceOrderForbidden Forbidden

swagger:response retrieveCancelServiceOrderForbidden
*/
type RetrieveCancelServiceOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderForbidden creates RetrieveCancelServiceOrderForbidden with default headers values
func NewRetrieveCancelServiceOrderForbidden() *RetrieveCancelServiceOrderForbidden {

	return &RetrieveCancelServiceOrderForbidden{}
}

// WithPayload adds the payload to the retrieve cancel service order forbidden response
func (o *RetrieveCancelServiceOrderForbidden) WithPayload(payload *models.Error) *RetrieveCancelServiceOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order forbidden response
func (o *RetrieveCancelServiceOrderForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveCancelServiceOrderNotFoundCode is the HTTP code returned for type RetrieveCancelServiceOrderNotFound
const RetrieveCancelServiceOrderNotFoundCode int = 404

/*
RetrieveCancelServiceOrderNotFound Not Found

swagger:response retrieveCancelServiceOrderNotFound
*/
type RetrieveCancelServiceOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderNotFound creates RetrieveCancelServiceOrderNotFound with default headers values
func NewRetrieveCancelServiceOrderNotFound() *RetrieveCancelServiceOrderNotFound {

	return &RetrieveCancelServiceOrderNotFound{}
}

// WithPayload adds the payload to the retrieve cancel service order not found response
func (o *RetrieveCancelServiceOrderNotFound) WithPayload(payload *models.Error) *RetrieveCancelServiceOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order not found response
func (o *RetrieveCancelServiceOrderNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveCancelServiceOrderMethodNotAllowedCode is the HTTP code returned for type RetrieveCancelServiceOrderMethodNotAllowed
const RetrieveCancelServiceOrderMethodNotAllowedCode int = 405

/*
RetrieveCancelServiceOrderMethodNotAllowed Method Not allowed

swagger:response retrieveCancelServiceOrderMethodNotAllowed
*/
type RetrieveCancelServiceOrderMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderMethodNotAllowed creates RetrieveCancelServiceOrderMethodNotAllowed with default headers values
func NewRetrieveCancelServiceOrderMethodNotAllowed() *RetrieveCancelServiceOrderMethodNotAllowed {

	return &RetrieveCancelServiceOrderMethodNotAllowed{}
}

// WithPayload adds the payload to the retrieve cancel service order method not allowed response
func (o *RetrieveCancelServiceOrderMethodNotAllowed) WithPayload(payload *models.Error) *RetrieveCancelServiceOrderMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order method not allowed response
func (o *RetrieveCancelServiceOrderMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveCancelServiceOrderConflictCode is the HTTP code returned for type RetrieveCancelServiceOrderConflict
const RetrieveCancelServiceOrderConflictCode int = 409

/*
RetrieveCancelServiceOrderConflict Conflict

swagger:response retrieveCancelServiceOrderConflict
*/
type RetrieveCancelServiceOrderConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderConflict creates RetrieveCancelServiceOrderConflict with default headers values
func NewRetrieveCancelServiceOrderConflict() *RetrieveCancelServiceOrderConflict {

	return &RetrieveCancelServiceOrderConflict{}
}

// WithPayload adds the payload to the retrieve cancel service order conflict response
func (o *RetrieveCancelServiceOrderConflict) WithPayload(payload *models.Error) *RetrieveCancelServiceOrderConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order conflict response
func (o *RetrieveCancelServiceOrderConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveCancelServiceOrderInternalServerErrorCode is the HTTP code returned for type RetrieveCancelServiceOrderInternalServerError
const RetrieveCancelServiceOrderInternalServerErrorCode int = 500

/*
RetrieveCancelServiceOrderInternalServerError Internal Server Error

swagger:response retrieveCancelServiceOrderInternalServerError
*/
type RetrieveCancelServiceOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveCancelServiceOrderInternalServerError creates RetrieveCancelServiceOrderInternalServerError with default headers values
func NewRetrieveCancelServiceOrderInternalServerError() *RetrieveCancelServiceOrderInternalServerError {

	return &RetrieveCancelServiceOrderInternalServerError{}
}

// WithPayload adds the payload to the retrieve cancel service order internal server error response
func (o *RetrieveCancelServiceOrderInternalServerError) WithPayload(payload *models.Error) *RetrieveCancelServiceOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve cancel service order internal server error response
func (o *RetrieveCancelServiceOrderInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveCancelServiceOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
