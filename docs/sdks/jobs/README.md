# Jobs
(*Ats.Jobs*)

### Available Operations

* [All](#all) - List Jobs
* [One](#one) - Get Job

## All

List Jobs

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


    operationSecurity := operations.JobsAllSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Jobs.All(ctx, operations.JobsAllRequest{
        PassThrough: map[string]interface{}{
            "search": "string",
        },
        XApideckAppID: "string",
        XApideckConsumerID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetJobsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.JobsAllRequest](../../pkg/models/operations/jobsallrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.JobsAllSecurity](../../pkg/models/operations/jobsallsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.JobsAllResponse](../../pkg/models/operations/jobsallresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## One

Get Job

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


    operationSecurity := operations.JobsOneSecurity{
            APIKey: "<your-apideck-api-key>",
        }

    ctx := context.Background()
    res, err := s.Ats.Jobs.One(ctx, operations.JobsOneRequest{
        ID: "<ID>",
        XApideckAppID: "string",
        XApideckConsumerID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.JobsOneRequest](../../pkg/models/operations/jobsonerequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.JobsOneSecurity](../../pkg/models/operations/jobsonesecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.JobsOneResponse](../../pkg/models/operations/jobsoneresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.BadRequestResponse      | 400                               | application/json                  |
| sdkerrors.UnauthorizedResponse    | 401                               | application/json                  |
| sdkerrors.PaymentRequiredResponse | 402                               | application/json                  |
| sdkerrors.NotFoundResponse        | 404                               | application/json                  |
| sdkerrors.UnprocessableResponse   | 422                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |
