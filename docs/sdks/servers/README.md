# Servers

## Overview

### Available Operations

* [List](#list) - List servers
* [Create](#create) - Create server
* [Get](#get) - Retrieve server
* [Update](#update) - Update server
* [Delete](#delete) - Remove server
* [GetDeployConfig](#getdeployconfig) - Retrieve deploy config
* [UpdateDeployConfig](#updatedeployconfig) - Update deploy config
* [Lock](#lock) - Lock server
* [Unlock](#unlock) - Unlock server
* [StartOutOfBandConnection](#startoutofbandconnection) - Create out-of-band connection
* [GetOutOfBand](#getoutofband) - List out-of-band connections
* [RunAction](#runaction) - Run power action
* [CreateIpmiSession](#createipmisession) - Create IPMI credentials
* [StartRescueMode](#startrescuemode) - Put server in rescue mode
* [ExitRescueMode](#exitrescuemode) - Exits rescue mode
* [ScheduleDeletion](#scheduledeletion) - Schedule server deletion
* [UnscheduleDeletion](#unscheduledeletion) - Unschedule server deletion
* [Reinstall](#reinstall) - Run Server Reinstall

## List

Returns a list of all servers belonging to the team.


### Example Usage: Filtered by multiple tags

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="Filtered by multiple tags" -->
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
        FilterTags: latitudeshgosdk.Pointer("tag_3lg8RjPJL7HK4jYBbLmxC0ZR0yVX,tag_5JgXW7Wyr6i5V85g2Y2MFWL8bXA"),
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="Success" -->
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
        FilterTags: latitudeshgosdk.Pointer("tag_pjAkRjVzw0tlYBA2WX1eHzW7w79,tag_yARk1KLJAvslWY7k5wNBCaKEV7e"),
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
### Example Usage: SuccessWithCredentials

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="SuccessWithCredentials" -->
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

    res, err := s.Servers.List(ctx, operations.GetServersRequest{})
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
### Example Usage: when no filters

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="when no filters" -->
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

    res, err := s.Servers.List(ctx, operations.GetServersRequest{})
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
### Example Usage: when project filter

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="when project filter" -->
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
### Example Usage: when region filter

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="when region filter" -->
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
        FilterRegion: latitudeshgosdk.Pointer("SAO"),
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
### Example Usage: when trying to filter by tag

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="when trying to filter by tag" -->
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
        FilterTags: latitudeshgosdk.Pointer("tag_0yrQNVQRLwHy0XwEGM6ESwLrW2PA"),
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
### Example Usage: with a `eql` filter

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="with a `eql` filter" -->
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
        FilterRAMEql: latitudeshgosdk.Pointer[int64](32),
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
### Example Usage: with a `gte` filter

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="with a `gte` filter" -->
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
        FilterRAMGte: latitudeshgosdk.Pointer[int64](40),
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
### Example Usage: with a `lte` filter

<!-- UsageSnippet language="go" operationID="get-servers" method="get" path="/servers" example="with a `lte` filter" -->
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
        FilterRAMLte: latitudeshgosdk.Pointer[int64](40),
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

Create server

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-server" method="post" path="/servers" example="Created" -->
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
                Project: latitudeshgosdk.Pointer("proj_lxWpD699qm6rk"),
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
### Example Usage: Payment Required

<!-- UsageSnippet language="go" operationID="create-server" method="post" path="/servers" example="Payment Required" -->
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
                Project: latitudeshgosdk.Pointer("proj_kjQwdEMXdYNVP"),
                Plan: operations.CreateServerPlanC2SmallX86.ToPointer(),
                Site: operations.CreateServerSiteSao.ToPointer(),
                OperatingSystem: operations.CreateServerOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
                Billing: operations.CreateServerBillingMonthly.ToPointer(),
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
### Example Usage: Unpermited parameter

<!-- UsageSnippet language="go" operationID="create-server" method="post" path="/servers" example="Unpermited parameter" -->
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
            Attributes: &operations.CreateServerServersAttributes{},
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
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="create-server" method="post" path="/servers" example="Unprocessable Entity" -->
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
                Project: latitudeshgosdk.Pointer("proj_RMLydp9XqQKr1"),
                Plan: operations.CreateServerPlanC2SmallX86.ToPointer(),
                Site: operations.CreateServerSiteSao.ToPointer(),
                OperatingSystem: operations.CreateServerOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
                SSHKeys: []string{
                    "ssh_93YjJOLydvZ87",
                },
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


### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-server" method="get" path="/servers/{server_id}" example="Success" -->
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

    res, err := s.Servers.Get(ctx, "sv_VE1Wd3aXDXnZJ", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```
### Example Usage: SuccessWithCredentials

<!-- UsageSnippet language="go" operationID="get-server" method="get" path="/servers/{server_id}" example="SuccessWithCredentials" -->
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

    res, err := s.Servers.Get(ctx, "<id>", nil)
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

Update server

### Example Usage: Locked server

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="Locked server" -->
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

    res, err := s.Servers.Update(ctx, "sv_g1mbDwrZqLv5B", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_g1mbDwrZqLv5B"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Billing: operations.UpdateServerServersBillingHourly.ToPointer(),
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
### Example Usage: Payment Required

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="Payment Required" -->
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

    res, err := s.Servers.Update(ctx, "sv_1ZJrdxe4qg4LV", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_1ZJrdxe4qg4LV"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Billing: operations.UpdateServerServersBillingMonthly.ToPointer(),
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
### Example Usage: Reserved project

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="Reserved project" -->
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

    res, err := s.Servers.Update(ctx, "sv_w49QDBaQqagKb", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_w49QDBaQqagKb"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Billing: operations.UpdateServerServersBillingHourly.ToPointer(),
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="Success" -->
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

    res, err := s.Servers.Update(ctx, "sv_yQrJdNAGO30gv", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_yQrJdNAGO30gv"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Project: latitudeshgosdk.Pointer("proj_yQrJdNMGO30gv"),
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
### Example Usage: Unpermitted parameter

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="Unpermitted parameter" -->
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
            ID: latitudeshgosdk.Pointer("sv_3YjJOLLNOvZ87"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Billing: operations.UpdateServerServersBillingMonthly.ToPointer(),
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
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="Unprocessable Entity" -->
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

    res, err := s.Servers.Update(ctx, "sv_vYAZqGNJdMQ94", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_vYAZqGNJdMQ94"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Billing: operations.UpdateServerServersBillingMonthly.ToPointer(),
                Project: latitudeshgosdk.Pointer("new-project"),
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
### Example Usage: when server has additional ips

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="when server has additional ips" -->
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

    res, err := s.Servers.Update(ctx, "sv_W6Q2D9lGqKLpr", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_W6Q2D9lGqKLpr"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Billing: operations.UpdateServerServersBillingYearly.ToPointer(),
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
### Example Usage: when server has vlan assignments

<!-- UsageSnippet language="go" operationID="update-server" method="patch" path="/servers/{server_id}" example="when server has vlan assignments" -->
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

    res, err := s.Servers.Update(ctx, "sv_LMmAD8k4qwop2", operations.UpdateServerServersRequestBody{
        Data: &operations.UpdateServerServersData{
            ID: latitudeshgosdk.Pointer("sv_LMmAD8k4qwop2"),
            Type: operations.UpdateServerServersTypeServers.ToPointer(),
            Attributes: &operations.UpdateServerServersAttributes{
                Billing: operations.UpdateServerServersBillingYearly.ToPointer(),
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

Remove server

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

Retrieve deploy config

### Example Usage

<!-- UsageSnippet language="go" operationID="get-server-deploy-config" method="get" path="/servers/{server_id}/deploy_config" example="Success" -->
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

Update deploy config

### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="update-server-deploy-config" method="patch" path="/servers/{server_id}/deploy_config" example="Forbidden" -->
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

    res, err := s.Servers.UpdateDeployConfig(ctx, "sv_kjQwdEmXdYNVP", operations.UpdateServerDeployConfigServersRequestBody{
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
### Example Usage: Not Acceptable

<!-- UsageSnippet language="go" operationID="update-server-deploy-config" method="patch" path="/servers/{server_id}/deploy_config" example="Not Acceptable" -->
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

    res, err := s.Servers.UpdateDeployConfig(ctx, "sv_5AEmq7xMDBkWX", operations.UpdateServerDeployConfigServersRequestBody{
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="update-server-deploy-config" method="patch" path="/servers/{server_id}/deploy_config" example="Success" -->
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

    res, err := s.Servers.UpdateDeployConfig(ctx, "sv_lkg1DeYLDvZE5", operations.UpdateServerDeployConfigServersRequestBody{
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
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="update-server-deploy-config" method="patch" path="/servers/{server_id}/deploy_config" example="Unprocessable Entity" -->
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

    res, err := s.Servers.UpdateDeployConfig(ctx, "sv_Gr47ql4vqAg0m", operations.UpdateServerDeployConfigServersRequestBody{
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

### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="server-lock" method="post" path="/servers/{server_id}/lock" example="Forbidden" -->
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

    res, err := s.Servers.Lock(ctx, "sv_VE1Wd3rXOXnZJ")
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="server-lock" method="post" path="/servers/{server_id}/lock" example="Not Found" -->
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

    res, err := s.Servers.Lock(ctx, "invalid")
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```
### Example Usage: Ok

<!-- UsageSnippet language="go" operationID="server-lock" method="post" path="/servers/{server_id}/lock" example="Ok" -->
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="server-lock" method="post" path="/servers/{server_id}/lock" example="Success" -->
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

    res, err := s.Servers.Lock(ctx, "sv_pbV0DgjKq4AWz")
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

### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="server-unlock" method="post" path="/servers/{server_id}/unlock" example="Forbidden" -->
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

    res, err := s.Servers.Unlock(ctx, "sv_A05EdQM4dvKYQ")
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="server-unlock" method="post" path="/servers/{server_id}/unlock" example="Not Found" -->
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

    res, err := s.Servers.Unlock(ctx, "invalid")
    if err != nil {
        log.Fatal(err)
    }
    if res.Server != nil {
        // handle response
    }
}
```
### Example Usage: Ok

<!-- UsageSnippet language="go" operationID="server-unlock" method="post" path="/servers/{server_id}/unlock" example="Ok" -->
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="server-unlock" method="post" path="/servers/{server_id}/unlock" example="Success" -->
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

    res, err := s.Servers.Unlock(ctx, "sv_e8pKq0xYqWAob")
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

Create out-of-band connection

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-server-out-of-band" method="post" path="/servers/{server_id}/out_of_band_connection" example="Created" -->
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

    res, err := s.Servers.StartOutOfBandConnection(ctx, "sv_8NkvdyGKDeLpx", operations.CreateServerOutOfBandServersRequestBody{
        Data: operations.CreateServerOutOfBandServersData{
            Type: operations.CreateServerOutOfBandServersTypeOutOfBand,
            Attributes: &operations.CreateServerOutOfBandServersAttributes{
                SSHKeyID: latitudeshgosdk.Pointer("ssh_3YjJOLMydvZ87"),
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
### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="create-server-out-of-band" method="post" path="/servers/{server_id}/out_of_band_connection" example="Forbidden" -->
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

    res, err := s.Servers.StartOutOfBandConnection(ctx, "sv_lkg1DeYLDvZE5", operations.CreateServerOutOfBandServersRequestBody{
        Data: operations.CreateServerOutOfBandServersData{
            Type: operations.CreateServerOutOfBandServersTypeOutOfBand,
            Attributes: &operations.CreateServerOutOfBandServersAttributes{
                SSHKeyID: latitudeshgosdk.Pointer("ssh_aKXgRdR3qv9k5"),
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
### Example Usage: Not found

<!-- UsageSnippet language="go" operationID="create-server-out-of-band" method="post" path="/servers/{server_id}/out_of_band_connection" example="Not found" -->
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

    res, err := s.Servers.StartOutOfBandConnection(ctx, "invalid-server", operations.CreateServerOutOfBandServersRequestBody{
        Data: operations.CreateServerOutOfBandServersData{
            Type: operations.CreateServerOutOfBandServersTypeOutOfBand,
            Attributes: &operations.CreateServerOutOfBandServersAttributes{
                SSHKeyID: latitudeshgosdk.Pointer("ssh_m1R3zq2bqWxyn"),
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
### Example Usage: when the server is being provisioned

<!-- UsageSnippet language="go" operationID="create-server-out-of-band" method="post" path="/servers/{server_id}/out_of_band_connection" example="when the server is being provisioned" -->
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

    res, err := s.Servers.StartOutOfBandConnection(ctx, "sv_0MK4O44ROa95w", operations.CreateServerOutOfBandServersRequestBody{
        Data: operations.CreateServerOutOfBandServersData{
            Type: operations.CreateServerOutOfBandServersTypeOutOfBand,
            Attributes: &operations.CreateServerOutOfBandServersAttributes{
                SSHKeyID: latitudeshgosdk.Pointer("ssh_vGMy1DbgON50m"),
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

List out-of-band connections

### Example Usage

<!-- UsageSnippet language="go" operationID="get-server-out-of-band" method="get" path="/servers/{server_id}/out_of_band_connection" example="Success" -->
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

    res, err := s.Servers.GetOutOfBand(ctx, "sv_GnzRD5lvqM5yw")
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


### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-server-action" method="post" path="/servers/{server_id}/actions" example="Created" -->
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

    res, err := s.Servers.RunAction(ctx, "sv_WVQJDMVBORbyE", operations.CreateServerActionServersRequestBody{
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
### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="create-server-action" method="post" path="/servers/{server_id}/actions" example="Forbidden" -->
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

    res, err := s.Servers.RunAction(ctx, "sv_e8pKq0xYqWAob", operations.CreateServerActionServersRequestBody{
        Data: operations.CreateServerActionServersData{
            Type: operations.CreateServerActionServersTypeActions,
            Attributes: &operations.CreateServerActionServersAttributes{
                Action: operations.CreateServerActionActionPowerOff,
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


### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-ipmi-session" method="post" path="/servers/{server_id}/remote_access" example="Created" -->
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

    res, err := s.Servers.CreateIpmiSession(ctx, "sv_Qkm7dXaRq8nZV")
    if err != nil {
        log.Fatal(err)
    }
    if res.IpmiSession != nil {
        // handle response
    }
}
```
### Example Usage: Not found

<!-- UsageSnippet language="go" operationID="create-ipmi-session" method="post" path="/servers/{server_id}/remote_access" example="Not found" -->
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

    res, err := s.Servers.CreateIpmiSession(ctx, "invalid")
    if err != nil {
        log.Fatal(err)
    }
    if res.IpmiSession != nil {
        // handle response
    }
}
```
### Example Usage: Unprocessable entity

<!-- UsageSnippet language="go" operationID="create-ipmi-session" method="post" path="/servers/{server_id}/remote_access" example="Unprocessable entity" -->
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

    res, err := s.Servers.CreateIpmiSession(ctx, "sv_yQrJdNMGO30gv")
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

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="server-start-rescue-mode" method="post" path="/servers/{server_id}/rescue_mode" example="Created" -->
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

    res, err := s.Servers.StartRescueMode(ctx, "sv_WeGoqAWNOP7nz")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerRescue != nil {
        // handle response
    }
}
```
### Example Usage: Not Acceptable

<!-- UsageSnippet language="go" operationID="server-start-rescue-mode" method="post" path="/servers/{server_id}/rescue_mode" example="Not Acceptable" -->
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

    res, err := s.Servers.StartRescueMode(ctx, "sv_GnzRD5lvqM5yw")
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

### Example Usage: Not Acceptable

<!-- UsageSnippet language="go" operationID="server-exit-rescue-mode" method="post" path="/servers/{server_id}/exit_rescue_mode" example="Not Acceptable" -->
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

    res, err := s.Servers.ExitRescueMode(ctx, "sv_WVQJDMVBORbyE")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerRescue != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="server-exit-rescue-mode" method="post" path="/servers/{server_id}/exit_rescue_mode" example="Success" -->
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

    res, err := s.Servers.ExitRescueMode(ctx, "sv_3YjJOLQNdvZ87")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerRescue != nil {
        // handle response
    }
}
```
### Example Usage: if the server is entering rescue mode

<!-- UsageSnippet language="go" operationID="server-exit-rescue-mode" method="post" path="/servers/{server_id}/exit_rescue_mode" example="if the server is entering rescue mode" -->
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

    res, err := s.Servers.ExitRescueMode(ctx, "sv_1R3zq2JxqWxyn")
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

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="server-schedule-deletion" method="post" path="/servers/{server_id}/schedule_deletion" example="Created" -->
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

    res, err := s.Servers.ScheduleDeletion(ctx, "sv_g1mbDwBZqLv5B")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerScheduleDeletion != nil {
        // handle response
    }
}
```
### Example Usage: when request deletion time is invalid

<!-- UsageSnippet language="go" operationID="server-schedule-deletion" method="post" path="/servers/{server_id}/schedule_deletion" example="when request deletion time is invalid" -->
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

    res, err := s.Servers.ScheduleDeletion(ctx, "sv_Qkm7dXaRq8nZV")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServerScheduleDeletion != nil {
        // handle response
    }
}
```
### Example Usage: when server is locked

<!-- UsageSnippet language="go" operationID="server-schedule-deletion" method="post" path="/servers/{server_id}/schedule_deletion" example="when server is locked" -->
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

    res, err := s.Servers.ScheduleDeletion(ctx, "sv_5xyZOnLMqWM0l")
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

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-server-reinstall" method="post" path="/servers/{server_id}/reinstall" example="Created" -->
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

    res, err := s.Servers.Reinstall(ctx, "sv_aNmodj6ydbE8W", operations.CreateServerReinstallServersRequestBody{
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
### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="create-server-reinstall" method="post" path="/servers/{server_id}/reinstall" example="Forbidden" -->
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

    res, err := s.Servers.Reinstall(ctx, "sv_LGXPdWK8dnNWk", operations.CreateServerReinstallServersRequestBody{
        Data: operations.CreateServerReinstallServersData{
            Type: operations.CreateServerReinstallServersTypeReinstalls,
            Attributes: &operations.CreateServerReinstallServersAttributes{
                OperatingSystem: operations.CreateServerReinstallServersOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
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
### Example Usage: Locked server

<!-- UsageSnippet language="go" operationID="create-server-reinstall" method="post" path="/servers/{server_id}/reinstall" example="Locked server" -->
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

    res, err := s.Servers.Reinstall(ctx, "sv_Gr47qlKvdAg0m", operations.CreateServerReinstallServersRequestBody{
        Data: operations.CreateServerReinstallServersData{
            Type: operations.CreateServerReinstallServersTypeReinstalls,
            Attributes: &operations.CreateServerReinstallServersAttributes{
                OperatingSystem: operations.CreateServerReinstallServersOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
                SSHKeys: []string{
                    "37",
                },
                UserData: latitudeshgosdk.Pointer("19"),
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
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="create-server-reinstall" method="post" path="/servers/{server_id}/reinstall" example="Not Found" -->
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

    res, err := s.Servers.Reinstall(ctx, "sv_0L6WO1m1OPlXy", operations.CreateServerReinstallServersRequestBody{
        Data: operations.CreateServerReinstallServersData{
            Type: operations.CreateServerReinstallServersTypeReinstalls,
            Attributes: &operations.CreateServerReinstallServersAttributes{
                OperatingSystem: operations.CreateServerReinstallServersOperatingSystemUbuntu2204X64Lts.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
                SSHKeys: []string{
                    "36",
                },
                UserData: latitudeshgosdk.Pointer("12"),
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
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="create-server-reinstall" method="post" path="/servers/{server_id}/reinstall" example="Unprocessable Entity" -->
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

    res, err := s.Servers.Reinstall(ctx, "sv_RMLydp9XqQKr1", operations.CreateServerReinstallServersRequestBody{
        Data: operations.CreateServerReinstallServersData{
            Type: operations.CreateServerReinstallServersTypeReinstalls,
            Attributes: &operations.CreateServerReinstallServersAttributes{
                OperatingSystem: operations.CreateServerReinstallServersOperatingSystemWindowsServer2019StdV1.ToPointer(),
                Hostname: latitudeshgosdk.Pointer("BRC1"),
                Raid: operations.CreateServerReinstallServersRaidRaid0.ToPointer(),
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