# Clusters

## Overview

Projects that have a GPU cluster only. Every operation answers 404 PROJECT_NOT_A_CLUSTER for a project without one.

### Available Operations

* [GetProjectClusterRequirements](#getprojectclusterrequirements) - Get cluster requirements
* [UpdateProjectClusterRequirements](#updateprojectclusterrequirements) - Change cluster requirements
* [SignProjectClusterRequirements](#signprojectclusterrequirements) - Sign cluster requirements
* [GetProjectClusterRequirementSignoffs](#getprojectclusterrequirementsignoffs) - List cluster requirement sign-offs
* [GetProjectClusterNodeGroups](#getprojectclusternodegroups) - List cluster node groups
* [CreateProjectClusterNodeGroup](#createprojectclusternodegroup) - Create a cluster node group
* [DeleteProjectClusterNodeGroup](#deleteprojectclusternodegroup) - Delete a cluster node group
* [UpdateProjectClusterNodeGroup](#updateprojectclusternodegroup) - Update a cluster node group

## GetProjectClusterRequirements

Shows a cluster project's requirements for review: the requirement blocks and their expected values, the node groups, the version to sign and who signed. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER. While the requirements are a draft, their content is empty.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-project-cluster-requirements" method="get" path="/projects/{project_id}/cluster/requirements" example="Success" -->
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

    res, err := s.Clusters.GetProjectClusterRequirements(ctx, "proj_lxWpD699qm6rk")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClusterRequirements != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID or slug of a project that has a GPU cluster       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetProjectClusterRequirementsResponse](../../models/operations/getprojectclusterrequirementsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateProjectClusterRequirements

Applies one change to the requirements: add or remove a requirement or the checks of a block, or set what a requirement expects. Signed requirements go back to with_customer to sign again. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER. Needs an owner, administrator or collaborator.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-project-cluster-requirements" method="patch" path="/projects/{project_id}/cluster/requirements" example="Success" -->
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

    res, err := s.Clusters.UpdateProjectClusterRequirements(ctx, "proj_lxWpD699qm6rk", components.ChangeClusterRequirements{
        Data: components.ChangeClusterRequirementsData{
            Type: components.ChangeClusterRequirementsTypeClusterRequirements,
            Attributes: components.ChangeClusterRequirementsAttributes{
                Add: latitudeshgosdk.Pointer("storage_fs"),
                Expected: latitudeshgosdk.Pointer("xfs"),
                Note: latitudeshgosdk.Pointer("Scratch"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClusterRequirements != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `projectID`                                                                                  | `string`                                                                                     | :heavy_check_mark:                                                                           | The ID or slug of a project that has a GPU cluster                                           |
| `changeClusterRequirements`                                                                  | [components.ChangeClusterRequirements](../../models/components/changeclusterrequirements.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateProjectClusterRequirementsResponse](../../models/operations/updateprojectclusterrequirementsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## SignProjectClusterRequirements

Signs the requirements as the authenticated user, for the requirements_version they were shown. A version that no longer matches the requirements is refused. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER. Needs an owner or administrator.

### Example Usage

<!-- UsageSnippet language="go" operationID="sign-project-cluster-requirements" method="post" path="/projects/{project_id}/cluster/requirements/sign" example="Created" -->
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

    res, err := s.Clusters.SignProjectClusterRequirements(ctx, "proj_lxWpD699qm6rk", components.SignClusterRequirements{
        Data: components.SignClusterRequirementsData{
            Type: components.SignClusterRequirementsTypeClusterRequirements,
            Attributes: components.SignClusterRequirementsAttributes{
                RequirementsVersion: "3f1c9b2e7a4d8c06b5e9f2a1d7c3e8b04a6f9d2c1b7e5a3f8d0c4b6e2a9f7d1c",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClusterRequirements != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `projectID`                                                                              | `string`                                                                                 | :heavy_check_mark:                                                                       | The ID or slug of a project that has a GPU cluster                                       |
| `signClusterRequirements`                                                                | [components.SignClusterRequirements](../../models/components/signclusterrequirements.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.SignProjectClusterRequirementsResponse](../../models/operations/signprojectclusterrequirementsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetProjectClusterRequirementSignoffs

Lists who signed the requirements and when, newest first. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-project-cluster-requirement-signoffs" method="get" path="/projects/{project_id}/cluster/requirements/signoffs" example="Success" -->
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

    res, err := s.Clusters.GetProjectClusterRequirementSignoffs(ctx, "proj_lxWpD699qm6rk")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClusterRequirementSignoffs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID or slug of a project that has a GPU cluster       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetProjectClusterRequirementSignoffsResponse](../../models/operations/getprojectclusterrequirementsignoffsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## GetProjectClusterNodeGroups

Lists the node groups of a cluster project. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER. While the cluster's requirements are a draft, the list is empty.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-project-cluster-node-groups" method="get" path="/projects/{project_id}/cluster/node_groups" example="Success" -->
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

    res, err := s.Clusters.GetProjectClusterNodeGroups(ctx, "proj_lxWpD699qm6rk")
    if err != nil {
        log.Fatal(err)
    }
    if res.ClusterNodeGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID or slug of a project that has a GPU cluster       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetProjectClusterNodeGroupsResponse](../../models/operations/getprojectclusternodegroupsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateProjectClusterNodeGroup

Adds a node group to a cluster project. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER. Signed requirements go back to the customer to sign again. Needs an owner, administrator or collaborator.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-project-cluster-node-group" method="post" path="/projects/{project_id}/cluster/node_groups" example="Created" -->
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

    res, err := s.Clusters.CreateProjectClusterNodeGroup(ctx, "proj_lxWpD699qm6rk", components.CreateClusterNodeGroup{
        Data: components.CreateClusterNodeGroupData{
            Type: components.CreateClusterNodeGroupTypeClusterNodeGroups,
            Attributes: components.CreateClusterNodeGroupAttributes{
                Name: "login",
                Kind: components.CreateClusterNodeGroupKindCPU,
                Plan: "g4-b300-large",
                ServerCount: 2,
                Os: latitudeshgosdk.Pointer("rocky_9_x64"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClusterNodeGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `projectID`                                                                            | `string`                                                                               | :heavy_check_mark:                                                                     | The ID or slug of a project that has a GPU cluster                                     |
| `createClusterNodeGroup`                                                               | [components.CreateClusterNodeGroup](../../models/components/createclusternodegroup.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateProjectClusterNodeGroupResponse](../../models/operations/createprojectclusternodegroupresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## DeleteProjectClusterNodeGroup

Removes a cluster project's node group. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER. Signed requirements go back to the customer to sign again. Needs an owner, administrator or collaborator.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-project-cluster-node-group" method="delete" path="/projects/{project_id}/cluster/node_groups/{id}" -->
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

    res, err := s.Clusters.DeleteProjectClusterNodeGroup(ctx, "proj_lxWpD699qm6rk", "7")
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | The ID or slug of a project that has a GPU cluster       |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The node group ID                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectClusterNodeGroupResponse](../../models/operations/deleteprojectclusternodegroupresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## UpdateProjectClusterNodeGroup

Changes the attributes sent of a cluster project's node group. Only available for projects that have a GPU cluster; any other project answers 404 PROJECT_NOT_A_CLUSTER. Signed requirements go back to the customer to sign again. Needs an owner, administrator or collaborator.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-project-cluster-node-group" method="patch" path="/projects/{project_id}/cluster/node_groups/{id}" example="Success" -->
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

    res, err := s.Clusters.UpdateProjectClusterNodeGroup(ctx, "proj_lxWpD699qm6rk", "7", components.UpdateClusterNodeGroup{
        Data: components.UpdateClusterNodeGroupData{
            Type: components.UpdateClusterNodeGroupTypeClusterNodeGroups,
            Attributes: components.UpdateClusterNodeGroupAttributes{
                ServerCount: latitudeshgosdk.Pointer[int64](32),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClusterNodeGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `projectID`                                                                            | `string`                                                                               | :heavy_check_mark:                                                                     | The ID or slug of a project that has a GPU cluster                                     |
| `id`                                                                                   | `string`                                                                               | :heavy_check_mark:                                                                     | The node group ID                                                                      |
| `updateClusterNodeGroup`                                                               | [components.UpdateClusterNodeGroup](../../models/components/updateclusternodegroup.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateProjectClusterNodeGroupResponse](../../models/operations/updateprojectclusternodegroupresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |