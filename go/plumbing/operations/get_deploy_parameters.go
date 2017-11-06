package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeployParams creates a new GetDeployParams object
// with the default values initialized.
func NewGetDeployParams() *GetDeployParams {
	var ()
	return &GetDeployParams{}
}

/*GetDeployParams contains all the parameters to send to the API endpoint
for the get deploy operation typically these are written to a http.Request
*/
type GetDeployParams struct {

	/*DeployID*/
	DeployID string
}

// WithDeployID adds the deployId to the get deploy params
func (o *GetDeployParams) WithDeployID(DeployID string) *GetDeployParams {
	o.DeployID = DeployID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
