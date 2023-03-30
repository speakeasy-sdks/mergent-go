// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/mergent-go/pkg/models/shared"
	"net/http"
)

type CreateTaskRequest struct {
	Request shared.TaskNewInput `request:"mediaType=application/json"`
}

type CreateTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Conflict
	Error *shared.Error
	// Created
	Task *shared.TaskOutput
}