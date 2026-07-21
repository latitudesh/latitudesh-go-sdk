# Phase

The current phase of the cluster lifecycle.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.PhasePending

// Open enum: custom values can be created with a direct type cast
custom := components.Phase("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `PhasePending`      | Pending             |
| `PhaseProvisioning` | Provisioning        |
| `PhaseProvisioned`  | Provisioned         |
| `PhaseUpgrading`    | Upgrading           |
| `PhaseDeleting`     | Deleting            |
| `PhaseFailed`       | Failed              |
| `PhaseUnknown`      | Unknown             |