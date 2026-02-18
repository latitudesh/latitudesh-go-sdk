# Storage

## Overview

### Available Operations

* [CreateFilesystem](#createfilesystem) - Create filesystem
* [ListFilesystems](#listfilesystems) - List filesystems
* [DeleteFilesystem](#deletefilesystem) - Delete filesystem
* [UpdateFilesystem](#updatefilesystem) - Update filesystem
* [GetStorageVolumes](#getstoragevolumes) - List volumes
* [PostStorageVolumes](#poststoragevolumes) - Create volume
* [GetStorageVolume](#getstoragevolume) - Retrieve volume
* [DeleteStorageVolumes](#deletestoragevolumes) - Delete volume
* [PostStorageVolumesMount](#poststoragevolumesmount) - Mount volume

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

    res, err := s.Storage.CreateFilesystem(ctx, operations.PostStorageFilesystemsStorageRequestBody{
        Data: operations.PostStorageFilesystemsStorageData{
            Type: operations.PostStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsStorageAttributes{
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

    res, err := s.Storage.CreateFilesystem(ctx, operations.PostStorageFilesystemsStorageRequestBody{
        Data: operations.PostStorageFilesystemsStorageData{
            Type: operations.PostStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsStorageAttributes{
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

    res, err := s.Storage.CreateFilesystem(ctx, operations.PostStorageFilesystemsStorageRequestBody{
        Data: operations.PostStorageFilesystemsStorageData{
            Type: operations.PostStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsStorageAttributes{
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

    res, err := s.Storage.CreateFilesystem(ctx, operations.PostStorageFilesystemsStorageRequestBody{
        Data: operations.PostStorageFilesystemsStorageData{
            Type: operations.PostStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsStorageAttributes{
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

    res, err := s.Storage.CreateFilesystem(ctx, operations.PostStorageFilesystemsStorageRequestBody{
        Data: operations.PostStorageFilesystemsStorageData{
            Type: operations.PostStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PostStorageFilesystemsStorageAttributes{
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.PostStorageFilesystemsStorageRequestBody](../../models/operations/poststoragefilesystemsstoragerequestbody.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

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

    res, err := s.Storage.ListFilesystems(ctx, latitudeshgosdk.Pointer("small-rubber-shirt"))
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
| `filterProject`                                          | **string*                                                | :heavy_minus_sign:                                       | The project ID or Slug to filter by                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageFilesystemsResponse](../../models/operations/getstoragefilesystemsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteFilesystem

Allows you to remove persistent storage from a project.

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

    res, err := s.Storage.DeleteFilesystem(ctx, "fs_123")
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
| `filesystemID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
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

    res, err := s.Storage.UpdateFilesystem(ctx, "fs_x1ZJrdx5qg4LV", operations.PatchStorageFilesystemsStorageRequestBody{
        Data: operations.PatchStorageFilesystemsStorageData{
            ID: "fs_x1ZJrdx5qg4LV",
            Type: operations.PatchStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PatchStorageFilesystemsStorageAttributes{
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

    res, err := s.Storage.UpdateFilesystem(ctx, "fs_7vYAZqGBdMQ94", operations.PatchStorageFilesystemsStorageRequestBody{
        Data: operations.PatchStorageFilesystemsStorageData{
            ID: "fs_7vYAZqGBdMQ94",
            Type: operations.PatchStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PatchStorageFilesystemsStorageAttributes{
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

    res, err := s.Storage.UpdateFilesystem(ctx, "fs_r0MK4O4kDa95w", operations.PatchStorageFilesystemsStorageRequestBody{
        Data: operations.PatchStorageFilesystemsStorageData{
            ID: "fs_r0MK4O4kDa95w",
            Type: operations.PatchStorageFilesystemsStorageTypeFilesystems,
            Attributes: operations.PatchStorageFilesystemsStorageAttributes{
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `filesystemID`                                                                                                               | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `requestBody`                                                                                                                | [operations.PatchStorageFilesystemsStorageRequestBody](../../models/operations/patchstoragefilesystemsstoragerequestbody.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.PatchStorageFilesystemsResponse](../../models/operations/patchstoragefilesystemsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetStorageVolumes

Lists all the volumes from a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-volumes" method="get" path="/storage/volumes" example="Success" -->
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

    res, err := s.Storage.GetStorageVolumes(ctx, latitudeshgosdk.Pointer("proj_WeGoqA5AqP7nz"))
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
| `filterProject`                                          | **string*                                                | :heavy_minus_sign:                                       | The project ID or Slug to filter by                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageVolumesResponse](../../models/operations/getstoragevolumesresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## PostStorageVolumes

Allows you to add persistent storage to a project. These volumes can be used to store data across your servers.

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="post-storage-volumes" method="post" path="/storage/volumes" example="Created" -->
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

    res, err := s.Storage.PostStorageVolumes(ctx, operations.PostStorageVolumesStorageRequestBody{
        Data: operations.PostStorageVolumesStorageData{
            Type: operations.PostStorageVolumesStorageTypeVolumes,
            Attributes: operations.PostStorageVolumesStorageAttributes{
                Project: "proj_enPbqoZ6dA2MQ",
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

<!-- UsageSnippet language="go" operationID="post-storage-volumes" method="post" path="/storage/volumes" example="Storage creation frozen" -->
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

    res, err := s.Storage.PostStorageVolumes(ctx, operations.PostStorageVolumesStorageRequestBody{
        Data: operations.PostStorageVolumesStorageData{
            Type: operations.PostStorageVolumesStorageTypeVolumes,
            Attributes: operations.PostStorageVolumesStorageAttributes{
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

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostStorageVolumesStorageRequestBody](../../models/operations/poststoragevolumesstoragerequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.PostStorageVolumesResponse](../../models/operations/poststoragevolumesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetStorageVolume

Shows details of a specific volume storage.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-volume" method="get" path="/storage/volumes/{id}" example="Success" -->
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

    res, err := s.Storage.GetStorageVolume(ctx, "vol_aKXgRdR3qv9k5")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The volume storage ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageVolumeResponse](../../models/operations/getstoragevolumeresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteStorageVolumes

Allows you to remove persistent storage from a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-storage-volumes" method="delete" path="/storage/volumes/{id}" -->
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

    res, err := s.Storage.DeleteStorageVolumes(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteStorageVolumesResponse](../../models/operations/deletestoragevolumesresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## PostStorageVolumesMount

Mounts volume storage by adding the client to an allowed list

### Example Usage

<!-- UsageSnippet language="go" operationID="post-storage-volumes-mount" method="post" path="/storage/volumes/{id}/mount" -->
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

    res, err := s.Storage.PostStorageVolumesMount(ctx, "<id>", operations.PostStorageVolumesMountRequestBody{
        Data: operations.PostStorageVolumesMountData{
            Type: operations.PostStorageVolumesMountTypeVolumes,
            Attributes: operations.PostStorageVolumesMountAttributes{
                Nqn: "nqn.2024-01.com.example:server01",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `id`                                                                                                           | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Volume storage ID                                                                                              |
| `requestBody`                                                                                                  | [operations.PostStorageVolumesMountRequestBody](../../models/operations/poststoragevolumesmountrequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PostStorageVolumesMountResponse](../../models/operations/poststoragevolumesmountresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |