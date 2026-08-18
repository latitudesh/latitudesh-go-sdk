# BaselinesPreview

## Overview

Preview. Available to teams with the `baselines_api` feature flag. The shape of these endpoints may change before general availability.

### Available Operations

* [GetBaselines](#getbaselines) - List baselines
* [CreateBaseline](#createbaseline) - Create baseline
* [GetBaseline](#getbaseline) - Retrieve baseline
* [DestroyBaseline](#destroybaseline) - Delete baseline

## GetBaselines

**Preview.** Available to teams with the `baselines_api` feature flag. The shape of this
endpoint may change before general availability.

List all baselines in the team. A baseline records the configuration you expect your
servers to be delivered with — plan, SSH keys, user data, disk layout and BIOS settings.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-baselines" method="get" path="/baselines" example="Success" -->
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

    res, err := s.BaselinesPreview.GetBaselines(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Baselines != nil {
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

**[*operations.GetBaselinesResponse](../../models/operations/getbaselinesresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateBaseline

**Preview.** Available to teams with the `baselines_api` feature flag.

Create a baseline in the team. A baseline can target all servers, a custom set (when the
plan is not yet known), or one or more specific platforms. When it targets platforms, the
disk layout is validated by the same rules a deploy applies — against the smallest of the
selected platforms — so a baseline that saves here can be dispatched verbatim.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-baseline" method="post" path="/baselines" example="Created" -->
```go
package main

import(
	"context"
	"os"
	latitudeshgosdk "github.com/latitudesh/latitudesh-go-sdk"
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := latitudeshgosdk.New(
        latitudeshgosdk.WithSecurity(os.Getenv("LATITUDESH_BEARER")),
    )

    res, err := s.BaselinesPreview.CreateBaseline(ctx, operations.CreateBaselineBaselinesPreviewRequestBody{
        Data: &operations.CreateBaselineData{
            Type: operations.CreateBaselineTypeBaselines.ToPointer(),
            Attributes: &operations.CreateBaselineAttributes{
                Name: "web-fleet-v3",
                Description: latitudeshgosdk.Pointer("Standard build for the public web tier"),
                TargetType: operations.CreateBaselineTargetTypePlatforms,
                OperatingSystem: "ubuntu_22_04_x64_lts",
                Platforms: []string{
                    "g3-l40s-small-76",
                },
                SSHKeyIds: []string{
                    "ssh_RLYV8DZ2D5QoE",
                },
                UserDataID: latitudeshgosdk.Pointer("ud_r0MK4O4kDa95w"),
                DiskLayout: []components.BaselineDiskLayoutGroup{
                    components.BaselineDiskLayoutGroup{
                        Role: components.BaselineDiskLayoutGroupRoleOs.ToPointer(),
                        Count: latitudeshgosdk.Pointer[int64](2),
                        RaidLevel: components.RaidLevelRaid1.ToPointer(),
                    },
                    components.BaselineDiskLayoutGroup{
                        Role: components.BaselineDiskLayoutGroupRoleStorage.ToPointer(),
                        Count: latitudeshgosdk.Pointer[int64](2),
                        Filesystem: components.FilesystemExt4.ToPointer(),
                        MountPoint: latitudeshgosdk.Pointer("/data"),
                    },
                },
                Bios: &operations.CreateBaselineBios{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Baseline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.CreateBaselineBaselinesPreviewRequestBody](../../models/operations/createbaselinebaselinespreviewrequestbody.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CreateBaselineResponse](../../models/operations/createbaselineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetBaseline

**Preview.** Available to teams with the `baselines_api` feature flag.

Retrieve a single baseline.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-baseline" method="get" path="/baselines/{baseline_id}" example="Success" -->
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

    res, err := s.BaselinesPreview.GetBaseline(ctx, "bl_6059EqYkOQj8p")
    if err != nil {
        log.Fatal(err)
    }
    if res.Baseline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `baselineID`                                             | `string`                                                 | :heavy_check_mark:                                       | Baseline ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetBaselineResponse](../../models/operations/getbaselineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DestroyBaseline

**Preview.** Available to teams with the `baselines_api` feature flag.

Delete a baseline.


### Example Usage

<!-- UsageSnippet language="go" operationID="destroy-baseline" method="delete" path="/baselines/{baseline_id}" -->
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

    res, err := s.BaselinesPreview.DestroyBaseline(ctx, "<id>")
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
| `baselineID`                                             | `string`                                                 | :heavy_check_mark:                                       | Baseline ID                                              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DestroyBaselineResponse](../../models/operations/destroybaselineresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |