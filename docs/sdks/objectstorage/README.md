# ObjectStorage

## Overview

### Available Operations

* [GetStorageUsage](#getstorageusage) - List storage usage
* [PostStorageAccessKeys](#poststorageaccesskeys) - Create access key
* [GetStorageAccessKeys](#getstorageaccesskeys) - List access keys
* [DeleteStorageAccessKeysUsername](#deletestorageaccesskeysusername) - Delete access key
* [GetStorageBucketAccessKeys](#getstoragebucketaccesskeys) - List bucket access keys
* [GetStorageBuckets](#getstoragebuckets) - List buckets
* [PostStorageBuckets](#poststoragebuckets) - Create bucket
* [GetStorageBucket](#getstoragebucket) - Retrieve bucket
* [DeleteStorageBuckets](#deletestoragebuckets) - Delete bucket
* [GetStorageBucketLifecycleRules](#getstoragebucketlifecyclerules) - List lifecycle rules
* [PostStorageBucketLifecycleRules](#poststoragebucketlifecyclerules) - Create lifecycle rule
* [GetStorageBucketLifecycleRule](#getstoragebucketlifecyclerule) - Retrieve lifecycle rule
* [DeleteStorageBucketLifecycleRule](#deletestoragebucketlifecyclerule) - Delete lifecycle rule
* [PatchStorageBucketLifecycleRule](#patchstoragebucketlifecyclerule) - Update lifecycle rule
* [GetStorageBucketMetrics](#getstoragebucketmetrics) - Retrieve bucket metrics

## GetStorageUsage

Returns daily object storage usage for a project. Each row reports the canonical usage in bytes for a single storage on a given day, plus the provider-reported raw value.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-usage" method="get" path="/storage/usage" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageUsage(ctx, "proj_5AEmq7wMqBkWX", latitudeshgosdk.Pointer("bkt_6VE1Wd37dXnZJ"), nil, nil)
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
| `filterProject`                                                        | `string`                                                               | :heavy_check_mark:                                                     | Project ID or Slug                                                     |
| `filterStorageID`                                                      | `*string`                                                              | :heavy_minus_sign:                                                     | Restrict the result to a single storage. Accepts the storage/bucket ID |
| `filterStartDate`                                                      | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to yesterday                                                  |
| `filterEndDate`                                                        | [*types.Date](../../types/date.md)                                     | :heavy_minus_sign:                                                     | Defaults to today; clamped to today when a future date is given        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.GetStorageUsageResponse](../../models/operations/getstorageusageresponse.md), error**

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

## GetStorageBucketLifecycleRules

Lists all lifecycle rules for a specific object storage bucket.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-bucket-lifecycle-rules" method="get" path="/storage/buckets/{bucket_id}/lifecycle_rules" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageBucketLifecycleRules(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LifecycleRules != nil {
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

**[*operations.GetStorageBucketLifecycleRulesResponse](../../models/operations/getstoragebucketlifecyclerulesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## PostStorageBucketLifecycleRules

Creates a new lifecycle rule for an object storage bucket. Lifecycle rules automate object expiration based on age.

### Example Usage: Create

<!-- UsageSnippet language="go" operationID="post-storage-bucket-lifecycle-rules" method="post" path="/storage/buckets/{bucket_id}/lifecycle_rules" example="Create" -->
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

    res, err := s.ObjectStorage.PostStorageBucketLifecycleRules(ctx, "<id>", operations.PostStorageBucketLifecycleRulesRequestBody{
        Data: operations.PostStorageBucketLifecycleRulesData{
            Type: operations.PostStorageBucketLifecycleRulesTypeLifecycleRules,
            Attributes: operations.PostStorageBucketLifecycleRulesAttributes{
                Name: "delete-old-logs",
                Prefix: latitudeshgosdk.Pointer("logs/"),
                ExpirationDays: 30,
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
### Example Usage: CreateWithAllOptions

<!-- UsageSnippet language="go" operationID="post-storage-bucket-lifecycle-rules" method="post" path="/storage/buckets/{bucket_id}/lifecycle_rules" example="CreateWithAllOptions" -->
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

    res, err := s.ObjectStorage.PostStorageBucketLifecycleRules(ctx, "<id>", operations.PostStorageBucketLifecycleRulesRequestBody{
        Data: operations.PostStorageBucketLifecycleRulesData{
            Type: operations.PostStorageBucketLifecycleRulesTypeLifecycleRules,
            Attributes: operations.PostStorageBucketLifecycleRulesAttributes{
                Name: "full-cleanup-rule",
                Prefix: latitudeshgosdk.Pointer("temp/"),
                ExpirationDays: 30,
                NoncurrentDays: latitudeshgosdk.Pointer[int64](14),
                AbortMpuDaysAfterInitiation: latitudeshgosdk.Pointer[int64](7),
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

<!-- UsageSnippet language="go" operationID="post-storage-bucket-lifecycle-rules" method="post" path="/storage/buckets/{bucket_id}/lifecycle_rules" example="Created" -->
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

    res, err := s.ObjectStorage.PostStorageBucketLifecycleRules(ctx, "<id>", operations.PostStorageBucketLifecycleRulesRequestBody{
        Data: operations.PostStorageBucketLifecycleRulesData{
            Type: operations.PostStorageBucketLifecycleRulesTypeLifecycleRules,
            Attributes: operations.PostStorageBucketLifecycleRulesAttributes{
                Name: "<value>",
                ExpirationDays: 69486,
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

<!-- UsageSnippet language="go" operationID="post-storage-bucket-lifecycle-rules" method="post" path="/storage/buckets/{bucket_id}/lifecycle_rules" example="FeatureNotEnabled" -->
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

    res, err := s.ObjectStorage.PostStorageBucketLifecycleRules(ctx, "<id>", operations.PostStorageBucketLifecycleRulesRequestBody{
        Data: operations.PostStorageBucketLifecycleRulesData{
            Type: operations.PostStorageBucketLifecycleRulesTypeLifecycleRules,
            Attributes: operations.PostStorageBucketLifecycleRulesAttributes{
                Name: "<value>",
                ExpirationDays: 69486,
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

<!-- UsageSnippet language="go" operationID="post-storage-bucket-lifecycle-rules" method="post" path="/storage/buckets/{bucket_id}/lifecycle_rules" example="InsufficientPermissions" -->
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

    res, err := s.ObjectStorage.PostStorageBucketLifecycleRules(ctx, "<id>", operations.PostStorageBucketLifecycleRulesRequestBody{
        Data: operations.PostStorageBucketLifecycleRulesData{
            Type: operations.PostStorageBucketLifecycleRulesTypeLifecycleRules,
            Attributes: operations.PostStorageBucketLifecycleRulesAttributes{
                Name: "<value>",
                ExpirationDays: 69486,
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `bucketID`                                                                                                                     | `string`                                                                                                                       | :heavy_check_mark:                                                                                                             | The object storage bucket ID                                                                                                   |
| `requestBody`                                                                                                                  | [operations.PostStorageBucketLifecycleRulesRequestBody](../../models/operations/poststoragebucketlifecyclerulesrequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.PostStorageBucketLifecycleRulesResponse](../../models/operations/poststoragebucketlifecyclerulesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.ErrorObject   | 500                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetStorageBucketLifecycleRule

Retrieves details of a specific lifecycle rule.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-storage-bucket-lifecycle-rule" method="get" path="/storage/buckets/{bucket_id}/lifecycle_rules/{id}" example="Success" -->
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

    res, err := s.ObjectStorage.GetStorageBucketLifecycleRule(ctx, "<id>", "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The lifecycle rule ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetStorageBucketLifecycleRuleResponse](../../models/operations/getstoragebucketlifecycleruleresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteStorageBucketLifecycleRule

Deletes a lifecycle rule from an object storage bucket.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-storage-bucket-lifecycle-rule" method="delete" path="/storage/buckets/{bucket_id}/lifecycle_rules/{id}" -->
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

    res, err := s.ObjectStorage.DeleteStorageBucketLifecycleRule(ctx, "<id>", "<id>")
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
| `bucketID`                                               | `string`                                                 | :heavy_check_mark:                                       | The object storage bucket ID                             |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The lifecycle rule ID                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteStorageBucketLifecycleRuleResponse](../../models/operations/deletestoragebucketlifecycleruleresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.ErrorObject   | 500                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## PatchStorageBucketLifecycleRule

Updates an existing lifecycle rule for an object storage bucket.

### Example Usage: FeatureNotEnabled

<!-- UsageSnippet language="go" operationID="patch-storage-bucket-lifecycle-rule" method="patch" path="/storage/buckets/{bucket_id}/lifecycle_rules/{id}" example="FeatureNotEnabled" -->
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

    res, err := s.ObjectStorage.PatchStorageBucketLifecycleRule(ctx, "<id>", "<id>", operations.PatchStorageBucketLifecycleRuleRequestBody{
        Data: operations.PatchStorageBucketLifecycleRuleData{
            Type: operations.PatchStorageBucketLifecycleRuleTypeLifecycleRules,
            Attributes: operations.PatchStorageBucketLifecycleRuleAttributes{
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
### Example Usage: InsufficientPermissions

<!-- UsageSnippet language="go" operationID="patch-storage-bucket-lifecycle-rule" method="patch" path="/storage/buckets/{bucket_id}/lifecycle_rules/{id}" example="InsufficientPermissions" -->
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

    res, err := s.ObjectStorage.PatchStorageBucketLifecycleRule(ctx, "<id>", "<id>", operations.PatchStorageBucketLifecycleRuleRequestBody{
        Data: operations.PatchStorageBucketLifecycleRuleData{
            Type: operations.PatchStorageBucketLifecycleRuleTypeLifecycleRules,
            Attributes: operations.PatchStorageBucketLifecycleRuleAttributes{
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="patch-storage-bucket-lifecycle-rule" method="patch" path="/storage/buckets/{bucket_id}/lifecycle_rules/{id}" example="Success" -->
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

    res, err := s.ObjectStorage.PatchStorageBucketLifecycleRule(ctx, "<id>", "<id>", operations.PatchStorageBucketLifecycleRuleRequestBody{
        Data: operations.PatchStorageBucketLifecycleRuleData{
            Type: operations.PatchStorageBucketLifecycleRuleTypeLifecycleRules,
            Attributes: operations.PatchStorageBucketLifecycleRuleAttributes{
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
### Example Usage: Update

<!-- UsageSnippet language="go" operationID="patch-storage-bucket-lifecycle-rule" method="patch" path="/storage/buckets/{bucket_id}/lifecycle_rules/{id}" example="Update" -->
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

    res, err := s.ObjectStorage.PatchStorageBucketLifecycleRule(ctx, "<id>", "<id>", operations.PatchStorageBucketLifecycleRuleRequestBody{
        Data: operations.PatchStorageBucketLifecycleRuleData{
            Type: operations.PatchStorageBucketLifecycleRuleTypeLifecycleRules,
            Attributes: operations.PatchStorageBucketLifecycleRuleAttributes{
                Name: "delete-old-logs",
                Enabled: latitudeshgosdk.Pointer(false),
                ExpirationDays: latitudeshgosdk.Pointer[int64](60),
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `bucketID`                                                                                                                     | `string`                                                                                                                       | :heavy_check_mark:                                                                                                             | The object storage bucket ID                                                                                                   |
| `id`                                                                                                                           | `string`                                                                                                                       | :heavy_check_mark:                                                                                                             | The lifecycle rule ID                                                                                                          |
| `requestBody`                                                                                                                  | [operations.PatchStorageBucketLifecycleRuleRequestBody](../../models/operations/patchstoragebucketlifecyclerulerequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.PatchStorageBucketLifecycleRuleResponse](../../models/operations/patchstoragebucketlifecycleruleresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
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