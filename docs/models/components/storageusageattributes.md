# StorageUsageAttributes


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Date`                                                        | [*types.Date](../../types/date.md)                            | :heavy_minus_sign:                                            | The day this usage row refers to                              |
| `StorageID`                                                   | `*int64`                                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `ProjectID`                                                   | `*int64`                                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `Provider`                                                    | `*string`                                                     | :heavy_minus_sign:                                            | Storage provider that reported this usage (e.g. vast, wasabi) |
| `Region`                                                      | `*string`                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Bytes`                                                       | `*int64`                                                      | :heavy_minus_sign:                                            | Canonical storage usage for the day, in bytes                 |
| `RawValue`                                                    | `*float64`                                                    | :heavy_minus_sign:                                            | The provider-reported usage value, in raw_unit                |
| `RawUnit`                                                     | `*string`                                                     | :heavy_minus_sign:                                            | The unit raw_value is expressed in                            |
| `IngestedAt`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |