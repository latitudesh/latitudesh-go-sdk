# Points


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Timestamp`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | ISO 8601 UTC timestamp (seconds precision) |
| `Value`                                    | `*float64`                                 | :heavy_minus_sign:                         | Sampled value, rounded to 2 decimal places |