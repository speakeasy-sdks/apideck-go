# UpdateApplicantResponse


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   | Example                                       |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Data`                                        | [UnifiedID](../../models/shared/unifiedid.md) | :heavy_check_mark:                            | N/A                                           |                                               |
| `Operation`                                   | *string*                                      | :heavy_check_mark:                            | Operation performed                           | update                                        |
| `Resource`                                    | *string*                                      | :heavy_check_mark:                            | Unified API resource name                     | Applicants                                    |
| `Service`                                     | *string*                                      | :heavy_check_mark:                            | Apideck ID of service provider                | lever                                         |
| `Status`                                      | *string*                                      | :heavy_check_mark:                            | HTTP Response Status                          | OK                                            |
| `StatusCode`                                  | *int64*                                       | :heavy_check_mark:                            | HTTP Response Status Code                     | 200                                           |