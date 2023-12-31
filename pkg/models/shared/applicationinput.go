// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/apideck-go/pkg/utils"
	"time"
)

type Stage struct {
	// Stage the candidate should be in. If omitted, the default stage for this job will be used.
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *Stage) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Stage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type Status string

const (
	StatusOpen      Status = "open"
	StatusRejected  Status = "rejected"
	StatusHired     Status = "hired"
	StatusConverted Status = "converted"
	StatusOther     Status = "other"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		fallthrough
	case "rejected":
		fallthrough
	case "hired":
		fallthrough
	case "converted":
		fallthrough
	case "other":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type ApplicationInput struct {
	ApplicantID *string `json:"applicant_id"`
	JobID       *string `json:"job_id"`
	Stage       *Stage  `json:"stage,omitempty"`
	Status      *Status `json:"status,omitempty"`
}

func (o *ApplicationInput) GetApplicantID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicantID
}

func (o *ApplicationInput) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *ApplicationInput) GetStage() *Stage {
	if o == nil {
		return nil
	}
	return o.Stage
}

func (o *ApplicationInput) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type Application struct {
	ApplicantID *string `json:"applicant_id"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The user who created the object.
	CreatedBy *string `json:"created_by,omitempty"`
	// A unique identifier for an object.
	ID     *string `json:"id,omitempty"`
	JobID  *string `json:"job_id"`
	Stage  *Stage  `json:"stage,omitempty"`
	Status *Status `json:"status,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

func (a Application) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Application) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Application) GetApplicantID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicantID
}

func (o *Application) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Application) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Application) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Application) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *Application) GetStage() *Stage {
	if o == nil {
		return nil
	}
	return o.Stage
}

func (o *Application) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Application) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Application) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}
