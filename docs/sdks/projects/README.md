# Projects
(*Projects*)

## Overview

### Available Operations

* [List](#list) - List all Projects
* [Create](#create) - Create a Project
* [Update](#update) - Update a Project
* [Delete](#delete) - Delete a Project

## List

Returns a list of all projects for the current team


### Example Usage

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

    res, err := s.Projects.List(ctx, operations.GetProjectsRequest{
        FilterTags: latitudeshgosdk.String("tag_ezo5vGxZXVU3ozExYeXaCMJzvp0"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Projects != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetProjectsRequest](../../models/operations/getprojectsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetProjectsResponse](../../models/operations/getprojectsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Create a Project

### Example Usage

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

    res, err := s.Projects.Create(ctx, operations.CreateProjectProjectsRequestBody{
        Data: &operations.CreateProjectProjectsData{
            Type: operations.CreateProjectProjectsTypeProjects,
            Attributes: &operations.CreateProjectProjectsAttributes{
                Name: "Mosciski Group",
                ProvisioningType: operations.CreateProjectProvisioningTypeOnDemand,
                Description: latitudeshgosdk.String("Three egg whites with spinach, mushrooms, caramelized onions, tomatoes and low-fat feta cheese. With herbed quinoa, and your choice of rye or whole-grain toast."),
                Environment: operations.CreateProjectEnvironmentDevelopment.ToPointer(),
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
| `request`                                                                                                  | [operations.CreateProjectProjectsRequestBody](../../models/operations/createprojectprojectsrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateProjectResponse](../../models/operations/createprojectresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 400, 403, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Update

Update a Project

### Example Usage

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

    res, err := s.Projects.Update(ctx, "proj_Z8rodm2Yq1jLB", &operations.UpdateProjectProjectsRequestBody{
        Data: operations.UpdateProjectProjectsData{
            Type: operations.UpdateProjectProjectsTypeProjects,
            Attributes: &operations.UpdateProjectProjectsAttributes{
                Tags: []string{
                    "tag_zNzZ1bp9Nos8E1YzEzGGce8vkMp",
                    "tag_L8KGx7Rp46sLPgmGPeARFEMLyxw",
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

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `projectID`                                                                                                 | *string*                                                                                                    | :heavy_check_mark:                                                                                          | The project ID or Slug                                                                                      |
| `requestBody`                                                                                               | [*operations.UpdateProjectProjectsRequestBody](../../models/operations/updateprojectprojectsrequestbody.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.UpdateProjectResponse](../../models/operations/updateprojectresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## Delete

Delete a Project

### Example Usage

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

    res, err := s.Projects.Delete(ctx, "invalid")
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The project ID or Slug                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectResponse](../../models/operations/deleteprojectresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |