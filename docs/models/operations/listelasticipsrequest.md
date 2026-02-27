# ListElasticIpsRequest


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `FilterProject`                                                     | **string*                                                           | :heavy_minus_sign:                                                  | The project ID or slug to filter by                                 |
| `FilterServer`                                                      | **string*                                                           | :heavy_minus_sign:                                                  | The server ID to filter by                                          |
| `FilterStatus`                                                      | [*operations.FilterStatus](../../models/operations/filterstatus.md) | :heavy_minus_sign:                                                  | The status to filter by                                             |
| `PageSize`                                                          | **int64*                                                            | :heavy_minus_sign:                                                  | Number of items to return per page                                  |
| `PageNumber`                                                        | **int64*                                                            | :heavy_minus_sign:                                                  | Page number to return (starts at 1)                                 |