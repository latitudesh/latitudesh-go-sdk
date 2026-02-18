# TeamMembers

## Overview

### Available Operations

* [PostTeamMembers](#postteammembers) - Create member
* [Delete](#delete) - Remove a member

## PostTeamMembers

Create member

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="post-team-members" method="post" path="/team/members" example="Created" -->
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

    res, err := s.TeamMembers.PostTeamMembers(ctx, operations.PostTeamMembersTeamMembersRequestBody{
        Data: operations.PostTeamMembersTeamMembersData{
            Type: operations.PostTeamMembersTeamMembersTypeMemberships,
            Attributes: &operations.PostTeamMembersTeamMembersAttributes{
                FirstName: latitudeshgosdk.Pointer("Maricela"),
                LastName: latitudeshgosdk.Pointer("Torphy"),
                Email: "maritza_schneider@mcglynn.test",
                Role: operations.PostTeamMembersRoleCollaborator,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Membership != nil {
        // handle response
    }
}
```
### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="post-team-members" method="post" path="/team/members" example="Forbidden" -->
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

    res, err := s.TeamMembers.PostTeamMembers(ctx, operations.PostTeamMembersTeamMembersRequestBody{
        Data: operations.PostTeamMembersTeamMembersData{
            Type: operations.PostTeamMembersTeamMembersTypeMemberships,
            Attributes: &operations.PostTeamMembersTeamMembersAttributes{
                FirstName: latitudeshgosdk.Pointer("Waltraud"),
                LastName: latitudeshgosdk.Pointer("Feil"),
                Email: "denis_senger@paucek.test",
                Role: operations.PostTeamMembersRoleCollaborator,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Membership != nil {
        // handle response
    }
}
```
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="post-team-members" method="post" path="/team/members" example="Unprocessable Entity" -->
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

    res, err := s.TeamMembers.PostTeamMembers(ctx, operations.PostTeamMembersTeamMembersRequestBody{
        Data: operations.PostTeamMembersTeamMembersData{
            Type: operations.PostTeamMembersTeamMembersTypeMemberships,
            Attributes: &operations.PostTeamMembersTeamMembersAttributes{
                FirstName: latitudeshgosdk.Pointer("Gabriel"),
                LastName: latitudeshgosdk.Pointer("Roob"),
                Email: "Jack.Schamberger-Kerluke@gmail.com",
                Role: operations.PostTeamMembersRoleCollaborator,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Membership != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostTeamMembersTeamMembersRequestBody](../../models/operations/postteammembersteammembersrequestbody.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PostTeamMembersResponse](../../models/operations/postteammembersresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Remove a member

### Example Usage

<!-- UsageSnippet language="go" operationID="destroy-team-member" method="delete" path="/team/members/{user_id}" -->
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

    res, err := s.TeamMembers.Delete(ctx, "user_0MoLqJEYd57pY")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The user ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DestroyTeamMemberResponse](../../models/operations/destroyteammemberresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |