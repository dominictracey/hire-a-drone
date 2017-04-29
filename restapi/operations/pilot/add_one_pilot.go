package pilot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/dominictracey/rugby-scores/models"
)

// AddOnePilotHandlerFunc turns a function with the right signature into a add one pilot handler
type AddOnePilotHandlerFunc func(AddOnePilotParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOnePilotHandlerFunc) Handle(params AddOnePilotParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// AddOnePilotHandler interface for that can handle valid add one pilot params
type AddOnePilotHandler interface {
	Handle(AddOnePilotParams, *models.Principal) middleware.Responder
}

// NewAddOnePilot creates a new http.Handler for the add one pilot operation
func NewAddOnePilot(ctx *middleware.Context, handler AddOnePilotHandler) *AddOnePilot {
	return &AddOnePilot{Context: ctx, Handler: handler}
}

/*AddOnePilot swagger:route POST /pilot pilot addOnePilot

AddOnePilot add one pilot API

*/
type AddOnePilot struct {
	Context *middleware.Context
	Handler AddOnePilotHandler
}

func (o *AddOnePilot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewAddOnePilotParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
