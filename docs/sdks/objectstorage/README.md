# ObjectStorage

## Overview

### Available Operations

* [IndexProjectStorageUsage](#indexprojectstorageusage) - List storage usage
* [PostStorageAccessKeys](#poststorageaccesskeys) - Create access key
* [GetStorageAccessKeys](#getstorageaccesskeys) - List access keys
* [DeleteStorageAccessKeysUsername](#deletestorageaccesskeysusername) - Delete access key
* [GetStorageBucketAccessKeys](#getstoragebucketaccesskeys) - List bucket access keys
* [GetStorageBuckets](#getstoragebuckets) - List buckets
* [PostStorageBuckets](#poststoragebuckets) - Create bucket
* [GetStorageBucket](#getstoragebucket) - Retrieve bucket
* [DeleteStorageBuckets](#deletestoragebuckets) - Delete bucket
* [GetStorageBucketMetrics](#getstoragebucketmetrics) - Retrieve bucket metrics

## IndexProjectStorageUsage

Returns daily object storage usage for the project. Each row reports the canonical usage in bytes for a single storage on a given day, plus the provider-reported raw value.


### Example Usage

<!-- UsageSnippet language="go" operationID="index-project-storage-usage" method="get" path="/projects/{project_id}/storage_usage" example="Success" -->
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

    res, err := s.ObjectStorage.IndexProjectStorageUsage(ctx, "proj_5AEmq7wMqBkWX", latitudeshgosdk.Pointer("bkt_6VE1Wd37dXnZJ"), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageUsage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `projectID`                                                            | `string`                                                               | :heavy_check_mark:                                                     | Project ID or Slug                                                     |
| `storageID`                                                            | `*string`                                                              | :heavy_minus_sign:                                                     | Restrict the result to a single storage. Accepts the storage/bucket ID |
| `startDate`                                                            | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to yesterday                                                  |
| `endDate`                                                              | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to today; clamped to today when a future date is given        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.IndexProjectStorageUsageResponse](../../models/operations/indexprojectstorageusageresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## PostStorageAccessKeys

Creates an object storage IAM access key for a project. The secret is returned only once, in this response, and cannot be retrieved again. The provider is selected by `storage_class`: `standard` provisions the key on Wasabi and `high_performance` provisions it on VAST.

### Example Usage: Created

<!-- UsageSnippet language="go" operationID="post-storage-access-keys" method="post" path="/storage/access_keys" example="Created" -->
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

    res, err := s.ObjectStorage.PostStorageAccessKeys(ctx, operations.PostStorageAccessKeysRequestBody{
        Data: operations.PostStorageAccessKeysData{
            Type: operations.PostStorageAccessKeysTypeAccessKeys,
            Attributes: operations.PostStorageAccessKeysAttributes{
                Project: "<value>",
                AccessKeyStorageClass: operations.AccessKeyStorageClassHighPerformance,
                Name: "<value>",
                AccessScope: operations.AccessScopeLimitedAccess,
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
### Example Usage: FullAccess

<!-- UsageSnippet language="go" operationID="post-storage-access-keys" method="post" path="/storage/access_keys" example="FullAccess" -->
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

    res, err := s.ObjectStorage.PostStorageAccessKeys(ctx, operations.PostStorageAccessKeysRequestBody{
        Data: operations.PostStorageAccessKeysData{
            Type: operations.PostStorageAccessKeysTypeAccessKeys,
            Attributes: operations.PostStorageAccessKeysAttributes{
                Project: "proj_6059EqYkOQj8p",
                AccessKeyStorageClass: operations.AccessKeyStorageClassStandard,
                Name: "my-access-key",
                AccessScope: operations.AccessScopeFullaccess,
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
### Example Usage: LimitedAccess

<!-- UsageSnippet language="go" operationID="post-storage-access-keys" method="post" path="/storage/access_keys" example="LimitedAccess" -->
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

    res, err := s.ObjectStorage.PostStorageAccessKeys(ctx, operations.PostStorageAccessKeysRequestBody{
        Data: operations.PostStorageAccessKeysData{
            Type: operations.PostStorageAccessKeysTypeAccessKeys,
            Attributes: operations.PostStorageAccessKeysAttributes{
                Project: "proj_6059EqYkOQj8p",
                AccessKeyStorageClass: operations.AccessKeyStorageClassStandard,
                Name: "my-limited-key",
                AccessScope: operations.AccessScopeLimitedAccess,
                Region: "DAL",
                BucketPermissions: []operations.BucketPermissions{
                    operations.BucketPermissions{
                        BucketID: "bucket_6VE1Wd37dXnZJ",
                        Permission: operations.PermissionReadonly,
                    },
                    operations.BucketPermissions{
                        BucketID: "bucket_7WF2Xe48eYoAK",
                        Permission: operations.PermissionRw,
                    },
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
### Example Usage: Unauthorized

<!-- UsageSnippet language="go" operationID="post-storage-access-keys" method="post" path="/storage/access_keys" example="Unauthorized" -->
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

    res, err := s.ObjectStorage.PostStorageAccessKeys(ctx, operations.PostStorageAccessKeysRequestBody{
        Data: operations.PostStorageAccessKeysData{
            Type: operations.PostStorageAccessKeysTypeAccessKeys,
            Attributes: operations.PostStorageAccessKeysAttributes{
                Project: "<value>",
                AccessKeyStorageClass: operations.AccessKeyStorageClassHighPerformance,
                Name: "<value>",
                AccessScope: operations.AccessScopeLimitedAccess,
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostStorageAccessKeysRequestBody](../../models/operations/poststorageaccesskeysrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PostStorageAccessKeysResponse](../../models/operations/poststorageaccesskeysresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetStorageAccessKeys

Lists object storage access keys for a project, grouped by storage class. Secrets are never returned by this endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-access-keys" method="get" path="/storage/access_keys" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageAccessKeys(ctx, "<value>")
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
| `project`                                                | `string`                                                 | :heavy_check_mark:                                       | Project ID or slug to list access keys for.              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageAccessKeysResponse](../../models/operations/getstorageaccesskeysresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteStorageAccessKeysUsername

Permanently deletes an object storage access key and its associated provider-side credentials. Deletion cannot be undone: it revokes the key's credentials and removes the key from the storage provider. For `standard` (Wasabi) keys, the IAM user is removed from each of the project's bucket policies, all of the user's access keys are revoked, and the IAM user is deleted. For `high_performance` (VAST) keys, the VAST user's S3 keys are revoked, its attached S3 policies are deleted, and the VAST user is removed.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-storage-access-keys-username" method="delete" path="/storage/access_keys/{username}/{storage_class}" -->
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

    res, err := s.ObjectStorage.DeleteStorageAccessKeysUsername(ctx, "Earline_Dooley27", operations.PathParamStorageClassHighPerformance, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `username`                                                                                                                                       | `string`                                                                                                                                         | :heavy_check_mark:                                                                                                                               | Name of the access key to delete.                                                                                                                |
| `storageClass`                                                                                                                                   | [operations.PathParamStorageClass](../../models/operations/pathparamstorageclass.md)                                                             | :heavy_check_mark:                                                                                                                               | Backend storage tier of the access key. `standard` targets Wasabi; `high_performance` targets VAST.                                              |
| `project`                                                                                                                                        | `string`                                                                                                                                         | :heavy_check_mark:                                                                                                                               | Project ID or slug the access key belongs to.                                                                                                    |
| `region`                                                                                                                                         | `*string`                                                                                                                                        | :heavy_minus_sign:                                                                                                                               | Region slug (e.g., `DAL`, `SAO2`). Required for `high_performance` (VAST) keys to select the VAST cluster; ignored for `standard` (Wasabi) keys. |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.DeleteStorageAccessKeysUsernameResponse](../../models/operations/deletestorageaccesskeysusernameresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetStorageBucketAccessKeys

Lists IAM access keys associated with an object storage bucket. Secrets are never returned by this endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-bucket-access-keys" method="get" path="/storage/buckets/{id}/access_keys" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageBucketAccessKeys(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The object storage (bucket) ID                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageBucketAccessKeysResponse](../../models/operations/getstoragebucketaccesskeysresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

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

## GetStorageBucketMetrics

Retrieves usage metrics for a specific object storage bucket, including storage consumption and estimated cost for the current billing period.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-bucket-metrics" method="get" path="/storage/buckets/{bucket_id}/metrics" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageBucketMetrics(ctx, "<id>")
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
| `bucketID`                                               | `string`                                                 | :heavy_check_mark:                                       | The object storage bucket ID                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageBucketMetricsResponse](../../models/operations/getstoragebucketmetricsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |