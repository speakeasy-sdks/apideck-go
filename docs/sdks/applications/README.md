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
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"context"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"log"
)

func main() {
    s := apideckgo.New(
        apideckgo.WithSecurity("<your-apideck-api-key>"),
    )

    ctx := context.Background()
    res, err := s.Ats.Applications.Add(ctx, operations.ApplicationsAddRequest{
        Application: shared.ApplicationInput{
            ApplicantID: "12345",
            JobID: "12345",
            Status: shared.StatusOpen.ToPointer(),
        },
        XApideckAppID: "<value>",
        XApideckConsumerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ApplicationsAddRequest](../../pkg/models/operations/applicationsaddrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ApplicationsAddResponse](../../pkg/models/operations/applicationsaddresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## All

List Applications

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"context"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"log"
)

func main() {
    s := apideckgo.New(
        apideckgo.WithSecurity("<your-apideck-api-key>"),
    )

    ctx := context.Background()
    res, err := s.Ats.Applications.All(ctx, operations.ApplicationsAllRequest{
        PassThrough: map[string]interface{}{
            "search": "San Francisco",
        },
        XApideckAppID: "<value>",
        XApideckConsumerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ApplicationsAllRequest](../../pkg/models/operations/applicationsallrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ApplicationsAllResponse](../../pkg/models/operations/applicationsallresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## Delete

Delete Application

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"context"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"log"
)

func main() {
    s := apideckgo.New(
        apideckgo.WithSecurity("<your-apideck-api-key>"),
    )

    ctx := context.Background()
    res, err := s.Ats.Applications.Delete(ctx, operations.ApplicationsDeleteRequest{
        ID: "<id>",
        XApideckAppID: "<value>",
        XApideckConsumerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ApplicationsDeleteRequest](../../pkg/models/operations/applicationsdeleterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ApplicationsDeleteResponse](../../pkg/models/operations/applicationsdeleteresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## One

Get Application

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"context"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"log"
)

func main() {
    s := apideckgo.New(
        apideckgo.WithSecurity("<your-apideck-api-key>"),
    )

    ctx := context.Background()
    res, err := s.Ats.Applications.One(ctx, operations.ApplicationsOneRequest{
        ID: "<id>",
        XApideckAppID: "<value>",
        XApideckConsumerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ApplicationsOneRequest](../../pkg/models/operations/applicationsonerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ApplicationsOneResponse](../../pkg/models/operations/applicationsoneresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## Update

Update Application

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"context"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"log"
)

func main() {
    s := apideckgo.New(
        apideckgo.WithSecurity("<your-apideck-api-key>"),
    )

    ctx := context.Background()
    res, err := s.Ats.Applications.Update(ctx, operations.ApplicationsUpdateRequest{
        Application: shared.ApplicationInput{
            ApplicantID: "12345",
            JobID: "12345",
            Status: shared.StatusOpen.ToPointer(),
        },
        ID: "<id>",
        XApideckAppID: "<value>",
        XApideckConsumerID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ApplicationsUpdateRequest](../../pkg/models/operations/applicationsupdaterequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ApplicationsUpdateResponse](../../pkg/models/operations/applicationsupdateresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |
