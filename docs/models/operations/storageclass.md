# StorageClass

Backend storage tier. `standard` provisions the key on Wasabi; `high_performance` provisions it on VAST.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.StorageClassStandard
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `StorageClassStandard`        | standard                      |
| `StorageClassHighPerformance` | high_performance              |