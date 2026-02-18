# Teams

## Overview

### Available Operations

* [Get](#get) - Retrieve team
* [Create](#create) - Create team
* [Update](#update) - Update team

## Get

Retrieve team

### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-team" method="get" path="/team" example="Success" -->
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

    res, err := s.Teams.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Teams != nil {
        // handle response
    }
}
```
### Example Usage: when team is older than one month

<!-- UsageSnippet language="go" operationID="get-team" method="get" path="/team" example="when team is older than one month" -->
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

    res, err := s.Teams.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Teams != nil {
        // handle response
    }
}
```
### Example Usage: when team is recent (created within last month)

<!-- UsageSnippet language="go" operationID="get-team" method="get" path="/team" example="when team is recent (created within last month)" -->
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

    res, err := s.Teams.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Teams != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTeamResponse](../../models/operations/getteamresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Create team

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="post-team" method="post" path="/team" example="Created" -->
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

    res, err := s.Teams.Create(ctx, operations.PostTeamTeamsRequestBody{
        Data: operations.PostTeamTeamsData{
            Type: operations.PostTeamTeamsTypeTeams,
            Attributes: &operations.PostTeamTeamsAttributes{
                Name: "Name",
                Currency: operations.PostTeamCurrencyUsd,
                Address: latitudeshgosdk.Pointer("Address"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: Not Acceptable

<!-- UsageSnippet language="go" operationID="post-team" method="post" path="/team" example="Not Acceptable" -->
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

    res, err := s.Teams.Create(ctx, operations.PostTeamTeamsRequestBody{
        Data: operations.PostTeamTeamsData{
            Type: operations.PostTeamTeamsTypeTeams,
            Attributes: &operations.PostTeamTeamsAttributes{
                Name: "Name",
                Currency: operations.PostTeamCurrencyUsd,
                Address: latitudeshgosdk.Pointer("Address"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="post-team" method="post" path="/team" example="Unprocessable Entity" -->
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

    res, err := s.Teams.Create(ctx, operations.PostTeamTeamsRequestBody{
        Data: operations.PostTeamTeamsData{
            Type: operations.PostTeamTeamsTypeTeams,
            Attributes: &operations.PostTeamTeamsAttributes{
                Name: "",
                Currency: operations.PostTeamCurrencyBrl,
                Address: latitudeshgosdk.Pointer(""),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PostTeamTeamsRequestBody](../../models/operations/postteamteamsrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PostTeamResponse](../../models/operations/postteamresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Update team

### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="patch-current-team" method="patch" path="/team/{team_id}" example="Forbidden" -->
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

    res, err := s.Teams.Update(ctx, "7107c56e-31be-45e4-9376-a905f204d08a", operations.PatchCurrentTeamTeamsRequestBody{
        Data: operations.PatchCurrentTeamTeamsData{
            ID: "team_8Wj589lRWNCK1gAje2pAfXmag5jg",
            Type: operations.PatchCurrentTeamTeamsTypeTeams,
            Attributes: &operations.PatchCurrentTeamTeamsAttributes{
                Address: latitudeshgosdk.Pointer("Address"),
                Name: latitudeshgosdk.Pointer("Name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="patch-current-team" method="patch" path="/team/{team_id}" example="Not Found" -->
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

    res, err := s.Teams.Update(ctx, "c7a799cc-bb1c-42e3-aaf9-a76b10f5c1ac", operations.PatchCurrentTeamTeamsRequestBody{
        Data: operations.PatchCurrentTeamTeamsData{
            ID: "c7a799cc-bb1c-42e3-aaf9-a76b10f5c1ac",
            Type: operations.PatchCurrentTeamTeamsTypeTeams,
            Attributes: &operations.PatchCurrentTeamTeamsAttributes{
                Address: latitudeshgosdk.Pointer("Address"),
                Name: latitudeshgosdk.Pointer("Name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="patch-current-team" method="patch" path="/team/{team_id}" example="Success" -->
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

    res, err := s.Teams.Update(ctx, "team_bVJM4y6m4VCyy101JzA3szlVGRb", operations.PatchCurrentTeamTeamsRequestBody{
        Data: operations.PatchCurrentTeamTeamsData{
            ID: "team_bVJM4y6m4VCyy101JzA3szlVGRb",
            Type: operations.PatchCurrentTeamTeamsTypeTeams,
            Attributes: &operations.PatchCurrentTeamTeamsAttributes{
                Address: latitudeshgosdk.Pointer("Address"),
                Name: latitudeshgosdk.Pointer("Name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `teamID`                                                                                                   | *string*                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `requestBody`                                                                                              | [operations.PatchCurrentTeamTeamsRequestBody](../../models/operations/patchcurrentteamteamsrequestbody.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PatchCurrentTeamResponse](../../models/operations/patchcurrentteamresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |