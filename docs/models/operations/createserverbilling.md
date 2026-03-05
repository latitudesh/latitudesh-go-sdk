# CreateServerBilling

The server billing type. Accepts `hourly` and `monthly` for on demand projects and `yearly` for reserved projects.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.CreateServerBillingHourly
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CreateServerBillingHourly`  | hourly                       |
| `CreateServerBillingMonthly` | monthly                      |
| `CreateServerBillingYearly`  | yearly                       |