# AtsJobs
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
	"context"
	"log"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
)

func main() {
    s := apideckgo.New()


    operationSecurity := "<your-apideck-api-key>"

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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.JobsAllRequest](../../models/operations/jobsallrequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.JobsAllSecurity](../../models/operations/jobsallsecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


### Response

**[*operations.JobsAllResponse](../../models/operations/jobsallresponse.md), error**


## One

Get Job

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


    operationSecurity := "<your-apideck-api-key>"

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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.JobsOneRequest](../../models/operations/jobsonerequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.JobsOneSecurity](../../models/operations/jobsonesecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


### Response

**[*operations.JobsOneResponse](../../models/operations/jobsoneresponse.md), error**

