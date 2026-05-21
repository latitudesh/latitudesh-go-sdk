# ObjectStorage

## Overview

### Available Operations

* [GetStorageObjects](#getstorageobjects) - List object storages
* [PostStorageObjects](#poststorageobjects) - Create object storage
* [GetStorageObject](#getstorageobject) - Retrieve object storage
* [DeleteStorageObjects](#deletestorageobjects) - Delete object storage

## GetStorageObjects

Lists all object storages from a team.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-objects" method="get" path="/storage/objects" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageObjects(ctx, nil)
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

**[*operations.GetStorageObjectsResponse](../../models/operations/getstorageobjectsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## PostStorageObjects

Creates a new object storage bucket for a project.

### Example Usage: Create

<!-- UsageSnippet language="go" operationID="post-storage-objects" method="post" path="/storage/objects" example="Create" -->
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

    res, err := s.ObjectStorage.PostStorageObjects(ctx, operations.PostStorageObjectsRequestBody{
        Data: operations.PostStorageObjectsData{
            Type: operations.PostStorageObjectsTypeObjects,
            Attributes: operations.PostStorageObjectsAttributes{
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

<!-- UsageSnippet language="go" operationID="post-storage-objects" method="post" path="/storage/objects" example="CreateScoped" -->
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

    res, err := s.ObjectStorage.PostStorageObjects(ctx, operations.PostStorageObjectsRequestBody{
        Data: operations.PostStorageObjectsData{
            Type: operations.PostStorageObjectsTypeObjects,
            Attributes: operations.PostStorageObjectsAttributes{
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

<!-- UsageSnippet language="go" operationID="post-storage-objects" method="post" path="/storage/objects" example="Created" -->
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

    res, err := s.ObjectStorage.PostStorageObjects(ctx, operations.PostStorageObjectsRequestBody{
        Data: operations.PostStorageObjectsData{
            Type: operations.PostStorageObjectsTypeObjects,
            Attributes: operations.PostStorageObjectsAttributes{
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

<!-- UsageSnippet language="go" operationID="post-storage-objects" method="post" path="/storage/objects" example="FeatureNotEnabled" -->
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

    res, err := s.ObjectStorage.PostStorageObjects(ctx, operations.PostStorageObjectsRequestBody{
        Data: operations.PostStorageObjectsData{
            Type: operations.PostStorageObjectsTypeObjects,
            Attributes: operations.PostStorageObjectsAttributes{
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

<!-- UsageSnippet language="go" operationID="post-storage-objects" method="post" path="/storage/objects" example="InsufficientPermissions" -->
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

    res, err := s.ObjectStorage.PostStorageObjects(ctx, operations.PostStorageObjectsRequestBody{
        Data: operations.PostStorageObjectsData{
            Type: operations.PostStorageObjectsTypeObjects,
            Attributes: operations.PostStorageObjectsAttributes{
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
| `request`                                                                                            | [operations.PostStorageObjectsRequestBody](../../models/operations/poststorageobjectsrequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PostStorageObjectsResponse](../../models/operations/poststorageobjectsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409, 422       | application/vnd.api+json |
| components.ErrorObject   | 500                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetStorageObject

Shows details of a specific object storage.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-object" method="get" path="/storage/objects/{id}" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageObject(ctx, "<id>")
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

**[*operations.GetStorageObjectResponse](../../models/operations/getstorageobjectresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteStorageObjects

Allows you to remove an object storage from a project.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-storage-objects" method="delete" path="/storage/objects/{id}" -->
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

    res, err := s.ObjectStorage.DeleteStorageObjects(ctx, "<id>")
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

**[*operations.DeleteStorageObjectsResponse](../../models/operations/deletestorageobjectsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.ErrorObject   | 500                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |