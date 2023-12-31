// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetApplicantsResponse struct {
	Data []Applicant `json:"data"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Operation performed
	Operation string `json:"operation"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// HTTP Response Status
	Status string `json:"status"`
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
}

func (o *GetApplicantsResponse) GetData() []Applicant {
	if o == nil {
		return []Applicant{}
	}
	return o.Data
}

func (o *GetApplicantsResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *GetApplicantsResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetApplicantsResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetApplicantsResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetApplicantsResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetApplicantsResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetApplicantsResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}
