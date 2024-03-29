// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/utils"
	"net/http"
)

type ApplicantsAllRequest struct {
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded.
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Apply filters
	Filter *shared.ApplicantsFilter `queryParam:"style=deepObject,explode=true,name=filter"`
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

func (a ApplicantsAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicantsAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicantsAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ApplicantsAllRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ApplicantsAllRequest) GetFilter() *shared.ApplicantsFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *ApplicantsAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ApplicantsAllRequest) GetPassThrough() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

func (o *ApplicantsAllRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ApplicantsAllRequest) GetXApideckAppID() string {
	if o == nil {
		return ""
	}
	return o.XApideckAppID
}

func (o *ApplicantsAllRequest) GetXApideckConsumerID() string {
	if o == nil {
		return ""
	}
	return o.XApideckConsumerID
}

func (o *ApplicantsAllRequest) GetXApideckServiceID() *string {
	if o == nil {
		return nil
	}
	return o.XApideckServiceID
}

type ApplicantsAllResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Applicants
	GetApplicantsResponse *shared.GetApplicantsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected error
	UnexpectedErrorResponse *shared.UnexpectedErrorResponse
}

func (o *ApplicantsAllResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ApplicantsAllResponse) GetGetApplicantsResponse() *shared.GetApplicantsResponse {
	if o == nil {
		return nil
	}
	return o.GetApplicantsResponse
}

func (o *ApplicantsAllResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ApplicantsAllResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ApplicantsAllResponse) GetUnexpectedErrorResponse() *shared.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
