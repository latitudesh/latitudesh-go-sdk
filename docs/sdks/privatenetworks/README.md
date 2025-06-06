# PrivateNetworks
(*PrivateNetworks*)

## Overview

### Available Operations

* [GetVirtualNetworks](#getvirtualnetworks) - List all Virtual Networks
* [CreateVirtualNetwork](#createvirtualnetwork) - Create a Virtual Network
* [UpdateVirtualNetwork](#updatevirtualnetwork) - Update a Virtual Network
* [GetVirtualNetwork](#getvirtualnetwork) - Retrieve a Virtual Network
* [GetVirtualNetworksAssignments](#getvirtualnetworksassignments) - List all servers assigned to virtual networks
* [AssignServerVirtualNetwork](#assignservervirtualnetwork) - Assign Virtual network
* [DeleteVirtualNetworksAssignments](#deletevirtualnetworksassignments) - Delete Virtual Network Assignment

## GetVirtualNetworks

Lists virtual networks assigned to a project


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

    res, err := s.PrivateNetworks.GetVirtualNetworks(ctx, operations.GetVirtualNetworksRequest{
        FilterLocation: latitudeshgosdk.String("SAO"),
        FilterProject: latitudeshgosdk.String("awesome-copper-clock"),
        FilterTags: latitudeshgosdk.String("tag_KLmjvaEPE7uL9G9E42pxTrEK96Jn"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualNetworks != nil {
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetVirtualNetworksRequest](../../models/operations/getvirtualnetworksrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetVirtualNetworksResponse](../../models/operations/getvirtualnetworksresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateVirtualNetwork

Creates a new Virtual Network.


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

    res, err := s.PrivateNetworks.CreateVirtualNetwork(ctx, operations.CreateVirtualNetworkPrivateNetworksRequestBody{
        Data: operations.CreateVirtualNetworkPrivateNetworksData{
            Type: operations.CreateVirtualNetworkPrivateNetworksTypeVirtualNetwork,
            Attributes: operations.CreateVirtualNetworkPrivateNetworksAttributes{
                Description: "São Paulo VLAN",
                Site: operations.CreateVirtualNetworkPrivateNetworksSiteMia.ToPointer(),
                Project: "enormous-paper-clock",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualNetwork != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.CreateVirtualNetworkPrivateNetworksRequestBody](../../models/operations/createvirtualnetworkprivatenetworksrequestbody.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.CreateVirtualNetworkResponse](../../models/operations/createvirtualnetworkresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 422                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UpdateVirtualNetwork

Update a Virtual Network.


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

    res, err := s.PrivateNetworks.UpdateVirtualNetwork(ctx, "vlan_zGr47qlMDAg0m", operations.UpdateVirtualNetworkPrivateNetworksRequestBody{
        Data: operations.UpdateVirtualNetworkPrivateNetworksData{
            Type: operations.UpdateVirtualNetworkPrivateNetworksTypeVirtualNetworks,
            Attributes: &operations.UpdateVirtualNetworkPrivateNetworksAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualNetwork != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `vlanID`                                                                                                                               | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The Virtual Network ID                                                                                                                 |
| `requestBody`                                                                                                                          | [operations.UpdateVirtualNetworkPrivateNetworksRequestBody](../../models/operations/updatevirtualnetworkprivatenetworksrequestbody.md) | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.UpdateVirtualNetworkResponse](../../models/operations/updatevirtualnetworkresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| components.VirtualNetworkError | 403                            | application/vnd.api+json       |
| components.APIError            | 4XX, 5XX                       | \*/\*                          |

## GetVirtualNetwork

Retrieve a Virtual Network.


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

    res, err := s.PrivateNetworks.GetVirtualNetwork(ctx, "vlan_W6Q2D9ordKLpr")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `vlanID`                                                 | *string*                                                 | :heavy_check_mark:                                       | Virtual Network ID                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVirtualNetworkResponse](../../models/operations/getvirtualnetworkresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetVirtualNetworksAssignments

Returns a list of all servers assigned to virtual networks.


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

    res, err := s.PrivateNetworks.GetVirtualNetworksAssignments(ctx, operations.GetVirtualNetworksAssignmentsRequest{
        FilterServer: latitudeshgosdk.String("217"),
        FilterVid: latitudeshgosdk.String("8"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualNetworkAssignments != nil {
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetVirtualNetworksAssignmentsRequest](../../models/operations/getvirtualnetworksassignmentsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetVirtualNetworksAssignmentsResponse](../../models/operations/getvirtualnetworksassignmentsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## AssignServerVirtualNetwork

Assign Virtual network

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

    res, err := s.PrivateNetworks.AssignServerVirtualNetwork(ctx, operations.AssignServerVirtualNetworkPrivateNetworksRequestBody{
        Data: &operations.AssignServerVirtualNetworkPrivateNetworksData{
            Type: operations.AssignServerVirtualNetworkPrivateNetworksTypeVirtualNetworkAssignment,
            Attributes: &operations.AssignServerVirtualNetworkPrivateNetworksAttributes{
                ServerID: "sv_pbV0DgQGd4AWz",
                VirtualNetworkID: "vlan_059EqYe2qQj8p",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualNetworkAssignment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `request`                                                                                                                                          | [operations.AssignServerVirtualNetworkPrivateNetworksRequestBody](../../models/operations/assignservervirtualnetworkprivatenetworksrequestbody.md) | :heavy_check_mark:                                                                                                                                 | The request object to use for the request.                                                                                                         |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.AssignServerVirtualNetworkResponse](../../models/operations/assignservervirtualnetworkresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 422                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteVirtualNetworksAssignments

Allow you to remove a Virtual Network assignment.


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

    res, err := s.PrivateNetworks.DeleteVirtualNetworksAssignments(ctx, "vnasg_LA73qk8WDaJ2o")
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
| `assignmentID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVirtualNetworksAssignmentsResponse](../../models/operations/deletevirtualnetworksassignmentsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 423                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |