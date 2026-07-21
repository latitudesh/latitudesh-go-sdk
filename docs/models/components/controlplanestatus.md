# ControlPlaneStatus

Current status of control plane nodes. 'ready' when control plane is operational, 'scaling' when nodes are being provisioned/removed, 'upgrading' while a Kubernetes version upgrade is rolling through the control plane, 'error' when a control plane node has failed.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.ControlPlaneStatusReady

// Open enum: custom values can be created with a direct type cast
custom := components.ControlPlaneStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ControlPlaneStatusReady`     | ready                         |
| `ControlPlaneStatusScaling`   | scaling                       |
| `ControlPlaneStatusUpgrading` | upgrading                     |
| `ControlPlaneStatusError`     | error                         |