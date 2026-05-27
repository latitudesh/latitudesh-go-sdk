# RetentionMode

Object Lock retention mode applied to new objects. `GOVERNANCE` allows privileged users to override the retention; `COMPLIANCE` cannot be overridden by anyone. Only applies when `locking` is `true`.

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