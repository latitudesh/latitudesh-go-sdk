# ManagedDatabases

## Overview

Managed database instances (PostgreSQL, ClickHouse)

### Available Operations

* [ShowManagedDatabaseMetrics](#showmanageddatabasemetrics) - Show managed database metrics
* [ListManagedDatabases](#listmanageddatabases) - List managed databases
* [CreateManagedDatabase](#createmanageddatabase) - Create a managed database
* [ShowManagedDatabase](#showmanageddatabase) - Show a managed database
* [DestroyManagedDatabase](#destroymanageddatabase) - Destroy a managed database
* [UpdateManagedDatabase](#updatemanageddatabase) - Update a managed database
* [ListManagedDatabaseBackups](#listmanageddatabasebackups) - List managed database backups

## ShowManagedDatabaseMetrics

Show managed database metrics

### Example Usage

<!-- UsageSnippet language="go" operationID="show-managed-database-metrics" method="get" path="/managed_databases/{managed_database_id}/metrics" -->
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

    res, err := s.ManagedDatabases.ShowManagedDatabaseMetrics(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `managedDatabaseID`                                                                                                                          | `string`                                                                                                                                     | :heavy_check_mark:                                                                                                                           | Managed database ID                                                                                                                          |
| `period`                                                                                                                                     | `*int64`                                                                                                                                     | :heavy_minus_sign:                                                                                                                           | Time window in seconds. One of 1800, 3600, 21600, 86400, 604800 (default 1800).                                                              |
| `queries`                                                                                                                                    | `*string`                                                                                                                                    | :heavy_minus_sign:                                                                                                                           | Comma-separated metrics to fetch. Defaults to all: cpuUsage, memoryUsage, tpsUsage, maxConnections, deadlocks, blockedQueries, databaseSize. |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.ShowManagedDatabaseMetricsResponse](../../models/operations/showmanageddatabasemetricsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ListManagedDatabases

List managed databases

### Example Usage

<!-- UsageSnippet language="go" operationID="list-managed-databases" method="get" path="/managed_databases" -->
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

    res, err := s.ManagedDatabases.ListManagedDatabases(ctx, "<id>", latitudeshgosdk.Pointer("postgres"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ManagedDatabases != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | The project slug to filter databases by                  |                                                          |
| `engine`                                                 | `*string`                                                | :heavy_minus_sign:                                       | Filter by database engine                                | postgres                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListManagedDatabasesResponse](../../models/operations/listmanageddatabasesresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateManagedDatabase

Create a managed database

### Example Usage

<!-- UsageSnippet language="go" operationID="create-managed-database" method="post" path="/managed_databases" -->
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

    res, err := s.ManagedDatabases.CreateManagedDatabase(ctx, components.ManagedDatabasePayload{
        Data: &components.ManagedDatabasePayloadData{
            Type: components.ManagedDatabasePayloadTypeManagedDatabases.ToPointer(),
            Attributes: &components.ManagedDatabasePayloadAttributes{
                Name: latitudeshgosdk.Pointer("my-postgres-db"),
                ProjectID: "proj_ABC123",
                Region: "ASH",
                Plan: "db.psql.small",
                Engine: "postgres",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ManagedDatabase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.ManagedDatabasePayload](../../models/components/manageddatabasepayload.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateManagedDatabaseResponse](../../models/operations/createmanageddatabaseresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ShowManagedDatabase

Show a managed database

### Example Usage

<!-- UsageSnippet language="go" operationID="show-managed-database" method="get" path="/managed_databases/{id}" -->
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

    res, err := s.ManagedDatabases.ShowManagedDatabase(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ManagedDatabase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Managed database ID                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ShowManagedDatabaseResponse](../../models/operations/showmanageddatabaseresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DestroyManagedDatabase

Destroy a managed database

### Example Usage

<!-- UsageSnippet language="go" operationID="destroy-managed-database" method="delete" path="/managed_databases/{id}" -->
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

    res, err := s.ManagedDatabases.DestroyManagedDatabase(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | Managed database ID                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DestroyManagedDatabaseResponse](../../models/operations/destroymanageddatabaseresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateManagedDatabase

Update a managed database

### Example Usage

<!-- UsageSnippet language="go" operationID="update-managed-database" method="patch" path="/managed_databases/{id}" -->
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

    res, err := s.ManagedDatabases.UpdateManagedDatabase(ctx, "<id>", components.ManagedDatabaseUpdatePayload{
        Data: &components.ManagedDatabaseUpdatePayloadData{
            Attributes: &components.ManagedDatabaseUpdatePayloadAttributes{
                TrustedSources: []string{
                    "203.0.113.0/24",
                    "198.51.100.0/24",
                },
                Parameters: map[string]any{
                    "shared_buffers": "256MB",
                    "work_mem": "16MB",
                    "effective_cache_size": "1GB",
                },
                Pooler: map[string]any{
                    "enabled": true,
                    "default_pool_size": 30,
                    "max_client_conn": 200,
                },
                Backup: map[string]any{
                    "enabled": true,
                    "schedule": "0 0 0 * * *",
                    "s3Endpoint": "https://s3.amazonaws.com",
                    "bucketName": "my-db-backups",
                    "path": "prod/postgres",
                    "retentionPolicy": "7",
                },
                AccessCredentials: &components.AccessCredentials{
                    AccessKeyID: latitudeshgosdk.Pointer("AKIAIOSFODNN7EXAMPLE"),
                    SecretAccessKey: latitudeshgosdk.Pointer("wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"),
                },
                Supabase: latitudeshgosdk.Pointer(true),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ManagedDatabase != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | `string`                                                                                           | :heavy_check_mark:                                                                                 | Managed database ID                                                                                |
| `managedDatabaseUpdatePayload`                                                                     | [components.ManagedDatabaseUpdatePayload](../../models/components/manageddatabaseupdatepayload.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateManagedDatabaseResponse](../../models/operations/updatemanageddatabaseresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ListManagedDatabaseBackups

List managed database backups

### Example Usage

<!-- UsageSnippet language="go" operationID="list-managed-database-backups" method="get" path="/managed_databases/{managed_database_id}/backups" -->
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

    res, err := s.ManagedDatabases.ListManagedDatabaseBackups(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `managedDatabaseID`                                                       | `string`                                                                  | :heavy_check_mark:                                                        | Managed database ID                                                       |
| `phase`                                                                   | `*string`                                                                 | :heavy_minus_sign:                                                        | Filter backups by phase. Use 'completed' to return only finished backups. |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |

### Response

**[*operations.ListManagedDatabaseBackupsResponse](../../models/operations/listmanageddatabasebackupsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |