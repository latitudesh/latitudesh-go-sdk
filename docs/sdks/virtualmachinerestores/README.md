# VirtualMachineRestores

## Overview

### Available Operations

* [ListForVirtualMachineBackup](#listforvirtualmachinebackup) - List a backup's restores
* [CreateForVirtualMachineBackup](#createforvirtualmachinebackup) - Create VM restore
* [ListForBackup](#listforbackup) - List a backup's restores (top-level backup path)
* [List](#list) - List all VM restores
* [Create](#create) - Create VM restore (flat)
* [Get](#get) - Get VM restore

## ListForVirtualMachineBackup

Lists the restores created from the given backup.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-virtual-machine-scoped-restores" method="get" path="/virtual_machines/{virtual_machine_id}/backups/{backup_id}/restores" -->
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

    res, err := s.VirtualMachineRestores.ListForVirtualMachineBackup(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineRestores != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `virtualMachineID`                                       | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `backupID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListVirtualMachineScopedRestoresResponse](../../models/operations/listvirtualmachinescopedrestoresresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateForVirtualMachineBackup

Restores a backup into a new Virtual Machine. Optionally accepts a `name` for the restored VM and a target `site` slug to restore into another region.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-virtual-machine-restore" method="post" path="/virtual_machines/{virtual_machine_id}/backups/{backup_id}/restores" -->
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

    res, err := s.VirtualMachineRestores.CreateForVirtualMachineBackup(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineRestore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |
| `virtualMachineID`                                                                                  | `string`                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 |
| `backupID`                                                                                          | `string`                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 |
| `virtualMachineRestorePayload`                                                                      | [*components.VirtualMachineRestorePayload](../../models/components/virtualmachinerestorepayload.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |

### Response

**[*operations.CreateVirtualMachineRestoreResponse](../../models/operations/createvirtualmachinerestoreresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListForBackup

Lists the restores created from the given backup, reached via the top-level backup path.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-backup-restores" method="get" path="/virtual_machine_backups/{backup_id}/restores" -->
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

    res, err := s.VirtualMachineRestores.ListForBackup(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineRestores != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `backupID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListBackupRestoresResponse](../../models/operations/listbackuprestoresresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## List

Lists every restore that belongs to the authenticated team.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-virtual-machine-restores" method="get" path="/virtual_machine_restores" -->
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

    res, err := s.VirtualMachineRestores.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineRestores != nil {
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

**[*operations.ListVirtualMachineRestoresResponse](../../models/operations/listvirtualmachinerestoresresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Create

Restores the backup referenced in the body into a new Virtual Machine. Optionally accepts a `name` for the restored VM and a target `site` slug to restore into another region.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-virtual-machine-restore-flat" method="post" path="/virtual_machine_restores" example="FeatureNotEnabled" -->
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

    res, err := s.VirtualMachineRestores.Create(ctx, components.VirtualMachineRestorePayload{})
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineRestore != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.VirtualMachineRestorePayload](../../models/components/virtualmachinerestorepayload.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.CreateVirtualMachineRestoreFlatResponse](../../models/operations/createvirtualmachinerestoreflatresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Get

Get VM restore

### Example Usage

<!-- UsageSnippet language="go" operationID="get-virtual-machine-restore" method="get" path="/virtual_machine_restores/{id}" -->
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

    res, err := s.VirtualMachineRestores.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachineRestore != nil {
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

**[*operations.GetVirtualMachineRestoreResponse](../../models/operations/getvirtualmachinerestoreresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |