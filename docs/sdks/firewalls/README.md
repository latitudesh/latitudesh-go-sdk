# Firewalls

## Overview

### Available Operations

* [GetAllFirewallAssignments](#getallfirewallassignments) - List firewall assignments
* [Create](#create) - Create firewall
* [List](#list) - List firewalls
* [Get](#get) - Retrieve firewall
* [Update](#update) - Update firewall
* [Delete](#delete) - Delete firewall
* [ListAssignments](#listassignments) - Firewall assignments
* [DeleteAssignment](#deleteassignment) - Delete assignment

## GetAllFirewallAssignments

Returns a list of all servers assigned to one or more firewalls.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-all-firewall-assignments" method="get" path="/firewalls/assignments" -->
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

    res, err := s.Firewalls.GetAllFirewallAssignments(ctx, latitudeshgosdk.Pointer("sv_Qk0Ryqv1dW36X"), latitudeshgosdk.Pointer[int64](20), latitudeshgosdk.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.FirewallAssignments != nil {
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
| `filterServer`                                           | **string*                                                | :heavy_minus_sign:                                       | The server ID to filter by                               |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Number of items to return per page                       |
| `pageNumber`                                             | **int64*                                                 | :heavy_minus_sign:                                       | Page number to return (starts at 1)                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAllFirewallAssignmentsResponse](../../models/operations/getallfirewallassignmentsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Create a firewall

### Example Usage

<!-- UsageSnippet language="go" operationID="create-firewall" method="post" path="/firewalls" -->
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

    res, err := s.Firewalls.Create(ctx, operations.CreateFirewallFirewallsRequestBody{
        Data: operations.CreateFirewallData{
            Type: operations.CreateFirewallTypeFirewalls,
            Attributes: &operations.CreateFirewallAttributes{
                Name: "my-firewall",
                Project: "heavy-duty-copper-watch",
                Rules: []operations.CreateFirewallRules{
                    operations.CreateFirewallRules{
                        From: latitudeshgosdk.Pointer("192.168.42.73"),
                        To: latitudeshgosdk.Pointer("192.168.43.51"),
                        Protocol: operations.CreateFirewallProtocolTCP.ToPointer(),
                        Port: latitudeshgosdk.Pointer("80"),
                        Description: latitudeshgosdk.Pointer("Allow HTTP traffic"),
                    },
                    operations.CreateFirewallRules{
                        From: latitudeshgosdk.Pointer("192.168.1.16"),
                        To: latitudeshgosdk.Pointer("192.168.1.30"),
                        Protocol: operations.CreateFirewallProtocolTCP.ToPointer(),
                        Port: latitudeshgosdk.Pointer("80"),
                    },
                    operations.CreateFirewallRules{
                        From: latitudeshgosdk.Pointer("192.168.1.10"),
                        To: latitudeshgosdk.Pointer("192.168.1.20"),
                        Protocol: operations.CreateFirewallProtocolUDP.ToPointer(),
                        Port: latitudeshgosdk.Pointer("3000-4000"),
                        Description: latitudeshgosdk.Pointer("Application ports"),
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Firewall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CreateFirewallFirewallsRequestBody](../../models/operations/createfirewallfirewallsrequestbody.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateFirewallResponse](../../models/operations/createfirewallresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## List

List firewalls

### Example Usage

<!-- UsageSnippet language="go" operationID="list-firewalls" method="get" path="/firewalls" -->
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

    res, err := s.Firewalls.List(ctx, latitudeshgosdk.Pointer("intelligent-marble-lamp"), latitudeshgosdk.Pointer[int64](20), latitudeshgosdk.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.Firewalls != nil {
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
| `filterProject`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Number of items to return per page                       |
| `pageNumber`                                             | **int64*                                                 | :heavy_minus_sign:                                       | Page number to return (starts at 1)                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListFirewallsResponse](../../models/operations/listfirewallsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Get

Returns a single firewall by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-firewall" method="get" path="/firewalls/{firewall_id}" -->
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

    res, err := s.Firewalls.Get(ctx, "fw_6A05EdQ1dvKYQ")
    if err != nil {
        log.Fatal(err)
    }
    if res.Firewall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `firewallID`                                             | *string*                                                 | :heavy_check_mark:                                       | The Firewall ID                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetFirewallResponse](../../models/operations/getfirewallresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Updates a firewall by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-firewall" method="patch" path="/firewalls/{firewall_id}" -->
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

    res, err := s.Firewalls.Update(ctx, "fw_r0MK4O4kDa95w", operations.UpdateFirewallFirewallsRequestBody{
        Data: operations.UpdateFirewallFirewallsData{
            Type: operations.UpdateFirewallFirewallsTypeFirewalls,
            Attributes: &operations.UpdateFirewallFirewallsAttributes{
                Name: latitudeshgosdk.Pointer("new-name"),
                Rules: []operations.UpdateFirewallFirewallsRules{
                    operations.UpdateFirewallFirewallsRules{
                        From: latitudeshgosdk.Pointer("192.168.42.72"),
                        To: latitudeshgosdk.Pointer("192.168.43.51"),
                        Protocol: operations.UpdateFirewallFirewallsProtocolTCP.ToPointer(),
                        Port: latitudeshgosdk.Pointer("80"),
                        Description: latitudeshgosdk.Pointer("Allow HTTP"),
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Firewall != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `firewallID`                                                                                                   | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The Firewall ID                                                                                                |
| `requestBody`                                                                                                  | [operations.UpdateFirewallFirewallsRequestBody](../../models/operations/updatefirewallfirewallsrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateFirewallResponse](../../models/operations/updatefirewallresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Delete firewall

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-firewall" method="delete" path="/firewalls/{firewall_id}" -->
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

    res, err := s.Firewalls.Delete(ctx, "fw_123")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `firewallID`                                             | *string*                                                 | :heavy_check_mark:                                       | The Firewall ID                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteFirewallResponse](../../models/operations/deletefirewallresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ListAssignments

Returns a list of all servers assigned to a particular firewall.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-firewall-assignments" method="get" path="/firewalls/{firewall_id}/assignments" -->
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

    res, err := s.Firewalls.ListAssignments(ctx, "fw_Qk0Ryqv1dW36X", latitudeshgosdk.Pointer[int64](20), latitudeshgosdk.Pointer[int64](1))
    if err != nil {
        log.Fatal(err)
    }
    if res.FirewallAssignments != nil {
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
| `firewallID`                                             | *string*                                                 | :heavy_check_mark:                                       | The Firewall ID                                          |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Number of items to return per page                       |
| `pageNumber`                                             | **int64*                                                 | :heavy_minus_sign:                                       | Page number to return (starts at 1)                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetFirewallAssignmentsResponse](../../models/operations/getfirewallassignmentsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteAssignment

Removes a server from a firewall by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-firewall-assignment" method="delete" path="/firewalls/{firewall_id}/assignments/{assignment_id}" -->
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

    res, err := s.Firewalls.DeleteAssignment(ctx, "fw_2695BdKrOevVo", "fwasg_6059EqYkOQj8p")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `firewallID`                                             | *string*                                                 | :heavy_check_mark:                                       | The Firewall ID                                          |
| `assignmentID`                                           | *string*                                                 | :heavy_check_mark:                                       | The Assignment ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteFirewallAssignmentResponse](../../models/operations/deletefirewallassignmentresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |