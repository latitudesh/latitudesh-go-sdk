# RetentionMode

Object lock retention mode. Requires `locking` to be true when not `NONE`.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.RetentionModeNone
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `RetentionModeNone`       | NONE                      |
| `RetentionModeCompliance` | COMPLIANCE                |
| `RetentionModeGovernance` | GOVERNANCE                |