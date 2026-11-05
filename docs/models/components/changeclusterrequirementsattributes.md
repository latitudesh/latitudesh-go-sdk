# ChangeClusterRequirementsAttributes

Exactly one of add, add_checks, remove, remove_checks or set. expected and note go with add and set.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Add`                                                 | `*string`                                             | :heavy_minus_sign:                                    | A requirement to add                                  | storage_fs                                            |
| `AddChecks`                                           | []`string`                                            | :heavy_minus_sign:                                    | Checks to add at once                                 |                                                       |
| `Remove`                                              | `*string`                                             | :heavy_minus_sign:                                    | A requirement id to remove                            |                                                       |
| `RemoveChecks`                                        | []`string`                                            | :heavy_minus_sign:                                    | Checks to remove at once                              |                                                       |
| `Set`                                                 | `*string`                                             | :heavy_minus_sign:                                    | A requirement id whose expected value or note changes |                                                       |
| `Expected`                                            | `*string`                                             | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `Note`                                                | `*string`                                             | :heavy_minus_sign:                                    | N/A                                                   |                                                       |