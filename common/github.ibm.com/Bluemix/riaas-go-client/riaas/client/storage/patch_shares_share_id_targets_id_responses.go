// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PatchSharesShareIDTargetsIDReader is a Reader for the PatchSharesShareIDTargetsID structure.
type PatchSharesShareIDTargetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSharesShareIDTargetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchSharesShareIDTargetsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchSharesShareIDTargetsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSharesShareIDTargetsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSharesShareIDTargetsIDOK creates a PatchSharesShareIDTargetsIDOK with default headers values
func NewPatchSharesShareIDTargetsIDOK() *PatchSharesShareIDTargetsIDOK {
	return &PatchSharesShareIDTargetsIDOK{}
}

/*PatchSharesShareIDTargetsIDOK handles this case with default header values.

The mount target was updated successfully.
*/
type PatchSharesShareIDTargetsIDOK struct {
	Payload *models.Sharetarget
}

func (o *PatchSharesShareIDTargetsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /shares/{share_id}/targets/{id}][%d] patchSharesShareIdTargetsIdOK  %+v", 200, o.Payload)
}

func (o *PatchSharesShareIDTargetsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sharetarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSharesShareIDTargetsIDBadRequest creates a PatchSharesShareIDTargetsIDBadRequest with default headers values
func NewPatchSharesShareIDTargetsIDBadRequest() *PatchSharesShareIDTargetsIDBadRequest {
	return &PatchSharesShareIDTargetsIDBadRequest{}
}

/*PatchSharesShareIDTargetsIDBadRequest handles this case with default header values.

error
*/
type PatchSharesShareIDTargetsIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchSharesShareIDTargetsIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /shares/{share_id}/targets/{id}][%d] patchSharesShareIdTargetsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSharesShareIDTargetsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSharesShareIDTargetsIDNotFound creates a PatchSharesShareIDTargetsIDNotFound with default headers values
func NewPatchSharesShareIDTargetsIDNotFound() *PatchSharesShareIDTargetsIDNotFound {
	return &PatchSharesShareIDTargetsIDNotFound{}
}

/*PatchSharesShareIDTargetsIDNotFound handles this case with default header values.

error
*/
type PatchSharesShareIDTargetsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchSharesShareIDTargetsIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /shares/{share_id}/targets/{id}][%d] patchSharesShareIdTargetsIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchSharesShareIDTargetsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchSharesShareIDTargetsIDBody ShareTargetPatch
swagger:model PatchSharesShareIDTargetsIDBody
*/
type PatchSharesShareIDTargetsIDBody struct {

	// The user-defined name for this file share mount target
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this patch shares share ID targets ID body
func (o *PatchSharesShareIDTargetsIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchSharesShareIDTargetsIDBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.Pattern("body"+"."+"name", "body", string(o.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchSharesShareIDTargetsIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchSharesShareIDTargetsIDBody) UnmarshalBinary(b []byte) error {
	var res PatchSharesShareIDTargetsIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
