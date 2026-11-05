# Series


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Timestamp`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                        | :heavy_minus_sign:                                                | Start of the bucket (UTC)                                         |
| `Models`                                                          | map[string][components.Models](../../models/components/models.md) | :heavy_minus_sign:                                                | Usage in the bucket keyed by model ID                             |