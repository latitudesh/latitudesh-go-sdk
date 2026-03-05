# Phase

The current phase of the cluster lifecycle

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.PhasePending
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `PhasePending`      | Pending             |
| `PhaseProvisioning` | Provisioning        |
| `PhaseProvisioned`  | Provisioned         |
| `PhaseDeleting`     | Deleting            |
| `PhaseFailed`       | Failed              |