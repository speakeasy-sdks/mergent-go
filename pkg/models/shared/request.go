// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Request struct {
	// The HTTP request body as a string.
	Body *string `json:"body,omitempty"`
	// The headers that will accompany any Task's HTTP request. For
	// example, you can use this to set Content-Type to "application/json"
	// or "application/octet-stream".
	//
	Headers map[string]interface{} `json:"headers,omitempty"`
	// The URL that the POST request will be sent to.
	//
	// For localhost development, use something like ngrok to get a publicly
	// accessible URL for your local service. See https://docs.mergent.co for
	// more info.
	//
	URL string `json:"url"`
}
