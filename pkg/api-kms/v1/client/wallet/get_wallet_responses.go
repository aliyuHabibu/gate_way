// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/oxygenpay/oxygen/pkg/api-kms/v1/model"
)

// GetWalletReader is a Reader for the GetWallet structure.
type GetWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWalletOK creates a GetWalletOK with default headers values
func NewGetWalletOK() *GetWalletOK {
	return &GetWalletOK{}
}

/* GetWalletOK describes a response with status code 200, with default header values.

Wallet
*/
type GetWalletOK struct {
	Payload *model.Wallet
}

func (o *GetWalletOK) Error() string {
	return fmt.Sprintf("[GET /wallet/{walletId}][%d] getWalletOK  %+v", 200, o.Payload)
}
func (o *GetWalletOK) GetPayload() *model.Wallet {
	return o.Payload
}

func (o *GetWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Wallet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWalletBadRequest creates a GetWalletBadRequest with default headers values
func NewGetWalletBadRequest() *GetWalletBadRequest {
	return &GetWalletBadRequest{}
}

/* GetWalletBadRequest describes a response with status code 400, with default header values.

Validation error / Not found
*/
type GetWalletBadRequest struct {
	Payload *model.ErrorResponse
}

func (o *GetWalletBadRequest) Error() string {
	return fmt.Sprintf("[GET /wallet/{walletId}][%d] getWalletBadRequest  %+v", 400, o.Payload)
}
func (o *GetWalletBadRequest) GetPayload() *model.ErrorResponse {
	return o.Payload
}

func (o *GetWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
