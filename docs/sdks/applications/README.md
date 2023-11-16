# Applications
(*Ats.Applications*)

### Available Operations

* [Add](#add) - Create Application
* [All](#all) - List Applications
* [Delete](#delete) - Delete Application
* [One](#one) - Get Application
* [Update](#update) - Update Application

## Add

Create Application

### Example Usage

```go
package main

import(
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"log"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicationsAddSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applications.Add(ctx, operations.ApplicationsAddRequest{
        Application: shared.ApplicationInput{
            ApplicantID: "12345",
            JobID: "12345",
            Stage: &shared.Stage{
                ID: apideckgo.String("12345"),
                Name: apideckgo.String("12345"),
            },
            Status: shared.StatusOpen.ToPointer(),
        },
        XApideckAppID: "string",
        XApideckConsumerID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ApplicationsAddRequest](../../pkg/models/operations/applicationsaddrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ApplicationsAddSecurity](../../pkg/models/operations/applicationsaddsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ApplicationsAddResponse](../../pkg/models/operations/applicationsaddresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## All

List Applications

### Example Usage

```go
package main

import(
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicationsAllSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applications.All(ctx, operations.ApplicationsAllRequest{
        PassThrough: map[string]interface{}{
            "search": "string",
        },
        XApideckAppID: "string",
        XApideckConsumerID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetApplicationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ApplicationsAllRequest](../../pkg/models/operations/applicationsallrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ApplicationsAllSecurity](../../pkg/models/operations/applicationsallsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ApplicationsAllResponse](../../pkg/models/operations/applicationsallresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## Delete

Delete Application

### Example Usage

```go
package main

import(
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicationsDeleteSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applications.Delete(ctx, operations.ApplicationsDeleteRequest{
        ID: "<ID>",
        XApideckAppID: "string",
        XApideckConsumerID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ApplicationsDeleteRequest](../../pkg/models/operations/applicationsdeleterequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.ApplicationsDeleteSecurity](../../pkg/models/operations/applicationsdeletesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.ApplicationsDeleteResponse](../../pkg/models/operations/applicationsdeleteresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## One

Get Application

### Example Usage

```go
package main

import(
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicationsOneSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applications.One(ctx, operations.ApplicationsOneRequest{
        ID: "<ID>",
        XApideckAppID: "string",
        XApideckConsumerID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ApplicationsOneRequest](../../pkg/models/operations/applicationsonerequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ApplicationsOneSecurity](../../pkg/models/operations/applicationsonesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ApplicationsOneResponse](../../pkg/models/operations/applicationsoneresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |

## Update

Update Application

### Example Usage

```go
package main

import(
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	"log"
)

func main() {
    s := apideckgo.New()


    operationSecurity := operations.ApplicationsUpdateSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Applications.Update(ctx, operations.ApplicationsUpdateRequest{
        Application: shared.ApplicationInput{
            ApplicantID: "12345",
            JobID: "12345",
            Stage: &shared.Stage{
                ID: apideckgo.String("12345"),
                Name: apideckgo.String("12345"),
            },
            Status: shared.StatusOpen.ToPointer(),
        },
        ID: "<ID>",
        XApideckAppID: "string",
        XApideckConsumerID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ApplicationsUpdateRequest](../../pkg/models/operations/applicationsupdaterequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.ApplicationsUpdateSecurity](../../pkg/models/operations/applicationsupdatesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.ApplicationsUpdateResponse](../../pkg/models/operations/applicationsupdateresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 400-600                           | */*                               |
