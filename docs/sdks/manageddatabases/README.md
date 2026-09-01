# ManagedDatabases

## Overview

### Available Operations

* [ShowManagedDatabaseMetrics](#showmanageddatabasemetrics) - Show managed database metrics

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