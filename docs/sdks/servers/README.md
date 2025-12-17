# Servers

## Overview

### Available Operations

* [List](#list) - List all Servers
* [Create](#create) - Deploy Server
* [Get](#get) - Retrieve a Server
* [Update](#update) - Update Server
* [Delete](#delete) - Remove Server
* [GetDeployConfig](#getdeployconfig) - Retrieve Deploy Config
* [UpdateDeployConfig](#updatedeployconfig) - Update Deploy Config
* [Lock](#lock) - Lock the server
* [Unlock](#unlock) - Unlock the server
* [StartOutOfBandConnection](#startoutofbandconnection) - Start Out of Band Connection
* [GetOutOfBand](#getoutofband) - List Out of Band Connections
* [RunAction](#runaction) - Run Server Action
* [CreateIpmiSession](#createipmisession) - Generate IPMI credentials
* [StartRescueMode](#startrescuemode) - Puts a Server in rescue mode
* [ExitRescueMode](#exitrescuemode) - Exits rescue mode for a Server
* [ScheduleDeletion](#scheduledeletion) - Schedule the server deletion
* [UnscheduleDeletion](#unscheduledeletion) - Unschedule the server deletion
* [Reinstall](#reinstall) - Run Server Reinstall

## List

Returns a list of all servers belonging to the team.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.List(ctx, operations.GetServersRequest{
        FilterProject: latitudeshgosdk.Pointer("proj_g1mbDwrZqLv5B"),
        FilterRegion: latitudeshgosdk.Pointer("SAO"),
        FilterRAMEql: latitudeshgosdk.Pointer[int64](32),
        FilterRAMGte: latitudeshgosdk.Pointer[int64](40),
        FilterRAMLte: latitudeshgosdk.Pointer[int64](40),
        FilterTags: latitudeshgosdk.Pointer("tag_Az0EY3zglei3jVBY1LroSWNyanye,tag_GXK6NGol1jF2xre0JrB0fK6wg0p"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Servers != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetServersRequest](../../models/operations/getserversrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetServersResponse](../../models/operations/getserversresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Deploy Server

### Example Usage

<!-- UsageSnippet language="go" operationID="create-server" method="post" path="/servers" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.Create(ctx, operations.CreateServerServersRequestBody{
        Data: &operations.CreateServerServersData{
            Type: operations.CreateServerServersTypeServers,
            Attributes: &operations.CreateServerServersAttributes{
                Project: latitudeshgosdk.Pointer("proj_A05EdQ50dvKYQ"),
                Plan: operations.CreateServerPlanC2SmallX86.ToPointer(),
                Site: operations.CreateServerSiteAsh.ToPointer(),
                OperatingSystem: operations.CreateServerOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.CreateServerServersRequestBody](../../models/operations/createserverserversrequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateServerResponse](../../models/operations/createserverresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Get

Returns a server that belongs to the team.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-server" method="get" path="/servers/{server_id}" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.Get(ctx, "sv_Gr47qleMDAg0m", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `serverID`                                                                                                                                               | *string*                                                                                                                                                 | :heavy_check_mark:                                                                                                                                       | The Server ID                                                                                                                                            |
| `extraFieldsServers`                                                                                                                                     | **string*                                                                                                                                                | :heavy_minus_sign:                                                                                                                                       | The `credentials` are provided as extra attributes that is lazy loaded. To request it, just set `extra_fields[servers]=credentials` in the query string. |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.GetServerResponse](../../models/operations/getserverresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Update Server

### Example Usage

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.Update(ctx, "sv_aNmodjGyqbE8W", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_aNmodjGyqbE8W"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Project: latitudeshgosdk.Pointer("proj_aNmodjoyqbE8W"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `serverID`                                                                                             | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `requestBody`                                                                                          | [operations.UpdateServerServersRequestBody](../../models/operations/updateserverserversrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateServerResponse](../../models/operations/updateserverresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Remove Server

### Example Usage

<!-- UsageSnippet language="go" operationID="destroy-server" method="delete" path="/servers/{server_id}" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.Delete(ctx, "sv_WeGoqAZNDP7nz", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | The server ID                                            |
| `reason`                                                 | **string*                                                | :heavy_minus_sign:                                       | The reason for deleting the server                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DestroyServerResponse](../../models/operations/destroyserverresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetDeployConfig

Retrieve Deploy Config

### Example Usage

<!-- UsageSnippet language="go" operationID="get-server-deploy-config" method="get" path="/servers/{server_id}/deploy_config" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.GetDeployConfig(ctx, "sv_VLMmAD8EOwop2")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeployConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | The Server ID                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetServerDeployConfigResponse](../../models/operations/getserverdeployconfigresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateDeployConfig

Update Deploy Config

### Example Usage

<!-- UsageSnippet language="go" operationID="update-server-deploy-config" method="patch" path="/servers/{server_id}/deploy_config" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.UpdateDeployConfig(ctx, "sv_0L6WO141DPlXy", operations.UpdateServerDeployConfigServersRequestBody{
        Type: operations.UpdateServerDeployConfigServersTypeDeployConfig,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeployConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `serverID`                                                                                                                     | *string*                                                                                                                       | :heavy_check_mark:                                                                                                             | The Server ID                                                                                                                  |
| `requestBody`                                                                                                                  | [operations.UpdateServerDeployConfigServersRequestBody](../../models/operations/updateserverdeployconfigserversrequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.UpdateServerDeployConfigResponse](../../models/operations/updateserverdeployconfigresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Lock

Locks the server. A locked server cannot be deleted or modified and no actions can be performed on it.

### Example Usage

<!-- UsageSnippet language="go" operationID="server-lock" method="post" path="/servers/{server_id}/lock" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.Lock(ctx, "sv_RMLydpoXOQKr1")
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServerLockResponse](../../models/operations/serverlockresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Unlock

Unlocks the server. A locked server cannot be deleted or modified and no actions can be performed on it.

### Example Usage

<!-- UsageSnippet language="go" operationID="server-unlock" method="post" path="/servers/{server_id}/unlock" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.Unlock(ctx, "sv_5AEmq7xMDBkWX")
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServerUnlockResponse](../../models/operations/serverunlockresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## StartOutOfBandConnection

Start Out of Band Connection

### Example Usage

<!-- UsageSnippet language="go" operationID="create-server-out-of-band" method="post" path="/servers/{server_id}/out_of_band_connection" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.StartOutOfBandConnection(ctx, "sv_059EqYX2dQj8p", operations.CreateServerOutOfBandServersRequestBody{
        Data: operations.CreateServerOutOfBandServersData{
            Type: operations.CreateServerOutOfBandServersTypeOutOfBand,
            Attributes: &operations.CreateServerOutOfBandServersAttributes{
                SSHKeyID: latitudeshgosdk.Pointer("ssh_w49QDB55qagKb"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OutOfBandConnection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `serverID`                                                                                                               | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `requestBody`                                                                                                            | [operations.CreateServerOutOfBandServersRequestBody](../../models/operations/createserveroutofbandserversrequestbody.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreateServerOutOfBandResponse](../../models/operations/createserveroutofbandresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetOutOfBand

List Out of Band Connections

### Example Usage

<!-- UsageSnippet language="go" operationID="get-server-out-of-band" method="get" path="/servers/{server_id}/out_of_band_connection" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.GetOutOfBand(ctx, "sv_vYAZqGyJOMQ94")
    if err != nil {
        log.Fatal(err)
    }
    if res.OutOfBandConnection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | The Server ID                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetServerOutOfBandResponse](../../models/operations/getserveroutofbandresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## RunAction

Performs an action on a given server:
- `power_on`
- `power_off`
- `reboot`


### Example Usage

<!-- UsageSnippet language="go" operationID="create-server-action" method="post" path="/servers/{server_id}/actions" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.RunAction(ctx, "sv_LA73qkJwdaJ2o", operations.CreateServerActionServersRequestBody{
        Data: operations.CreateServerActionServersData{
            Type: operations.CreateServerActionServersTypeActions,
            Attributes: &operations.CreateServerActionServersAttributes{
                Action: operations.CreateServerActionActionReboot,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerAction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `serverID`                                                                                                         | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `requestBody`                                                                                                      | [operations.CreateServerActionServersRequestBody](../../models/operations/createserveractionserversrequestbody.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateServerActionResponse](../../models/operations/createserveractionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateIpmiSession

Generates IPMI credentials for a given server. Remote access creates a VPN connection to the internal network of your server so you can connect to its IPMI.
You will have to use a VPN client such as https://openvpn.net to connect. See `VPN Sessions` API to create a VPN connection.

Related guide: https://docs.latitude.sh/docs/ipmi


### Example Usage

<!-- UsageSnippet language="go" operationID="create-ipmi-session" method="post" path="/servers/{server_id}/remote_access" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.CreateIpmiSession(ctx, "sv_e8pKq0xYqWAob")
    if err != nil {
        log.Fatal(err)
    }
    if res.IpmiSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateIpmiSessionResponse](../../models/operations/createipmisessionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## StartRescueMode

Starts rescue mode on a given server.

### Example Usage

<!-- UsageSnippet language="go" operationID="server-start-rescue-mode" method="post" path="/servers/{server_id}/rescue_mode" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.StartRescueMode(ctx, "sv_695BdK25OevVo")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerRescue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServerStartRescueModeResponse](../../models/operations/serverstartrescuemoderesponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ExitRescueMode

Exits rescue mode on a given server.

### Example Usage

<!-- UsageSnippet language="go" operationID="server-exit-rescue-mode" method="post" path="/servers/{server_id}/exit_rescue_mode" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.ExitRescueMode(ctx, "sv_wg3ZDr0Wd5QlP")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerRescue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServerExitRescueModeResponse](../../models/operations/serverexitrescuemoderesponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ScheduleDeletion

Schedules the server to be removed at the end of the billing cycle.

### Example Usage

<!-- UsageSnippet language="go" operationID="server-schedule-deletion" method="post" path="/servers/{server_id}/schedule_deletion" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.ScheduleDeletion(ctx, "sv_GMy1Db2NDN50m")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerScheduleDeletion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServerScheduleDeletionResponse](../../models/operations/serverscheduledeletionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UnscheduleDeletion

Unschedules the server removal at the end of the billing cycle.

### Example Usage

<!-- UsageSnippet language="go" operationID="server-unschedule-deletion" method="delete" path="/servers/{server_id}/schedule_deletion" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.UnscheduleDeletion(ctx, "sv_Z8rodmJGq1jLB")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serverID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ServerUnscheduleDeletionResponse](../../models/operations/serverunscheduledeletionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Reinstall

Run Server Reinstall

### Example Usage

<!-- UsageSnippet language="go" operationID="create-server-reinstall" method="post" path="/servers/{server_id}/reinstall" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Servers.Reinstall(ctx, "sv_Z8rodmJGq1jLB", operations.CreateServerReinstallServersRequestBody{
        Data: operations.CreateServerReinstallServersData{
            Type: operations.CreateServerReinstallServersTypeReinstalls,
            Attributes: &operations.CreateServerReinstallServersAttributes{
                OperatingSystem: operations.CreateServerReinstallServersOperatingSystemIpxe.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
                Ipxe: latitudeshgosdk.Pointer("https://some-host.com/image.ipxe"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `serverID`                                                                                                               | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `requestBody`                                                                                                            | [operations.CreateServerReinstallServersRequestBody](../../models/operations/createserverreinstallserversrequestbody.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CreateServerReinstallResponse](../../models/operations/createserverreinstallresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |