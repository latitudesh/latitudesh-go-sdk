# ListLksNodesRequest


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ClusterID`                                                                       | `string`                                                                          | :heavy_check_mark:                                                                | The cluster ID (format: `lksc_<hash>`).                                           |
| `NodePoolID`                                                                      | `string`                                                                          | :heavy_check_mark:                                                                | The node pool ID (format: `lksnp_<hash>`).                                        |
| `FilterStatus`                                                                    | `*string`                                                                         | :heavy_minus_sign:                                                                | Only return nodes in this status: `pending`, `ready`, `not_ready`, or `deleting`. |