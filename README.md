# apideck Go SDK 

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/apideck-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/apideck-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/apideck-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/types"
	"log"
)

func main() {
	s := apideckgo.New()

	operationSecurity := "<your-apideck-api-key>"

	ctx := context.Background()
	res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
		ApplicantInput: shared.ApplicantInput{
			Addresses: []shared.Address{
				shared.Address{
					City:         apideckgo.String("San Francisco"),
					ContactName:  apideckgo.String("Elon Musk"),
					Country:      apideckgo.String("US"),
					County:       apideckgo.String("Santa Clara"),
					Email:        apideckgo.String("elon@musk.com"),
					Fax:          apideckgo.String("122-111-1111"),
					ID:           apideckgo.String("123"),
					Latitude:     apideckgo.String("40.759211"),
					Line1:        apideckgo.String("Main street"),
					Line2:        apideckgo.String("apt #"),
					Line3:        apideckgo.String("Suite #"),
					Line4:        apideckgo.String("delivery instructions"),
					Longitude:    apideckgo.String("-73.984638"),
					Name:         apideckgo.String("HQ US"),
					Notes:        apideckgo.String("Address notes or delivery instructions."),
					PhoneNumber:  apideckgo.String("111-111-1111"),
					PostalCode:   apideckgo.String("94104"),
					RowVersion:   apideckgo.String("1-12345"),
					Salutation:   apideckgo.String("Mr"),
					State:        apideckgo.String("CA"),
					StreetNumber: apideckgo.String("25"),
					String:       apideckgo.String("25 Spring Street, Blackburn, VIC 3130"),
					Type:         shared.AddressTypePrimary.ToPointer(),
					Website:      apideckgo.String("https://elonmusk.com"),
				},
			},
			Anonymized: apideckgo.Bool(true),
			ApplicationIds: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Applications: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Archived:      apideckgo.Bool(false),
			Birthday:      types.MustDateFromString("2000-08-12"),
			Confidential:  apideckgo.Bool(false),
			CoordinatorID: apideckgo.String("12345"),
			CoverLetter:   apideckgo.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
			CustomFields: []shared.CustomField{
				shared.CustomField{
					Description: apideckgo.String("Employee Level"),
					ID:          "2389328923893298",
					Name:        apideckgo.String("employee_level"),
					Value: shared.CreateCustomFieldValueBoolean(
						true,
					),
				},
			},
			Deleted: apideckgo.Bool(true),
			Emails: []shared.Email{
				shared.Email{
					Email: "elon@musk.com",
					ID:    apideckgo.String("123"),
					Type:  shared.EmailTypePrimary.ToPointer(),
				},
			},
			FirstName: apideckgo.String("Elon"),
			Followers: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Headline:   apideckgo.String("PepsiCo, Inc, Central Perk"),
			Initials:   apideckgo.String("EM"),
			LastName:   apideckgo.String("Musk"),
			MiddleName: apideckgo.String("D."),
			Name:       apideckgo.String("Elon Musk"),
			OwnerID:    apideckgo.String("54321"),
			PhoneNumbers: []shared.PhoneNumber{
				shared.PhoneNumber{
					AreaCode:    apideckgo.String("323"),
					CountryCode: apideckgo.String("1"),
					Extension:   apideckgo.String("105"),
					ID:          apideckgo.String("12345"),
					Number:      "111-111-1111",
					Type:        shared.PhoneNumberTypePrimary.ToPointer(),
				},
			},
			PhotoURL:    apideckgo.String("https://unavatar.io/elon-musk"),
			PositionID:  apideckgo.String("123"),
			RecordURL:   apideckgo.String("https://app.intercom.io/contacts/12345"),
			RecruiterID: apideckgo.String("12345"),
			SocialLinks: []shared.ApplicantSocialLinks{
				shared.ApplicantSocialLinks{
					ID:   apideckgo.String("12345"),
					Type: apideckgo.String("twitter"),
					URL:  "https://www.twitter.com/apideck",
				},
			},
			Sources: []string{
				"Job site",
			},
			StageID: apideckgo.String("12345"),
			Tags: []string{
				"New",
			},
			Title: apideckgo.String("CEO"),
			Websites: []shared.ApplicantWebsites{
				shared.ApplicantWebsites{
					ID:   apideckgo.String("12345"),
					Type: shared.ApplicantWebsitesTypePrimary.ToPointer(),
					URL:  "http://example.com",
				},
			},
		},
		XApideckAppID:      "string",
		XApideckConsumerID: "string",
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
## Available Resources and Operations



### [Ats.Applicants](docs/sdks/atsapplicants/README.md)

* [Add](docs/sdks/atsapplicants/README.md#add) - Create Applicant
* [All](docs/sdks/atsapplicants/README.md#all) - List Applicants
* [Delete](docs/sdks/atsapplicants/README.md#delete) - Delete Applicant
* [One](docs/sdks/atsapplicants/README.md#one) - Get Applicant
* [Update](docs/sdks/atsapplicants/README.md#update) - Update Applicant

### [Ats.Applications](docs/sdks/atsapplications/README.md)

* [Add](docs/sdks/atsapplications/README.md#add) - Create Application
* [All](docs/sdks/atsapplications/README.md#all) - List Applications
* [Delete](docs/sdks/atsapplications/README.md#delete) - Delete Application
* [One](docs/sdks/atsapplications/README.md#one) - Get Application
* [Update](docs/sdks/atsapplications/README.md#update) - Update Application

### [Ats.Jobs](docs/sdks/atsjobs/README.md)

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



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


## Example

```go
package main

import (
	"context"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/types"
	"log"
)

func main() {
	s := apideckgo.New()

	operationSecurity := "<your-apideck-api-key>"

	ctx := context.Background()
	res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
		ApplicantInput: shared.ApplicantInput{
			Addresses: []shared.Address{
				shared.Address{
					City:         apideckgo.String("San Francisco"),
					ContactName:  apideckgo.String("Elon Musk"),
					Country:      apideckgo.String("US"),
					County:       apideckgo.String("Santa Clara"),
					Email:        apideckgo.String("elon@musk.com"),
					Fax:          apideckgo.String("122-111-1111"),
					ID:           apideckgo.String("123"),
					Latitude:     apideckgo.String("40.759211"),
					Line1:        apideckgo.String("Main street"),
					Line2:        apideckgo.String("apt #"),
					Line3:        apideckgo.String("Suite #"),
					Line4:        apideckgo.String("delivery instructions"),
					Longitude:    apideckgo.String("-73.984638"),
					Name:         apideckgo.String("HQ US"),
					Notes:        apideckgo.String("Address notes or delivery instructions."),
					PhoneNumber:  apideckgo.String("111-111-1111"),
					PostalCode:   apideckgo.String("94104"),
					RowVersion:   apideckgo.String("1-12345"),
					Salutation:   apideckgo.String("Mr"),
					State:        apideckgo.String("CA"),
					StreetNumber: apideckgo.String("25"),
					String:       apideckgo.String("25 Spring Street, Blackburn, VIC 3130"),
					Type:         shared.AddressTypePrimary.ToPointer(),
					Website:      apideckgo.String("https://elonmusk.com"),
				},
			},
			Anonymized: apideckgo.Bool(true),
			ApplicationIds: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Applications: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Archived:      apideckgo.Bool(false),
			Birthday:      types.MustDateFromString("2000-08-12"),
			Confidential:  apideckgo.Bool(false),
			CoordinatorID: apideckgo.String("12345"),
			CoverLetter:   apideckgo.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
			CustomFields: []shared.CustomField{
				shared.CustomField{
					Description: apideckgo.String("Employee Level"),
					ID:          "2389328923893298",
					Name:        apideckgo.String("employee_level"),
					Value: shared.CreateCustomFieldValueBoolean(
						true,
					),
				},
			},
			Deleted: apideckgo.Bool(true),
			Emails: []shared.Email{
				shared.Email{
					Email: "elon@musk.com",
					ID:    apideckgo.String("123"),
					Type:  shared.EmailTypePrimary.ToPointer(),
				},
			},
			FirstName: apideckgo.String("Elon"),
			Followers: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Headline:   apideckgo.String("PepsiCo, Inc, Central Perk"),
			Initials:   apideckgo.String("EM"),
			LastName:   apideckgo.String("Musk"),
			MiddleName: apideckgo.String("D."),
			Name:       apideckgo.String("Elon Musk"),
			OwnerID:    apideckgo.String("54321"),
			PhoneNumbers: []shared.PhoneNumber{
				shared.PhoneNumber{
					AreaCode:    apideckgo.String("323"),
					CountryCode: apideckgo.String("1"),
					Extension:   apideckgo.String("105"),
					ID:          apideckgo.String("12345"),
					Number:      "111-111-1111",
					Type:        shared.PhoneNumberTypePrimary.ToPointer(),
				},
			},
			PhotoURL:    apideckgo.String("https://unavatar.io/elon-musk"),
			PositionID:  apideckgo.String("123"),
			RecordURL:   apideckgo.String("https://app.intercom.io/contacts/12345"),
			RecruiterID: apideckgo.String("12345"),
			SocialLinks: []shared.ApplicantSocialLinks{
				shared.ApplicantSocialLinks{
					ID:   apideckgo.String("12345"),
					Type: apideckgo.String("twitter"),
					URL:  "https://www.twitter.com/apideck",
				},
			},
			Sources: []string{
				"Job site",
			},
			StageID: apideckgo.String("12345"),
			Tags: []string{
				"New",
			},
			Title: apideckgo.String("CEO"),
			Websites: []shared.ApplicantWebsites{
				shared.ApplicantWebsites{
					ID:   apideckgo.String("12345"),
					Type: shared.ApplicantWebsitesTypePrimary.ToPointer(),
					URL:  "http://example.com",
				},
			},
		},
		XApideckAppID:      "string",
		XApideckConsumerID: "string",
	}, operationSecurity)
	if err != nil {

		var e *BadRequestResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *UnauthorizedResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *PaymentRequiredResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *NotFoundResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *UnprocessableResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://unify.apideck.com` | None |

For example:


```go
package main

import (
	"context"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/types"
	"log"
)

func main() {
	s := apideckgo.New(
		apideckgo.WithServerIndex(0),
	)

	operationSecurity := "<your-apideck-api-key>"

	ctx := context.Background()
	res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
		ApplicantInput: shared.ApplicantInput{
			Addresses: []shared.Address{
				shared.Address{
					City:         apideckgo.String("San Francisco"),
					ContactName:  apideckgo.String("Elon Musk"),
					Country:      apideckgo.String("US"),
					County:       apideckgo.String("Santa Clara"),
					Email:        apideckgo.String("elon@musk.com"),
					Fax:          apideckgo.String("122-111-1111"),
					ID:           apideckgo.String("123"),
					Latitude:     apideckgo.String("40.759211"),
					Line1:        apideckgo.String("Main street"),
					Line2:        apideckgo.String("apt #"),
					Line3:        apideckgo.String("Suite #"),
					Line4:        apideckgo.String("delivery instructions"),
					Longitude:    apideckgo.String("-73.984638"),
					Name:         apideckgo.String("HQ US"),
					Notes:        apideckgo.String("Address notes or delivery instructions."),
					PhoneNumber:  apideckgo.String("111-111-1111"),
					PostalCode:   apideckgo.String("94104"),
					RowVersion:   apideckgo.String("1-12345"),
					Salutation:   apideckgo.String("Mr"),
					State:        apideckgo.String("CA"),
					StreetNumber: apideckgo.String("25"),
					String:       apideckgo.String("25 Spring Street, Blackburn, VIC 3130"),
					Type:         shared.AddressTypePrimary.ToPointer(),
					Website:      apideckgo.String("https://elonmusk.com"),
				},
			},
			Anonymized: apideckgo.Bool(true),
			ApplicationIds: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Applications: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Archived:      apideckgo.Bool(false),
			Birthday:      types.MustDateFromString("2000-08-12"),
			Confidential:  apideckgo.Bool(false),
			CoordinatorID: apideckgo.String("12345"),
			CoverLetter:   apideckgo.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
			CustomFields: []shared.CustomField{
				shared.CustomField{
					Description: apideckgo.String("Employee Level"),
					ID:          "2389328923893298",
					Name:        apideckgo.String("employee_level"),
					Value: shared.CreateCustomFieldValueBoolean(
						true,
					),
				},
			},
			Deleted: apideckgo.Bool(true),
			Emails: []shared.Email{
				shared.Email{
					Email: "elon@musk.com",
					ID:    apideckgo.String("123"),
					Type:  shared.EmailTypePrimary.ToPointer(),
				},
			},
			FirstName: apideckgo.String("Elon"),
			Followers: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Headline:   apideckgo.String("PepsiCo, Inc, Central Perk"),
			Initials:   apideckgo.String("EM"),
			LastName:   apideckgo.String("Musk"),
			MiddleName: apideckgo.String("D."),
			Name:       apideckgo.String("Elon Musk"),
			OwnerID:    apideckgo.String("54321"),
			PhoneNumbers: []shared.PhoneNumber{
				shared.PhoneNumber{
					AreaCode:    apideckgo.String("323"),
					CountryCode: apideckgo.String("1"),
					Extension:   apideckgo.String("105"),
					ID:          apideckgo.String("12345"),
					Number:      "111-111-1111",
					Type:        shared.PhoneNumberTypePrimary.ToPointer(),
				},
			},
			PhotoURL:    apideckgo.String("https://unavatar.io/elon-musk"),
			PositionID:  apideckgo.String("123"),
			RecordURL:   apideckgo.String("https://app.intercom.io/contacts/12345"),
			RecruiterID: apideckgo.String("12345"),
			SocialLinks: []shared.ApplicantSocialLinks{
				shared.ApplicantSocialLinks{
					ID:   apideckgo.String("12345"),
					Type: apideckgo.String("twitter"),
					URL:  "https://www.twitter.com/apideck",
				},
			},
			Sources: []string{
				"Job site",
			},
			StageID: apideckgo.String("12345"),
			Tags: []string{
				"New",
			},
			Title: apideckgo.String("CEO"),
			Websites: []shared.ApplicantWebsites{
				shared.ApplicantWebsites{
					ID:   apideckgo.String("12345"),
					Type: shared.ApplicantWebsitesTypePrimary.ToPointer(),
					URL:  "http://example.com",
				},
			},
		},
		XApideckAppID:      "string",
		XApideckConsumerID: "string",
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateApplicantResponse != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


```go
package main

import (
	"context"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/types"
	"log"
)

func main() {
	s := apideckgo.New(
		apideckgo.WithServerURL("https://unify.apideck.com"),
	)

	operationSecurity := "<your-apideck-api-key>"

	ctx := context.Background()
	res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
		ApplicantInput: shared.ApplicantInput{
			Addresses: []shared.Address{
				shared.Address{
					City:         apideckgo.String("San Francisco"),
					ContactName:  apideckgo.String("Elon Musk"),
					Country:      apideckgo.String("US"),
					County:       apideckgo.String("Santa Clara"),
					Email:        apideckgo.String("elon@musk.com"),
					Fax:          apideckgo.String("122-111-1111"),
					ID:           apideckgo.String("123"),
					Latitude:     apideckgo.String("40.759211"),
					Line1:        apideckgo.String("Main street"),
					Line2:        apideckgo.String("apt #"),
					Line3:        apideckgo.String("Suite #"),
					Line4:        apideckgo.String("delivery instructions"),
					Longitude:    apideckgo.String("-73.984638"),
					Name:         apideckgo.String("HQ US"),
					Notes:        apideckgo.String("Address notes or delivery instructions."),
					PhoneNumber:  apideckgo.String("111-111-1111"),
					PostalCode:   apideckgo.String("94104"),
					RowVersion:   apideckgo.String("1-12345"),
					Salutation:   apideckgo.String("Mr"),
					State:        apideckgo.String("CA"),
					StreetNumber: apideckgo.String("25"),
					String:       apideckgo.String("25 Spring Street, Blackburn, VIC 3130"),
					Type:         shared.AddressTypePrimary.ToPointer(),
					Website:      apideckgo.String("https://elonmusk.com"),
				},
			},
			Anonymized: apideckgo.Bool(true),
			ApplicationIds: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Applications: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Archived:      apideckgo.Bool(false),
			Birthday:      types.MustDateFromString("2000-08-12"),
			Confidential:  apideckgo.Bool(false),
			CoordinatorID: apideckgo.String("12345"),
			CoverLetter:   apideckgo.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
			CustomFields: []shared.CustomField{
				shared.CustomField{
					Description: apideckgo.String("Employee Level"),
					ID:          "2389328923893298",
					Name:        apideckgo.String("employee_level"),
					Value: shared.CreateCustomFieldValueBoolean(
						true,
					),
				},
			},
			Deleted: apideckgo.Bool(true),
			Emails: []shared.Email{
				shared.Email{
					Email: "elon@musk.com",
					ID:    apideckgo.String("123"),
					Type:  shared.EmailTypePrimary.ToPointer(),
				},
			},
			FirstName: apideckgo.String("Elon"),
			Followers: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Headline:   apideckgo.String("PepsiCo, Inc, Central Perk"),
			Initials:   apideckgo.String("EM"),
			LastName:   apideckgo.String("Musk"),
			MiddleName: apideckgo.String("D."),
			Name:       apideckgo.String("Elon Musk"),
			OwnerID:    apideckgo.String("54321"),
			PhoneNumbers: []shared.PhoneNumber{
				shared.PhoneNumber{
					AreaCode:    apideckgo.String("323"),
					CountryCode: apideckgo.String("1"),
					Extension:   apideckgo.String("105"),
					ID:          apideckgo.String("12345"),
					Number:      "111-111-1111",
					Type:        shared.PhoneNumberTypePrimary.ToPointer(),
				},
			},
			PhotoURL:    apideckgo.String("https://unavatar.io/elon-musk"),
			PositionID:  apideckgo.String("123"),
			RecordURL:   apideckgo.String("https://app.intercom.io/contacts/12345"),
			RecruiterID: apideckgo.String("12345"),
			SocialLinks: []shared.ApplicantSocialLinks{
				shared.ApplicantSocialLinks{
					ID:   apideckgo.String("12345"),
					Type: apideckgo.String("twitter"),
					URL:  "https://www.twitter.com/apideck",
				},
			},
			Sources: []string{
				"Job site",
			},
			StageID: apideckgo.String("12345"),
			Tags: []string{
				"New",
			},
			Title: apideckgo.String("CEO"),
			Websites: []shared.ApplicantWebsites{
				shared.ApplicantWebsites{
					ID:   apideckgo.String("12345"),
					Type: shared.ApplicantWebsitesTypePrimary.ToPointer(),
					URL:  "http://example.com",
				},
			},
		},
		XApideckAppID:      "string",
		XApideckConsumerID: "string",
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateApplicantResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

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
