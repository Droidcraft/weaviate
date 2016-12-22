package registration_tickets


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveRegistrationTicketsUpdateHandlerFunc turns a function with the right signature into a weave registration tickets update handler
type WeaveRegistrationTicketsUpdateHandlerFunc func(WeaveRegistrationTicketsUpdateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveRegistrationTicketsUpdateHandlerFunc) Handle(params WeaveRegistrationTicketsUpdateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveRegistrationTicketsUpdateHandler interface for that can handle valid weave registration tickets update params
type WeaveRegistrationTicketsUpdateHandler interface {
	Handle(WeaveRegistrationTicketsUpdateParams, interface{}) middleware.Responder
}

// NewWeaveRegistrationTicketsUpdate creates a new http.Handler for the weave registration tickets update operation
func NewWeaveRegistrationTicketsUpdate(ctx *middleware.Context, handler WeaveRegistrationTicketsUpdateHandler) *WeaveRegistrationTicketsUpdate {
	return &WeaveRegistrationTicketsUpdate{Context: ctx, Handler: handler}
}

/*WeaveRegistrationTicketsUpdate swagger:route PUT /registrationTickets/{registrationTicketId} registrationTickets weaveRegistrationTicketsUpdate

Updates an existing registration ticket.

*/
type WeaveRegistrationTicketsUpdate struct {
	Context *middleware.Context
	Handler WeaveRegistrationTicketsUpdateHandler
}

func (o *WeaveRegistrationTicketsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveRegistrationTicketsUpdateParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
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