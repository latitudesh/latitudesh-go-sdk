# KubernetesClusterDataPhase

The current phase of the cluster lifecycle

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.KubernetesClusterDataPhasePending
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `KubernetesClusterDataPhasePending`      | Pending                                  |
| `KubernetesClusterDataPhaseProvisioning` | Provisioning                             |
| `KubernetesClusterDataPhaseProvisioned`  | Provisioned                              |
| `KubernetesClusterDataPhaseDeleting`     | Deleting                                 |
| `KubernetesClusterDataPhaseFailed`       | Failed                                   |