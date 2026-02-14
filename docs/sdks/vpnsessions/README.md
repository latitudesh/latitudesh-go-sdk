# VpnSessions

## Overview

### Available Operations

* [List](#list) - List VPN sessions
* [Create](#create) - Create VPN session
* [RefreshPassword](#refreshpassword) - Refresh VPN session
* [Delete](#delete) - Delete VPN session

## List

List VPN sessions

### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-vpn-sessions" method="get" path="/vpn_sessions" example="Success" -->
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

    res, err := s.VpnSessions.List(ctx, operations.FilterLocationSao.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: when there is a site filter

<!-- UsageSnippet language="go" operationID="get-vpn-sessions" method="get" path="/vpn_sessions" example="when there is a site filter" -->
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

    res, err := s.VpnSessions.List(ctx, operations.FilterLocationSao.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `filterLocation`                                                        | [*operations.FilterLocation](../../models/operations/filterlocation.md) | :heavy_minus_sign:                                                      | N/A                                                                     |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |

### Response

**[*operations.GetVpnSessionsResponse](../../models/operations/getvpnsessionsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Creates a new VPN Session.
`NOTE:` The VPN credentials are only listed ONCE upon creation. They can however be refreshed or deleted.


### Example Usage: Created

<!-- UsageSnippet language="go" operationID="post-vpn-session" method="post" path="/vpn_sessions" example="Created" -->
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

    res, err := s.VpnSessions.Create(ctx, operations.PostVpnSessionVpnSessionsRequestBody{
        Data: &operations.PostVpnSessionVpnSessionsData{
            Attributes: &operations.PostVpnSessionVpnSessionsAttributes{
                Site: operations.PostVpnSessionVpnSessionsSiteSao.ToPointer(),
                ServerID: latitudeshgosdk.Pointer("sv_LMmAD8wyqwop2"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VpnSessionWithPassword != nil {
        // handle response
    }
}
```
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="post-vpn-session" method="post" path="/vpn_sessions" example="Unprocessable Entity" -->
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

    res, err := s.VpnSessions.Create(ctx, operations.PostVpnSessionVpnSessionsRequestBody{
        Data: &operations.PostVpnSessionVpnSessionsData{
            Attributes: &operations.PostVpnSessionVpnSessionsAttributes{
                Site: operations.PostVpnSessionVpnSessionsSiteSyd.ToPointer(),
                ServerID: latitudeshgosdk.Pointer("sv_GnzRD5BYOM5yw"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VpnSessionWithPassword != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostVpnSessionVpnSessionsRequestBody](../../models/operations/postvpnsessionvpnsessionsrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.PostVpnSessionResponse](../../models/operations/postvpnsessionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## RefreshPassword

Refreshing an existing VPN Session will create new credentials for that session


### Example Usage: Success

<!-- UsageSnippet language="go" operationID="put-vpn-session" method="patch" path="/vpn_sessions/{vpn_session_id}/refresh_password" example="Success" -->
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

    res, err := s.VpnSessions.RefreshPassword(ctx, "vpn_pRMLydp0dQKr1")
    if err != nil {
        log.Fatal(err)
    }
    if res.VpnSessionWithPassword != nil {
        // handle response
    }
}
```
### Example Usage: VPN Session Password Reset

<!-- UsageSnippet language="go" operationID="put-vpn-session" method="patch" path="/vpn_sessions/{vpn_session_id}/refresh_password" example="VPN Session Password Reset" -->
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

    res, err := s.VpnSessions.RefreshPassword(ctx, "vpn_6VE1Wd37dXnZJ")
    if err != nil {
        log.Fatal(err)
    }
    if res.VpnSessionWithPassword != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `vpnSessionID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PutVpnSessionResponse](../../models/operations/putvpnsessionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Deletes an existing VPN Session.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-vpn-session" method="delete" path="/vpn_sessions/{vpn_session_id}" -->
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

    res, err := s.VpnSessions.Delete(ctx, "invalid")
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
| `vpnSessionID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVpnSessionResponse](../../models/operations/deletevpnsessionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |