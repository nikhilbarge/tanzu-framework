// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateVSphereRegionalClusterHandlerFunc turns a function with the right signature into a create v sphere regional cluster handler
type CreateVSphereRegionalClusterHandlerFunc func(CreateVSphereRegionalClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateVSphereRegionalClusterHandlerFunc) Handle(params CreateVSphereRegionalClusterParams) middleware.Responder {
	return fn(params)
}

// CreateVSphereRegionalClusterHandler interface for that can handle valid create v sphere regional cluster params
type CreateVSphereRegionalClusterHandler interface {
	Handle(CreateVSphereRegionalClusterParams) middleware.Responder
}

// NewCreateVSphereRegionalCluster creates a new http.Handler for the create v sphere regional cluster operation
func NewCreateVSphereRegionalCluster(ctx *middleware.Context, handler CreateVSphereRegionalClusterHandler) *CreateVSphereRegionalCluster {
	return &CreateVSphereRegionalCluster{Context: ctx, Handler: handler}
}

/*
CreateVSphereRegionalCluster swagger:route POST /api/providers/vsphere/create vsphere createVSphereRegionalCluster

Create vSphere regional cluster
*/
type CreateVSphereRegionalCluster struct {
	Context *middleware.Context
	Handler CreateVSphereRegionalClusterHandler
}

func (o *CreateVSphereRegionalCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateVSphereRegionalClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
