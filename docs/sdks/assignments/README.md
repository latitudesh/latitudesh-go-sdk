# Firewalls.Assignments

## Overview

### Available Operations

* [Create](#create) - Assign server to firewall

## Create

Assigns a server to a firewall by its ID.

### Example Usage: Conflict

<!-- UsageSnippet language="go" operationID="create-firewall-assignment" method="post" path="/firewalls/{firewall_id}/assignments" example="Conflict" -->
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

    res, err := s.Firewalls.Assignments.Create(ctx, "fw_lQraYDPeOpjwW", operations.CreateFirewallAssignmentFirewallsAssignmentsRequestBody{
        Data: operations.CreateFirewallAssignmentFirewallsAssignmentsData{
            Type: operations.CreateFirewallAssignmentFirewallsAssignmentsTypeFirewallAssignments,
            Attributes: &operations.CreateFirewallAssignmentFirewallsAssignmentsAttributes{
                ServerID: "sv_aKXgRdR3qv9k5",
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
### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-firewall-assignment" method="post" path="/firewalls/{firewall_id}/assignments" example="Created" -->
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

    res, err := s.Firewalls.Assignments.Create(ctx, "fw_Ee8pKq05DWAob", operations.CreateFirewallAssignmentFirewallsAssignmentsRequestBody{
        Data: operations.CreateFirewallAssignmentFirewallsAssignmentsData{
            Type: operations.CreateFirewallAssignmentFirewallsAssignmentsTypeFirewallAssignments,
            Attributes: &operations.CreateFirewallAssignmentFirewallsAssignmentsAttributes{
                ServerID: "sv_aKXgRdR3qv9k5",
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
### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="create-firewall-assignment" method="post" path="/firewalls/{firewall_id}/assignments" example="Forbidden" -->
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

    res, err := s.Firewalls.Assignments.Create(ctx, "fw_5LA73qkjdaJ2o", operations.CreateFirewallAssignmentFirewallsAssignmentsRequestBody{
        Data: operations.CreateFirewallAssignmentFirewallsAssignmentsData{
            Type: operations.CreateFirewallAssignmentFirewallsAssignmentsTypeFirewallAssignments,
            Attributes: &operations.CreateFirewallAssignmentFirewallsAssignmentsAttributes{
                ServerID: "sv_byQrJdNJd30gv",
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
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="create-firewall-assignment" method="post" path="/firewalls/{firewall_id}/assignments" example="Not Found" -->
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

    res, err := s.Firewalls.Assignments.Create(ctx, "fw_RLYV8DZ2D5QoE", operations.CreateFirewallAssignmentFirewallsAssignmentsRequestBody{
        Data: operations.CreateFirewallAssignmentFirewallsAssignmentsData{
            Type: operations.CreateFirewallAssignmentFirewallsAssignmentsTypeFirewallAssignments,
            Attributes: &operations.CreateFirewallAssignmentFirewallsAssignmentsAttributes{
                ServerID: "sv_93YjJOLydvZ87",
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
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="create-firewall-assignment" method="post" path="/firewalls/{firewall_id}/assignments" example="Unprocessable Entity" -->
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

    res, err := s.Firewalls.Assignments.Create(ctx, "fw_Ee8pKq05DWAob", operations.CreateFirewallAssignmentFirewallsAssignmentsRequestBody{
        Data: operations.CreateFirewallAssignmentFirewallsAssignmentsData{
            Type: operations.CreateFirewallAssignmentFirewallsAssignmentsTypeFirewallAssignments,
            Attributes: &operations.CreateFirewallAssignmentFirewallsAssignmentsAttributes{
                ServerID: "sv_NGnzRD5ADM5yw",
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |