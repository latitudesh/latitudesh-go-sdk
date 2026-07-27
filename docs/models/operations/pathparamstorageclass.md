# PathParamStorageClass

Backend storage tier of the access key. `standard` targets Wasabi; `high_performance` targets VAST.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.PathParamStorageClassStandard
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `PathParamStorageClassStandard`        | standard                               |
| `PathParamStorageClassHighPerformance` | high_performance                       |