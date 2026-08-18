# ListFirewallsRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `FilterProject`                                                          | `*string`                                                                | :heavy_minus_sign:                                                       | N/A                                                                      |
| `FilterTags`                                                             | `*string`                                                                | :heavy_minus_sign:                                                       | Comma-separated tag IDs. Returns firewalls that have all the given tags. |
| `PageSize`                                                               | `*int64`                                                                 | :heavy_minus_sign:                                                       | Number of items to return per page                                       |
| `PageNumber`                                                             | `*int64`                                                                 | :heavy_minus_sign:                                                       | Page number to return (starts at 1)                                      |