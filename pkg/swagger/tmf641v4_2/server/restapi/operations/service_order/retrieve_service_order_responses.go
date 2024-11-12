// Code generated by go-swagger; DO NOT EDIT.

package service_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
)

// RetrieveServiceOrderOKCode is the HTTP code returned for type RetrieveServiceOrderOK
const RetrieveServiceOrderOKCode int = 200

/*
RetrieveServiceOrderOK Success

swagger:response retrieveServiceOrderOK
*/
type RetrieveServiceOrderOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceOrder `json:"body,omitempty"`
}

// NewRetrieveServiceOrderOK creates RetrieveServiceOrderOK with default headers values
func NewRetrieveServiceOrderOK() *RetrieveServiceOrderOK {

	return &RetrieveServiceOrderOK{}
}

// WithPayload adds the payload to the retrieve service order o k response
func (o *RetrieveServiceOrderOK) WithPayload(payload *models.ServiceOrder) *RetrieveServiceOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order o k response
func (o *RetrieveServiceOrderOK) SetPayload(payload *models.ServiceOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveServiceOrderBadRequestCode is the HTTP code returned for type RetrieveServiceOrderBadRequest
const RetrieveServiceOrderBadRequestCode int = 400

/*
RetrieveServiceOrderBadRequest Bad Request

swagger:response retrieveServiceOrderBadRequest
*/
type RetrieveServiceOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveServiceOrderBadRequest creates RetrieveServiceOrderBadRequest with default headers values
func NewRetrieveServiceOrderBadRequest() *RetrieveServiceOrderBadRequest {

	return &RetrieveServiceOrderBadRequest{}
}

// WithPayload adds the payload to the retrieve service order bad request response
func (o *RetrieveServiceOrderBadRequest) WithPayload(payload *models.Error) *RetrieveServiceOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order bad request response
func (o *RetrieveServiceOrderBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveServiceOrderUnauthorizedCode is the HTTP code returned for type RetrieveServiceOrderUnauthorized
const RetrieveServiceOrderUnauthorizedCode int = 401

/*
RetrieveServiceOrderUnauthorized Unauthorized

swagger:response retrieveServiceOrderUnauthorized
*/
type RetrieveServiceOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveServiceOrderUnauthorized creates RetrieveServiceOrderUnauthorized with default headers values
func NewRetrieveServiceOrderUnauthorized() *RetrieveServiceOrderUnauthorized {

	return &RetrieveServiceOrderUnauthorized{}
}

// WithPayload adds the payload to the retrieve service order unauthorized response
func (o *RetrieveServiceOrderUnauthorized) WithPayload(payload *models.Error) *RetrieveServiceOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order unauthorized response
func (o *RetrieveServiceOrderUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveServiceOrderForbiddenCode is the HTTP code returned for type RetrieveServiceOrderForbidden
const RetrieveServiceOrderForbiddenCode int = 403

/*
RetrieveServiceOrderForbidden Forbidden

swagger:response retrieveServiceOrderForbidden
*/
type RetrieveServiceOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveServiceOrderForbidden creates RetrieveServiceOrderForbidden with default headers values
func NewRetrieveServiceOrderForbidden() *RetrieveServiceOrderForbidden {

	return &RetrieveServiceOrderForbidden{}
}

// WithPayload adds the payload to the retrieve service order forbidden response
func (o *RetrieveServiceOrderForbidden) WithPayload(payload *models.Error) *RetrieveServiceOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order forbidden response
func (o *RetrieveServiceOrderForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveServiceOrderNotFoundCode is the HTTP code returned for type RetrieveServiceOrderNotFound
const RetrieveServiceOrderNotFoundCode int = 404

/*
RetrieveServiceOrderNotFound Not Found

swagger:response retrieveServiceOrderNotFound
*/
type RetrieveServiceOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveServiceOrderNotFound creates RetrieveServiceOrderNotFound with default headers values
func NewRetrieveServiceOrderNotFound() *RetrieveServiceOrderNotFound {

	return &RetrieveServiceOrderNotFound{}
}

// WithPayload adds the payload to the retrieve service order not found response
func (o *RetrieveServiceOrderNotFound) WithPayload(payload *models.Error) *RetrieveServiceOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order not found response
func (o *RetrieveServiceOrderNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveServiceOrderMethodNotAllowedCode is the HTTP code returned for type RetrieveServiceOrderMethodNotAllowed
const RetrieveServiceOrderMethodNotAllowedCode int = 405

/*
RetrieveServiceOrderMethodNotAllowed Method Not allowed

swagger:response retrieveServiceOrderMethodNotAllowed
*/
type RetrieveServiceOrderMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveServiceOrderMethodNotAllowed creates RetrieveServiceOrderMethodNotAllowed with default headers values
func NewRetrieveServiceOrderMethodNotAllowed() *RetrieveServiceOrderMethodNotAllowed {

	return &RetrieveServiceOrderMethodNotAllowed{}
}

// WithPayload adds the payload to the retrieve service order method not allowed response
func (o *RetrieveServiceOrderMethodNotAllowed) WithPayload(payload *models.Error) *RetrieveServiceOrderMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order method not allowed response
func (o *RetrieveServiceOrderMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveServiceOrderConflictCode is the HTTP code returned for type RetrieveServiceOrderConflict
const RetrieveServiceOrderConflictCode int = 409

/*
RetrieveServiceOrderConflict Conflict

swagger:response retrieveServiceOrderConflict
*/
type RetrieveServiceOrderConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveServiceOrderConflict creates RetrieveServiceOrderConflict with default headers values
func NewRetrieveServiceOrderConflict() *RetrieveServiceOrderConflict {

	return &RetrieveServiceOrderConflict{}
}

// WithPayload adds the payload to the retrieve service order conflict response
func (o *RetrieveServiceOrderConflict) WithPayload(payload *models.Error) *RetrieveServiceOrderConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order conflict response
func (o *RetrieveServiceOrderConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveServiceOrderInternalServerErrorCode is the HTTP code returned for type RetrieveServiceOrderInternalServerError
const RetrieveServiceOrderInternalServerErrorCode int = 500

/*
RetrieveServiceOrderInternalServerError Internal Server Error

swagger:response retrieveServiceOrderInternalServerError
*/
type RetrieveServiceOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveServiceOrderInternalServerError creates RetrieveServiceOrderInternalServerError with default headers values
func NewRetrieveServiceOrderInternalServerError() *RetrieveServiceOrderInternalServerError {

	return &RetrieveServiceOrderInternalServerError{}
}

// WithPayload adds the payload to the retrieve service order internal server error response
func (o *RetrieveServiceOrderInternalServerError) WithPayload(payload *models.Error) *RetrieveServiceOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service order internal server error response
func (o *RetrieveServiceOrderInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}