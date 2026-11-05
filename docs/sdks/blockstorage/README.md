# BlockStorage

## Overview

### Available Operations

* [ListVolumes](#listvolumes) - List volumes
* [CreateVolume](#createvolume) - Create volume
* [RetrieveVolume](#retrievevolume) - Retrieve volume
* [DeleteVolume](#deletevolume) - Delete volume
* [UpdateVolume](#updatevolume) - Update volume
* [~~MountVolume~~](#mountvolume) - Mount volume (deprecated) :warning: **Deprecated**
* [MapVolume](#mapvolume) - Map volume
* [UnmapVolume](#unmapvolume) - Unmap volume

## ListVolumes

Lists all the volumes from a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-volumes" method="get" path="/storage/volumes" example="Success" -->
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

    res, err := s.BlockStorage.ListVolumes(ctx, latitudeshgosdk.Pointer("proj_WeGoqA5AqP7nz"))
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
| `filterProject`                                          | `*string`                                                | :heavy_minus_sign:                                       | The project ID or Slug to filter by                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListVolumesResponse](../../models/operations/listvolumesresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateVolume

Allows you to add persistent storage to a project. These volumes can be used to store data across your servers.

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-volume" method="post" path="/storage/volumes" example="Created" -->
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

    res, err := s.BlockStorage.CreateVolume(ctx, operations.CreateVolumeBlockStorageRequestBody{
        Data: operations.CreateVolumeBlockStorageData{
            Type: operations.CreateVolumeBlockStorageTypeVolumes,
            Attributes: operations.CreateVolumeBlockStorageAttributes{
                Project: "proj_enPbqoZ6dA2MQ",
                Name: "my-data",
                Region: "DAL",
                SizeInGb: 1500,
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
### Example Usage: LocationUnavailable

<!-- UsageSnippet language="go" operationID="create-volume" method="post" path="/storage/volumes" example="LocationUnavailable" -->
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

    res, err := s.BlockStorage.CreateVolume(ctx, operations.CreateVolumeBlockStorageRequestBody{
        Data: operations.CreateVolumeBlockStorageData{
            Type: operations.CreateVolumeBlockStorageTypeVolumes,
            Attributes: operations.CreateVolumeBlockStorageAttributes{
                Project: "<value>",
                Name: "<value>",
                Region: "<value>",
                SizeInGb: 185826,
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateVolumeBlockStorageRequestBody](../../models/operations/createvolumeblockstoragerequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateVolumeResponse](../../models/operations/createvolumeresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 422                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## RetrieveVolume

Shows details of a specific volume.

### Example Usage

<!-- UsageSnippet language="go" operationID="retrieve-volume" method="get" path="/storage/volumes/{id}" example="Success" -->
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

    res, err := s.BlockStorage.RetrieveVolume(ctx, "vol_aKXgRdR3qv9k5")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The volume ID                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RetrieveVolumeResponse](../../models/operations/retrievevolumeresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteVolume

Allows you to remove a volume from a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-volume" method="delete" path="/storage/volumes/{id}" -->
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

    res, err := s.BlockStorage.DeleteVolume(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteVolumeResponse](../../models/operations/deletevolumeresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateVolume

Increases the size of a high performance volume. Shrinking is not supported. Billing is prorated to the new size.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-volume" method="patch" path="/storage/volumes/{id}" -->
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

    res, err := s.BlockStorage.UpdateVolume(ctx, "<id>", operations.UpdateVolumeBlockStorageRequestBody{
        Data: operations.UpdateVolumeBlockStorageData{
            Type: operations.UpdateVolumeBlockStorageTypeVolumes,
            Attributes: operations.UpdateVolumeBlockStorageAttributes{
                SizeInGb: 873925,
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `id`                                                                                                             | `string`                                                                                                         | :heavy_check_mark:                                                                                               | Volume ID                                                                                                        |
| `requestBody`                                                                                                    | [operations.UpdateVolumeBlockStorageRequestBody](../../models/operations/updatevolumeblockstoragerequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateVolumeResponse](../../models/operations/updatevolumeresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~MountVolume~~

Mounts a volume by adding the client to an allowed list

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="mount-volume" method="post" path="/storage/volumes/{id}/mount" -->
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

    res, err := s.BlockStorage.MountVolume(ctx, "<id>", operations.MountVolumeBlockStorageRequestBody{
        Data: operations.MountVolumeBlockStorageData{
            Type: operations.MountVolumeBlockStorageTypeVolumes,
            Attributes: operations.MountVolumeBlockStorageAttributes{
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
| `id`                                                                                                           | `string`                                                                                                       | :heavy_check_mark:                                                                                             | Volume ID                                                                                                      |
| `requestBody`                                                                                                  | [operations.MountVolumeBlockStorageRequestBody](../../models/operations/mountvolumeblockstoragerequestbody.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.MountVolumeResponse](../../models/operations/mountvolumeresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## MapVolume

Maps a high performance volume to a server over NVMe-TCP. Teams with multi-host mapping enabled can map a volume to several servers, once per server; otherwise a volume maps to one server. Every mapped server writes to the same NVMe namespace, so use a cluster-aware filesystem when more than one server mounts the volume.

### Example Usage

<!-- UsageSnippet language="go" operationID="map-volume" method="post" path="/storage/volumes/{id}/map" -->
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

    res, err := s.BlockStorage.MapVolume(ctx, "<id>", operations.MapVolumeBlockStorageRequestBody{
        Data: operations.MapVolumeBlockStorageData{
            Type: operations.MapVolumeBlockStorageTypeVolumes,
            Attributes: operations.MapVolumeBlockStorageAttributes{
                ServerID: "sv_abcd1234",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `id`                                                                                                       | `string`                                                                                                   | :heavy_check_mark:                                                                                         | Volume ID                                                                                                  |
| `requestBody`                                                                                              | [operations.MapVolumeBlockStorageRequestBody](../../models/operations/mapvolumeblockstoragerequestbody.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.MapVolumeResponse](../../models/operations/mapvolumeresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UnmapVolume

Unmaps a high performance volume from one of the servers it is mapped to.

### Example Usage

<!-- UsageSnippet language="go" operationID="unmap-volume" method="post" path="/storage/volumes/{id}/unmap" -->
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

    res, err := s.BlockStorage.UnmapVolume(ctx, "<id>", &operations.UnmapVolumeBlockStorageRequestBody{
        Data: &operations.UnmapVolumeBlockStorageData{
            Attributes: &operations.UnmapVolumeBlockStorageAttributes{
                ServerID: latitudeshgosdk.Pointer("sv_abcd1234"),
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

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |
| `id`                                                                                                            | `string`                                                                                                        | :heavy_check_mark:                                                                                              | Volume ID                                                                                                       |
| `requestBody`                                                                                                   | [*operations.UnmapVolumeBlockStorageRequestBody](../../models/operations/unmapvolumeblockstoragerequestbody.md) | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `opts`                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                        | :heavy_minus_sign:                                                                                              | The options for this request.                                                                                   |

### Response

**[*operations.UnmapVolumeResponse](../../models/operations/unmapvolumeresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |