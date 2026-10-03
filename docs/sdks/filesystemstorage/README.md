# FilesystemStorage

## Overview

### Available Operations

* [CreateFilesystem](#createfilesystem) - Create filesystem
* [ListFilesystems](#listfilesystems) - List filesystems
* [DeleteFilesystem](#deletefilesystem) - Delete filesystem
* [UpdateFilesystem](#updatefilesystem) - Update filesystem

## CreateFilesystem

Allows you to add persistent storage to a project. These filesystems can be used to store data across your servers.

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-filesystem" method="post" path="/storage/filesystems" example="Created" -->
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

    res, err := s.FilesystemStorage.CreateFilesystem(ctx, operations.CreateFilesystemFilesystemStorageRequestBody{
        Data: operations.CreateFilesystemFilesystemStorageData{
            Type: operations.CreateFilesystemFilesystemStorageTypeFilesystems,
            Attributes: operations.CreateFilesystemFilesystemStorageAttributes{
                Project: "proj_lkg1De6ROvZE5",
                Name: "my-data",
                Region: "NYC",
                Protocols: []operations.CreateFilesystemProtocols{
                    operations.CreateFilesystemProtocolsNfs3,
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
### Example Usage: Storage creation frozen

<!-- UsageSnippet language="go" operationID="create-filesystem" method="post" path="/storage/filesystems" example="Storage creation frozen" -->
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

    res, err := s.FilesystemStorage.CreateFilesystem(ctx, operations.CreateFilesystemFilesystemStorageRequestBody{
        Data: operations.CreateFilesystemFilesystemStorageData{
            Type: operations.CreateFilesystemFilesystemStorageTypeFilesystems,
            Attributes: operations.CreateFilesystemFilesystemStorageAttributes{
                Project: "<value>",
                Name: "<value>",
                Region: "<value>",
                Protocols: []operations.CreateFilesystemProtocols{
                    operations.CreateFilesystemProtocolsNfs4,
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.CreateFilesystemFilesystemStorageRequestBody](../../models/operations/createfilesystemfilesystemstoragerequestbody.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CreateFilesystemResponse](../../models/operations/createfilesystemresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListFilesystems

Lists all the filesystems from a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-filesystems" method="get" path="/storage/filesystems" example="Success" -->
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

    res, err := s.FilesystemStorage.ListFilesystems(ctx, latitudeshgosdk.Pointer("small-rubber-shirt"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Filesystems != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `filterProject`                                          | `*string`                                                | :heavy_minus_sign:                                       | The project ID or Slug to filter by                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListFilesystemsResponse](../../models/operations/listfilesystemsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteFilesystem

Allows you to remove a filesystem from a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-filesystem" method="delete" path="/storage/filesystems/{filesystem_id}" -->
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

    res, err := s.FilesystemStorage.DeleteFilesystem(ctx, "<id>")
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
| `filesystemID`                                           | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteFilesystemResponse](../../models/operations/deletefilesystemresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateFilesystem

Allow you to upgrade the size of a filesystem.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-filesystem" method="patch" path="/storage/filesystems/{filesystem_id}" example="Success" -->
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

    res, err := s.FilesystemStorage.UpdateFilesystem(ctx, "fs_7vYAZqGBdMQ94", operations.UpdateFilesystemFilesystemStorageRequestBody{
        Data: operations.UpdateFilesystemFilesystemStorageData{
            ID: latitudeshgosdk.Pointer("fs_7vYAZqGBdMQ94"),
            Type: operations.UpdateFilesystemFilesystemStorageTypeFilesystems,
            Attributes: operations.UpdateFilesystemFilesystemStorageAttributes{
                SizeInGb: 1501,
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `filesystemID`                                                                                                                     | `string`                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `requestBody`                                                                                                                      | [operations.UpdateFilesystemFilesystemStorageRequestBody](../../models/operations/updatefilesystemfilesystemstoragerequestbody.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.UpdateFilesystemResponse](../../models/operations/updatefilesystemresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |