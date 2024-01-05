// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/apideck-go/pkg/types"
	"github.com/speakeasy-sdks/apideck-go/pkg/utils"
	"time"
)

type SocialLinks struct {
	// Unique identifier of the social link
	ID *string `json:"id,omitempty"`
	// Type of the social link, e.g. twitter
	Type *string `json:"type,omitempty"`
	// URL of the social link, e.g. https://www.twitter.com/apideck
	URL string `json:"url"`
}

func (o *SocialLinks) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SocialLinks) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SocialLinks) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

// ApplicantType - The type of website
type ApplicantType string

const (
	ApplicantTypePrimary   ApplicantType = "primary"
	ApplicantTypeSecondary ApplicantType = "secondary"
	ApplicantTypeWork      ApplicantType = "work"
	ApplicantTypePersonal  ApplicantType = "personal"
	ApplicantTypeOther     ApplicantType = "other"
)

func (e ApplicantType) ToPointer() *ApplicantType {
	return &e
}

func (e *ApplicantType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "primary":
		fallthrough
	case "secondary":
		fallthrough
	case "work":
		fallthrough
	case "personal":
		fallthrough
	case "other":
		*e = ApplicantType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApplicantType: %v", v)
	}
}

type Websites struct {
	// Unique identifier for the website
	ID *string `json:"id,omitempty"`
	// The type of website
	Type *ApplicantType `json:"type,omitempty"`
	// The website URL
	URL string `json:"url"`
}

func (o *Websites) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Websites) GetType() *ApplicantType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Websites) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type ApplicantInput struct {
	Addresses      []Address `json:"addresses,omitempty"`
	Anonymized     *bool     `json:"anonymized,omitempty"`
	ApplicationIds []string  `json:"application_ids,omitempty"`
	Applications   []string  `json:"applications,omitempty"`
	Archived       *bool     `json:"archived,omitempty"`
	// The date of birth of the person.
	Birthday      *types.Date   `json:"birthday,omitempty"`
	Confidential  *bool         `json:"confidential,omitempty"`
	CoordinatorID *string       `json:"coordinator_id,omitempty"`
	CoverLetter   *string       `json:"cover_letter,omitempty"`
	CustomFields  []CustomField `json:"custom_fields,omitempty"`
	// Flag to indicate if the object is deleted.
	Deleted *bool   `json:"deleted,omitempty"`
	Emails  []Email `json:"emails,omitempty"`
	// The first name of the person.
	FirstName *string  `json:"first_name,omitempty"`
	Followers []string `json:"followers,omitempty"`
	// Typically a list of previous companies where the contact has worked or schools that the contact has attended
	Headline *string `json:"headline,omitempty"`
	// The initials of the person, usually derived from their first, middle, and last names.
	Initials *string `json:"initials,omitempty"`
	// The last name of the person.
	LastName *string `json:"last_name,omitempty"`
	// Middle name of the person.
	MiddleName *string `json:"middle_name,omitempty"`
	// The name of an applicant.
	Name         *string       `json:"name,omitempty"`
	OwnerID      *string       `json:"owner_id,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	// The URL of the photo of a person.
	PhotoURL *string `json:"photo_url,omitempty"`
	// The PositionId the applicant applied for.
	PositionID  *string       `json:"position_id,omitempty"`
	RecordURL   *string       `json:"record_url,omitempty"`
	RecruiterID *string       `json:"recruiter_id,omitempty"`
	SocialLinks []SocialLinks `json:"social_links,omitempty"`
	Sources     []string      `json:"sources,omitempty"`
	StageID     *string       `json:"stage_id,omitempty"`
	Tags        []string      `json:"tags,omitempty"`
	// The job title of the person.
	Title    *string    `json:"title,omitempty"`
	Websites []Websites `json:"websites,omitempty"`
}

func (a ApplicantInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicantInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicantInput) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *ApplicantInput) GetAnonymized() *bool {
	if o == nil {
		return nil
	}
	return o.Anonymized
}

func (o *ApplicantInput) GetApplicationIds() []string {
	if o == nil {
		return nil
	}
	return o.ApplicationIds
}

func (o *ApplicantInput) GetApplications() []string {
	if o == nil {
		return nil
	}
	return o.Applications
}

func (o *ApplicantInput) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *ApplicantInput) GetBirthday() *types.Date {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *ApplicantInput) GetConfidential() *bool {
	if o == nil {
		return nil
	}
	return o.Confidential
}

func (o *ApplicantInput) GetCoordinatorID() *string {
	if o == nil {
		return nil
	}
	return o.CoordinatorID
}

func (o *ApplicantInput) GetCoverLetter() *string {
	if o == nil {
		return nil
	}
	return o.CoverLetter
}

func (o *ApplicantInput) GetCustomFields() []CustomField {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *ApplicantInput) GetDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.Deleted
}

func (o *ApplicantInput) GetEmails() []Email {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *ApplicantInput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *ApplicantInput) GetFollowers() []string {
	if o == nil {
		return nil
	}
	return o.Followers
}

func (o *ApplicantInput) GetHeadline() *string {
	if o == nil {
		return nil
	}
	return o.Headline
}

func (o *ApplicantInput) GetInitials() *string {
	if o == nil {
		return nil
	}
	return o.Initials
}

func (o *ApplicantInput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *ApplicantInput) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *ApplicantInput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ApplicantInput) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *ApplicantInput) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *ApplicantInput) GetPhotoURL() *string {
	if o == nil {
		return nil
	}
	return o.PhotoURL
}

func (o *ApplicantInput) GetPositionID() *string {
	if o == nil {
		return nil
	}
	return o.PositionID
}

func (o *ApplicantInput) GetRecordURL() *string {
	if o == nil {
		return nil
	}
	return o.RecordURL
}

func (o *ApplicantInput) GetRecruiterID() *string {
	if o == nil {
		return nil
	}
	return o.RecruiterID
}

func (o *ApplicantInput) GetSocialLinks() []SocialLinks {
	if o == nil {
		return nil
	}
	return o.SocialLinks
}

func (o *ApplicantInput) GetSources() []string {
	if o == nil {
		return nil
	}
	return o.Sources
}

func (o *ApplicantInput) GetStageID() *string {
	if o == nil {
		return nil
	}
	return o.StageID
}

func (o *ApplicantInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ApplicantInput) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ApplicantInput) GetWebsites() []Websites {
	if o == nil {
		return nil
	}
	return o.Websites
}

type Applicant struct {
	Addresses      []Address `json:"addresses,omitempty"`
	Anonymized     *bool     `json:"anonymized,omitempty"`
	ApplicationIds []string  `json:"application_ids,omitempty"`
	Applications   []string  `json:"applications,omitempty"`
	Archived       *bool     `json:"archived,omitempty"`
	// The date of birth of the person.
	Birthday      *types.Date `json:"birthday,omitempty"`
	Confidential  *bool       `json:"confidential,omitempty"`
	CoordinatorID *string     `json:"coordinator_id,omitempty"`
	CoverLetter   *string     `json:"cover_letter,omitempty"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The user who created the object.
	CreatedBy    *string       `json:"created_by,omitempty"`
	CustomFields []CustomField `json:"custom_fields,omitempty"`
	CvURL        *string       `json:"cv_url,omitempty"`
	// Flag to indicate if the object is deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// The time at which the object was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The user who deleted the object.
	DeletedBy *string `json:"deleted_by,omitempty"`
	Emails    []Email `json:"emails,omitempty"`
	// The first name of the person.
	FirstName *string  `json:"first_name,omitempty"`
	Followers []string `json:"followers,omitempty"`
	// Typically a list of previous companies where the contact has worked or schools that the contact has attended
	Headline *string `json:"headline,omitempty"`
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// The initials of the person, usually derived from their first, middle, and last names.
	Initials          *string    `json:"initials,omitempty"`
	JobURL            *string    `json:"job_url,omitempty"`
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
	// The last name of the person.
	LastName *string `json:"last_name,omitempty"`
	// Middle name of the person.
	MiddleName *string `json:"middle_name,omitempty"`
	// The name of an applicant.
	Name         *string       `json:"name,omitempty"`
	OwnerID      *string       `json:"owner_id,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	// The URL of the photo of a person.
	PhotoURL *string `json:"photo_url,omitempty"`
	// The PositionId the applicant applied for.
	PositionID  *string       `json:"position_id,omitempty"`
	RecordURL   *string       `json:"record_url,omitempty"`
	RecruiterID *string       `json:"recruiter_id,omitempty"`
	RejectedAt  *time.Time    `json:"rejected_at,omitempty"`
	SocialLinks []SocialLinks `json:"social_links,omitempty"`
	SourceID    *string       `json:"source_id,omitempty"`
	SourcedBy   *string       `json:"sourced_by,omitempty"`
	Sources     []string      `json:"sources,omitempty"`
	StageID     *string       `json:"stage_id,omitempty"`
	Tags        []string      `json:"tags,omitempty"`
	// The job title of the person.
	Title *string `json:"title,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string    `json:"updated_by,omitempty"`
	Websites  []Websites `json:"websites,omitempty"`
}

