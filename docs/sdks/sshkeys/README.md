# SSHKeys
(*SSHKeys*)

## Overview

### Available Operations

* [List](#list) - List all Project SSH Keys
* [Get](#get) - Retrieve a Project SSH Key
* [Update](#update) - Update a Project SSH Key
* [Delete](#delete) - Delete a Project SSH Key

## List

List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.


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

    res, err := s.SSHKeys.List(ctx, "proj_pbV0Dg2Gq4AWz", latitudeshgosdk.String("tag_9KjPQYXN1RsEzzEJ7JJKfo7ykalQ"))
    if err != nil {
        log.Fatal(err)
    }
    if res.SSHKey != nil {
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

## Get

List all SSH Keys in the project. These keys can be used to access servers after deploy and reinstall actions.


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

    res, err := s.SSHKeys.Get(ctx, "proj_QraYDP41DpjwW", "ssh_j0L6WO1QOPlXy")
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

## Update

Allow you update SSH Key in a project. These keys can be used to access servers after deploy and reinstall actions.


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

    res, err := s.SSHKeys.Update(ctx, "proj_GMy1DbW0ON50m", "ssh_5AEmq71XOBkWX", operations.PutProjectSSHKeySSHKeysRequestBody{
        Data: operations.PutProjectSSHKeySSHKeysData{
            Type: operations.PutProjectSSHKeySSHKeysTypeSSHKeys,
            Attributes: &operations.PutProjectSSHKeySSHKeysAttributes{
                Tags: []string{
                    "tag_V8rnwNgREksbr9yA4RagcZMkWlP",
                    "tag_KJJ8erRzEbijlJkLjWVVIJ4Gm7A",
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

## Delete

Allow you remove SSH Keys in a project. Remove a SSH Key from the project won't revoke the SSH Keys access for previously deploy and reinstall actions.


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

    res, err := s.SSHKeys.Delete(ctx, "proj_WVQJDMAvDRbyE", "ssh_kjQwdEGNDYNVP")
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