# SandboxesPreview

## Overview

Preview. Available to teams with the `sandboxes_api` feature flag. The shape of these endpoints may change before general availability.

### Available Operations

* [GetSandboxes](#getsandboxes) - List sandboxes
* [CreateSandbox](#createsandbox) - Create sandbox
* [GetSandbox](#getsandbox) - Retrieve sandbox
* [DestroySandbox](#destroysandbox) - Delete sandbox
* [ExecSandbox](#execsandbox) - Run a command in a sandbox

## GetSandboxes

**Preview.** Available to teams with the `sandboxes_api` feature flag. The shape of this
endpoint may change before general availability.

List the sandboxes of a project, up to 100 per page.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-sandboxes" method="get" path="/sandboxes" example="Success" -->
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

    res, err := s.SandboxesPreview.GetSandboxes(ctx, "proj_6059EqYkOQj8p", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Sandboxes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `projectID`                                                   | `string`                                                      | :heavy_check_mark:                                            | ID of the project whose sandboxes are listed                  |
| `pageToken`                                                   | `*string`                                                     | :heavy_minus_sign:                                            | Token returned in `meta.next_page_token` by the previous page |
| `opts`                                                        | [][operations.Option](../../models/operations/option.md)      | :heavy_minus_sign:                                            | The options for this request.                                 |

### Response

**[*operations.GetSandboxesResponse](../../models/operations/getsandboxesresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateSandbox

**Preview.** Available to teams with the `sandboxes_api` feature flag. The shape of this
endpoint may change before general availability.

Create a sandbox in a project.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-sandbox" method="post" path="/sandboxes" example="Created" -->
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

    res, err := s.SandboxesPreview.CreateSandbox(ctx, operations.CreateSandboxSandboxesPreviewRequestBody{
        Data: &operations.CreateSandboxSandboxesPreviewData{
            Type: operations.CreateSandboxSandboxesPreviewTypeSandboxes.ToPointer(),
            Attributes: &operations.CreateSandboxSandboxesPreviewAttributes{
                ProjectID: "proj_6059EqYkOQj8p",
                Name: latitudeshgosdk.Pointer("sbx-demo"),
                Image: latitudeshgosdk.Pointer("python:3.12"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sandbox != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.CreateSandboxSandboxesPreviewRequestBody](../../models/operations/createsandboxsandboxespreviewrequestbody.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.CreateSandboxResponse](../../models/operations/createsandboxresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetSandbox

**Preview.** Available to teams with the `sandboxes_api` feature flag.

Retrieve a single sandbox.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-sandbox" method="get" path="/sandboxes/{sandbox_id}" example="Success" -->
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

    res, err := s.SandboxesPreview.GetSandbox(ctx, "sbx-demo")
    if err != nil {
        log.Fatal(err)
    }
    if res.Sandbox != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `sandboxID`                                              | `string`                                                 | :heavy_check_mark:                                       | Sandbox name                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSandboxResponse](../../models/operations/getsandboxresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DestroySandbox

**Preview.** Available to teams with the `sandboxes_api` feature flag.

Delete a sandbox. Deletion completes in the background.


### Example Usage

<!-- UsageSnippet language="go" operationID="destroy-sandbox" method="delete" path="/sandboxes/{sandbox_id}" -->
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

    res, err := s.SandboxesPreview.DestroySandbox(ctx, "<id>")
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
| `sandboxID`                                              | `string`                                                 | :heavy_check_mark:                                       | Sandbox name                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DestroySandboxResponse](../../models/operations/destroysandboxresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ExecSandbox

**Preview.** Available to teams with the `sandboxes_api` feature flag. The shape of this
endpoint may change before general availability.

Run a command in a sandbox and wait for its output.


### Example Usage

<!-- UsageSnippet language="go" operationID="exec-sandbox" method="post" path="/sandboxes/{sandbox_id}/exec" example="Success" -->
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

    res, err := s.SandboxesPreview.ExecSandbox(ctx, "sbx-demo", operations.ExecSandboxSandboxesPreviewRequestBody{
        Data: &operations.ExecSandboxSandboxesPreviewData{
            Attributes: &operations.ExecSandboxSandboxesPreviewAttributes{
                Command: []string{
                    "python",
                    "-V",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SandboxExec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `sandboxID`                                                                                                            | `string`                                                                                                               | :heavy_check_mark:                                                                                                     | Sandbox name                                                                                                           |
| `requestBody`                                                                                                          | [operations.ExecSandboxSandboxesPreviewRequestBody](../../models/operations/execsandboxsandboxespreviewrequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ExecSandboxResponse](../../models/operations/execsandboxresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |