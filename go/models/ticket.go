package models

import "github.com/go-openapi/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Ticket ticket

swagger:model ticket
*/
type Ticket struct {

	/* authorized
	 */
	Authorized bool `json:"authorized,omitempty"`

	/* client id
	 */
	ClientID string `json:"client_id,omitempty"`

	/* created at
	 */
	CreatedAt string `json:"created_at,omitempty"`

	/* id
	 */
	ID string `json:"id,omitempty"`
}

// Validate validates this ticket
func (m *Ticket) Validate(formats strfmt.Registry) error {
	return nil
}
