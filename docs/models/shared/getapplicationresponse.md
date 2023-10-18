# GetApplicationResponse


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Data`                                            | [Application](../../models/shared/application.md) | :heavy_check_mark:                                | N/A                                               |                                                   |
| `Operation`                                       | *string*                                          | :heavy_check_mark:                                | Operation performed                               | one                                               |
| `Resource`                                        | *string*                                          | :heavy_check_mark:                                | Unified API resource name                         | Applications                                      |
| `Service`                                         | *string*                                          | :heavy_check_mark:                                | Apideck ID of service provider                    | sap-successfactors                                |
| `Status`                                          | *string*                                          | :heavy_check_mark:                                | HTTP Response Status                              | OK                                                |
| `StatusCode`                                      | *int64*                                           | :heavy_check_mark:                                | HTTP Response Status Code                         | 200                                               |