# ~~Projects.SshKeys~~

> [!WARNING]
> This SDK is **DEPRECATED**

## Overview

### Available Operations

* [~~PostProjectSSHKey~~](#postprojectsshkey) - Create a Project SSH Key :warning: **Deprecated**

## ~~PostProjectSSHKey~~

Allow you create SSH Keys in a project. These keys can be used to access servers after deploy and reinstall actions.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="post-project-ssh-key" method="post" path="/projects/{project_id}/ssh_keys" -->
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

    res, err := s.Projects.SSHKeys.PostProjectSSHKey(ctx, "proj_RMLydp70OQKr1", operations.PostProjectSSHKeyProjectsSSHKeysRequestBody{
        Data: operations.PostProjectSSHKeyProjectsSSHKeysData{
            Type: operations.PostProjectSSHKeyProjectsSSHKeysTypeSSHKeys,
            Attributes: &operations.PostProjectSSHKeyProjectsSSHKeysAttributes{
                Name: latitudeshgosdk.Pointer("SSH Key"),
                PublicKey: latitudeshgosdk.Pointer("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOLFnjGP3Jsh1usHNS2EILgfqZNC9pOvNqBZqxH+qNAdZdQCzy2csMuiq+ZwLA8Mm4Vo5CvSgBHs/kuZRUKyTl+79YUMZIj8PhHzL4XbdqX1ZnAIklHWcJaveB0+UXLEPKGzFIFq+FkuwtiXQsVe5NnSpIDYgpzhqEs38NsnXvsubKphGUdARDhaxvMdUUl4YsAtLHKMzSyIvE6xwfTtIVwA9bZt/8GoBzrn9px9PEcf25Rgd2NhOYs3WYcZuwvRmfcFdi2vGhVqTPqL9n16R/n5jknxHYrTyqWNxJdpdvg2YqXpN7vnFNoOjYFD6EahJ0pF/+WL4tPCIkLfoaVaSx"),
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `projectID`                                                                                                                      | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | Project ID or Slug                                                                                                               |
| `requestBody`                                                                                                                    | [operations.PostProjectSSHKeyProjectsSSHKeysRequestBody](../../models/operations/postprojectsshkeyprojectssshkeysrequestbody.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.PostProjectSSHKeyResponse](../../models/operations/postprojectsshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |