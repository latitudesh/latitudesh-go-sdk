# AccessKey

The newly created access key. The secret is included only in this create response and cannot be retrieved again. Field names depend on the provider: `standard` (Wasabi) returns `access_key_id` and `secret_access_key`; `high_performance` (VAST) returns `access_key` and `secret_key`.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AccessKeyID`                                              | `*string`                                                  | :heavy_minus_sign:                                         | Access key ID (standard / Wasabi).                         |
| `SecretAccessKey`                                          | `*string`                                                  | :heavy_minus_sign:                                         | Secret access key (standard / Wasabi). Returned only once. |
| `AccessKey`                                                | `*string`                                                  | :heavy_minus_sign:                                         | Access key ID (high_performance / VAST).                   |
| `SecretKey`                                                | `*string`                                                  | :heavy_minus_sign:                                         | Secret key (high_performance / VAST). Returned only once.  |
| `Name`                                                     | `*string`                                                  | :heavy_minus_sign:                                         | Access key name.                                           |
| `Status`                                                   | `*string`                                                  | :heavy_minus_sign:                                         | Access key status (e.g., `Active`).                        |
| `Username`                                                 | `*string`                                                  | :heavy_minus_sign:                                         | Underlying IAM user the key belongs to.                    |