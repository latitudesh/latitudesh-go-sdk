# SSHKeys

## Overview

### Available Operations

* [~~List~~](#list) - List all Project SSH Keys :warning: **Deprecated**
* [~~Get~~](#get) - Retrieve a Project SSH Key :warning: **Deprecated**
* [~~ModifyProjectKey~~](#modifyprojectkey) - Update a Project SSH Key :warning: **Deprecated**
* [~~RemoveFromProject~~](#removefromproject) - Delete a Project SSH Key :warning: **Deprecated**
* [ListAll](#listall) - List all SSH Keys
* [Create](#create) - Create a SSH Key
* [Retrieve](#retrieve) - Retrieve a SSH Key
* [Update](#update) - Update a SSH Key
* [Delete](#delete) - Delete a SSH Key

## ~~List~~

List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-project-ssh-keys" method="get" path="/projects/{project_id}/ssh_keys" -->
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

    res, err := s.SSHKeys.List(ctx, "proj_LMmAD8kEqwop2", latitudeshgosdk.Pointer("tag_K9kzXA45BEsjobKPNVEQCrrKbw4o,tag_my9K5XGPYBIr1ve257anIW4KRX2e"))
    if err != nil {
        log.Fatal(err)
    }
    if res.SSHKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |
| `projectID`                                                                                                                 | *string*                                                                                                                    | :heavy_check_mark:                                                                                                          | Project ID or Slug                                                                                                          |
| `filterTags`                                                                                                                | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | The tags ids to filter by, separated by comma, e.g. `filter[tags]=tag_1,tag_2`will return ssh keys with `tag_1` AND `tag_2` |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |

### Response

**[*operations.GetProjectSSHKeysResponse](../../models/operations/getprojectsshkeysresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~Get~~

List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-project-ssh-key" method="get" path="/projects/{project_id}/ssh_keys/{ssh_key_id}" -->
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

    res, err := s.SSHKeys.Get(ctx, "proj_g1mbDwrBqLv5B", "ssh_zGr47qlMDAg0m")
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project ID or Slug                                       |
| `sshKeyID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetProjectSSHKeyResponse](../../models/operations/getprojectsshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~ModifyProjectKey~~

Allow you update SSH Key in a project. These keys can be used to access servers after deploy and reinstall actions.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="put-project-ssh-key" method="patch" path="/projects/{project_id}/ssh_keys/{ssh_key_id}" -->
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

    res, err := s.SSHKeys.ModifyProjectKey(ctx, "proj_W6Q2D9lGqKLpr", "ssh_zlkg1DegdvZE5", operations.PutProjectSSHKeySSHKeysRequestBody{
        Data: operations.PutProjectSSHKeySSHKeysData{
            ID: latitudeshgosdk.Pointer("ssh_zlkg1DegdvZE5"),
            Type: operations.PutProjectSSHKeySSHKeysTypeSSHKeys,
            Attributes: &operations.PutProjectSSHKeySSHKeysAttributes{
                Tags: []string{
                    "tag_7YpWMjKZ8vujwmEV7blQhG87aaB",
                    "tag_E8ZoPkZPXXsgN2L46WVvTeWog1rz",
                },
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `projectID`                                                                                                    | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Project ID or Slug                                                                                             |
| `sshKeyID`                                                                                                     | *string*                                                                                                       | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `requestBody`                                                                                                  | [operations.PutProjectSSHKeySSHKeysRequestBody](../../models/operations/putprojectsshkeysshkeysrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PutProjectSSHKeyResponse](../../models/operations/putprojectsshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~RemoveFromProject~~

Allow you remove SSH Keys in a project. Remove a SSH Key from the project won't revoke the SSH Keys access for previously deploy and reinstall actions.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-project-ssh-key" method="delete" path="/projects/{project_id}/ssh_keys/{ssh_key_id}" -->
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

    res, err := s.SSHKeys.RemoveFromProject(ctx, "proj_LMmAD8k4qwop2", "ssh_7vYAZqGBdMQ94")
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project ID or Slug                                       |
| `sshKeyID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectSSHKeyResponse](../../models/operations/deleteprojectsshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ListAll

List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ssh-keys" method="get" path="/ssh_keys" -->
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

    res, err := s.SSHKeys.ListAll(ctx, nil, nil, latitudeshgosdk.Pointer("tag_lYXenaknWwsJ4rryVbNoUbn6p4V,tag_ARJPX1YRrgTKQ4xpvX5YuWNG2nwW"))
    if err != nil {
        log.Fatal(err)
    }
    if res.SSHKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |
| `filterProject`                                                                                                             | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | Project ID or slug                                                                                                          |
| `filterScope`                                                                                                               | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | Filter by scope: `project` (has projects), `team` (no projects), or empty (all)                                             |
| `filterTags`                                                                                                                | **string*                                                                                                                   | :heavy_minus_sign:                                                                                                          | The tags ids to filter by, separated by comma, e.g. `filter[tags]=tag_1,tag_2`will return ssh keys with `tag_1` AND `tag_2` |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |

### Response

**[*operations.GetSSHKeysResponse](../../models/operations/getsshkeysresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Allows you create SSH Keys. These keys can be used to access servers after deploy and reinstall actions.


### Example Usage

<!-- UsageSnippet language="go" operationID="post-ssh-key" method="post" path="/ssh_keys" -->
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

    res, err := s.SSHKeys.Create(ctx, operations.PostSSHKeySSHKeysRequestBody{
        Data: operations.PostSSHKeySSHKeysData{
            Type: operations.PostSSHKeySSHKeysTypeSSHKeys,
            Attributes: &operations.PostSSHKeySSHKeysAttributes{
                Name: latitudeshgosdk.Pointer("SSH Key"),
                Project: latitudeshgosdk.Pointer("proj_mw49QDB5qagKb"),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostSSHKeySSHKeysRequestBody](../../models/operations/postsshkeysshkeysrequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.PostSSHKeyResponse](../../models/operations/postsshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Retrieve

List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ssh-key" method="get" path="/ssh_keys/{ssh_key_id}" -->
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

    res, err := s.SSHKeys.Retrieve(ctx, "ssh_0MK4O4Zkqa95w")
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
| `sshKeyID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSSHKeyResponse](../../models/operations/getsshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Allows you update SSH Key in a project. These keys can be used to access servers after deploy and reinstall actions.


### Example Usage

<!-- UsageSnippet language="go" operationID="put-ssh-key" method="patch" path="/ssh_keys/{ssh_key_id}" -->
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

    res, err := s.SSHKeys.Update(ctx, "ssh_8Nkvdyn1qeLpx", operations.PutSSHKeySSHKeysRequestBody{
        Data: operations.PutSSHKeySSHKeysData{
            ID: latitudeshgosdk.Pointer("ssh_8Nkvdyn1qeLpx"),
            Type: operations.PutSSHKeySSHKeysTypeSSHKeys,
            Attributes: &operations.PutSSHKeySSHKeysAttributes{
                Tags: []string{
                    "tag_BR6a9QPmXWCPVzLrmEYAslyroMK",
                    "tag_yxVkGz7rxMsl9M54lr71fp5XZg60",
                },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `sshKeyID`                                                                                       | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `requestBody`                                                                                    | [operations.PutSSHKeySSHKeysRequestBody](../../models/operations/putsshkeysshkeysrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.PutSSHKeyResponse](../../models/operations/putsshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Allows you remove SSH Keys in a project. Remove a SSH Key from the project won't revoke the SSH Keys access for previously deploy and reinstall actions.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ssh-key" method="delete" path="/ssh_keys/{ssh_key_id}" -->
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

    res, err := s.SSHKeys.Delete(ctx, "ssh_lxWpD6xKdm6rk")
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
| `sshKeyID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSSHKeyResponse](../../models/operations/deletesshkeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |