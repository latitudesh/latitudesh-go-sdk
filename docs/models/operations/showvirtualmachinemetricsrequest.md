# ShowVirtualMachineMetricsRequest


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `VirtualMachineID`                                     | `string`                                               | :heavy_check_mark:                                     | N/A                                                    |
| `Metric`                                               | [operations.Metric](../../models/operations/metric.md) | :heavy_check_mark:                                     | N/A                                                    |
| `Range`                                                | [*operations.Range](../../models/operations/range.md)  | :heavy_minus_sign:                                     | N/A                                                    |
| `ForceRefresh`                                         | `*bool`                                                | :heavy_minus_sign:                                     | N/A                                                    |