# IPAddresses
(*IPAddresses*)

## Overview

### Available Operations

* [List](#list) - List IPs
* [Get](#get) - Retrieve an IP

## List

List all Management and Additional IP Addresses.
 • Management IPs are IPs that are used for the management IP of a device.
   This is a public IP address that a device is born and dies with. It never changes during the lifecycle of the device.
 • Additional IPs are individual IPs that can be added to a device as an additional IP that can be used.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ips" method="get" path="/ips" -->
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

    res, err := s.IPAddresses.List(ctx, operations.GetIpsRequest{
        FilterServer: latitudeshgosdk.Pointer("46"),
        FilterProject: latitudeshgosdk.Pointer("64"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IPAddresses != nil {
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [operations.GetIpsRequest](../../models/operations/getipsrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.GetIpsResponse](../../models/operations/getipsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Get

Retrieve an IP Address

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ip" method="get" path="/ips/{ip_id}" -->
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

    res, err := s.IPAddresses.Get(ctx, "ip_059EqY7kOQj8p", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IPAddress != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `ipID`                                                                                                                                                                   | *string*                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                       | The IP Address ID                                                                                                                                                        |
| `extraFieldsIPAddresses`                                                                                                                                                 | **string*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                       | The `region` and `server` are provided as extra attributes that are lazy loaded. To request it, just set `extra_fields[ip_addresses]=region,server` in the query string. |
| `opts`                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.GetIPResponse](../../models/operations/getipresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |