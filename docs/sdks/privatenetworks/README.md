# PrivateNetworks
(*PrivateNetworks*)

## Overview

### Available Operations

* [List](#list) - List all Virtual Networks
* [Create](#create) - Create a Virtual Network
* [Update](#update) - Update a Virtual Network
* [Get](#get) - Retrieve a Virtual Network
* [ListAssignments](#listassignments) - List all servers assigned to virtual networks
* [Assign](#assign) - Assign Virtual network
* [RemoveAssignment](#removeassignment) - Delete Virtual Network Assignment

## List

Lists virtual networks assigned to a project


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

    res, err := s.PrivateNetworks.List(ctx, latitudeshgosdk.String("SAO"), latitudeshgosdk.String("lightweight-iron-keyboard"), latitudeshgosdk.String("tag_jVnnao8eYWSQW5EKKR4QH5bQllx"))
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualNetworks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |
| `filterLocation`                                                                                                            | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | The location slug to filter by                                                                                              |
| `filterProject`                                                                                                             | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | The project id or slug to filter by                                                                                         |
| `filterTags`                                                                                                                | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | The tags ids to filter by, separated by comma, e.g. `filter[tags]=tag_1,tag_2`will return ssh keys with `tag_1` AND `tag_2` |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |

### Response

**[*operations.GetVirtualNetworksResponse](../../models/operations/getvirtualnetworksresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

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

    res, err := s.PrivateNetworks.Create(ctx, operations.CreateVirtualNetworkPrivateNetworksRequestBody{
        Data: operations.CreateVirtualNetworkPrivateNetworksData{
            Type: operations.CreateVirtualNetworkPrivateNetworksTypeVirtualNetwork,
            Attributes: operations.CreateVirtualNetworkPrivateNetworksAttributes{
                Description: "São Paulo VLAN",
                Site: operations.CreateVirtualNetworkPrivateNetworksSiteMia.ToPointer(),
                Project: "aerodynamic-marble-bench",
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

## Update

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

    res, err := s.PrivateNetworks.Update(ctx, "vlan_pRMLydp0dQKr1", operations.UpdateVirtualNetworkPrivateNetworksRequestBody{
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

## Get

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

    res, err := s.PrivateNetworks.Get(ctx, "vlan_W6Q2D9ordKLpr")
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

## ListAssignments

Returns a list of all servers assigned to virtual networks.


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

    res, err := s.PrivateNetworks.ListAssignments(ctx, latitudeshgosdk.String("208"), latitudeshgosdk.String("4"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualNetworkAssignments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `filterServer`                                           | **string*                                                | :heavy_minus_sign:                                       | The server ID to filter by                               |
| `filterVid`                                              | **string*                                                | :heavy_minus_sign:                                       | The vlan ID to filter by                                 |
| `filterVirtualNetworkID`                                 | **string*                                                | :heavy_minus_sign:                                       | The virtual network ID to filter by                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVirtualNetworksAssignmentsResponse](../../models/operations/getvirtualnetworksassignmentsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Assign

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

    res, err := s.PrivateNetworks.Assign(ctx, operations.AssignServerVirtualNetworkPrivateNetworksRequestBody{
        Data: &operations.AssignServerVirtualNetworkPrivateNetworksData{
            Type: operations.AssignServerVirtualNetworkPrivateNetworksTypeVirtualNetworkAssignment,
            Attributes: &operations.AssignServerVirtualNetworkPrivateNetworksAttributes{
                ServerID: "sv_aNmodj6ZdbE8W",
                VirtualNetworkID: vlan_g1mbDwG0DLv5B,
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

## RemoveAssignment

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

    res, err := s.PrivateNetworks.RemoveAssignment(ctx, "vnasg_695BdKagDevVo")
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