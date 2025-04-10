# TeamMembers
(*TeamMembers*)

## Overview

### Available Operations

* [Add](#add) - Add a Team Member

## Add

Add a Team Member

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

    res, err := s.TeamMembers.Add(ctx, operations.PostTeamMembersTeamMembersRequestBody{
        Data: operations.PostTeamMembersTeamMembersData{
            Type: operations.PostTeamMembersTeamMembersTypeMemberships,
            Attributes: &operations.PostTeamMembersTeamMembersAttributes{
                FirstName: latitudeshgosdk.String("Dianne"),
                LastName: latitudeshgosdk.String("Bahringer"),
                Email: "theron_keeling@balistreri.test",
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

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 422                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |