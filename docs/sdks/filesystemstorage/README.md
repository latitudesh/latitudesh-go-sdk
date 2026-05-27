# FilesystemStorage

## Overview

### Available Operations

* [CreateFilesystem](#createfilesystem) - Create filesystem
* [ListFilesystems](#listfilesystems) - List filesystems
* [DeleteFilesystem](#deletefilesystem) - Delete filesystem
* [UpdateFilesystem](#updatefilesystem) - Update filesystem

## CreateFilesystem

Allows you to add persistent storage to a project. These filesystems can be used to store data across your servers.

### Example Usage: Conflict

<!-- UsageSnippet language="go" operationID="post-storage-filesystems" method="post" path="/storage/filesystems" example="Conflict" -->
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

    res, err := s.FilesystemStorage.CreateFilesystem(ctx, operations.PostStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PostStorageFilesystemsFilesystemStorageData{
            Type: operations.PostStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsFilesystemStorageAttributes{
                Project: "proj_0L6WO19lOPlXy",
                Name: "my-data",
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
### Example Usage: Created

<!-- UsageSnippet language="go" operationID="post-storage-filesystems" method="post" path="/storage/filesystems" example="Created" -->
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

    res, err := s.FilesystemStorage.CreateFilesystem(ctx, operations.PostStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PostStorageFilesystemsFilesystemStorageData{
            Type: operations.PostStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsFilesystemStorageAttributes{
                Project: "proj_lkg1De6ROvZE5",
                Name: "my-data",
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
### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="post-storage-filesystems" method="post" path="/storage/filesystems" example="Forbidden" -->
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

    res, err := s.FilesystemStorage.CreateFilesystem(ctx, operations.PostStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PostStorageFilesystemsFilesystemStorageData{
            Type: operations.PostStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsFilesystemStorageAttributes{
                Project: "proj_3YjJOLejdvZ87",
                Name: "my-data",
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

<!-- UsageSnippet language="go" operationID="post-storage-filesystems" method="post" path="/storage/filesystems" example="Storage creation frozen" -->
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

    res, err := s.FilesystemStorage.CreateFilesystem(ctx, operations.PostStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PostStorageFilesystemsFilesystemStorageData{
            Type: operations.PostStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsFilesystemStorageAttributes{
                Project: "<value>",
                Name: "<value>",
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

<!-- UsageSnippet language="go" operationID="post-storage-filesystems" method="post" path="/storage/filesystems" example="Unprocessable Entity" -->
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

    res, err := s.FilesystemStorage.CreateFilesystem(ctx, operations.PostStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PostStorageFilesystemsFilesystemStorageData{
            Type: operations.PostStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsFilesystemStorageAttributes{
                Project: "proj_pRMLydp0dQKr1",
                Name: "test storage @",
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.PostStorageFilesystemsFilesystemStorageRequestBody](../../models/operations/poststoragefilesystemsfilesystemstoragerequestbody.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.PostStorageFilesystemsResponse](../../models/operations/poststoragefilesystemsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListFilesystems

Lists all the filesystems from a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-filesystems" method="get" path="/storage/filesystems" -->
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
    if res != nil {
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

**[*operations.GetStorageFilesystemsResponse](../../models/operations/getstoragefilesystemsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteFilesystem

Allows you to remove a filesystem from a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-storage-filesystems" method="delete" path="/storage/filesystems/{filesystem_id}" -->
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

    res, err := s.FilesystemStorage.DeleteFilesystem(ctx, "fs_123")
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

**[*operations.DeleteStorageFilesystemsResponse](../../models/operations/deletestoragefilesystemsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateFilesystem

Allow you to upgrade the size of a filesystem.

### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="patch-storage-filesystems" method="patch" path="/storage/filesystems/{filesystem_id}" example="Forbidden" -->
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

    res, err := s.FilesystemStorage.UpdateFilesystem(ctx, "fs_x1ZJrdx5qg4LV", operations.PatchStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PatchStorageFilesystemsFilesystemStorageData{
            ID: "fs_x1ZJrdx5qg4LV",
            Type: operations.PatchStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PatchStorageFilesystemsFilesystemStorageAttributes{
                SizeInGb: latitudeshgosdk.Pointer[int64](1501),
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

<!-- UsageSnippet language="go" operationID="patch-storage-filesystems" method="patch" path="/storage/filesystems/{filesystem_id}" example="Success" -->
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

    res, err := s.FilesystemStorage.UpdateFilesystem(ctx, "fs_7vYAZqGBdMQ94", operations.PatchStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PatchStorageFilesystemsFilesystemStorageData{
            ID: "fs_7vYAZqGBdMQ94",
            Type: operations.PatchStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PatchStorageFilesystemsFilesystemStorageAttributes{
                SizeInGb: latitudeshgosdk.Pointer[int64](1501),
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
### Example Usage: Validation Error

<!-- UsageSnippet language="go" operationID="patch-storage-filesystems" method="patch" path="/storage/filesystems/{filesystem_id}" example="Validation Error" -->
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

    res, err := s.FilesystemStorage.UpdateFilesystem(ctx, "fs_r0MK4O4kDa95w", operations.PatchStorageFilesystemsFilesystemStorageRequestBody{
        Data: operations.PatchStorageFilesystemsFilesystemStorageData{
            ID: "fs_r0MK4O4kDa95w",
            Type: operations.PatchStorageFilesystemsFilesystemStorageTypeFilesystems,
            Attributes: operations.PatchStorageFilesystemsFilesystemStorageAttributes{
                SizeInGb: latitudeshgosdk.Pointer[int64](1499),
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `filesystemID`                                                                                                                                   | `string`                                                                                                                                         | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `requestBody`                                                                                                                                    | [operations.PatchStorageFilesystemsFilesystemStorageRequestBody](../../models/operations/patchstoragefilesystemsfilesystemstoragerequestbody.md) | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.PatchStorageFilesystemsResponse](../../models/operations/patchstoragefilesystemsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |