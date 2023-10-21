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