# TeamMembers

## Overview

### Available Operations

* [PostTeamMembers](#postteammembers) - Add a Team Member
* [Delete](#delete) - Remove a Team Member

## PostTeamMembers

Add a Team Member

### Example Usage

<!-- UsageSnippet language="go" operationID="post-team-members" method="post" path="/team/members" -->
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
                FirstName: latitudeshgosdk.Pointer("Bernard"),
                LastName: latitudeshgosdk.Pointer("Cremin"),
                Email: "ernest@carter-lehner.example",
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

Remove a Team Member

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