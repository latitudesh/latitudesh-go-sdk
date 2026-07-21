# KubernetesClusterDataPhase

The current phase of the cluster lifecycle. 'Upgrading' is reported while a Kubernetes version upgrade is rolling through the cluster.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.KubernetesClusterDataPhasePending

// Open enum: custom values can be created with a direct type cast
custom := components.KubernetesClusterDataPhase("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `KubernetesClusterDataPhasePending`      | Pending                                  |
| `KubernetesClusterDataPhaseProvisioning` | Provisioning                             |
| `KubernetesClusterDataPhaseProvisioned`  | Provisioned                              |
| `KubernetesClusterDataPhaseUpgrading`    | Upgrading                                |
| `KubernetesClusterDataPhaseDeleting`     | Deleting                                 |
| `KubernetesClusterDataPhaseFailed`       | Failed                                   |
| `KubernetesClusterDataPhaseUnknown`      | Unknown                                  |