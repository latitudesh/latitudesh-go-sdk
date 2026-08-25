# DeploymentStrategy

How the app is delivered: cloud-init install on a stock OS image (user_data) or a pre-built disk image (image)

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.DeploymentStrategyUserData
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `DeploymentStrategyUserData` | user_data                    |
| `DeploymentStrategyImage`    | image                        |