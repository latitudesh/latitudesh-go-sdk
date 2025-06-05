# Firewalls
(*Firewalls*)

## Overview

### Available Operations

* [CreateFirewall](#createfirewall) - Create a firewall
* [ListFirewalls](#listfirewalls) - List firewalls
* [GetFirewall](#getfirewall) - Retrieve Firewall
* [UpdateFirewall](#updatefirewall) - Update Firewall
* [DeleteFirewall](#deletefirewall) - Delete Firewall
* [GetFirewallAssignments](#getfirewallassignments) - Firewall Assignments
* [DeleteFirewallAssignment](#deletefirewallassignment) - Delete Firewall Assignment

## CreateFirewall

Create a firewall

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

    res, err := s.Firewalls.CreateFirewall(ctx, operations.CreateFirewallFirewallsRequestBody{
        Data: operations.CreateFirewallData{
            Type: operations.CreateFirewallTypeFirewalls,
            Attributes: &operations.CreateFirewallAttributes{
                Name: "my-firewall",
                Project: "sleek-steel-shirt",
                Rules: []operations.CreateFirewallRules{
                    operations.CreateFirewallRules{
                        From: latitudeshgosdk.String("192.168.42.72"),
                        To: latitudeshgosdk.String("192.168.43.51"),
                        Protocol: operations.CreateFirewallProtocolTCP.ToPointer(),
                        Port: latitudeshgosdk.String("80"),
                    },
                    operations.CreateFirewallRules{
                        From: latitudeshgosdk.String("192.168.1.16"),
                        To: latitudeshgosdk.String("192.168.1.30"),
                        Protocol: operations.CreateFirewallProtocolTCP.ToPointer(),
                        Port: latitudeshgosdk.String("80"),
                    },
                    operations.CreateFirewallRules{
                        From: latitudeshgosdk.String("192.168.1.10"),
                        To: latitudeshgosdk.String("192.168.1.20"),
                        Protocol: operations.CreateFirewallProtocolUDP.ToPointer(),
                        Port: latitudeshgosdk.String("3000-4000"),
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 422                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListFirewalls

List firewalls

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

    res, err := s.Firewalls.ListFirewalls(ctx, latitudeshgosdk.String("lightweight-silk-table"), latitudeshgosdk.Int64(20), latitudeshgosdk.Int64(1))
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

## GetFirewall

Retrieve a firewall

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

    res, err := s.Firewalls.GetFirewall(ctx, "fw_xkjQwdENqYNVP")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UpdateFirewall

Update a firewall

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

    res, err := s.Firewalls.UpdateFirewall(ctx, "fw_VaNmodjeObE8W", operations.UpdateFirewallFirewallsRequestBody{
        Data: operations.UpdateFirewallFirewallsData{
            Type: operations.UpdateFirewallFirewallsTypeFirewalls,
            Attributes: &operations.UpdateFirewallFirewallsAttributes{
                Rules: []operations.UpdateFirewallFirewallsRules{},
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404, 422                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteFirewall

Delete a firewall

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

    res, err := s.Firewalls.DeleteFirewall(ctx, "fw_123")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404, 422                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetFirewallAssignments

List servers assigned to a firewall

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

    res, err := s.Firewalls.GetFirewallAssignments(ctx, "fw_93YjJOLydvZ87", latitudeshgosdk.Int64(20), latitudeshgosdk.Int64(1))
    if err != nil {
        log.Fatal(err)
    }
    if res.FirewallServer != nil {
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 404                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteFirewallAssignment

Remove a server from a firewall

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

    res, err := s.Firewalls.DeleteFirewallAssignment(ctx, "fw_2695BdKrOevVo", "fwasg_6059EqYkOQj8p")
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |