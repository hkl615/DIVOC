// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// GetFacilityGroupsOKCode is the HTTP code returned for type GetFacilityGroupsOK
const GetFacilityGroupsOKCode int = 200

/*GetFacilityGroupsOK OK

swagger:response getFacilityGroupsOK
*/
type GetFacilityGroupsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserGroup `json:"body,omitempty"`
}

// NewGetFacilityGroupsOK creates GetFacilityGroupsOK with default headers values
func NewGetFacilityGroupsOK() *GetFacilityGroupsOK {

	return &GetFacilityGroupsOK{}
}

// WithPayload adds the payload to the get facility groups o k response
func (o *GetFacilityGroupsOK) WithPayload(payload []*models.UserGroup) *GetFacilityGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get facility groups o k response
func (o *GetFacilityGroupsOK) SetPayload(payload []*models.UserGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFacilityGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UserGroup, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetFacilityGroupsBadRequestCode is the HTTP code returned for type GetFacilityGroupsBadRequest
const GetFacilityGroupsBadRequestCode int = 400

/*GetFacilityGroupsBadRequest Invalid input

swagger:response getFacilityGroupsBadRequest
*/
type GetFacilityGroupsBadRequest struct {
}

// NewGetFacilityGroupsBadRequest creates GetFacilityGroupsBadRequest with default headers values
func NewGetFacilityGroupsBadRequest() *GetFacilityGroupsBadRequest {

	return &GetFacilityGroupsBadRequest{}
}

// WriteResponse to the client
func (o *GetFacilityGroupsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetFacilityGroupsUnauthorizedCode is the HTTP code returned for type GetFacilityGroupsUnauthorized
const GetFacilityGroupsUnauthorizedCode int = 401

/*GetFacilityGroupsUnauthorized Unauthorized

swagger:response getFacilityGroupsUnauthorized
*/
type GetFacilityGroupsUnauthorized struct {
}

// NewGetFacilityGroupsUnauthorized creates GetFacilityGroupsUnauthorized with default headers values
func NewGetFacilityGroupsUnauthorized() *GetFacilityGroupsUnauthorized {

	return &GetFacilityGroupsUnauthorized{}
}

// WriteResponse to the client
func (o *GetFacilityGroupsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}