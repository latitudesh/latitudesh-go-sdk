# PostStorageBucketsStorageClass

Backend storage tier. `standard` is the default S3-compatible tier. `high_performance` is a lower-latency, higher-throughput tier available in select regions only.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.PostStorageBucketsStorageClassStandard
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `PostStorageBucketsStorageClassStandard`        | standard                                        |
| `PostStorageBucketsStorageClassHighPerformance` | high_performance                                |