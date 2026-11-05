# Inference

## Overview

### Available Operations

* [GetInferenceModels](#getinferencemodels) - List available inference models
* [GetInferenceAPIKeys](#getinferenceapikeys) - List inference API keys
* [PostInferenceAPIKey](#postinferenceapikey) - Create an inference API key
* [DeleteInferenceAPIKey](#deleteinferenceapikey) - Revoke an inference API key
* [GetInferenceUsage](#getinferenceusage) - Get inference usage for a project

## GetInferenceModels

Lists the inference models available through the inference gateway. The id is the canonical model ID — send it verbatim as the "model" value in inference requests. Only enabled models are returned; availability is global unless a model is explicitly scoped by region.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-inference-models" method="get" path="/inference/models" example="Success" -->
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

    res, err := s.Inference.GetInferenceModels(ctx, "proj_AoW6vRnwkvLn0")
    if err != nil {
        log.Fatal(err)
    }
    if res.InferenceModels != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | Project ID or Slug                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInferenceModelsResponse](../../models/operations/getinferencemodelsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403, 404            | application/vnd.api+json |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetInferenceAPIKeys

Lists the project's active inference API keys. The full token is never returned here — only its last characters.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-inference-api-keys" method="get" path="/inference/api_keys" example="Success" -->
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

    res, err := s.Inference.GetInferenceAPIKeys(ctx, "proj_AoW6vRnwkvLn0")
    if err != nil {
        log.Fatal(err)
    }
    if res.InferenceAPIKeys != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | Project ID or Slug                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInferenceAPIKeysResponse](../../models/operations/getinferenceapikeysresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403, 404            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## PostInferenceAPIKey

Creates an inference API key for the project. The full token is returned ONLY in this response — store it securely.
All active keys for the project can authenticate at the inference control plane. Synchronization is asynchronous: sync_status confirms that the refresh receiver observed the key; it does not acknowledge every region.


### Example Usage

<!-- UsageSnippet language="go" operationID="post-inference-api-key" method="post" path="/inference/api_keys" example="Created" -->
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

    res, err := s.Inference.PostInferenceAPIKey(ctx, components.CreateInferenceAPIKey{
        Data: components.CreateInferenceAPIKeyData{
            Type: components.CreateInferenceAPIKeyTypeInferenceAPIKeys,
            Attributes: components.CreateInferenceAPIKeyAttributes{
                ProjectID: "proj_AoW6vRnwkvLn0",
                Name: "Production inference key",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.CreateInferenceAPIKey](../../models/components/createinferenceapikey.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostInferenceAPIKeyResponse](../../models/operations/postinferenceapikeyresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteInferenceAPIKey

Revokes the inference API key and immediately removes it from active listings while retaining its audit record.
Control-plane removal is asynchronous. Other active keys remain valid. The key stops authorizing new requests after the control plane refreshes or its identity snapshot expires. Revoking an already revoked key succeeds.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-inference-api-key" method="delete" path="/inference/api_keys/{api_key_id}" example="Success" -->
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

    res, err := s.Inference.DeleteInferenceAPIKey(ctx, "proj_AoW6vRnwkvLn0", "aitok_6VE1Wd37dXnZJ")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | Project ID or Slug                                       |
| `apiKeyID`                                               | `string`                                                 | :heavy_check_mark:                                       | Inference API key ID                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteInferenceAPIKeyResponse](../../models/operations/deleteinferenceapikeyresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetInferenceUsage

Returns the project's inference requests, tokens and estimated cost for the window, per model and as a time series. The estimate uses current model prices.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-inference-usage" method="get" path="/inference/usage" example="Success" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/types"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.Inference.GetInferenceUsage(ctx, "proj_AoW6vRnwkvLn0", types.MustTimeFromString("2026-10-01T00:00:00Z"), types.MustTimeFromString("2026-10-02T00:00:00Z"), operations.GranularitySixh)
    if err != nil {
        log.Fatal(err)
    }
    if res.InferenceUsage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `projectID`                                                                           | `string`                                                                              | :heavy_check_mark:                                                                    | Project ID or Slug                                                                    |
| `from`                                                                                | [time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_check_mark:                                                                    | Start of the window (inclusive), truncated to whole seconds                           |
| `to`                                                                                  | [time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_check_mark:                                                                    | End of the window (exclusive), truncated to whole seconds; at most 31 days after from |
| `granularity`                                                                         | [operations.Granularity](../../models/operations/granularity.md)                      | :heavy_check_mark:                                                                    | Size of each series bucket                                                            |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*operations.GetInferenceUsageResponse](../../models/operations/getinferenceusageresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |