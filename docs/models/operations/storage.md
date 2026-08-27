# Storage

Storage consumption metrics


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Consumed`                                          | `*int64`                                            | :heavy_minus_sign:                                  | Billed storage usage for the current period         |
| `Current`                                           | `*int64`                                            | :heavy_minus_sign:                                  | Latest recorded storage usage (last datapoint)      |
| `Unit`                                              | [*operations.Unit](../../models/operations/unit.md) | :heavy_minus_sign:                                  | Unit of measurement for storage                     |