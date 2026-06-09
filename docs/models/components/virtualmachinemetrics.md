# VirtualMachineMetrics


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Metric`                                                     | [*components.Metric](../../models/components/metric.md)      | :heavy_minus_sign:                                           | N/A                                                          |
| `Range`                                                      | [*components.Range](../../models/components/range.md)        | :heavy_minus_sign:                                           | N/A                                                          |
| `Step`                                                       | `*string`                                                    | :heavy_minus_sign:                                           | Sampling interval between adjacent points (e.g. "15s", "1m") |
| `Unit`                                                       | [*components.Unit](../../models/components/unit.md)          | :heavy_minus_sign:                                           | Unit applied to every point value                            |
| `Points`                                                     | [][components.Points](../../models/components/points.md)     | :heavy_minus_sign:                                           | N/A                                                          |