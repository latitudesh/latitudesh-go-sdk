# ApiKeys

## Overview

### Available Operations

* [List](#list) - List API keys
* [Create](#create) - Create API key
* [Update](#update) - Rotate API key
* [Delete](#delete) - Delete API key
* [PatchAPIKey](#patchapikey) - Update API key settings

## List

Returns a list of all API keys.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-api-keys" method="get" path="/auth/api_keys" -->
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

    res, err := s.APIKeys.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAPIKeysResponse](../../models/operations/getapikeysresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Create a new API Key that is tied to the current user account. The created API key is only listed ONCE upon creation. It can however be regenerated or deleted.


### Example Usage

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.APIKeys.Create(ctx, components.CreateAPIKey{
        Data: &components.Data{
            Type: components.CreateAPIKeyTypeAPIKeys,
            Attributes: &components.CreateAPIKeyAttributes{
                Name: latitudeshgosdk.Pointer("App Token"),
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.CreateAPIKey](../../models/components/createapikey.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.PostAPIKeyResponse](../../models/operations/postapikeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Rotate (regenerate) an API key's token and optionally update its settings. This generates a new token and invalidates the previous one. To update settings without rotating the token, use the PATCH endpoint instead.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-api-key" method="put" path="/auth/api_keys/{api_key_id}" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.APIKeys.Update(ctx, "tok_zlkg1DegdvZE5", components.UpdateAPIKey{
        Data: &components.UpdateAPIKeyData{
            ID: latitudeshgosdk.Pointer("tok_zlkg1DegdvZE5"),
            Type: components.UpdateAPIKeyTypeAPIKeys,
            Attributes: &components.UpdateAPIKeyAttributes{
                Name: latitudeshgosdk.Pointer("App Token"),
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `apiKeyID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `updateAPIKey`                                                     | [components.UpdateAPIKey](../../models/components/updateapikey.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.UpdateAPIKeyResponse](../../models/operations/updateapikeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Delete an existing API Key. Once deleted, the API Key can no longer be used to access the API.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-api-key" method="delete" path="/auth/api_keys/{api_key_id}" -->
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

    res, err := s.APIKeys.Delete(ctx, "tok_x1ZJrdx5qg4LV")
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
| `apiKeyID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAPIKeyResponse](../../models/operations/deleteapikeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## PatchAPIKey

Update an API key's settings (name, read_only, allowed_ips) without regenerating the token. To rotate the token, use the PUT endpoint instead.


### Example Usage

<!-- UsageSnippet language="go" operationID="patch-api-key" method="patch" path="/auth/api_keys/{api_key_id}" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.APIKeys.PatchAPIKey(ctx, "tok_zlkg1DegdvZE5", components.UpdateAPIKey{
        Data: &components.UpdateAPIKeyData{
            ID: latitudeshgosdk.Pointer("tok_zlkg1DegdvZE5"),
            Type: components.UpdateAPIKeyTypeAPIKeys,
            Attributes: &components.UpdateAPIKeyAttributes{
                Name: latitudeshgosdk.Pointer("Updated Token Name"),
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `apiKeyID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `updateAPIKey`                                                     | [components.UpdateAPIKey](../../models/components/updateapikey.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.PatchAPIKeyResponse](../../models/operations/patchapikeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |