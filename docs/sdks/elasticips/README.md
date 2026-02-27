# ElasticIPs

## Overview

### Available Operations

* [ListElasticIps](#listelasticips) - List Elastic IPs
* [CreateElasticIP](#createelasticip) - Create an Elastic IP
* [GetElasticIP](#getelasticip) - Retrieve an Elastic IP
* [DeleteElasticIP](#deleteelasticip) - Release an Elastic IP
* [UpdateElasticIP](#updateelasticip) - Move an Elastic IP

## ListElasticIps

List all Elastic IPs for the authenticated team. Elastic IPs are static public IP addresses that can be assigned to servers and moved between servers within the same project.

**Note:** This feature requires the `elastic_ips` feature flag to be enabled for your team. When the flag is disabled, the endpoint returns an empty list.


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

    res, err := s.ElasticIPs.ListElasticIps(ctx, operations.ListElasticIpsRequest{})
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

Creates a new Elastic IP and assigns it to the specified server. The IP is provisioned asynchronously—the response will show status `provisioning` and the `id` will be `null` until provisioning completes.

**Note:** This feature requires the `elastic_ips` feature flag to be enabled for your team. Currently only IPv4 /32 addresses in routed mode are supported.


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

    res, err := s.ElasticIPs.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
            Attributes: components.CreateElasticIPAttributes{
                ProjectID: "<id>",
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

    res, err := s.ElasticIPs.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
            Attributes: components.CreateElasticIPAttributes{
                ProjectID: "proj_AoW6vRnwkvLn0",
                ServerID: "sv_2GmAlJ6BXlK1a",
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

    res, err := s.ElasticIPs.CreateElasticIP(ctx, components.CreateElasticIP{
        Data: components.CreateElasticIPData{
            Type: components.CreateElasticIPTypeElasticIps,
            Attributes: components.CreateElasticIPAttributes{
                ProjectID: "<id>",
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
| `request`                                                                | [components.CreateElasticIP](../../models/components/createelasticip.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateElasticIPResponse](../../models/operations/createelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 422                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetElasticIP

Returns a single Elastic IP by its ID.

**Note:** This feature requires the `elastic_ips` feature flag to be enabled for your team.


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

    res, err := s.ElasticIPs.GetElasticIP(ctx, "eip_KeQbB4BoO6x10")
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
| `elasticIPID`                                            | *string*                                                 | :heavy_check_mark:                                       | The Elastic IP ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetElasticIPResponse](../../models/operations/getelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteElasticIP

Releases an Elastic IP, returning it to the available pool. The IP will transition to `releasing` status before being fully removed.

**Note:** This feature requires the `elastic_ips` feature flag to be enabled for your team. Only Elastic IPs with status `active` or `error` can be released.


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

    res, err := s.ElasticIPs.DeleteElasticIP(ctx, "<id>")
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
| `elasticIPID`                                            | *string*                                                 | :heavy_check_mark:                                       | The Elastic IP ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteElasticIPResponse](../../models/operations/deleteelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UpdateElasticIP

Moves an Elastic IP to a different server within the same project. The reassignment is performed asynchronously.

**Note:** This feature requires the `elastic_ips` feature flag to be enabled for your team. The Elastic IP must be in `active` status and the target server must belong to the same project.


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

    res, err := s.ElasticIPs.UpdateElasticIP(ctx, "<id>", components.UpdateElasticIP{
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

    res, err := s.ElasticIPs.UpdateElasticIP(ctx, "<id>", components.UpdateElasticIP{
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

    res, err := s.ElasticIPs.UpdateElasticIP(ctx, "eip_KeQbB4BoO6x10", components.UpdateElasticIP{
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
| `elasticIPID`                                                            | *string*                                                                 | :heavy_check_mark:                                                       | The Elastic IP ID                                                        |
| `updateElasticIP`                                                        | [components.UpdateElasticIP](../../models/components/updateelasticip.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.UpdateElasticIPResponse](../../models/operations/updateelasticipresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |