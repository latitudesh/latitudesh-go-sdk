# ControlPlaneStatus

Current status of control plane nodes. 'ready' when control plane is operational, 'scaling' when nodes are being provisioned/removed, 'error' when a control plane node has failed.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.ControlPlaneStatusReady
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ControlPlaneStatusReady`   | ready                       |
| `ControlPlaneStatusScaling` | scaling                     |
| `ControlPlaneStatusError`   | error                       |