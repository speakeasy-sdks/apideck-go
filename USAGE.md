<!-- Start SDK Example Usage [usage] -->
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

	operationSecurity := operations.ApplicantsAddSecurity{
		APIKey: "<your-apideck-api-key>",
	}

	ctx := context.Background()
	res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
		Applicant: shared.ApplicantInput{
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
			Birthday:      types.MustNewDateFromString("2000-08-12"),
			Confidential:  apideckgo.Bool(false),
			CoordinatorID: apideckgo.String("12345"),
			CoverLetter:   apideckgo.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
			Deleted:       apideckgo.Bool(true),
			FirstName:     apideckgo.String("Elon"),
			Followers: []string{
				"a0d636c6-43b3-4bde-8c70-85b707d992f4",
				"a98lfd96-43b3-4bde-8c70-85b707d992e6",
			},
			Headline:    apideckgo.String("PepsiCo, Inc, Central Perk"),
			Initials:    apideckgo.String("EM"),
			LastName:    apideckgo.String("Musk"),
			MiddleName:  apideckgo.String("D."),
			Name:        apideckgo.String("Elon Musk"),
			OwnerID:     apideckgo.String("54321"),
			PhotoURL:    apideckgo.String("https://unavatar.io/elon-musk"),
			PositionID:  apideckgo.String("123"),
			RecordURL:   apideckgo.String("https://app.intercom.io/contacts/12345"),
			RecruiterID: apideckgo.String("12345"),
			Sources: []string{
				"Job site",
			},
			StageID: apideckgo.String("12345"),
			Tags: []string{
				"New",
			},
			Title: apideckgo.String("CEO"),
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
<!-- End SDK Example Usage [usage] -->