# VirtualMachinePayloadBilling

Billing cycle for the VM. The supported set is validated per-project (on_demand vs reserved). Defaults to the project's default billing when omitted.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.VirtualMachinePayloadBillingHourly
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `VirtualMachinePayloadBillingHourly`  | hourly                                |
| `VirtualMachinePayloadBillingMonthly` | monthly                               |
| `VirtualMachinePayloadBillingYearly`  | yearly                                |