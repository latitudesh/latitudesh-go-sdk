# ObjectStorage

## Overview

### Available Operations

* [GetStorageBuckets](#getstoragebuckets) - List object storages
* [PostStorageBuckets](#poststoragebuckets) - Create object storage
* [GetStorageBucket](#getstoragebucket) - Retrieve object storage
* [DeleteStorageBuckets](#deletestoragebuckets) - Delete object storage

## GetStorageBuckets

Lists all object storages from a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-buckets" method="get" path="/storage/buckets" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageBuckets(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ObjectStorages != nil {
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

**[*operations.GetStorageBucketsResponse](../../models/operations/getstoragebucketsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## PostStorageBuckets

Creates a new object storage bucket for a project.

### Example Usage: Create

<!-- UsageSnippet language="go" operationID="post-storage-buckets" method="post" path="/storage/buckets" example="Create" -->
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

    res, err := s.ObjectStorage.PostStorageBuckets(ctx, operations.PostStorageBucketsRequestBody{
        Data: operations.PostStorageBucketsData{
            Type: operations.PostStorageBucketsTypeObjects,
            Attributes: operations.PostStorageBucketsAttributes{
                Project: "proj_6059EqYkOQj8p",
                Name: "my-bucket",
                Region: "DAL",
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
### Example Usage: CreateScoped

<!-- UsageSnippet language="go" operationID="post-storage-buckets" method="post" path="/storage/buckets" example="CreateScoped" -->
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

    res, err := s.ObjectStorage.PostStorageBuckets(ctx, operations.PostStorageBucketsRequestBody{
        Data: operations.PostStorageBucketsData{
            Type: operations.PostStorageBucketsTypeObjects,
            Attributes: operations.PostStorageBucketsAttributes{
                Project: "proj_6059EqYkOQj8p",
                Name: "customer-bucket",
                Region: "DAL",
                Scoped: latitudeshgosdk.Pointer(true),
                Customer: latitudeshgosdk.Pointer("acme-corp"),
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

<!-- UsageSnippet language="go" operationID="post-storage-buckets" method="post" path="/storage/buckets" example="Created" -->
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

    res, err := s.ObjectStorage.PostStorageBuckets(ctx, operations.PostStorageBucketsRequestBody{
        Data: operations.PostStorageBucketsData{
            Type: operations.PostStorageBucketsTypeObjects,
            Attributes: operations.PostStorageBucketsAttributes{
                Project: "<value>",
                Name: "<value>",
                Region: "<value>",
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
### Example Usage: FeatureNotEnabled

<!-- UsageSnippet language="go" operationID="post-storage-buckets" method="post" path="/storage/buckets" example="FeatureNotEnabled" -->
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

    res, err := s.ObjectStorage.PostStorageBuckets(ctx, operations.PostStorageBucketsRequestBody{
        Data: operations.PostStorageBucketsData{
            Type: operations.PostStorageBucketsTypeObjects,
            Attributes: operations.PostStorageBucketsAttributes{
                Project: "<value>",
                Name: "<value>",
                Region: "<value>",
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
### Example Usage: InsufficientPermissions

<!-- UsageSnippet language="go" operationID="post-storage-buckets" method="post" path="/storage/buckets" example="InsufficientPermissions" -->
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

    res, err := s.ObjectStorage.PostStorageBuckets(ctx, operations.PostStorageBucketsRequestBody{
        Data: operations.PostStorageBucketsData{
            Type: operations.PostStorageBucketsTypeObjects,
            Attributes: operations.PostStorageBucketsAttributes{
                Project: "<value>",
                Name: "<value>",
                Region: "<value>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PostStorageBucketsRequestBody](../../models/operations/poststoragebucketsrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PostStorageBucketsResponse](../../models/operations/poststoragebucketsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409, 422       | application/vnd.api+json |
| components.ErrorObject   | 500                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetStorageBucket

Shows details of a specific object storage.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-bucket" method="get" path="/storage/buckets/{id}" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageBucket(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The object storage ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageBucketResponse](../../models/operations/getstoragebucketresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteStorageBuckets

Allows you to remove an object storage from a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-storage-buckets" method="delete" path="/storage/buckets/{id}" -->
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

    res, err := s.ObjectStorage.DeleteStorageBuckets(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The object storage ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteStorageBucketsResponse](../../models/operations/deletestoragebucketsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409            | application/vnd.api+json |
| components.ErrorObject   | 500                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |