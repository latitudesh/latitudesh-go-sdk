# VirtualMachines

## Overview

### Available Operations

* [Create](#create) - Create VM
* [List](#list) - List VMs
* [Get](#get) - Retrieve VM
* [Delete](#delete) - Destroy VM
* [UpdateVirtualMachine](#updatevirtualmachine) - Update VM
* [CreateVirtualMachineAction](#createvirtualmachineaction) - Run VM power action

## Create

Creates a new Virtual Machine.


### Example Usage: Conflict

<!-- UsageSnippet language="go" operationID="create-virtual-machine" method="post" path="/virtual_machines" example="Conflict" -->
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

    res, err := s.VirtualMachines.Create(ctx, components.VirtualMachinePayload{
        Data: &components.VirtualMachinePayloadData{
            Type: components.VirtualMachinePayloadTypeVirtualMachines.ToPointer(),
            Attributes: &components.VirtualMachinePayloadAttributes{
                Name: latitudeshgosdk.Pointer("my-new-vm"),
                Project: latitudeshgosdk.Pointer("intelligent-paper-table"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachine != nil {
        // handle response
    }
}
```
### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-virtual-machine" method="post" path="/virtual_machines" example="Created" -->
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

    res, err := s.VirtualMachines.Create(ctx, components.VirtualMachinePayload{
        Data: &components.VirtualMachinePayloadData{
            Type: components.VirtualMachinePayloadTypeVirtualMachines.ToPointer(),
            Attributes: &components.VirtualMachinePayloadAttributes{
                Name: latitudeshgosdk.Pointer("my-new-vm"),
                Project: latitudeshgosdk.Pointer("lightweight-leather-lamp"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachine != nil {
        // handle response
    }
}
```
### Example Usage: CreatedWithOS

<!-- UsageSnippet language="go" operationID="create-virtual-machine" method="post" path="/virtual_machines" example="CreatedWithOS" -->
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

    res, err := s.VirtualMachines.Create(ctx, components.VirtualMachinePayload{
        Data: &components.VirtualMachinePayloadData{
            Type: components.VirtualMachinePayloadTypeVirtualMachines.ToPointer(),
            Attributes: &components.VirtualMachinePayloadAttributes{
                Name: latitudeshgosdk.Pointer("my-new-vm"),
                Project: latitudeshgosdk.Pointer("lightweight-leather-lamp"),
                OperatingSystem: latitudeshgosdk.Pointer("ubuntu_24_04_x64_lts"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachine != nil {
        // handle response
    }
}
```
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="create-virtual-machine" method="post" path="/virtual_machines" example="Unprocessable Entity" -->
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

    res, err := s.VirtualMachines.Create(ctx, components.VirtualMachinePayload{
        Data: &components.VirtualMachinePayloadData{
            Type: components.VirtualMachinePayloadTypeVirtualMachines.ToPointer(),
            Attributes: &components.VirtualMachinePayloadAttributes{
                Name: latitudeshgosdk.Pointer("My new-vm"),
                Project: latitudeshgosdk.Pointer("mediocre-wool-knife"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachine != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.VirtualMachinePayload](../../models/components/virtualmachinepayload.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateVirtualMachineResponse](../../models/operations/createvirtualmachineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## List

Show all Team's Virtual Machines.


### Example Usage

<!-- UsageSnippet language="go" operationID="index-virtual-machine" method="get" path="/virtual_machines" -->
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

    res, err := s.VirtualMachines.List(ctx, nil, latitudeshgosdk.Pointer("credentials"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                          | Type                                                                                                                                                               | Required                                                                                                                                                           | Description                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                              | :heavy_check_mark:                                                                                                                                                 | The context to use for the request.                                                                                                                                |
| `filterProject`                                                                                                                                                    | **string*                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                 | The project ID or Slug to filter by                                                                                                                                |
| `extraFieldsVirtualMachines`                                                                                                                                       | **string*                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                 | The `credentials` are provided as extra attributes that are lazy loaded. To request it, just set `extra_fields[virtual_machines]=credentials` in the query string. |
| `opts`                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                 | The options for this request.                                                                                                                                      |

### Response

**[*operations.IndexVirtualMachineResponse](../../models/operations/indexvirtualmachineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Get

Show a Virtual Machine.


### Example Usage

<!-- UsageSnippet language="go" operationID="show-virtual-machine" method="get" path="/virtual_machines/{virtual_machine_id}" -->
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

    res, err := s.VirtualMachines.Get(ctx, "vm_7vYAZqGBdMQ94")
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
| `virtualMachineID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ShowVirtualMachineResponse](../../models/operations/showvirtualmachineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Destroys a Virtual Machine.


### Example Usage

<!-- UsageSnippet language="go" operationID="destroy-virtual-machine" method="delete" path="/virtual_machines/{virtual_machine_id}" -->
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

    res, err := s.VirtualMachines.Delete(ctx, "<id>")
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
| `virtualMachineID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DestroyVirtualMachineResponse](../../models/operations/destroyvirtualmachineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateVirtualMachine

Updates a Virtual Machine's display name (hostname).


### Example Usage

<!-- UsageSnippet language="go" operationID="update-virtual-machine" method="patch" path="/virtual_machines/{virtual_machine_id}" example="Success" -->
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

    res, err := s.VirtualMachines.UpdateVirtualMachine(ctx, "vm_7vYAZqGBdMQ94", components.VirtualMachineUpdatePayload{
        Data: components.VirtualMachineUpdatePayloadData{
            Type: components.VirtualMachineUpdatePayloadTypeVirtualMachines,
            ID: latitudeshgosdk.Pointer("vm_7vYAZqGBdMQ94"),
            Attributes: components.VirtualMachineUpdatePayloadAttributes{
                Name: "my-updated-vm",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachine != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `virtualMachineID`                                                                               | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `virtualMachineUpdatePayload`                                                                    | [components.VirtualMachineUpdatePayload](../../models/components/virtualmachineupdatepayload.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.UpdateVirtualMachineResponse](../../models/operations/updatevirtualmachineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateVirtualMachineAction

Performs a power action on a given virtual machine:
- `power_on` - Starts the virtual machine
- `power_off` - Stops the virtual machine
- `reboot` - Restarts the virtual machine


### Example Usage

<!-- UsageSnippet language="go" operationID="create-virtual-machine-action" method="post" path="/virtual_machines/{virtual_machine_id}/actions" example="Created" -->
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

    res, err := s.VirtualMachines.CreateVirtualMachineAction(ctx, "vm_5LA73qkjdaJ2o", operations.CreateVirtualMachineActionVirtualMachinesRequestBody{
        ID: "vm_5LA73qkjdaJ2o",
        Type: operations.CreateVirtualMachineActionVirtualMachinesTypeVirtualMachines,
        Attributes: operations.CreateVirtualMachineActionVirtualMachinesAttributes{
            Action: operations.CreateVirtualMachineActionVirtualMachinesActionReboot,
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `virtualMachineID`                                                                                                                                 | *string*                                                                                                                                           | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `requestBody`                                                                                                                                      | [operations.CreateVirtualMachineActionVirtualMachinesRequestBody](../../models/operations/createvirtualmachineactionvirtualmachinesrequestbody.md) | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.CreateVirtualMachineActionResponse](../../models/operations/createvirtualmachineactionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |