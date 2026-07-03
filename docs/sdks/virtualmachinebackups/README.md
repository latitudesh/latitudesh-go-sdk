# VirtualMachineBackups

## Overview

### Available Operations

* [ListVirtualMachineScopedBackups](#listvirtualmachinescopedbackups) - List a VM's backups
* [CreateVirtualMachineBackup](#createvirtualmachinebackup) - Create VM backup
* [ListVirtualMachineBackups](#listvirtualmachinebackups) - List all VM backups
* [CreateVirtualMachineBackupTopLevel](#createvirtualmachinebackuptoplevel) - Create VM backup (top-level)
* [GetVirtualMachineBackup](#getvirtualmachinebackup) - Get VM backup
* [DeleteVirtualMachineBackup](#deletevirtualmachinebackup) - Delete VM backup

## ListVirtualMachineScopedBackups

Lists the backups of the given Virtual Machine.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-virtual-machine-scoped-backups" method="get" path="/virtual_machines/{virtual_machine_id}/backups" -->
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

    res, err := s.VirtualMachineBackups.ListVirtualMachineScopedBackups(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineBackups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `virtualMachineID`                                       | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListVirtualMachineScopedBackupsResponse](../../models/operations/listvirtualmachinescopedbackupsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateVirtualMachineBackup

Triggers a backup of the given Virtual Machine.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-virtual-machine-backup" method="post" path="/virtual_machines/{virtual_machine_id}/backups" -->
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

    res, err := s.VirtualMachineBackups.CreateVirtualMachineBackup(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineBackup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `virtualMachineID`                                       | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateVirtualMachineBackupResponse](../../models/operations/createvirtualmachinebackupresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListVirtualMachineBackups

Lists every backup that belongs to the authenticated team.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-virtual-machine-backups" method="get" path="/virtual_machine_backups" -->
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

    res, err := s.VirtualMachineBackups.ListVirtualMachineBackups(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineBackups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListVirtualMachineBackupsResponse](../../models/operations/listvirtualmachinebackupsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateVirtualMachineBackupTopLevel

Triggers a backup of the Virtual Machine referenced in the body.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-virtual-machine-backup-top-level" method="post" path="/virtual_machine_backups" example="FeatureNotEnabled" -->
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

    res, err := s.VirtualMachineBackups.CreateVirtualMachineBackupTopLevel(ctx, components.VirtualMachineBackupPayload{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineBackup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.VirtualMachineBackupPayload](../../models/components/virtualmachinebackuppayload.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateVirtualMachineBackupTopLevelResponse](../../models/operations/createvirtualmachinebackuptoplevelresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetVirtualMachineBackup

Get VM backup

### Example Usage

<!-- UsageSnippet language="go" operationID="get-virtual-machine-backup" method="get" path="/virtual_machine_backups/{id}" -->
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

    res, err := s.VirtualMachineBackups.GetVirtualMachineBackup(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineBackup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVirtualMachineBackupResponse](../../models/operations/getvirtualmachinebackupresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteVirtualMachineBackup

Archives and deletes a Virtual Machine backup. Work runs asynchronously and returns 202 Accepted. Only `Ready` or `Failed` backups can be deleted, and not while a restore from the backup is in progress.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-virtual-machine-backup" method="delete" path="/virtual_machine_backups/{id}" -->
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

    res, err := s.VirtualMachineBackups.DeleteVirtualMachineBackup(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineBackup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVirtualMachineBackupResponse](../../models/operations/deletevirtualmachinebackupresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |