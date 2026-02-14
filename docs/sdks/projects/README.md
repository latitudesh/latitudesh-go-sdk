# Projects

## Overview

### Available Operations

* [List](#list) - List projects
* [Create](#create) - Create project
* [Update](#update) - Update project
* [Delete](#delete) - Delete project

## List

Returns a list of all projects for the current team


### Example Usage: Filtered by multiple tags

<!-- UsageSnippet language="go" operationID="get-projects" method="get" path="/projects" example="Filtered by multiple tags" -->
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
        FilterTags: latitudeshgosdk.Pointer("tag_bg6v8Kz52LHyg0okmVREI6mnKM2,tag_PzXjAnK7MLfpMPzprQ3wSn4rG5P"),
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
### Example Usage: Filtered by tag

<!-- UsageSnippet language="go" operationID="get-projects" method="get" path="/projects" example="Filtered by tag" -->
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
        FilterTags: latitudeshgosdk.Pointer("tag_R3YGrW8m0NSAm0l5Wp6XTnnww9r"),
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-projects" method="get" path="/projects" example="Success" -->
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
        FilterTags: latitudeshgosdk.Pointer("tag_GXeww714mRF2gZ05lnKgU8emo5RE,tag_QQkaK9JnV6tWwPG3pmLviXveVK0Y"),
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

Create project

### Example Usage: Bad Request

<!-- UsageSnippet language="go" operationID="create-project" method="post" path="/projects" example="Bad Request" -->
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
                Name: "Collins LLC",
                ProvisioningType: operations.CreateProjectProvisioningTypeOnDemand,
                Description: latitudeshgosdk.Pointer("Granny Smith apples mixed with brown sugar and butter filling, in a flaky all-butter crust, with ice cream."),
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
### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-project" method="post" path="/projects" example="Created" -->
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
                Name: "Bailey and Sons",
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
### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="create-project" method="post" path="/projects" example="Forbidden" -->
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
                Name: "Morissette, Zieme and Botsford",
                ProvisioningType: operations.CreateProjectProvisioningTypeOnDemand,
                Description: latitudeshgosdk.Pointer("Three egg omelet with Roquefort cheese, chives, and ham. With a side of roasted potatoes, and your choice of toast or croissant."),
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
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="create-project" method="post" path="/projects" example="Success" -->
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
                Name: "Cormier-Corkery",
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
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="create-project" method="post" path="/projects" example="Unprocessable Entity" -->
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
                Name: "123",
                ProvisioningType: operations.CreateProjectProvisioningTypeReserved,
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

Update project

### Example Usage: Forbidden

<!-- UsageSnippet language="go" operationID="update-project" method="patch" path="/projects/{project_id}" example="Forbidden" -->
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

    res, err := s.Projects.Update(ctx, "proj_0MoLqJmGd57pY", &operations.UpdateProjectProjectsRequestBody{
        Data: operations.UpdateProjectProjectsData{
            ID: latitudeshgosdk.Pointer("proj_0MoLqJmGd57pY"),
            Type: operations.UpdateProjectProjectsTypeProjects,
            Attributes: &operations.UpdateProjectProjectsAttributes{
                Name: latitudeshgosdk.Pointer("Moore-Durgan"),
                Description: latitudeshgosdk.Pointer("Sequi occaecati eaque exercitationem."),
                Environment: operations.UpdateProjectProjectsEnvironmentProduction.ToPointer(),
                BandwidthAlert: latitudeshgosdk.Pointer(true),
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
### Example Usage: Not found

<!-- UsageSnippet language="go" operationID="update-project" method="patch" path="/projects/{project_id}" example="Not found" -->
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

    res, err := s.Projects.Update(ctx, "invalid-id", &operations.UpdateProjectProjectsRequestBody{
        Data: operations.UpdateProjectProjectsData{
            ID: latitudeshgosdk.Pointer("invalid-id"),
            Type: operations.UpdateProjectProjectsTypeProjects,
            Attributes: &operations.UpdateProjectProjectsAttributes{
                Name: latitudeshgosdk.Pointer("Moore-Durgan"),
                Description: latitudeshgosdk.Pointer("Sequi occaecati eaque exercitationem."),
                Environment: operations.UpdateProjectProjectsEnvironmentProduction.ToPointer(),
                BandwidthAlert: latitudeshgosdk.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="update-project" method="patch" path="/projects/{project_id}" example="Success" -->
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

    res, err := s.Projects.Update(ctx, "proj_Gr47qleMDAg0m", &operations.UpdateProjectProjectsRequestBody{
        Data: operations.UpdateProjectProjectsData{
            ID: latitudeshgosdk.Pointer("proj_Gr47qleMDAg0m"),
            Type: operations.UpdateProjectProjectsTypeProjects,
            Attributes: &operations.UpdateProjectProjectsAttributes{
                Tags: []string{
                    "tag_VgrmvzlEGJhbGYv0z8YzHLa9PKV",
                    "tag_PEAMyKnQZEHpGAWKMpB6F7EVYyYj",
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
### Example Usage: Tag not updated

<!-- UsageSnippet language="go" operationID="update-project" method="patch" path="/projects/{project_id}" example="Tag not updated" -->
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
            ID: latitudeshgosdk.Pointer("proj_Z8rodm2Yq1jLB"),
            Type: operations.UpdateProjectProjectsTypeProjects,
            Attributes: &operations.UpdateProjectProjectsAttributes{
                Tags: []string{
                    "invalid-tag",
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
### Example Usage: Unprocessable Entity

<!-- UsageSnippet language="go" operationID="update-project" method="patch" path="/projects/{project_id}" example="Unprocessable Entity" -->
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

    res, err := s.Projects.Update(ctx, "proj_5xyZOnoNdWM0l", &operations.UpdateProjectProjectsRequestBody{
        Data: operations.UpdateProjectProjectsData{
            ID: latitudeshgosdk.Pointer("proj_5xyZOnoNdWM0l"),
            Type: operations.UpdateProjectProjectsTypeProjects,
            Attributes: &operations.UpdateProjectProjectsAttributes{
                Name: latitudeshgosdk.Pointer("123"),
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

Delete project

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