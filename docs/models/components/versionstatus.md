# VersionStatus

The cluster's version status relative to available upgrades

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.VersionStatusUpToDate
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `VersionStatusUpToDate`         | up_to_date                      |
| `VersionStatusUpgradeAvailable` | upgrade_available               |
| `VersionStatusUnsupported`      | unsupported                     |
| `VersionStatusUnknown`          | unknown                         |