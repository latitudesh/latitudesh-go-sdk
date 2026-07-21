# WorkerStatus

Current status of worker nodes. 'idle' when 0 workers, 'ready' when all workers are ready, 'scaling' when workers are being provisioned/removed, 'upgrading' while a Kubernetes version upgrade is rolling through the workers, 'error' when a worker has failed.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.WorkerStatusIdle

// Open enum: custom values can be created with a direct type cast
custom := components.WorkerStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `WorkerStatusIdle`      | idle                    |
| `WorkerStatusReady`     | ready                   |
| `WorkerStatusScaling`   | scaling                 |
| `WorkerStatusUpgrading` | upgrading               |
| `WorkerStatusError`     | error                   |