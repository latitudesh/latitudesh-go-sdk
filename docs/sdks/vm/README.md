# Plans.Vm

## Overview

### Available Operations

* [List](#list) - List all Virtual Machines Plans

## List

List all Virtual Machines Plans

### Example Usage

<!-- UsageSnippet language="go" operationID="get-vm-plans" method="get" path="/plans/virtual_machines" -->
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

    res, err := s.Plans.VM.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMachinePlans != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `filterGpu`                                              | **bool*                                                  | :heavy_minus_sign:                                       | Filter plans by GPU availability                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVMPlansResponse](../../models/operations/getvmplansresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |