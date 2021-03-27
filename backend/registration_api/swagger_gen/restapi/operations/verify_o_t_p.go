// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VerifyOTPHandlerFunc turns a function with the right signature into a verify o t p handler
type VerifyOTPHandlerFunc func(VerifyOTPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyOTPHandlerFunc) Handle(params VerifyOTPParams) middleware.Responder {
	return fn(params)
}

// VerifyOTPHandler interface for that can handle valid verify o t p params
type VerifyOTPHandler interface {
	Handle(VerifyOTPParams) middleware.Responder
}

// NewVerifyOTP creates a new http.Handler for the verify o t p operation
func NewVerifyOTP(ctx *middleware.Context, handler VerifyOTPHandler) *VerifyOTP {
	return &VerifyOTP{Context: ctx, Handler: handler}
}

/* VerifyOTP swagger:route POST /verifyOTP verifyOTP

Verify OTP

*/
type VerifyOTP struct {
	Context *middleware.Context
	Handler VerifyOTPHandler
}

func (o *VerifyOTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyOTPParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// VerifyOTPBody verify o t p body
//
// swagger:model VerifyOTPBody
type VerifyOTPBody struct {

	// otp
	Otp string `json:"otp,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`
}

// Validate validates this verify o t p body
func (o *VerifyOTPBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this verify o t p body based on context it is used
func (o *VerifyOTPBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VerifyOTPBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VerifyOTPBody) UnmarshalBinary(b []byte) error {
	var res VerifyOTPBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// VerifyOTPOKBody verify o t p o k body
//
// swagger:model VerifyOTPOKBody
type VerifyOTPOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this verify o t p o k body
func (o *VerifyOTPOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this verify o t p o k body based on context it is used
func (o *VerifyOTPOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VerifyOTPOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VerifyOTPOKBody) UnmarshalBinary(b []byte) error {
	var res VerifyOTPOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}