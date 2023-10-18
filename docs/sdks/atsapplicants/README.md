# AtsApplicants
(*Ats.Applicants*)

### Available Operations

* [Add](#add) - Create Applicant
* [All](#all) - List Applicants
* [Delete](#delete) - Delete Applicant
* [One](#one) - Get Applicant
* [Update](#update) - Update Applicant

## Add

Create Applicant

### Example Usage

```go
package main

import(
	"context"
	"log"
	apidecksamplesdk "github.com/speakeasy-sdks/apideck-sample-sdk"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/types"
)

func main() {
    s := apidecksamplesdk.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
        ApplicantInput: shared.ApplicantInput{
            Addresses: []shared.Address{
                shared.Address{
                    City: apidecksamplesdk.String("San Francisco"),
                    ContactName: apidecksamplesdk.String("Elon Musk"),
                    Country: apidecksamplesdk.String("US"),
                    County: apidecksamplesdk.String("Santa Clara"),
                    Email: apidecksamplesdk.String("elon@musk.com"),
                    Fax: apidecksamplesdk.String("122-111-1111"),
                    ID: apidecksamplesdk.String("123"),
                    Latitude: apidecksamplesdk.String("40.759211"),
                    Line1: apidecksamplesdk.String("Main street"),
                    Line2: apidecksamplesdk.String("apt #"),
                    Line3: apidecksamplesdk.String("Suite #"),
                    Line4: apidecksamplesdk.String("delivery instructions"),
                    Longitude: apidecksamplesdk.String("-73.984638"),
                    Name: apidecksamplesdk.String("HQ US"),
                    Notes: apidecksamplesdk.String("Address notes or delivery instructions."),
                    PhoneNumber: apidecksamplesdk.String("111-111-1111"),
                    PostalCode: apidecksamplesdk.String("94104"),
                    RowVersion: apidecksamplesdk.String("1-12345"),
                    Salutation: apidecksamplesdk.String("Mr"),
                    State: apidecksamplesdk.String("CA"),
                    StreetNumber: apidecksamplesdk.String("25"),
                    String: apidecksamplesdk.String("25 Spring Street, Blackburn, VIC 3130"),
                    Type: shared.AddressTypePrimary.ToPointer(),
                    Website: apidecksamplesdk.String("https://elonmusk.com"),
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
            Archived: apidecksamplesdk.Bool(false),
            Birthday: types.MustDateFromString("2000-08-12"),
            Confidential: apidecksamplesdk.Bool(false),
            CoordinatorID: apidecksamplesdk.String("12345"),
            CoverLetter: apidecksamplesdk.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
            CustomFields: []shared.CustomField{
                shared.CustomField{
                    Description: apidecksamplesdk.String("Employee Level"),
                    ID: "2389328923893298",
                    Name: apidecksamplesdk.String("employee_level"),
                    Value: shared.CreateCustomFieldValueBoolean(
                    true,
                    ),
                },
            },
            Deleted: apidecksamplesdk.Bool(true),
            Emails: []shared.Email{
                shared.Email{
                    Email: "elon@musk.com",
                    ID: apidecksamplesdk.String("123"),
                    Type: shared.EmailTypePrimary.ToPointer(),
                },
            },
            FirstName: apidecksamplesdk.String("Elon"),
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Headline: apidecksamplesdk.String("PepsiCo, Inc, Central Perk"),
            Initials: apidecksamplesdk.String("EM"),
            LastName: apidecksamplesdk.String("Musk"),
            MiddleName: apidecksamplesdk.String("D."),
            Name: apidecksamplesdk.String("Elon Musk"),
            OwnerID: apidecksamplesdk.String("54321"),
            PhoneNumbers: []shared.PhoneNumber{
                shared.PhoneNumber{
                    AreaCode: apidecksamplesdk.String("323"),
                    CountryCode: apidecksamplesdk.String("1"),
                    Extension: apidecksamplesdk.String("105"),
                    ID: apidecksamplesdk.String("12345"),
                    Number: "111-111-1111",
                    Type: shared.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            PhotoURL: apidecksamplesdk.String("https://unavatar.io/elon-musk"),
            PositionID: apidecksamplesdk.String("123"),
            RecordURL: apidecksamplesdk.String("https://app.intercom.io/contacts/12345"),
            RecruiterID: apidecksamplesdk.String("12345"),
            SocialLinks: []shared.ApplicantSocialLinks{
                shared.ApplicantSocialLinks{
                    ID: apidecksamplesdk.String("12345"),
                    Type: apidecksamplesdk.String("twitter"),
                    URL: "https://www.twitter.com/apideck",
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
                    ID: apidecksamplesdk.String("12345"),
                    Type: shared.ApplicantWebsitesTypePrimary.ToPointer(),
                    URL: "http://example.com",
                },
            },
        },
        XApideckAppID: "Small",
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

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ApplicantsAddRequest](../../models/operations/applicantsaddrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ApplicantsAddSecurity](../../models/operations/applicantsaddsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ApplicantsAddResponse](../../models/operations/applicantsaddresponse.md), error**


## All

List Applicants

### Example Usage

```go
package main

import(
	"context"
	"log"
	apidecksamplesdk "github.com/speakeasy-sdks/apideck-sample-sdk"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/shared"
)

func main() {
    s := apidecksamplesdk.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Ats.Applicants.All(ctx, operations.ApplicantsAllRequest{
        Filter: &shared.ApplicantsFilter{
            JobID: apidecksamplesdk.String("1234"),
        },
        PassThrough: &shared.PassThroughQuery{
            AdditionalProperties: map[string]interface{}{
                "search": "deposit",
            },
        },
        XApideckAppID: "Mobility",
        XApideckConsumerID: "Mobility",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetApplicantsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ApplicantsAllRequest](../../models/operations/applicantsallrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ApplicantsAllSecurity](../../models/operations/applicantsallsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ApplicantsAllResponse](../../models/operations/applicantsallresponse.md), error**


## Delete

Delete Applicant

### Example Usage

```go
package main

import(
	"context"
	"log"
	apidecksamplesdk "github.com/speakeasy-sdks/apideck-sample-sdk"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/operations"
)

func main() {
    s := apidecksamplesdk.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Ats.Applicants.Delete(ctx, operations.ApplicantsDeleteRequest{
        ID: "<ID>",
        XApideckAppID: "roughly",
        XApideckConsumerID: "EXE",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteApplicantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ApplicantsDeleteRequest](../../models/operations/applicantsdeleterequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ApplicantsDeleteSecurity](../../models/operations/applicantsdeletesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ApplicantsDeleteResponse](../../models/operations/applicantsdeleteresponse.md), error**


## One

Get Applicant

### Example Usage

```go
package main

import(
	"context"
	"log"
	apidecksamplesdk "github.com/speakeasy-sdks/apideck-sample-sdk"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/operations"
)

func main() {
    s := apidecksamplesdk.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Ats.Applicants.One(ctx, operations.ApplicantsOneRequest{
        ID: "<ID>",
        XApideckAppID: "primary",
        XApideckConsumerID: "Fall",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetApplicantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ApplicantsOneRequest](../../models/operations/applicantsonerequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ApplicantsOneSecurity](../../models/operations/applicantsonesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ApplicantsOneResponse](../../models/operations/applicantsoneresponse.md), error**


## Update

Update Applicant

### Example Usage

```go
package main

import(
	"context"
	"log"
	apidecksamplesdk "github.com/speakeasy-sdks/apideck-sample-sdk"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/types"
)

func main() {
    s := apidecksamplesdk.New()


    operationSecurity := ""

    ctx := context.Background()
    res, err := s.Ats.Applicants.Update(ctx, operations.ApplicantsUpdateRequest{
        ApplicantInput: shared.ApplicantInput{
            Addresses: []shared.Address{
                shared.Address{
                    City: apidecksamplesdk.String("San Francisco"),
                    ContactName: apidecksamplesdk.String("Elon Musk"),
                    Country: apidecksamplesdk.String("US"),
                    County: apidecksamplesdk.String("Santa Clara"),
                    Email: apidecksamplesdk.String("elon@musk.com"),
                    Fax: apidecksamplesdk.String("122-111-1111"),
                    ID: apidecksamplesdk.String("123"),
                    Latitude: apidecksamplesdk.String("40.759211"),
                    Line1: apidecksamplesdk.String("Main street"),
                    Line2: apidecksamplesdk.String("apt #"),
                    Line3: apidecksamplesdk.String("Suite #"),
                    Line4: apidecksamplesdk.String("delivery instructions"),
                    Longitude: apidecksamplesdk.String("-73.984638"),
                    Name: apidecksamplesdk.String("HQ US"),
                    Notes: apidecksamplesdk.String("Address notes or delivery instructions."),
                    PhoneNumber: apidecksamplesdk.String("111-111-1111"),
                    PostalCode: apidecksamplesdk.String("94104"),
                    RowVersion: apidecksamplesdk.String("1-12345"),
                    Salutation: apidecksamplesdk.String("Mr"),
                    State: apidecksamplesdk.String("CA"),
                    StreetNumber: apidecksamplesdk.String("25"),
                    String: apidecksamplesdk.String("25 Spring Street, Blackburn, VIC 3130"),
                    Type: shared.AddressTypePrimary.ToPointer(),
                    Website: apidecksamplesdk.String("https://elonmusk.com"),
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
            Archived: apidecksamplesdk.Bool(false),
            Birthday: types.MustDateFromString("2000-08-12"),
            Confidential: apidecksamplesdk.Bool(false),
            CoordinatorID: apidecksamplesdk.String("12345"),
            CoverLetter: apidecksamplesdk.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
            CustomFields: []shared.CustomField{
                shared.CustomField{
                    Description: apidecksamplesdk.String("Employee Level"),
                    ID: "2389328923893298",
                    Name: apidecksamplesdk.String("employee_level"),
                    Value: shared.CreateCustomFieldValueBoolean(
                    true,
                    ),
                },
            },
            Deleted: apidecksamplesdk.Bool(true),
            Emails: []shared.Email{
                shared.Email{
                    Email: "elon@musk.com",
                    ID: apidecksamplesdk.String("123"),
                    Type: shared.EmailTypePrimary.ToPointer(),
                },
            },
            FirstName: apidecksamplesdk.String("Elon"),
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Headline: apidecksamplesdk.String("PepsiCo, Inc, Central Perk"),
            Initials: apidecksamplesdk.String("EM"),
            LastName: apidecksamplesdk.String("Musk"),
            MiddleName: apidecksamplesdk.String("D."),
            Name: apidecksamplesdk.String("Elon Musk"),
            OwnerID: apidecksamplesdk.String("54321"),
            PhoneNumbers: []shared.PhoneNumber{
                shared.PhoneNumber{
                    AreaCode: apidecksamplesdk.String("323"),
                    CountryCode: apidecksamplesdk.String("1"),
                    Extension: apidecksamplesdk.String("105"),
                    ID: apidecksamplesdk.String("12345"),
                    Number: "111-111-1111",
                    Type: shared.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            PhotoURL: apidecksamplesdk.String("https://unavatar.io/elon-musk"),
            PositionID: apidecksamplesdk.String("123"),
            RecordURL: apidecksamplesdk.String("https://app.intercom.io/contacts/12345"),
            RecruiterID: apidecksamplesdk.String("12345"),
            SocialLinks: []shared.ApplicantSocialLinks{
                shared.ApplicantSocialLinks{
                    ID: apidecksamplesdk.String("12345"),
                    Type: apidecksamplesdk.String("twitter"),
                    URL: "https://www.twitter.com/apideck",
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
                    ID: apidecksamplesdk.String("12345"),
                    Type: shared.ApplicantWebsitesTypePrimary.ToPointer(),
                    URL: "http://example.com",
                },
            },
        },
        ID: "<ID>",
        XApideckAppID: "South",
        XApideckConsumerID: "complexity",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateApplicantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ApplicantsUpdateRequest](../../models/operations/applicantsupdaterequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ApplicantsUpdateSecurity](../../models/operations/applicantsupdatesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ApplicantsUpdateResponse](../../models/operations/applicantsupdateresponse.md), error**

