# Environment

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.EnvironmentDevelopment

// Open enum: custom values can be created with a direct type cast
custom := components.Environment("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `EnvironmentDevelopment` | Development              |
| `EnvironmentStaging`     | Staging                  |
| `EnvironmentProduction`  | Production               |