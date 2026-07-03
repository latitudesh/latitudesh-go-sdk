# VirtualMachineUpdatePayloadBilling

Target billing cycle. Upgrades only (hourly → monthly → yearly); downgrades and reserved-project changes return 422.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.VirtualMachineUpdatePayloadBillingHourly
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `VirtualMachineUpdatePayloadBillingHourly`  | hourly                                      |
| `VirtualMachineUpdatePayloadBillingMonthly` | monthly                                     |
| `VirtualMachineUpdatePayloadBillingYearly`  | yearly                                      |