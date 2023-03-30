// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/mergent-go/pkg/models/shared"
	"net/http"
)

type ListTasksResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Not Found
	Error *shared.Error
	// OK
	Tasks []shared.TaskOutput
}
