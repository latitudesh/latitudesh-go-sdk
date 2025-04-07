# Assignments
(*Firewalls.Assignments*)

## Overview

### Available Operations

* [Create](#create) - Firewall Assignment

## Create

Assign a server to a firewall

### Example Usage

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

    res, err := s.Firewalls.Assignments.Create(ctx, "fw_Av9BVDavORm1W", operations.CreateFirewallAssignmentFirewallsAssignmentsRequestBody{
        Data: operations.CreateFirewallAssignmentFirewallsAssignmentsData{
            Type: operations.CreateFirewallAssignmentFirewallsAssignmentsTypeFirewallAssignments,
            Attributes: &operations.CreateFirewallAssignmentFirewallsAssignmentsAttributes{
                ServerID: "sv_2695BdKrOevVo",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FirewallServer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                | Type                                                                                                                                                     | Required                                                                                                                                                 | Description                                                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                    | :heavy_check_mark:                                                                                                                                       | The context to use for the request.                                                                                                                      |
| `firewallID`                                                                                                                                             | *string*                                                                                                                                                 | :heavy_check_mark:                                                                                                                                       | The Firewall ID                                                                                                                                          |
| `requestBody`                                                                                                                                            | [operations.CreateFirewallAssignmentFirewallsAssignmentsRequestBody](../../models/operations/createfirewallassignmentfirewallsassignmentsrequestbody.md) | :heavy_check_mark:                                                                                                                                       | N/A                                                                                                                                                      |
| `opts`                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                       | The options for this request.                                                                                                                            |

### Response

**[*operations.CreateFirewallAssignmentResponse](../../models/operations/createfirewallassignmentresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409, 422       | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |