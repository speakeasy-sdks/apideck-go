# github.com/speakeasy-sdks/apideck-sample-sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/apideck-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/apideck-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
# SDK Installation

```bash
go get github.com/speakeasy-sdks/apideck-sample-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	apidecksamplesdk "github.com/speakeasy-sdks/apideck-sample-sdk"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/types"
	"log"
)

func main() {
	s := apidecksamplesdk.New()

	operationSecurity := ""

	ctx := context.Background()
	res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
		ApplicantInput: shared.ApplicantInput{
			Addresses: []shared.Address{
				shared.Address{
					City:         apidecksamplesdk.String("San Francisco"),
					ContactName:  apidecksamplesdk.String("Elon Musk"),
					Country:      apidecksamplesdk.String("US"),
					County:       apidecksamplesdk.String("Santa Clara"),
					Email:        apidecksamplesdk.String("elon@musk.com"),
					Fax:          apidecksamplesdk.String("122-111-1111"),
					ID:           apidecksamplesdk.String("123"),
					Latitude:     apidecksamplesdk.String("40.759211"),
					Line1:        apidecksamplesdk.String("Main street"),
					Line2:        apidecksamplesdk.String("apt #"),
					Line3:        apidecksamplesdk.String("Suite #"),
					Line4:        apidecksamplesdk.String("delivery instructions"),
					Longitude:    apidecksamplesdk.String("-73.984638"),
					Name:         apidecksamplesdk.String("HQ US"),
					Notes:        apidecksamplesdk.String("Address notes or delivery instructions."),
					PhoneNumber:  apidecksamplesdk.String("111-111-1111"),
					PostalCode:   apidecksamplesdk.String("94104"),
					RowVersion:   apidecksamplesdk.String("1-12345"),
					Salutation:   apidecksamplesdk.String("Mr"),
					State:        apidecksamplesdk.String("CA"),
					StreetNumber: apidecksamplesdk.String("25"),
					String:       apidecksamplesdk.String("25 Spring Street, Blackburn, VIC 3130"),
					Type:         shared.AddressTypePrimary.ToPointer(),
					Website:      apidecksamplesdk.String("https://elonmusk.com"),
				},
			},
			Anonymized: apidecksamplesdk.Bool(true),
			ApplicationIds: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Applications: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Archived:      apidecksamplesdk.Bool(false),
			Birthday:      types.MustDateFromString("2000-08-12"),
			Confidential:  apidecksamplesdk.Bool(false),
			CoordinatorID: apidecksamplesdk.String("12345"),
			CoverLetter:   apidecksamplesdk.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
			CustomFields: []shared.CustomField{
				shared.CustomField{
					Description: apidecksamplesdk.String("Employee Level"),
					ID:          "2389328923893298",
					Name:        apidecksamplesdk.String("employee_level"),
					Value: shared.CreateCustomFieldValueBoolean(
						true,
					),
				},
			},
			Deleted: apidecksamplesdk.Bool(true),
			Emails: []shared.Email{
				shared.Email{
					Email: "elon@musk.com",
					ID:    apidecksamplesdk.String("123"),
					Type:  shared.EmailTypePrimary.ToPointer(),
				},
			},
			FirstName: apidecksamplesdk.String("Elon"),
			Followers: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Headline:   apidecksamplesdk.String("PepsiCo, Inc, Central Perk"),
			Initials:   apidecksamplesdk.String("EM"),
			LastName:   apidecksamplesdk.String("Musk"),
			MiddleName: apidecksamplesdk.String("D."),
			Name:       apidecksamplesdk.String("Elon Musk"),
			OwnerID:    apidecksamplesdk.String("54321"),
			PhoneNumbers: []shared.PhoneNumber{
				shared.PhoneNumber{
					AreaCode:    apidecksamplesdk.String("323"),
					CountryCode: apidecksamplesdk.String("1"),
					Extension:   apidecksamplesdk.String("105"),
					ID:          apidecksamplesdk.String("12345"),
					Number:      "111-111-1111",
					Type:        shared.PhoneNumberTypePrimary.ToPointer(),
				},
			},
			PhotoURL:    apidecksamplesdk.String("https://unavatar.io/elon-musk"),
			PositionID:  apidecksamplesdk.String("123"),
			RecordURL:   apidecksamplesdk.String("https://app.intercom.io/contacts/12345"),
			RecruiterID: apidecksamplesdk.String("12345"),
			SocialLinks: []shared.ApplicantSocialLinks{
				shared.ApplicantSocialLinks{
					ID:   apidecksamplesdk.String("12345"),
					Type: apidecksamplesdk.String("twitter"),
					URL:  "https://www.twitter.com/apideck",
				},
			},
			Sources: []string{
				"Job site",
			},
			StageID: apidecksamplesdk.String("12345"),
			Tags: []string{
				"New",
			},
			Title: apidecksamplesdk.String("CEO"),
			Websites: []shared.ApplicantWebsites{
				shared.ApplicantWebsites{
					ID:   apidecksamplesdk.String("12345"),
					Type: shared.ApplicantWebsitesTypePrimary.ToPointer(),
					URL:  "http://example.com",
				},
			},
		},
		XApideckAppID:      "Small",
		XApideckConsumerID: "West",
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateApplicantResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
# Available Resources and Operations



## [Ats.Applicants](docs/sdks/atsapplicants/README.md)

* [Add](docs/sdks/atsapplicants/README.md#add) - Create Applicant
* [All](docs/sdks/atsapplicants/README.md#all) - List Applicants
* [Delete](docs/sdks/atsapplicants/README.md#delete) - Delete Applicant
* [One](docs/sdks/atsapplicants/README.md#one) - Get Applicant
* [Update](docs/sdks/atsapplicants/README.md#update) - Update Applicant

## [Ats.Applications](docs/sdks/atsapplications/README.md)

* [Add](docs/sdks/atsapplications/README.md#add) - Create Application
* [All](docs/sdks/atsapplications/README.md#all) - List Applications
* [Delete](docs/sdks/atsapplications/README.md#delete) - Delete Application
* [One](docs/sdks/atsapplications/README.md#one) - Get Application
* [Update](docs/sdks/atsapplications/README.md#update) - Update Application

## [Ats.Jobs](docs/sdks/atsjobs/README.md)

* [All](docs/sdks/atsjobs/README.md#all) - List Jobs
* [One](docs/sdks/atsjobs/README.md#one) - Get Job
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->



<!-- End Dev Containers -->

<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
