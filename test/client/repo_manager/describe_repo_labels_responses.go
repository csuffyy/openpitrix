// Code generated by go-swagger; DO NOT EDIT.

package repo_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeRepoLabelsReader is a Reader for the DescribeRepoLabels structure.
type DescribeRepoLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeRepoLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeRepoLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeRepoLabelsOK creates a DescribeRepoLabelsOK with default headers values
func NewDescribeRepoLabelsOK() *DescribeRepoLabelsOK {
	return &DescribeRepoLabelsOK{}
}

/*DescribeRepoLabelsOK handles this case with default header values.

DescribeRepoLabelsOK describe repo labels o k
*/
type DescribeRepoLabelsOK struct {
	Payload *models.OpenpitrixDescribeRepoLabelsResponse
}

func (o *DescribeRepoLabelsOK) Error() string {
	return fmt.Sprintf("[GET /v1/repos/labels][%d] describeRepoLabelsOK  %+v", 200, o.Payload)
}

func (o *DescribeRepoLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeRepoLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}