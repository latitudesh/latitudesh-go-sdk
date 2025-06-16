# Servers
(*Servers*)

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
        FilterProject: latitudeshgosdk.String("proj_g1mbDwrZqLv5B"),
        FilterRegion: latitudeshgosdk.String("SAO"),
        FilterRAMEql: latitudeshgosdk.Int64(32),
        FilterRAMGte: latitudeshgosdk.Int64(40),
        FilterRAMLte: latitudeshgosdk.Int64(40),
        FilterTags: latitudeshgosdk.String("tag_0yrQNVQRLwHy0XwEGM6ESwLrW2PA"),
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 422                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Create

Deploy Server

### Example Usage

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
                Project: latitudeshgosdk.String("proj_W6Q2D93GdKLpr"),
                Plan: operations.CreateServerPlanC2SmallX86.ToPointer(),
                Site: operations.CreateServerSiteSao.ToPointer(),
                OperatingSystem: operations.CreateServerOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.String("BRC1"),
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 400, 402, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Get

Returns a server that belongs to the team.


### Example Usage

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

    res, err := s.Servers.Get(ctx, "sv_aNmodjGeqbE8W", nil)
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

    res, err := s.Servers.Update(ctx, "sv_3YjJOLLNOvZ87", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.String("sv_3YjJOLLNOvZ87"),
            Type: operations.UpdateServerServersRequestApplicationJSONTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersRequestApplicationJSONAttributes{},
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `serverID`                                                                                             | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `requestBody`                                                                                          | [operations.UpdateServerServersRequestBody](../../models/operations/updateserverserversrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateServerResponse](../../models/operations/updateserverresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ServerError   | 400, 422                 | application/vnd.api+json |
| components.ErrorObject   | 402, 423                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Delete

Remove Server

### Example Usage

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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 406, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetDeployConfig

Retrieve Deploy Config

### Example Usage

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

    res, err := s.Servers.GetDeployConfig(ctx, "sv_pRMLydp0dQKr1")
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

    res, err := s.Servers.UpdateDeployConfig(ctx, "sv_g1mbDweZdLv5B", operations.UpdateServerDeployConfigServersRequestBody{
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

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| components.ErrorObject       | 403, 406                     | application/vnd.api+json     |
| components.DeployConfigError | 422                          | application/vnd.api+json     |
| components.APIError          | 4XX, 5XX                     | \*/\*                        |

## Lock

Locks the server. A locked server cannot be deleted or modified and no actions can be performed on it.

### Example Usage

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

    res, err := s.Servers.Lock(ctx, "sv_059EqYX2dQj8p")
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

    res, err := s.Servers.Unlock(ctx, "sv_aNmodjoyqbE8W")
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

    res, err := s.Servers.StartOutOfBandConnection(ctx, "sv_z2A3DVpQdnawP", operations.CreateServerOutOfBandServersRequestBody{
        Data: operations.CreateServerOutOfBandServersData{
            Type: operations.CreateServerOutOfBandServersTypeOutOfBand,
            Attributes: &operations.CreateServerOutOfBandServersAttributes{
                SSHKeyID: latitudeshgosdk.String("ssh_NGnzRD5ADM5yw"),
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetOutOfBand

List Out of Band Connections

### Example Usage

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

    res, err := s.Servers.GetOutOfBand(ctx, "sv_1ZJrdx34Og4LV")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## RunAction

Performs an action on a given server:
- `power_on`
- `power_off`
- `reboot`


### Example Usage

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

    res, err := s.Servers.RunAction(ctx, "sv_LYV8DZAQq5QoE", operations.CreateServerActionServersRequestBody{
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateIpmiSession

Generates IPMI credentials for a given server. Remote access creates a VPN connection to the internal network of your server so you can connect to its IPMI.
You will have to use a VPN client such as https://openvpn.net to connect. See `VPN Sessions` API to create a VPN connection.

Related guide: https://docs.latitude.sh/docs/ipmi


### Example Usage

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

    res, err := s.Servers.CreateIpmiSession(ctx, "sv_8NkvdyGKDeLpx")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## StartRescueMode

Starts rescue mode on a given server.

### Example Usage

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

    res, err := s.Servers.StartRescueMode(ctx, "sv_k0Ryqv9adW36X")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 406                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ExitRescueMode

Exits rescue mode on a given server.

### Example Usage

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

    res, err := s.Servers.ExitRescueMode(ctx, "sv_KXgRdRRodv9k5")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 406                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ScheduleDeletion

Schedules the server to be removed at the end of the billing cycle.

### Example Usage

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

    res, err := s.Servers.ScheduleDeletion(ctx, "sv_enPbqoBJdA2MQ")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 406, 423            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UnscheduleDeletion

Unschedules the server removal at the end of the billing cycle.

### Example Usage

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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Reinstall

Run Server Reinstall

### Example Usage

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

    res, err := s.Servers.Reinstall(ctx, "sv_WeGoqAWNOP7nz", operations.CreateServerReinstallServersRequestBody{
        Data: operations.CreateServerReinstallServersData{
            Type: operations.CreateServerReinstallServersTypeReinstalls,
            Attributes: &operations.CreateServerReinstallServersAttributes{
                OperatingSystem: operations.CreateServerReinstallServersOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.String("BRC1"),
                Partitions: []operations.CreateServerReinstallServersPartitions{
                    operations.CreateServerReinstallServersPartitions{
                        SizeInGb: latitudeshgosdk.Int64(300),
                        Path: latitudeshgosdk.String("/"),
                        FilesystemType: latitudeshgosdk.String("ext4"),
                    },
                },
                SSHKeys: []string{
                    "35",
                },
                UserData: latitudeshgosdk.Int64(10),
                Raid: operations.CreateServerReinstallServersRaidRaid1.ToPointer(),
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.ServerError   | 423                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |