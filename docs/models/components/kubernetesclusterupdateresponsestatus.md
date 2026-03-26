# KubernetesClusterUpdateResponseStatus

The update status. 'scaling' indicates worker nodes are being added or removed. 'unchanged' indicates the requested worker_count matches the current count.

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
| `KubernetesClusterUpdateResponseStatusUnchanged` | unchanged                                        |