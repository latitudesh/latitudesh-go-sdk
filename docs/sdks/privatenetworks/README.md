# PrivateNetworks

## Overview

### Available Operations

* [List](#list) - List all Virtual Networks
* [Create](#create) - Create a Virtual Network
* [Update](#update) - Update a Virtual Network
* [Get](#get) - Retrieve a Virtual Network
* [ListAssignments](#listassignments) - List all servers assigned to virtual networks
* [Assign](#assign) - Assign Virtual network
* [DeleteAssignment](#deleteassignment) - Delete Virtual Network Assignment

## List

Lists virtual networks assigned to a project


### Example Usage

<!-- UsageSnippet language="go" operationID="get-virtual-networks" method="get" path="/virtual_networks" -->
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

    res, err := s.PrivateNetworks.List(ctx, operations.GetVirtualNetworksRequest{
        FilterLocation: latitudeshgosdk.Pointer("SAO"),
        FilterProject: latitudeshgosdk.Pointer("awesome-copper-clock"),
        FilterTags: latitudeshgosdk.Pointer("tag_BZWAJKePr2Fx9kRyyaARImQlXmW,tag_X8yMgb8AZPFrX72lQgrwhBVnPN2"),
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

## Create

Creates a new Virtual Network.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-virtual-network" method="post" path="/virtual_networks" -->
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
                Project: "ergonomic-steel-bag",
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Update a Virtual Network.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-virtual-network" method="patch" path="/virtual_networks/{vlan_id}" -->
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

    res, err := s.PrivateNetworks.Update(ctx, "vlan_VaNmodjeObE8W", operations.UpdateVirtualNetworkPrivateNetworksRequestBody{
        Data: operations.UpdateVirtualNetworkPrivateNetworksData{
            ID: "vlan_VaNmodjeObE8W",
            Type: operations.UpdateVirtualNetworkPrivateNetworksTypeVirtualNetworks,
            Attributes: &operations.UpdateVirtualNetworkPrivateNetworksAttributes{
                Tags: []string{
                    "tag_RjLvG6oe84IAw7BxxEGaFAXK4l4",
                    "tag_lpPQ21kXEYfb9az3jRoVIVw4RBk",
                },
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
| `vlanID`                                                                                                                               | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The Virtual Network ID                                                                                                                 |
| `requestBody`                                                                                                                          | [operations.UpdateVirtualNetworkPrivateNetworksRequestBody](../../models/operations/updatevirtualnetworkprivatenetworksrequestbody.md) | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.UpdateVirtualNetworkResponse](../../models/operations/updatevirtualnetworkresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Get

Retrieve a Virtual Network.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-virtual-network" method="get" path="/virtual_networks/{vlan_id}" -->
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

<!-- UsageSnippet language="go" operationID="get-virtual-networks-assignments" method="get" path="/virtual_networks/assignments" -->
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

    res, err := s.PrivateNetworks.ListAssignments(ctx, operations.GetVirtualNetworksAssignmentsRequest{
        FilterServer: latitudeshgosdk.Pointer("217"),
        FilterVid: latitudeshgosdk.Pointer("8"),
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

## Assign

Assign Virtual network

### Example Usage

<!-- UsageSnippet language="go" operationID="assign-server-virtual-network" method="post" path="/virtual_networks/assignments" -->
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
                ServerID: "sv_5xyZOn5vDWM0l",
                VirtualNetworkID: "vlan_Z8rodmpGO1jLB",
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteAssignment

Allow you to remove a Virtual Network assignment.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-virtual-networks-assignments" method="delete" path="/virtual_networks/assignments/{assignment_id}" -->
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

    res, err := s.PrivateNetworks.DeleteAssignment(ctx, "vnasg_LA73qk8WDaJ2o")
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |