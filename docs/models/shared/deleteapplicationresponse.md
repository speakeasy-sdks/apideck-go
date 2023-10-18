# DeleteApplicationResponse


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   | Example                                       |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Data`                                        | [UnifiedID](../../models/shared/unifiedid.md) | :heavy_check_mark:                            | N/A                                           |                                               |
| `Operation`                                   | *string*                                      | :heavy_check_mark:                            | Operation performed                           | delete                                        |
| `Resource`                                    | *string*                                      | :heavy_check_mark:                            | Unified API resource name                     | Applications                                  |
| `Service`                                     | *string*                                      | :heavy_check_mark:                            | Apideck ID of service provider                | sap-successfactors                            |
| `Status`                                      | *string*                                      | :heavy_check_mark:                            | HTTP Response Status                          | OK                                            |
| `StatusCode`                                  | *int64*                                       | :heavy_check_mark:                            | HTTP Response Status Code                     | 200                                           |