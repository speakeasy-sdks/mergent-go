// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/mergent-go/pkg/models/shared"
	"net/http"
)

type CreateScheduleResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Unprocessable Entity
	Error *shared.Error
	// OK
	Schedule *shared.Schedule
}
