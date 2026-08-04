# ElasticIps

## Overview

### Available Operations

* [ListElasticIps](#listelasticips) - List Elastic IPs
* [CreateElasticIP](#createelasticip) - Create an Elastic IP
* [GetElasticIP](#getelasticip) - Retrieve an Elastic IP
* [DeleteElasticIP](#deleteelasticip) - Release an Elastic IP
* [UpdateElasticIP](#updateelasticip) - Move an Elastic IP
* [ListElasticIPBgpSessions](#listelasticipbgpsessions) - List BGP sessions
* [CreateElasticIPBgpSession](#createelasticipbgpsession) - Create a BGP session
* [GetElasticIPBgpSession](#getelasticipbgpsession) - Retrieve a BGP session
* [DeleteElasticIPBgpSession](#deleteelasticipbgpsession) - Delete a BGP session

## ListElasticIps

List all Elastic IPs for the authenticated team. Elastic IPs are static public IP addresses that can be assigned to servers and moved between servers within the same project.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-elastic-ips" method="get" path="/elastic_ips" example="Success" -->
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

    res, err := s.ElasticIps.ListElasticIps(ctx, operations.ListElasticIpsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIps != nil {
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListElasticIpsRequest](../../models/operations/listelasticipsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListElasticIpsResponse](../../models/operations/listelasticipsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateElasticIP

Creates a new Elastic IP and assigns it to the specified server. The IP is provisioned asynchronously—the response will show status `configuring` and the `id` will be `null` until provisioning completes. Currently only IPv4 /32 addresses in routed mode are supported.


### Example Usage: Accepted

<!-- UsageSnippet language="go" operationID="create-elastic-ip" method="post" path="/elastic_ips" example="Accepted" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: Create

<!-- UsageSnippet language="go" operationID="create-elastic-ip" method="post" path="/elastic_ips" example="Create" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
            Attributes: &components.CreateElasticIPAttributes{
                ServerID: latitudeshgosdk.Pointer("sv_2GmAlJ6BXlK1a"),
                ProjectID: "proj_AoW6vRnwkvLn0",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: FeatureNotEnabled

<!-- UsageSnippet language="go" operationID="create-elastic-ip" method="post" path="/elastic_ips" example="FeatureNotEnabled" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
            Attributes: &components.CreateElasticIPAttributes{
                ServerID: latitudeshgosdk.Pointer("<id>"),
                ProjectID: "<id>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: IpAllocationFailed

<!-- UsageSnippet language="go" operationID="create-elastic-ip" method="post" path="/elastic_ips" example="IpAllocationFailed" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: ServerNetworkIncompatible

<!-- UsageSnippet language="go" operationID="create-elastic-ip" method="post" path="/elastic_ips" example="ServerNetworkIncompatible" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: ServerNotInProject

<!-- UsageSnippet language="go" operationID="create-elastic-ip" method="post" path="/elastic_ips" example="ServerNotInProject" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: SiteNotSupported

<!-- UsageSnippet language="go" operationID="create-elastic-ip" method="post" path="/elastic_ips" example="SiteNotSupported" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.CreateElasticIP](../../models/components/createelasticip.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateElasticIPResponse](../../models/operations/createelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 422                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetElasticIP

Returns a single Elastic IP by its ID.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-elastic-ip" method="get" path="/elastic_ips/{elastic_ip_id}" example="Success" -->
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

    res, err := s.ElasticIps.GetElasticIP(ctx, "eip_KeQbB4BoO6x10")
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `elasticIPID`                                            | `string`                                                 | :heavy_check_mark:                                       | The Elastic IP ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetElasticIPResponse](../../models/operations/getelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteElasticIP

Releases an Elastic IP, returning it to the available pool. The IP will transition to `releasing` status before being fully removed. Only Elastic IPs with status `active` or `error` can be released.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-elastic-ip" method="delete" path="/elastic_ips/{elastic_ip_id}" -->
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

    res, err := s.ElasticIps.DeleteElasticIP(ctx, "<id>")
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
| `elasticIPID`                                            | `string`                                                 | :heavy_check_mark:                                       | The Elastic IP ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteElasticIPResponse](../../models/operations/deleteelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404, 422                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UpdateElasticIP

Moves an Elastic IP to a different server within the same project and site. The reassignment is performed asynchronously. The Elastic IP must be in `active` status, the target server must belong to the same project, and the target server must be in the same site as the currently assigned server.


### Example Usage: FeatureNotEnabled

<!-- UsageSnippet language="go" operationID="update-elastic-ip" method="patch" path="/elastic_ips/{elastic_ip_id}" example="FeatureNotEnabled" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.UpdateElasticIP(ctx, "<id>", components.UpdateElasticIP{
        Data: components.UpdateElasticIPData{
            Type: components.UpdateElasticIPTypeElasticIps,
            Attributes: components.UpdateElasticIPAttributes{
                ServerID: "<id>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: Move

<!-- UsageSnippet language="go" operationID="update-elastic-ip" method="patch" path="/elastic_ips/{elastic_ip_id}" example="Move" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.UpdateElasticIP(ctx, "<id>", components.UpdateElasticIP{
        Data: components.UpdateElasticIPData{
            Type: components.UpdateElasticIPTypeElasticIps,
            Attributes: components.UpdateElasticIPAttributes{
                ServerID: "sv_oDEBlwBGRO2me",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="update-elastic-ip" method="patch" path="/elastic_ips/{elastic_ip_id}" example="Success" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.UpdateElasticIP(ctx, "eip_KeQbB4BoO6x10", components.UpdateElasticIP{
        Data: components.UpdateElasticIPData{
            Type: components.UpdateElasticIPTypeElasticIps,
            Attributes: components.UpdateElasticIPAttributes{
                ServerID: "<id>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ElasticIP != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `elasticIPID`                                                            | `string`                                                                 | :heavy_check_mark:                                                       | The Elastic IP ID                                                        |
| `updateElasticIP`                                                        | [components.UpdateElasticIP](../../models/components/updateelasticip.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.UpdateElasticIPResponse](../../models/operations/updateelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404, 422                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListElasticIPBgpSessions

List the BGP sessions announcing an elastic IP

### Example Usage

<!-- UsageSnippet language="go" operationID="list-elastic-ip-bgp-sessions" method="get" path="/elastic_ips/{elastic_ip_id}/bgp_sessions" example="Success" -->
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

    res, err := s.ElasticIps.ListElasticIPBgpSessions(ctx, "ueip_v9BVDaGvdRm1W")
    if err != nil {
        log.Fatal(err)
    }
    if res.BgpSessions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `elasticIPID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListElasticIPBgpSessionsResponse](../../models/operations/listelasticipbgpsessionsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateElasticIPBgpSession

Announce an elastic IP from a server over BGP

### Example Usage

<!-- UsageSnippet language="go" operationID="create-elastic-ip-bgp-session" method="post" path="/elastic_ips/{elastic_ip_id}/bgp_sessions" example="Accepted" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.ElasticIps.CreateElasticIPBgpSession(ctx, "ueip_059EqY7kOQj8p", components.CreateBgpSession{
        Data: components.CreateBgpSessionData{
            Type: components.CreateBgpSessionTypeBgpSessions,
            Attributes: &components.CreateBgpSessionAttributes{
                ServerID: "sv_A05EdQp4OvKYQ",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BgpSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `elasticIPID`                                                              | `string`                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `createBgpSession`                                                         | [components.CreateBgpSession](../../models/components/createbgpsession.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateElasticIPBgpSessionResponse](../../models/operations/createelasticipbgpsessionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetElasticIPBgpSession

Retrieve a BGP session announcing an elastic IP

### Example Usage

<!-- UsageSnippet language="go" operationID="get-elastic-ip-bgp-session" method="get" path="/elastic_ips/{elastic_ip_id}/bgp_sessions/{id}" example="Success" -->
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

    res, err := s.ElasticIps.GetElasticIPBgpSession(ctx, "ueip_e8pKq015DWAob", "bgps_w49QDB9PqagKb")
    if err != nil {
        log.Fatal(err)
    }
    if res.BgpSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `elasticIPID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetElasticIPBgpSessionResponse](../../models/operations/getelasticipbgpsessionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteElasticIPBgpSession

Stop announcing an elastic IP from a server

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-elastic-ip-bgp-session" method="delete" path="/elastic_ips/{elastic_ip_id}/bgp_sessions/{id}" -->
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

    res, err := s.ElasticIps.DeleteElasticIPBgpSession(ctx, "<id>", "<id>", nil)
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
| `elasticIPID`                                            | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `force`                                                  | `*bool`                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteElasticIPBgpSessionResponse](../../models/operations/deleteelasticipbgpsessionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |