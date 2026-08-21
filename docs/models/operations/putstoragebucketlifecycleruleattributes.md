# PutStorageBucketLifecycleRuleAttributes


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Name`                                                                | `string`                                                              | :heavy_check_mark:                                                    | Name of the lifecycle rule                                            |
| `Enabled`                                                             | `*bool`                                                               | :heavy_minus_sign:                                                    | Whether the rule is active                                            |
| `Prefix`                                                              | `*string`                                                             | :heavy_minus_sign:                                                    | Object key prefix to filter which objects the rule applies to         |
| `ExpirationDays`                                                      | `*int64`                                                              | :heavy_minus_sign:                                                    | Number of days after object creation when the object expires          |
| `NoncurrentDays`                                                      | `*int64`                                                              | :heavy_minus_sign:                                                    | Number of days after which noncurrent object versions expire          |
| `AbortMpuDaysAfterInitiation`                                         | `*int64`                                                              | :heavy_minus_sign:                                                    | Number of days after initiation to abort incomplete multipart uploads |