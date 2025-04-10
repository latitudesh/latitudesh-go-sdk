# Products


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Resource`                                                    | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Name`                                                        | **string*                                                     | :heavy_minus_sign:                                            | N/A                                                           |
| `Start`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |
| `End`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |
| `Unit`                                                        | [*components.Unit](../../models/components/unit.md)           | :heavy_minus_sign:                                            | N/A                                                           |
| `UnitPrice`                                                   | **float64*                                                    | :heavy_minus_sign:                                            | The unit price of the product in cents                        |
| `UsageType`                                                   | [*components.UsageType](../../models/components/usagetype.md) | :heavy_minus_sign:                                            | N/A                                                           |
| `Quantity`                                                    | **float64*                                                    | :heavy_minus_sign:                                            | N/A                                                           |
| `Price`                                                       | **float64*                                                    | :heavy_minus_sign:                                            | The total usage price of the product in cents                 |
| `Metadata`                                                    | [*components.Metadata](../../models/components/metadata.md)   | :heavy_minus_sign:                                            | N/A                                                           |