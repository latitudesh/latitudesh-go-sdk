# AccessKeyStorageClass

Backend storage tier. `standard` provisions the key on Wasabi; `high_performance` provisions it on VAST.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.AccessKeyStorageClassStandard
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AccessKeyStorageClassStandard`        | standard                               |
| `AccessKeyStorageClassHighPerformance` | high_performance                       |