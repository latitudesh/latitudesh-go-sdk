# KubernetesClusters

## Overview

### Available Operations

* [ListKubernetesClusters](#listkubernetesclusters) - List Kubernetes Clusters
* [CreateKubernetesCluster](#createkubernetescluster) - Create a Kubernetes Cluster
* [ListKubernetesAvailableVersions](#listkubernetesavailableversions) - List Available Kubernetes Versions
* [GetKubernetesCluster](#getkubernetescluster) - Get a Kubernetes Cluster
* [DeleteKubernetesCluster](#deletekubernetescluster) - Delete a Kubernetes Cluster
* [UpdateKubernetesCluster](#updatekubernetescluster) - Update Kubernetes Cluster
* [GetKubernetesClusterKubeconfig](#getkubernetesclusterkubeconfig) - Get Kubernetes Cluster Kubeconfig

## ListKubernetesClusters

Lists all Kubernetes clusters for a project.


### Example Usage: EmptyList

<!-- UsageSnippet language="go" operationID="list-kubernetes-clusters" method="get" path="/kubernetes_clusters" example="EmptyList" -->
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

    res, err := s.KubernetesClusters.ListKubernetesClusters(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusters != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="list-kubernetes-clusters" method="get" path="/kubernetes_clusters" example="Success" -->
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

    res, err := s.KubernetesClusters.ListKubernetesClusters(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | The project ID to filter clusters by                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListKubernetesClustersResponse](../../models/operations/listkubernetesclustersresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 400, 401, 403            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateKubernetesCluster

Creates a new managed Kubernetes cluster. Maximum of 1 cluster per project.

Cluster names must follow Kubernetes naming rules: lowercase alphanumeric characters or hyphens, must start and end with an alphanumeric character, and be at most 63 characters long.


### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-kubernetes-cluster" method="post" path="/kubernetes_clusters" example="Created" -->
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

    res, err := s.KubernetesClusters.CreateKubernetesCluster(ctx, components.CreateKubernetesCluster{
        Data: components.CreateKubernetesClusterData{
            Type: components.CreateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.CreateKubernetesClusterAttributes{
                Name: latitudeshgosdk.Pointer("my-cluster"),
                ProjectID: "proj_6059EqYkOQj8p",
                Site: "SAN3",
                Plan: "c2-small-x86",
                SSHKeys: []string{
                    "ssh_VkE1DwV37dnZJ",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterCreateResponse != nil {
        // handle response
    }
}
```
### Example Usage: InvalidSshKeys

<!-- UsageSnippet language="go" operationID="create-kubernetes-cluster" method="post" path="/kubernetes_clusters" example="InvalidSshKeys" -->
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

    res, err := s.KubernetesClusters.CreateKubernetesCluster(ctx, components.CreateKubernetesCluster{
        Data: components.CreateKubernetesClusterData{
            Type: components.CreateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.CreateKubernetesClusterAttributes{
                ProjectID: "<id>",
                Site: "<value>",
                Plan: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterCreateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ValidationError

<!-- UsageSnippet language="go" operationID="create-kubernetes-cluster" method="post" path="/kubernetes_clusters" example="ValidationError" -->
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

    res, err := s.KubernetesClusters.CreateKubernetesCluster(ctx, components.CreateKubernetesCluster{
        Data: components.CreateKubernetesClusterData{
            Type: components.CreateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.CreateKubernetesClusterAttributes{
                ProjectID: "<id>",
                Site: "<value>",
                Plan: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.CreateKubernetesCluster](../../models/components/createkubernetescluster.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateKubernetesClusterResponse](../../models/operations/createkubernetesclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 400, 403, 422            | application/vnd.api+json |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListKubernetesAvailableVersions

Returns the list of available Kubernetes versions for cluster creation and upgrades. Versions are sourced from the RKE2 release channels and cached for 24 hours.

Each version object includes:
- `version`: The full version string (e.g., `v1.35.3+rke2r1`)
- `minor`: The minor version number (e.g., `1.35`)

The API returns the latest 5 supported minor versions. When upgrading clusters, you can only upgrade one minor version at a time (e.g., from 1.34 to 1.35).


### Example Usage

<!-- UsageSnippet language="go" operationID="list-kubernetes-available-versions" method="get" path="/kubernetes_clusters/available_versions" example="Success" -->
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

    res, err := s.KubernetesClusters.ListKubernetesAvailableVersions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesAvailableVersions != nil {
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

**[*operations.ListKubernetesAvailableVersionsResponse](../../models/operations/listkubernetesavailableversionsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403                 | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetKubernetesCluster

Retrieves detailed information about a Kubernetes cluster including its status, control plane, worker node details, and individual node information.


### Example Usage: Provisioning

<!-- UsageSnippet language="go" operationID="get-kubernetes-cluster" method="get" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="Provisioning" -->
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

    res, err := s.KubernetesClusters.GetKubernetesCluster(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesCluster != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-kubernetes-cluster" method="get" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="Success" -->
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

    res, err := s.KubernetesClusters.GetKubernetesCluster(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `kubernetesClusterID`                                                                                     | `string`                                                                                                  | :heavy_check_mark:                                                                                        | The cluster ID (format: kc_<hash>) or cluster name. Both formats are accepted for backward compatibility. |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |

### Response

**[*operations.GetKubernetesClusterResponse](../../models/operations/getkubernetesclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403, 404            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteKubernetesCluster

Deletes a Kubernetes cluster. This action is irreversible and will destroy all cluster resources.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-kubernetes-cluster" method="delete" path="/kubernetes_clusters/{kubernetes_cluster_id}" -->
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

    res, err := s.KubernetesClusters.DeleteKubernetesCluster(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `kubernetesClusterID`                                                                                     | `string`                                                                                                  | :heavy_check_mark:                                                                                        | The cluster ID (format: kc_<hash>) or cluster name. Both formats are accepted for backward compatibility. |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |

### Response

**[*operations.DeleteKubernetesClusterResponse](../../models/operations/deletekubernetesclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403, 404, 422       | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UpdateKubernetesCluster

Updates a Kubernetes cluster by scaling nodes or upgrading the Kubernetes version. The cluster must be in `Provisioned` phase to accept updates.

## Scaling Operations

Exactly one of `worker_count` or `control_plane_count` must be provided per request. You cannot scale workers and control plane nodes in the same request.

When scaling up, the API validates that sufficient server stock is available for the requested delta (e.g., scaling from 2 to 5 workers checks for 3 available servers).

When scaling from 0 workers, you must provide a `worker_plan` since there is no existing configuration to inherit the plan from.

Control plane scaling has a minimum of 1 node. You cannot scale control plane nodes to zero.

## Version Upgrades

Provide a `kubernetes_version` parameter to upgrade the cluster to a new Kubernetes version. Version upgrades follow these rules:

- **No downgrades**: You cannot downgrade to a lower version than currently installed
- **One minor version at a time**: You can only upgrade one minor version at a time (e.g., from 1.34 to 1.35, not from 1.34 to 1.36)
- **Mutually exclusive**: Version upgrades cannot be combined with scaling operations in the same request
- **Available versions only**: The target version must be in the list returned by `GET /kubernetes_clusters/available_versions`

Returns 202 Accepted when an update operation is triggered. Poll the GET endpoint to monitor progress. Returns 200 OK if no change is needed (no-op).


### Example Usage: ControlPlaneUnchanged

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ControlPlaneUnchanged" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: InvalidControlPlaneCountType

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="InvalidControlPlaneCountType" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: InvalidWorkerCountType

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="InvalidWorkerCountType" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: MissingParameter

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="MissingParameter" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: MissingScalingParameter

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="MissingScalingParameter" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: MissingWorkerCount

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="MissingWorkerCount" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](204185),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: MutualExclusionViolation

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="MutualExclusionViolation" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleControlPlaneDown

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleControlPlaneDown" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                ControlPlaneCount: latitudeshgosdk.Pointer[int64](1),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleControlPlaneUp

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleControlPlaneUp" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                ControlPlaneCount: latitudeshgosdk.Pointer[int64](3),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleDown

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleDown" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](2),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleFromZero

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleFromZero" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](3),
                WorkerPlan: latitudeshgosdk.Pointer("c3-small-x86"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleToZero

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleToZero" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](0),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleUp

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleUp" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](5),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleWorkersDown

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleWorkersDown" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](2),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleWorkersFromZero

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleWorkersFromZero" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](3),
                WorkerPlan: latitudeshgosdk.Pointer("c3-small-x86"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleWorkersToZero

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleWorkersToZero" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](0),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScaleWorkersUp

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScaleWorkersUp" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](5),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: ScalingMutualExclusion

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="ScalingMutualExclusion" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: Unchanged

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="Unchanged" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                WorkerCount: latitudeshgosdk.Pointer[int64](204185),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpgradeVersion

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="UpgradeVersion" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{
                KubernetesVersion: latitudeshgosdk.Pointer("v1.35.0+rke2r1"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: VersionWithScaling

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="VersionWithScaling" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```
### Example Usage: WorkersUnchanged

<!-- UsageSnippet language="go" operationID="update-kubernetes-cluster" method="patch" path="/kubernetes_clusters/{kubernetes_cluster_id}" example="WorkersUnchanged" -->
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

    res, err := s.KubernetesClusters.UpdateKubernetesCluster(ctx, "<id>", components.UpdateKubernetesCluster{
        Data: components.UpdateKubernetesClusterData{
            Type: components.UpdateKubernetesClusterTypeKubernetesClusters,
            Attributes: components.UpdateKubernetesClusterAttributes{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `kubernetesClusterID`                                                                                     | `string`                                                                                                  | :heavy_check_mark:                                                                                        | The cluster ID (format: kc_<hash>) or cluster name. Both formats are accepted for backward compatibility. |
| `updateKubernetesCluster`                                                                                 | [components.UpdateKubernetesCluster](../../models/components/updatekubernetescluster.md)                  | :heavy_check_mark:                                                                                        | N/A                                                                                                       |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |

### Response

**[*operations.UpdateKubernetesClusterResponse](../../models/operations/updatekubernetesclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 400, 403, 404, 422       | application/vnd.api+json |
| components.ErrorObject   | 503                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetKubernetesClusterKubeconfig

Retrieves the kubeconfig file for a Kubernetes cluster. The kubeconfig is only available once the cluster is fully provisioned.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-kubernetes-cluster-kubeconfig" method="get" path="/kubernetes_clusters/{kubernetes_cluster_id}/kubeconfig" example="Success" -->
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

    res, err := s.KubernetesClusters.GetKubernetesClusterKubeconfig(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.KubernetesClusterKubeconfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `kubernetesClusterID`                                                                                     | `string`                                                                                                  | :heavy_check_mark:                                                                                        | The cluster ID (format: kc_<hash>) or cluster name. Both formats are accepted for backward compatibility. |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |

### Response

**[*operations.GetKubernetesClusterKubeconfigResponse](../../models/operations/getkubernetesclusterkubeconfigresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403, 404            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |