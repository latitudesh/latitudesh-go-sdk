# SyncStatus

Control-plane refresh observation. Synced confirms the receiver observed an active key; revoked confirms it observed the key absent. Other regions reconcile independently.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.SyncStatusPending
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `SyncStatusPending` | pending             |
| `SyncStatusSynced`  | synced              |
| `SyncStatusFailed`  | failed              |
| `SyncStatusRevoked` | revoked             |