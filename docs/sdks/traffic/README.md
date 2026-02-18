# Traffic

## Overview

### Available Operations

* [Get](#get) - Retrieve traffic
* [GetQuota](#getquota) - Retrieve traffic quota

## Get

Retrieve traffic

### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-traffic-consumption" method="get" path="/traffic" example="Success" -->
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

    res, err := s.Traffic.Get(ctx, "2025-12-14T15:57:10Z", "2026-01-14T15:57:10Z", latitudeshgosdk.Pointer("sv_A05EdQ50dvKYQ"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Traffic != nil {
        // handle response
    }
}
```
### Example Usage: when the period filter is the deprecated format

<!-- UsageSnippet language="go" operationID="get-traffic-consumption" method="get" path="/traffic" example="when the period filter is the deprecated format" -->
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

    res, err := s.Traffic.Get(ctx, "2025-04-28T10:40:31Z", "2025-04-28T18:40:31Z", nil, latitudeshgosdk.Pointer("308"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Traffic != nil {
        // handle response
    }
}
```
### Example Usage: when the period is less than 1 day

<!-- UsageSnippet language="go" operationID="get-traffic-consumption" method="get" path="/traffic" example="when the period is less than 1 day" -->
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

    res, err := s.Traffic.Get(ctx, "2025-04-28T10:40:31Z", "2025-04-28T18:40:31Z", latitudeshgosdk.Pointer("198"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Traffic != nil {
        // handle response
    }
}
```
### Example Usage: with invalid filter

<!-- UsageSnippet language="go" operationID="get-traffic-consumption" method="get" path="/traffic" example="with invalid filter" -->
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

    res, err := s.Traffic.Get(ctx, "invalid", "invalid", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Traffic != nil {
        // handle response
    }
}
```
### Example Usage: with project filter

<!-- UsageSnippet language="go" operationID="get-traffic-consumption" method="get" path="/traffic" example="with project filter" -->
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

    res, err := s.Traffic.Get(ctx, "2025-04-06T21:00:00Z", "2025-05-06T21:00:00Z", nil, latitudeshgosdk.Pointer("proj_AW6Q2D9lqKLpr"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Traffic != nil {
        // handle response
    }
}
```
### Example Usage: with server filter

<!-- UsageSnippet language="go" operationID="get-traffic-consumption" method="get" path="/traffic" example="with server filter" -->
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

    res, err := s.Traffic.Get(ctx, "2025-04-06T21:00:00Z", "2025-05-06T21:00:00Z", latitudeshgosdk.Pointer("sv_mw49QDB5qagKb"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Traffic != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `filterDateGte`                                                                                                                  | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The start timestamp to retrieve the traffic. You must provide in ISO8601 format. Example: filter[date][gte]=2024-04-01T00:00:00Z |
| `filterDateLte`                                                                                                                  | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The end timestamp to retrieve the traffic. You must provide in ISO8601 format. Example: filter[date][gte]=2024-04-31T23:59:59Z   |
| `filterServer`                                                                                                                   | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | The server id to filter by                                                                                                       |
| `filterProject`                                                                                                                  | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | The project id to filter by                                                                                                      |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.GetTrafficConsumptionResponse](../../models/operations/gettrafficconsumptionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetQuota

Retrieve traffic quota

### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-traffic-quota" method="get" path="/traffic/quota" example="Success" -->
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

    res, err := s.Traffic.GetQuota(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TrafficQuota != nil {
        // handle response
    }
}
```
### Example Usage: when `billing_method` is `95th percentile`

<!-- UsageSnippet language="go" operationID="get-traffic-quota" method="get" path="/traffic/quota" example="when `billing_method` is `95th percentile`" -->
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

    res, err := s.Traffic.GetQuota(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TrafficQuota != nil {
        // handle response
    }
}
```
### Example Usage: when `billing_method` is `Normal`

<!-- UsageSnippet language="go" operationID="get-traffic-quota" method="get" path="/traffic/quota" example="when `billing_method` is `Normal`" -->
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

    res, err := s.Traffic.GetQuota(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TrafficQuota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `filterProject`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTrafficQuotaResponse](../../models/operations/gettrafficquotaresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |