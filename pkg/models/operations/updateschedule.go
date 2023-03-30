// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/mergent-go/pkg/models/shared"
	"net/http"
)

type UpdateScheduleRequest struct {
	ScheduleInput shared.ScheduleInput `request:"mediaType=application/json"`
	// Schedule ID
	ScheduleID string `pathParam:"style=simple,explode=false,name=schedule_id"`
}

type UpdateScheduleResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Not Found
	Error *shared.Error
	// OK
	Schedule *shared.Schedule
}
