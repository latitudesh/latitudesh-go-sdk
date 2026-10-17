# CreateLksNodePoolMode

on_demand provisions new servers of `plan` from stock. allocated builds the pool from `server_ids`, servers already in the cluster project: each one is reinstalled with the LKS node image (erasing its disks) and, when it leaves the pool, reinstalled with a clean OS (no data, keys or user data are restored).

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.CreateLksNodePoolModeOnDemand
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CreateLksNodePoolModeOnDemand`  | on_demand                        |
| `CreateLksNodePoolModeAllocated` | allocated                        |