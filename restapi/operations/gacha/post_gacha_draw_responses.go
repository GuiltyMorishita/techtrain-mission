// Code generated by go-swagger; DO NOT EDIT.

package gacha

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/GuiltyMorishita/techtrain-mission/models"
)

// PostGachaDrawOKCode is the HTTP code returned for type PostGachaDrawOK
const PostGachaDrawOKCode int = 200

/*PostGachaDrawOK A successful response.

swagger:response postGachaDrawOK
*/
type PostGachaDrawOK struct {

	/*
	  In: Body
	*/
	Payload *models.GachaDrawResponse `json:"body,omitempty"`
}

// NewPostGachaDrawOK creates PostGachaDrawOK with default headers values
func NewPostGachaDrawOK() *PostGachaDrawOK {

	return &PostGachaDrawOK{}
}

// WithPayload adds the payload to the post gacha draw o k response
func (o *PostGachaDrawOK) WithPayload(payload *models.GachaDrawResponse) *PostGachaDrawOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post gacha draw o k response
func (o *PostGachaDrawOK) SetPayload(payload *models.GachaDrawResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostGachaDrawOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
