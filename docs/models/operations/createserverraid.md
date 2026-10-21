# CreateServerRaid

RAID mode for the server. Set to 'raid-0' for RAID 0, 'raid-1' for RAID 1, 'raid-10' for RAID 10 on the data disks (only at locations with the `raid_ten` feature), or omit/null for no RAID configuration

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.CreateServerRaidRaid0
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `CreateServerRaidRaid0`  | raid-0                   |
| `CreateServerRaidRaid1`  | raid-1                   |
| `CreateServerRaidRaid10` | raid-10                  |