# ClusterNodeGroupDataAttributes


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Name`                                              | `*string`                                           | :heavy_minus_sign:                                  | N/A                                                 | gpu-workers                                         |
| `Kind`                                              | [*components.Kind](../../models/components/kind.md) | :heavy_minus_sign:                                  | N/A                                                 | gpu                                                 |
| `Plan`                                              | `*string`                                           | :heavy_minus_sign:                                  | The plan slug; null while no plan is set            | g4-b300-large                                       |
| `PlanName`                                          | `*string`                                           | :heavy_minus_sign:                                  | N/A                                                 | g4.b300.large                                       |
| `ServerCount`                                       | `*int64`                                            | :heavy_minus_sign:                                  | N/A                                                 | 16                                                  |
| `Os`                                                | `*string`                                           | :heavy_minus_sign:                                  | The operating system slug; null while none is set   | ubuntu_24_04_x64_lts                                |