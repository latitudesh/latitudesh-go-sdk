# KubernetesClusterUpdateResponseStatus

The update status. 'scaling' indicates nodes are being added or removed. 'upgrading' indicates a version upgrade is in progress. 'unchanged' indicates no change was needed.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.KubernetesClusterUpdateResponseStatusScaling
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `KubernetesClusterUpdateResponseStatusScaling`   | scaling                                          |
| `KubernetesClusterUpdateResponseStatusUpgrading` | upgrading                                        |
| `KubernetesClusterUpdateResponseStatusUnchanged` | unchanged                                        |