func (a Applicant) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Applicant) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Applicant) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *Applicant) GetAnonymized() *bool {
	if o == nil {
		return nil
	}
	return o.Anonymized
}

func (o *Applicant) GetApplicationIds() []string {
	if o == nil {
		return nil
	}
	return o.ApplicationIds
}

func (o *Applicant) GetApplications() []string {
	if o == nil {
		return nil
	}
	return o.Applications
}

func (o *Applicant) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *Applicant) GetBirthday() *types.Date {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *Applicant) GetConfidential() *bool {
	if o == nil {
		return nil
	}
	return o.Confidential
}

func (o *Applicant) GetCoordinatorID() *string {
	if o == nil {
		return nil
	}
	return o.CoordinatorID
}

func (o *Applicant) GetCoverLetter() *string {
	if o == nil {
		return nil
	}
	return o.CoverLetter
}

func (o *Applicant) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Applicant) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Applicant) GetCustomFields() []CustomField {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *Applicant) GetCvURL() *string {
	if o == nil {
		return nil
	}
	return o.CvURL
}

func (o *Applicant) GetDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.Deleted
}

func (o *Applicant) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Applicant) GetDeletedBy() *string {
	if o == nil {
		return nil
	}
	return o.DeletedBy
}

func (o *Applicant) GetEmails() []Email {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *Applicant) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *Applicant) GetFollowers() []string {
	if o == nil {
		return nil
	}
	return o.Followers
}

func (o *Applicant) GetHeadline() *string {
	if o == nil {
		return nil
	}
	return o.Headline
}

func (o *Applicant) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Applicant) GetInitials() *string {
	if o == nil {
		return nil
	}
	return o.Initials
}

func (o *Applicant) GetJobURL() *string {
	if o == nil {
		return nil
	}
	return o.JobURL
}

func (o *Applicant) GetLastInteractionAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastInteractionAt
}

func (o *Applicant) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *Applicant) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *Applicant) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Applicant) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *Applicant) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *Applicant) GetPhotoURL() *string {
	if o == nil {
		return nil
	}
	return o.PhotoURL
}

func (o *Applicant) GetPositionID() *string {
	if o == nil {
		return nil
	}
	return o.PositionID
}

func (o *Applicant) GetRecordURL() *string {
	if o == nil {
		return nil
	}
	return o.RecordURL
}

func (o *Applicant) GetRecruiterID() *string {
	if o == nil {
		return nil
	}
	return o.RecruiterID
}

func (o *Applicant) GetRejectedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RejectedAt
}

func (o *Applicant) GetSocialLinks() []SocialLinks {
	if o == nil {
		return nil
	}
	return o.SocialLinks
}

func (o *Applicant) GetSourceID() *string {
	if o == nil {
		return nil
	}
	return o.SourceID
}

func (o *Applicant) GetSourcedBy() *string {
	if o == nil {
		return nil
	}
	return o.SourcedBy
}

func (o *Applicant) GetSources() []string {
	if o == nil {
		return nil
	}
	return o.Sources
}

func (o *Applicant) GetStageID() *string {
	if o == nil {
		return nil
	}
	return o.StageID
}

func (o *Applicant) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Applicant) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Applicant) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Applicant) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *Applicant) GetWebsites() []Websites {
	if o == nil {
		return nil
	}
	return o.Websites
}