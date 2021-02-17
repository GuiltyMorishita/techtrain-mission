// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/GuiltyMorishita/techtrain-mission/models"
)

// GetUserGetOKCode is the HTTP code returned for type GetUserGetOK
const GetUserGetOKCode int = 200

/*GetUserGetOK A successful response.

swagger:response getUserGetOK
*/
type GetUserGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserGetResponse `json:"body,omitempty"`
}

// NewGetUserGetOK creates GetUserGetOK with default headers values
func NewGetUserGetOK() *GetUserGetOK {

	return &GetUserGetOK{}
}

// WithPayload adds the payload to the get user get o k response
func (o *GetUserGetOK) WithPayload(payload *models.UserGetResponse) *GetUserGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user get o k response
func (o *GetUserGetOK) SetPayload(payload *models.UserGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
