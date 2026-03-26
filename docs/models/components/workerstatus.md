# WorkerStatus

Current status of worker nodes. 'idle' when 0 workers, 'ready' when all workers are ready, 'scaling' when workers are being provisioned/removed, 'error' when a worker has failed.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.WorkerStatusIdle
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `WorkerStatusIdle`    | idle                  |
| `WorkerStatusReady`   | ready                 |
| `WorkerStatusScaling` | scaling               |
| `WorkerStatusError`   | error                 |