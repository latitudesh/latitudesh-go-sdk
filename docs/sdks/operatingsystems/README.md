# OperatingSystems
(*OperatingSystems*)

## Overview

### Available Operations

* [ListPlans](#listplans) - List all operating systems available

## ListPlans

Lists all operating systems available to deploy and reinstall.


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

    res, err := s.OperatingSystems.ListPlans(ctx, latitudeshgosdk.Int64(20), latitudeshgosdk.Int64(1))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Number of items to return per page                       |
| `pageNumber`                                             | **int64*                                                 | :heavy_minus_sign:                                       | Page number to return (starts at 1)                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPlansOperatingSystemResponse](../../models/operations/getplansoperatingsystemresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |