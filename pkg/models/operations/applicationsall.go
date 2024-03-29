// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/utils"
	"net/http"
)

type ApplicationsAllRequest struct {
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Optional unmapped key/values that will be passed through to downstream as query parameters. Ie: ?pass_through[search]=leads becomes ?search=leads
	PassThrough map[string]interface{} `queryParam:"style=deepObject,explode=true,name=pass_through"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// The ID of your Unify application
	XApideckAppID string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// ID of the consumer which you want to get or push data from
	XApideckConsumerID string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	XApideckServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
}

func (a ApplicationsAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationsAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationsAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ApplicationsAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ApplicationsAllRequest) GetPassThrough() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

func (o *ApplicationsAllRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ApplicationsAllRequest) GetXApideckAppID() string {
	if o == nil {
		return ""
	}
	return o.XApideckAppID
}

func (o *ApplicationsAllRequest) GetXApideckConsumerID() string {
	if o == nil {
		return ""
	}
	return o.XApideckConsumerID
}

func (o *ApplicationsAllRequest) GetXApideckServiceID() *string {
	if o == nil {
		return nil
	}
	return o.XApideckServiceID
}

type ApplicationsAllResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Applications
	GetApplicationsResponse *shared.GetApplicationsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected error
	UnexpectedErrorResponse *shared.UnexpectedErrorResponse
}

func (o *ApplicationsAllResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ApplicationsAllResponse) GetGetApplicationsResponse() *shared.GetApplicationsResponse {
	if o == nil {
		return nil
	}
	return o.GetApplicationsResponse
}

func (o *ApplicationsAllResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ApplicationsAllResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ApplicationsAllResponse) GetUnexpectedErrorResponse() *shared.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
