# AtsApplications
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
	"context"
	"log"
	apideckgo "github.com/speakeasy-sdks/apideck-go"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/apideck-go/pkg/models/shared"
)

func main() {
    s := apideckgo.New()


    operationSecurity := "<your-apideck-api-key>"

    ctx := context.Background()
    res, err := s.Ats.Applications.Add(ctx, operations.ApplicationsAddRequest{
        ApplicationInput: shared.ApplicationInput{
            ApplicantID: "12345",
            JobID: "12345",
            Stage: &shared.ApplicationStage{
                ID: apideckgo.String("12345"),
                Name: apideckgo.String("12345"),
            },
            Status: shared.ApplicationStatusOpen.ToPointer(),
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ApplicationsAddRequest](../../models/operations/applicationsaddrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ApplicationsAddSecurity](../../models/operations/applicationsaddsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ApplicationsAddResponse](../../models/operations/applicationsaddresponse.md), error**


## All

List Applications

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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ApplicationsAllRequest](../../models/operations/applicationsallrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ApplicationsAllSecurity](../../models/operations/applicationsallsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ApplicationsAllResponse](../../models/operations/applicationsallresponse.md), error**


## Delete

Delete Application

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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ApplicationsDeleteRequest](../../models/operations/applicationsdeleterequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ApplicationsDeleteSecurity](../../models/operations/applicationsdeletesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ApplicationsDeleteResponse](../../models/operations/applicationsdeleteresponse.md), error**


## One

Get Application

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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ApplicationsOneRequest](../../models/operations/applicationsonerequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ApplicationsOneSecurity](../../models/operations/applicationsonesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ApplicationsOneResponse](../../models/operations/applicationsoneresponse.md), error**


## Update

Update Application

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


    operationSecurity := "<your-apideck-api-key>"

    ctx := context.Background()
    res, err := s.Ats.Applications.Update(ctx, operations.ApplicationsUpdateRequest{
        ApplicationInput: shared.ApplicationInput{
            ApplicantID: "12345",
            JobID: "12345",
            Stage: &shared.ApplicationStage{
                ID: apideckgo.String("12345"),
                Name: apideckgo.String("12345"),
            },
            Status: shared.ApplicationStatusOpen.ToPointer(),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ApplicationsUpdateRequest](../../models/operations/applicationsupdaterequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ApplicationsUpdateSecurity](../../models/operations/applicationsupdatesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ApplicationsUpdateResponse](../../models/operations/applicationsupdateresponse.md), error**

