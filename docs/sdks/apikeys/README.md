# ApiKeys

## Overview

### Available Operations

* [List](#list) - List API keys
* [Create](#create) - Create API key
* [Update](#update) - Rotate API key
* [Delete](#delete) - Delete API key
* [UpdateAPIKey](#updateapikey) - Update API key settings

## List

Returns a list of all API keys.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-api-keys" method="get" path="/auth/api_keys" example="Success" -->
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
    if res.APIKeys != nil {
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


### Example Usage: API Key Created

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" example="API Key Created" -->
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
### Example Usage: Bad Request

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" example="Bad Request" -->
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
                Name: latitudeshgosdk.Pointer("foobar"),
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

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" example="Created" -->
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
### Example Usage: IPRestrictedKey

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" example="IPRestrictedKey" -->
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
                Name: latitudeshgosdk.Pointer("Office Network Key"),
                AllowedIps: []string{
                    "192.168.1.100",
                    "10.0.0.0/24",
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
### Example Usage: ReadOnlyIPRestrictedKey

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" example="ReadOnlyIPRestrictedKey" -->
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
                Name: latitudeshgosdk.Pointer("Monitoring Dashboard Key"),
                ReadOnly: latitudeshgosdk.Pointer(true),
                AllowedIps: []string{
                    "203.0.113.50",
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
### Example Usage: ReadOnlyKey

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" example="ReadOnlyKey" -->
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
                Name: latitudeshgosdk.Pointer("Read-Only Analytics Key"),
                ReadOnly: latitudeshgosdk.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="post-api-key" method="post" path="/auth/api_keys" example="Unprocessable Entity" -->
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
            Attributes: &components.CreateAPIKeyAttributes{},
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

Rotate an existing API Key, generating a new token. This invalidates the previous key.
Use PATCH to update settings without rotating the token.


### Example Usage

<!-- UsageSnippet language="go" operationID="rotate-api-key" method="put" path="/auth/api_keys/{api_key_id}" example="Success" -->
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

    res, err := s.APIKeys.Update(ctx, "tok_x1ZJrdx5qg4LV", components.UpdateAPIKey{
        Data: &components.UpdateAPIKeyData{
            ID: latitudeshgosdk.Pointer("tok_x1ZJrdx5qg4LV"),
            Type: components.UpdateAPIKeyTypeAPIKeys,
            Attributes: &components.UpdateAPIKeyAttributes{
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
| `apiKeyID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `updateAPIKey`                                                     | [components.UpdateAPIKey](../../models/components/updateapikey.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.RotateAPIKeyResponse](../../models/operations/rotateapikeyresponse.md), error**

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

    res, err := s.APIKeys.Delete(ctx, "tok_lQraYDPeOpjwW")
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

## UpdateAPIKey

Update API Key settings (name, read_only, allowed_ips) without rotating the token.
Use PUT to rotate the token.


### Example Usage: API Key Not Found

<!-- UsageSnippet language="go" operationID="update-api-key" method="patch" path="/auth/api_keys/{api_key_id}" example="API Key Not Found" -->
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

    res, err := s.APIKeys.UpdateAPIKey(ctx, "invalid-api-key", components.UpdateAPIKey{
        Data: &components.UpdateAPIKeyData{
            ID: latitudeshgosdk.Pointer("invalid-api-key"),
            Type: components.UpdateAPIKeyTypeAPIKeys,
            Attributes: &components.UpdateAPIKeyAttributes{
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
### Example Usage: API Key Updated

<!-- UsageSnippet language="go" operationID="update-api-key" method="patch" path="/auth/api_keys/{api_key_id}" example="API Key Updated" -->
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

    res, err := s.APIKeys.UpdateAPIKey(ctx, "tok_pRMLydp0dQKr1", components.UpdateAPIKey{
        Data: &components.UpdateAPIKeyData{
            ID: latitudeshgosdk.Pointer("tok_pRMLydp0dQKr1"),
            Type: components.UpdateAPIKeyTypeAPIKeys,
            Attributes: &components.UpdateAPIKeyAttributes{
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
### Example Usage: Bad Request

<!-- UsageSnippet language="go" operationID="update-api-key" method="patch" path="/auth/api_keys/{api_key_id}" example="Bad Request" -->
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

    res, err := s.APIKeys.UpdateAPIKey(ctx, "tok_w5AEmq7XDBkWX", components.UpdateAPIKey{
        Data: &components.UpdateAPIKeyData{
            ID: latitudeshgosdk.Pointer("tok_w5AEmq7XDBkWX"),
            Type: components.UpdateAPIKeyTypeAPIKeys,
            Attributes: &components.UpdateAPIKeyAttributes{
                Name: latitudeshgosdk.Pointer("Name of the API Key"),
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

<!-- UsageSnippet language="go" operationID="update-api-key" method="patch" path="/auth/api_keys/{api_key_id}" example="Success" -->
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

    res, err := s.APIKeys.UpdateAPIKey(ctx, "tok_zlkg1DegdvZE5", components.UpdateAPIKey{
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
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: Success - Update name without rotating token

<!-- UsageSnippet language="go" operationID="update-api-key" method="patch" path="/auth/api_keys/{api_key_id}" example="Success - Update name without rotating token" -->
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

    res, err := s.APIKeys.UpdateAPIKey(ctx, "tok_lpbV0DgRq4AWz", components.UpdateAPIKey{
        Data: &components.UpdateAPIKeyData{
            ID: latitudeshgosdk.Pointer("tok_lpbV0DgRq4AWz"),
            Type: components.UpdateAPIKeyTypeAPIKeys,
            Attributes: &components.UpdateAPIKeyAttributes{
                Name: latitudeshgosdk.Pointer("Updated Name"),
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

**[*operations.UpdateAPIKeyResponse](../../models/operations/updateapikeyresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |