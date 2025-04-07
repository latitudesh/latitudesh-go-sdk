# Traffic
(*Traffic*)

## Overview

### Available Operations

* [Get](#get) - Retrieve Traffic consumption
* [GetQuota](#getquota) - Retrieve Traffic Quota

## Get

Retrieve Traffic consumption

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

    res, err := s.Traffic.Get(ctx, "2025-01-05T21:14:55Z", "2025-02-05T21:14:55Z", latitudeshgosdk.Int64(187), latitudeshgosdk.Int64(301))
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
| `filterServer`                                                                                                                   | **int64*                                                                                                                         | :heavy_minus_sign:                                                                                                               | N/A                                                                                                                              |
| `filterProject`                                                                                                                  | **int64*                                                                                                                         | :heavy_minus_sign:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.GetTrafficConsumptionResponse](../../models/operations/gettrafficconsumptionresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetQuota

Retrieve Traffic Quota

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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |