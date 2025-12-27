# Billing

## Overview

### Available Operations

* [ListUsage](#listusage) - Retrieve billing usage

## ListUsage

Returns the billing usage of a project


### Example Usage

<!-- UsageSnippet language="go" operationID="get-billing-usage" method="get" path="/billing/usage" -->
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

    res, err := s.Billing.ListUsage(ctx, "proj_6059EqYkOQj8p", []string{
        "si_lvdub7r3",
        "si_utp5nfrf",
    }, latitudeshgosdk.Pointer("plan.name"))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingUsage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `filterProject`                                                                                | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `filterProducts`                                                                               | []*string*                                                                                     | :heavy_minus_sign:                                                                             | Allows to filter the billing usage for specific products. It accepts an array of product ids.<br/> |
| `filterPlan`                                                                                   | **string*                                                                                      | :heavy_minus_sign:                                                                             | Accepts a plan name and allows to filter the usage for that plan.<br/>                         |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetBillingUsageResponse](../../models/operations/getbillingusageresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |