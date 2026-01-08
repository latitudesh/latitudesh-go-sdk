# Traffic

## Overview

### Available Operations

* [Get](#get) - Retrieve traffic
* [GetQuota](#getquota) - Retrieve traffic quota

## Get

Retrieve traffic

### Example Usage

<!-- UsageSnippet language="go" operationID="get-traffic-consumption" method="get" path="/traffic" -->
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

    res, err := s.Traffic.Get(ctx, "2025-06-16T18:45:37Z", "2025-07-16T18:45:37Z", latitudeshgosdk.Pointer("sv_kjQwdEMXdYNVP"), latitudeshgosdk.Pointer("proj_AW6Q2D9lqKLpr"))
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

### Example Usage

<!-- UsageSnippet language="go" operationID="get-traffic-quota" method="get" path="/traffic/quota" -->
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