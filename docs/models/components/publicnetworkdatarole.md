# PublicNetworkDataRole

gateway: reserved for the network gateway; server: a server on the network; elastic_ip: an elastic IP; reserved: held in IPAM but not by a server; available: free to use

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.PublicNetworkDataRoleGateway
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PublicNetworkDataRoleGateway`   | gateway                          |
| `PublicNetworkDataRoleServer`    | server                           |
| `PublicNetworkDataRoleElasticIP` | elastic_ip                       |
| `PublicNetworkDataRoleReserved`  | reserved                         |
| `PublicNetworkDataRoleAvailable` | available                        |