# UserProfile

## Overview

### Available Operations

* [Get](#get) - Retrieve profile
* [Update](#update) - Update profile
* [ListTeams](#listteams) - List user teams

## Get

Retrieve the current user profile


### Example Usage

<!-- UsageSnippet language="go" operationID="get-user-profile" method="get" path="/user/profile" example="Success" -->
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

    res, err := s.UserProfile.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.GetUserProfileResponse](../../models/operations/getuserprofileresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Update the current user profile


### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="patch-user-profile" method="patch" path="/user/profile/{id}" example="Forbidden" -->
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

    res, err := s.UserProfile.Update(ctx, "user_kLn528op4zIMZ61MzRapu7XnAkL", operations.PatchUserProfileUserProfileRequestBody{
        Data: operations.PatchUserProfileUserProfileData{
            ID: "user_kLn528op4zIMZ61MzRapu7XnAkL",
            Type: operations.PatchUserProfileUserProfileTypeUsers,
            Attributes: &operations.PatchUserProfileUserProfileAttributes{
                Role: operations.PatchUserProfileUserProfileRoleCollaborator.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="patch-user-profile" method="patch" path="/user/profile/{id}" example="Success" -->
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

    res, err := s.UserProfile.Update(ctx, "user_8PMvy2GL72uKKyXyANEnsjMQYMP", operations.PatchUserProfileUserProfileRequestBody{
        Data: operations.PatchUserProfileUserProfileData{
            ID: "user_8PMvy2GL72uKKyXyANEnsjMQYMP",
            Type: operations.PatchUserProfileUserProfileTypeUsers,
            Attributes: &operations.PatchUserProfileUserProfileAttributes{
                Role: operations.PatchUserProfileUserProfileRoleCollaborator.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="patch-user-profile" method="patch" path="/user/profile/{id}" example="Unprocessable Entity" -->
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

    res, err := s.UserProfile.Update(ctx, "user_WeGoqA4ndP7nz", operations.PatchUserProfileUserProfileRequestBody{
        Data: operations.PatchUserProfileUserProfileData{
            ID: "user_WeGoqA4ndP7nz",
            Type: operations.PatchUserProfileUserProfileTypeUsers,
            Attributes: &operations.PatchUserProfileUserProfileAttributes{},
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `requestBody`                                                                                                          | [operations.PatchUserProfileUserProfileRequestBody](../../models/operations/patchuserprofileuserprofilerequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.PatchUserProfileResponse](../../models/operations/patchuserprofileresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ListTeams

Returns a list of all teams the user belongs to


### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-user-teams" method="get" path="/user/teams" example="Success" -->
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

    res, err := s.UserProfile.ListTeams(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserTeams != nil {
        // handle response
    }
}
```
### Example Usage: when team is older than one month

<!-- UsageSnippet language="go" operationID="get-user-teams" method="get" path="/user/teams" example="when team is older than one month" -->
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

    res, err := s.UserProfile.ListTeams(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserTeams != nil {
        // handle response
    }
}
```
### Example Usage: when team is recent (created within last month)

<!-- UsageSnippet language="go" operationID="get-user-teams" method="get" path="/user/teams" example="when team is recent (created within last month)" -->
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

    res, err := s.UserProfile.ListTeams(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserTeams != nil {
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

**[*operations.GetUserTeamsResponse](../../models/operations/getuserteamsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |