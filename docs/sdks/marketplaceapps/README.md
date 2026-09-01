# MarketplaceApps

## Overview

### Available Operations

* [ListMarketplaceApps](#listmarketplaceapps) - List marketplace apps
* [GetMarketplaceApp](#getmarketplaceapp) - Retrieve a marketplace app

## ListMarketplaceApps

Lists the marketplace apps available to deploy on Virtual Machines. Only published apps are returned.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-marketplace-apps" method="get" path="/marketplace_apps" example="Success" -->
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

    res, err := s.MarketplaceApps.ListMarketplaceApps(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketplaceApps != nil {
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

**[*operations.ListMarketplaceAppsResponse](../../models/operations/listmarketplaceappsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetMarketplaceApp

Retrieve a single marketplace app by ID or slug. Only published apps are visible.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-marketplace-app" method="get" path="/marketplace_apps/{id}" example="Success" -->
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

    res, err := s.MarketplaceApps.GetMarketplaceApp(ctx, "openclaw")
    if err != nil {
        log.Fatal(err)
    }
    if res.MarketplaceApp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The marketplace app ID or slug                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMarketplaceAppResponse](../../models/operations/getmarketplaceappresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |