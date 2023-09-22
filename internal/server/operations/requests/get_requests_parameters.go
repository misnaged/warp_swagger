// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRequestsParams creates a new GetRequestsParams object
// with the default values initialized.
func NewGetRequestsParams() GetRequestsParams {

	var (
		// initialize parameters with default values

		limitDefault  = int64(1)
		offsetDefault = int64(0)
	)

	return GetRequestsParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetRequestsParams contains all the bound params for the get requests operation
// typically these are obtained from a http.Request
//
// swagger:parameters getRequests
type GetRequestsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*API key for filtering analytics.
	  In: query
	*/
	APIKey *string
	/*Chain identifier for filtering analytics.
	  In: query
	*/
	Chain *string
	/*Start date for filtering analytics. If not included, retrieves all analytics from the beginning.
	  In: query
	*/
	From *int64
	/*Number of records to retrieve per page for pagination. Default is set by the server.
	  In: query
	  Default: 1
	*/
	Limit *int64
	/*Starting point for records retrieval for pagination. Default is 0.
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*Organization identifier.
	  Required: true
	  In: path
	*/
	Org string
	/*End date for filtering analytics. If not included, retrieves analytics up to the latest.
	  In: query
	*/
	To *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRequestsParams() beforehand.
func (o *GetRequestsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAPIKey, qhkAPIKey, _ := qs.GetOK("apiKey")
	if err := o.bindAPIKey(qAPIKey, qhkAPIKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qChain, qhkChain, _ := qs.GetOK("chain")
	if err := o.bindChain(qChain, qhkChain, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	rOrg, rhkOrg, _ := route.Params.GetOK("org")
	if err := o.bindOrg(rOrg, rhkOrg, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAPIKey binds and validates parameter APIKey from query.
func (o *GetRequestsParams) bindAPIKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.APIKey = &raw

	return nil
}

// bindChain binds and validates parameter Chain from query.
func (o *GetRequestsParams) bindChain(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Chain = &raw

	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetRequestsParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("from", "query", "int64", raw)
	}
	o.From = &value

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetRequestsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRequestsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetRequestsParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRequestsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	return nil
}

// bindOrg binds and validates parameter Org from path.
func (o *GetRequestsParams) bindOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Org = raw

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetRequestsParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("to", "query", "int64", raw)
	}
	o.To = &value

	return nil
}
