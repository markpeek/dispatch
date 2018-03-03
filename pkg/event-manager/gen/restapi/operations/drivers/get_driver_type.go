///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDriverTypeHandlerFunc turns a function with the right signature into a get driver type handler
type GetDriverTypeHandlerFunc func(GetDriverTypeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDriverTypeHandlerFunc) Handle(params GetDriverTypeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDriverTypeHandler interface for that can handle valid get driver type params
type GetDriverTypeHandler interface {
	Handle(GetDriverTypeParams, interface{}) middleware.Responder
}

// NewGetDriverType creates a new http.Handler for the get driver type operation
func NewGetDriverType(ctx *middleware.Context, handler GetDriverTypeHandler) *GetDriverType {
	return &GetDriverType{Context: ctx, Handler: handler}
}

/*GetDriverType swagger:route GET /drivertypes/{driverTypeName} drivers getDriverType

Find driver type by Name

Returns a single driver type

*/
type GetDriverType struct {
	Context *middleware.Context
	Handler GetDriverTypeHandler
}

func (o *GetDriverType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDriverTypeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
