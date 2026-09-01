# ShowManagedDatabaseMetricsResponseBody

OK


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `From`                                                              | [time.Time](https://pkg.go.dev/time#Time)                           | :heavy_check_mark:                                                  | N/A                                                                 |
| `To`                                                                | [time.Time](https://pkg.go.dev/time#Time)                           | :heavy_check_mark:                                                  | N/A                                                                 |
| `Metrics`                                                           | map[string][operations.Metrics](../../models/operations/metrics.md) | :heavy_check_mark:                                                  | N/A                                                                 |