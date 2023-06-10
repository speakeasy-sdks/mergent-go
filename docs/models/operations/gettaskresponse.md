# GetTaskResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | N/A                                                     |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | N/A                                                     |
| `Error`                                                 | [*shared.Error](../../models/shared/error.md)           | :heavy_minus_sign:                                      | Not Found                                               |
| `Task`                                                  | [*shared.TaskOutput](../../models/shared/taskoutput.md) | :heavy_minus_sign:                                      | OK                                                      |