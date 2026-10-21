# RaidLevel

RAID level for the group. Only valid for the 'os' and 'storage' roles; 'raid-10' is storage-only, needs at least 4 disks, and is applied only at locations with the `raid_ten` feature

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.RaidLevelRaid0
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `RaidLevelRaid0`  | raid-0            |
| `RaidLevelRaid1`  | raid-1            |
| `RaidLevelRaid10` | raid-10           |