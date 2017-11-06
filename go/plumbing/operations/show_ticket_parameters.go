package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowTicketParams creates a new ShowTicketParams object
// with the default values initialized.
func NewShowTicketParams() *ShowTicketParams {
	var ()
	return &ShowTicketParams{}
}

/*ShowTicketParams contains all the parameters to send to the API endpoint
for the show ticket operation typically these are written to a http.Request
*/
type ShowTicketParams struct {

	/*TicketID*/
	TicketID string
}

// WithTicketID adds the ticketId to the show ticket params
func (o *ShowTicketParams) WithTicketID(TicketID string) *ShowTicketParams {
	o.TicketID = TicketID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ShowTicketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param ticket_id
	if err := r.SetPathParam("ticket_id", o.TicketID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
