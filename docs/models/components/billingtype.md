# BillingType

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.BillingTypeYearly

// Open enum: custom values can be created with a direct type cast
custom := components.BillingType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `BillingTypeYearly`  | Yearly               |
| `BillingTypeMonthly` | Monthly              |
| `BillingTypeHourly`  | Hourly               |
| `BillingTypeNormal`  | Normal               |
| `BillingTypeCustom`  | Custom               |