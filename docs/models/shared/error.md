# Error

Unprocessable Entity


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Errors`                                                                                   | [][Error](../../models/shared/error.md)                                                    | :heavy_minus_sign:                                                                         | If multiple errors occured (e.g., with param validation), the list of errors that occured. |
| `Message`                                                                                  | **string*                                                                                  | :heavy_minus_sign:                                                                         | A human-readable message providing more details about the error(s).                        |
| `Param`                                                                                    | **string*                                                                                  | :heavy_minus_sign:                                                                         | If the error is parameter-specific, the parameter related to the error.                    |