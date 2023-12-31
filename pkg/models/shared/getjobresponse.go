// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetJobResponse struct {
	Data Job `json:"data"`
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

func (o *GetJobResponse) GetData() Job {
	if o == nil {
		return Job{}
	}
	return o.Data
}

func (o *GetJobResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetJobResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetJobResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetJobResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetJobResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}
