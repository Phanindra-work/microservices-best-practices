// Code generated by go-swagger; DO NOT EDIT.

package teacher

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/microservice/models/apimodels"
)

// SuspendStudentNoContentCode is the HTTP code returned for type SuspendStudentNoContent
const SuspendStudentNoContentCode int = 204

/*SuspendStudentNoContent No content

swagger:response suspendStudentNoContent
*/
type SuspendStudentNoContent struct {
}

// NewSuspendStudentNoContent creates SuspendStudentNoContent with default headers values
func NewSuspendStudentNoContent() *SuspendStudentNoContent {

	return &SuspendStudentNoContent{}
}

// WriteResponse to the client
func (o *SuspendStudentNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// SuspendStudentBadRequestCode is the HTTP code returned for type SuspendStudentBadRequest
const SuspendStudentBadRequestCode int = 400

/*SuspendStudentBadRequest Bad request <br>


swagger:response suspendStudentBadRequest
*/
type SuspendStudentBadRequest struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewSuspendStudentBadRequest creates SuspendStudentBadRequest with default headers values
func NewSuspendStudentBadRequest() *SuspendStudentBadRequest {

	return &SuspendStudentBadRequest{}
}

// WithPayload adds the payload to the suspend student bad request response
func (o *SuspendStudentBadRequest) WithPayload(payload *apimodels.ErrorResponse) *SuspendStudentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the suspend student bad request response
func (o *SuspendStudentBadRequest) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SuspendStudentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SuspendStudentUnauthorizedCode is the HTTP code returned for type SuspendStudentUnauthorized
const SuspendStudentUnauthorizedCode int = 401

/*SuspendStudentUnauthorized Unauthorized

swagger:response suspendStudentUnauthorized
*/
type SuspendStudentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *apimodels.ErrorResponse `json:"body,omitempty"`
}

// NewSuspendStudentUnauthorized creates SuspendStudentUnauthorized with default headers values
func NewSuspendStudentUnauthorized() *SuspendStudentUnauthorized {

	return &SuspendStudentUnauthorized{}
}

// WithPayload adds the payload to the suspend student unauthorized response
func (o *SuspendStudentUnauthorized) WithPayload(payload *apimodels.ErrorResponse) *SuspendStudentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the suspend student unauthorized response
func (o *SuspendStudentUnauthorized) SetPayload(payload *apimodels.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SuspendStudentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SuspendStudentInternalServerErrorCode is the HTTP code returned for type SuspendStudentInternalServerError
const SuspendStudentInternalServerErrorCode int = 500

/*SuspendStudentInternalServerError Internal error

swagger:response suspendStudentInternalServerError
*/
type SuspendStudentInternalServerError struct {
}

// NewSuspendStudentInternalServerError creates SuspendStudentInternalServerError with default headers values
func NewSuspendStudentInternalServerError() *SuspendStudentInternalServerError {

	return &SuspendStudentInternalServerError{}
}

// WriteResponse to the client
func (o *SuspendStudentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
