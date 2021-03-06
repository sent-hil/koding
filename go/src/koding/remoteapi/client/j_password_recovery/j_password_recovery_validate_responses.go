package j_password_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JPasswordRecoveryValidateReader is a Reader for the JPasswordRecoveryValidate structure.
type JPasswordRecoveryValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JPasswordRecoveryValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJPasswordRecoveryValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJPasswordRecoveryValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJPasswordRecoveryValidateOK creates a JPasswordRecoveryValidateOK with default headers values
func NewJPasswordRecoveryValidateOK() *JPasswordRecoveryValidateOK {
	return &JPasswordRecoveryValidateOK{}
}

/*JPasswordRecoveryValidateOK handles this case with default header values.

Request processed successfully
*/
type JPasswordRecoveryValidateOK struct {
	Payload *models.DefaultResponse
}

func (o *JPasswordRecoveryValidateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JPasswordRecovery.validate][%d] jPasswordRecoveryValidateOK  %+v", 200, o.Payload)
}

func (o *JPasswordRecoveryValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJPasswordRecoveryValidateUnauthorized creates a JPasswordRecoveryValidateUnauthorized with default headers values
func NewJPasswordRecoveryValidateUnauthorized() *JPasswordRecoveryValidateUnauthorized {
	return &JPasswordRecoveryValidateUnauthorized{}
}

/*JPasswordRecoveryValidateUnauthorized handles this case with default header values.

Unauthorized request
*/
type JPasswordRecoveryValidateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JPasswordRecoveryValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JPasswordRecovery.validate][%d] jPasswordRecoveryValidateUnauthorized  %+v", 401, o.Payload)
}

func (o *JPasswordRecoveryValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
