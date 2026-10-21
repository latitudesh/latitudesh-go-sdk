# Raid

Expected RAID level with the default disk layout (data array on 4+ disk servers, the OS on 2-3 disk servers). Mutually exclusive with disk_layout

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.RaidRaid0
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `RaidRaid0`  | raid-0       |
| `RaidRaid1`  | raid-1       |
| `RaidRaid10` | raid-10      |