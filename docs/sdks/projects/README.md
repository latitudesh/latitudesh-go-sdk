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

<!-- UsageSnippet language="go" operationID="get-projects" method="get" path="/projects" -->
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
        FilterTags: latitudeshgosdk.Pointer("tag_GGXAB6PNQbU285ZYLP3wt3ZQyXW,tag_xR2B8J4W81TBL6YRGPJaCQRkVyn"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Projects != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
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

<!-- UsageSnippet language="go" operationID="create-project" method="post" path="/projects" -->
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
                Name: "Kshlerin, Hand and Beer",
                ProvisioningType: operations.CreateProjectProvisioningTypeOnDemand,
                Description: latitudeshgosdk.Pointer("Thick slices of French toast bread, brown sugar, half-and-half and vanilla, topped with powdered sugar. With two eggs served any style, and your choice of smoked bacon or smoked ham."),
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Update a Project

### Example Usage

<!-- UsageSnippet language="go" operationID="update-project" method="patch" path="/projects/{project_id}" -->
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

    res, err := s.Projects.Update(ctx, "proj_WVQJDMMwDRbyE", &operations.UpdateProjectProjectsRequestBody{
        Data: operations.UpdateProjectProjectsData{
            ID: latitudeshgosdk.Pointer("proj_WVQJDMMwDRbyE"),
            Type: operations.UpdateProjectProjectsTypeProjects,
            Attributes: &operations.UpdateProjectProjectsAttributes{
                Tags: []string{
                    "tag_xNrKWa1ZaMtBKwaZnYaMTBE2XQb",
                    "tag_Ge9E4oboPgILP5e9aKNAUB4Gaw5",
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Delete a Project

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-project" method="delete" path="/projects/{project_id}" -->
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

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |