# GetStorageUsageRequest


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `FilterProject`                                                        | `string`                                                               | :heavy_check_mark:                                                     | Project ID or Slug                                                     |
| `FilterStorageID`                                                      | `*string`                                                              | :heavy_minus_sign:                                                     | Restrict the result to a single storage. Accepts the storage/bucket ID |
| `FilterStartDate`                                                      | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to yesterday                                                  |
| `FilterEndDate`                                                        | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to today; clamped to today when a future date is given        |