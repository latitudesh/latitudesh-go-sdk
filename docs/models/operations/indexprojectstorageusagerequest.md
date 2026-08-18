# IndexProjectStorageUsageRequest


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ProjectID`                                                            | `string`                                                               | :heavy_check_mark:                                                     | Project ID or Slug                                                     |
| `StorageID`                                                            | `*string`                                                              | :heavy_minus_sign:                                                     | Restrict the result to a single storage. Accepts the storage/bucket ID |
| `StartDate`                                                            | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to yesterday                                                  |
| `EndDate`                                                              | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to today; clamped to today when a future date is given        |