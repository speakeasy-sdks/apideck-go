# Applicants
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
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/types"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicantsAddSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applicants.Add(ctx, operations.ApplicantsAddRequest{
        Applicant: shared.ApplicantInput{
            Addresses: []shared.Address{
                shared.Address{
                    City: apideckgo.String("San Francisco"),
                    ContactName: apideckgo.String("Elon Musk"),
                    Country: apideckgo.String("US"),
                    County: apideckgo.String("Santa Clara"),
                    Email: apideckgo.String("elon@musk.com"),
                    Fax: apideckgo.String("122-111-1111"),
                    ID: apideckgo.String("123"),
                    Latitude: apideckgo.String("40.759211"),
                    Line1: apideckgo.String("Main street"),
                    Line2: apideckgo.String("apt #"),
                    Line3: apideckgo.String("Suite #"),
                    Line4: apideckgo.String("delivery instructions"),
                    Longitude: apideckgo.String("-73.984638"),
                    Name: apideckgo.String("HQ US"),
                    Notes: apideckgo.String("Address notes or delivery instructions."),
                    PhoneNumber: apideckgo.String("111-111-1111"),
                    PostalCode: apideckgo.String("94104"),
                    RowVersion: apideckgo.String("1-12345"),
                    Salutation: apideckgo.String("Mr"),
                    State: apideckgo.String("CA"),
                    StreetNumber: apideckgo.String("25"),
                    String: apideckgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Type: shared.TypePrimary.ToPointer(),
                    Website: apideckgo.String("https://elonmusk.com"),
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
            Archived: apideckgo.Bool(false),
            Birthday: types.MustDateFromString("2000-08-12"),
            Confidential: apideckgo.Bool(false),
            CoordinatorID: apideckgo.String("12345"),
            CoverLetter: apideckgo.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
            CustomFields: []shared.CustomField{
                shared.CustomField{
                    Description: apideckgo.String("Employee Level"),
                    ID: "2389328923893298",
                    Name: apideckgo.String("employee_level"),
                    Value: shared.CreateValueBoolean(
                    true,
                    ),
                },
            },
            Deleted: apideckgo.Bool(true),
            Emails: []shared.Email{
                shared.Email{
                    Email: "elon@musk.com",
                    ID: apideckgo.String("123"),
                    Type: shared.EmailTypePrimary.ToPointer(),
                },
            },
            FirstName: apideckgo.String("Elon"),
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Headline: apideckgo.String("PepsiCo, Inc, Central Perk"),
            Initials: apideckgo.String("EM"),
            LastName: apideckgo.String("Musk"),
            MiddleName: apideckgo.String("D."),
            Name: apideckgo.String("Elon Musk"),
            OwnerID: apideckgo.String("54321"),
            PhoneNumbers: []shared.PhoneNumber{
                shared.PhoneNumber{
                    AreaCode: apideckgo.String("323"),
                    CountryCode: apideckgo.String("1"),
                    Extension: apideckgo.String("105"),
                    ID: apideckgo.String("12345"),
                    Number: "111-111-1111",
                    Type: shared.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            PhotoURL: apideckgo.String("https://unavatar.io/elon-musk"),
            PositionID: apideckgo.String("123"),
            RecordURL: apideckgo.String("https://app.intercom.io/contacts/12345"),
            RecruiterID: apideckgo.String("12345"),
            SocialLinks: []shared.SocialLinks{
                shared.SocialLinks{
                    ID: apideckgo.String("12345"),
                    Type: apideckgo.String("twitter"),
                    URL: "https://www.twitter.com/apideck",
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
            Websites: []shared.Websites{
                shared.Websites{
                    ID: apideckgo.String("12345"),
                    Type: shared.ApplicantTypePrimary.ToPointer(),
                    URL: "http://example.com",
                },
            },
        },
        XApideckAppID: "string",
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

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ApplicantsAddRequest](../../pkg/models/operations/applicantsaddrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ApplicantsAddSecurity](../../pkg/models/operations/applicantsaddsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ApplicantsAddResponse](../../pkg/models/operations/applicantsaddresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## All

List Applicants

### Example Usage

```go
package main

import(
	"context"
	"log"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicantsAllSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applicants.All(ctx, operations.ApplicantsAllRequest{
        Filter: &shared.ApplicantsFilter{
            JobID: apideckgo.String("1234"),
        },
        PassThrough: map[string]interface{}{
            "search": "string",
        },
        XApideckAppID: "string",
        XApideckConsumerID: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ApplicantsAllRequest](../../pkg/models/operations/applicantsallrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ApplicantsAllSecurity](../../pkg/models/operations/applicantsallsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ApplicantsAllResponse](../../pkg/models/operations/applicantsallresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## Delete

Delete Applicant

### Example Usage

```go
package main

import(
	"context"
	"log"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicantsDeleteSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applicants.Delete(ctx, operations.ApplicantsDeleteRequest{
        ID: "<ID>",
        XApideckAppID: "string",
        XApideckConsumerID: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ApplicantsDeleteRequest](../../pkg/models/operations/applicantsdeleterequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ApplicantsDeleteSecurity](../../pkg/models/operations/applicantsdeletesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ApplicantsDeleteResponse](../../pkg/models/operations/applicantsdeleteresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## One

Get Applicant

### Example Usage

```go
package main

import(
	"context"
	"log"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicantsOneSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applicants.One(ctx, operations.ApplicantsOneRequest{
        ID: "<ID>",
        XApideckAppID: "string",
        XApideckConsumerID: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ApplicantsOneRequest](../../pkg/models/operations/applicantsonerequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ApplicantsOneSecurity](../../pkg/models/operations/applicantsonesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ApplicantsOneResponse](../../pkg/models/operations/applicantsoneresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## Update

Update Applicant

### Example Usage

```go
package main

import(
	"context"
	"log"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-go/pkg/types"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicantsUpdateSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applicants.Update(ctx, operations.ApplicantsUpdateRequest{
        Applicant: shared.ApplicantInput{
            Addresses: []shared.Address{
                shared.Address{
                    City: apideckgo.String("San Francisco"),
                    ContactName: apideckgo.String("Elon Musk"),
                    Country: apideckgo.String("US"),
                    County: apideckgo.String("Santa Clara"),
                    Email: apideckgo.String("elon@musk.com"),
                    Fax: apideckgo.String("122-111-1111"),
                    ID: apideckgo.String("123"),
                    Latitude: apideckgo.String("40.759211"),
                    Line1: apideckgo.String("Main street"),
                    Line2: apideckgo.String("apt #"),
                    Line3: apideckgo.String("Suite #"),
                    Line4: apideckgo.String("delivery instructions"),
                    Longitude: apideckgo.String("-73.984638"),
                    Name: apideckgo.String("HQ US"),
                    Notes: apideckgo.String("Address notes or delivery instructions."),
                    PhoneNumber: apideckgo.String("111-111-1111"),
                    PostalCode: apideckgo.String("94104"),
                    RowVersion: apideckgo.String("1-12345"),
                    Salutation: apideckgo.String("Mr"),
                    State: apideckgo.String("CA"),
                    StreetNumber: apideckgo.String("25"),
                    String: apideckgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Type: shared.TypePrimary.ToPointer(),
                    Website: apideckgo.String("https://elonmusk.com"),
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
            Archived: apideckgo.Bool(false),
            Birthday: types.MustDateFromString("2000-08-12"),
            Confidential: apideckgo.Bool(false),
            CoordinatorID: apideckgo.String("12345"),
            CoverLetter: apideckgo.String("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
            CustomFields: []shared.CustomField{
                shared.CustomField{
                    Description: apideckgo.String("Employee Level"),
                    ID: "2389328923893298",
                    Name: apideckgo.String("employee_level"),
                    Value: shared.CreateValueBoolean(
                    true,
                    ),
                },
            },
            Deleted: apideckgo.Bool(true),
            Emails: []shared.Email{
                shared.Email{
                    Email: "elon@musk.com",
                    ID: apideckgo.String("123"),
                    Type: shared.EmailTypePrimary.ToPointer(),
                },
            },
            FirstName: apideckgo.String("Elon"),
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Headline: apideckgo.String("PepsiCo, Inc, Central Perk"),
            Initials: apideckgo.String("EM"),
            LastName: apideckgo.String("Musk"),
            MiddleName: apideckgo.String("D."),
            Name: apideckgo.String("Elon Musk"),
            OwnerID: apideckgo.String("54321"),
            PhoneNumbers: []shared.PhoneNumber{
                shared.PhoneNumber{
                    AreaCode: apideckgo.String("323"),
                    CountryCode: apideckgo.String("1"),
                    Extension: apideckgo.String("105"),
                    ID: apideckgo.String("12345"),
                    Number: "111-111-1111",
                    Type: shared.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            PhotoURL: apideckgo.String("https://unavatar.io/elon-musk"),
            PositionID: apideckgo.String("123"),
            RecordURL: apideckgo.String("https://app.intercom.io/contacts/12345"),
            RecruiterID: apideckgo.String("12345"),
            SocialLinks: []shared.SocialLinks{
                shared.SocialLinks{
                    ID: apideckgo.String("12345"),
                    Type: apideckgo.String("twitter"),
                    URL: "https://www.twitter.com/apideck",
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
            Websites: []shared.Websites{
                shared.Websites{
                    ID: apideckgo.String("12345"),
                    Type: shared.ApplicantTypePrimary.ToPointer(),
                    URL: "http://example.com",
                },
            },
        },
        ID: "<ID>",
        XApideckAppID: "string",
        XApideckConsumerID: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ApplicantsUpdateRequest](../../pkg/models/operations/applicantsupdaterequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ApplicantsUpdateSecurity](../../pkg/models/operations/applicantsupdatesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ApplicantsUpdateResponse](../../pkg/models/operations/applicantsupdateresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |
