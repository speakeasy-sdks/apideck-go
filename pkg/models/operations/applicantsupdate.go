// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/utils"
	"net/http"
)

type ApplicantsUpdateSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *ApplicantsUpdateSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type ApplicantsUpdateRequest struct {
	Applicant shared.ApplicantInput `request:"mediaType=application/json"`
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// The ID of your Unify application
	XApideckAppID string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// ID of the consumer which you want to get or push data from
	XApideckConsumerID string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	XApideckServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
}

func (a ApplicantsUpdateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicantsUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicantsUpdateRequest) GetApplicant() shared.ApplicantInput {
	if o == nil {
		return shared.ApplicantInput{}
	}
	return o.Applicant
}

func (o *ApplicantsUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ApplicantsUpdateRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ApplicantsUpdateRequest) GetXApideckAppID() string {
	if o == nil {
		return ""
	}
	return o.XApideckAppID
}

func (o *ApplicantsUpdateRequest) GetXApideckConsumerID() string {
	if o == nil {
		return ""
	}
	return o.XApideckConsumerID
}

func (o *ApplicantsUpdateRequest) GetXApideckServiceID() *string {
	if o == nil {
		return nil
	}
	return o.XApideckServiceID
}

type ApplicantsUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected error
	UnexpectedErrorResponse *shared.UnexpectedErrorResponse
	// Applicants
	UpdateApplicantResponse *shared.UpdateApplicantResponse
}

func (o *ApplicantsUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ApplicantsUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ApplicantsUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ApplicantsUpdateResponse) GetUnexpectedErrorResponse() *shared.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}

func (o *ApplicantsUpdateResponse) GetUpdateApplicantResponse() *shared.UpdateApplicantResponse {
	if o == nil {
		return nil
	}
	return o.UpdateApplicantResponse
}
