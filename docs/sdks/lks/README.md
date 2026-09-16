# Lks

## Overview

### Available Operations

* [ListLksClusters](#listlksclusters) - List LKS clusters
* [CreateLksCluster](#createlkscluster) - Create an LKS cluster
* [GetLksCluster](#getlkscluster) - Get an LKS cluster
* [DeleteLksCluster](#deletelkscluster) - Delete an LKS cluster
* [UpdateLksCluster](#updatelkscluster) - Update an LKS cluster
* [GetLksClusterKubeconfig](#getlksclusterkubeconfig) - Get the cluster kubeconfig
* [ListLksNodePools](#listlksnodepools) - List node pools
* [CreateLksNodePool](#createlksnodepool) - Create a node pool
* [GetLksNodePool](#getlksnodepool) - Get a node pool
* [DeleteLksNodePool](#deletelksnodepool) - Delete a node pool
* [UpdateLksNodePool](#updatelksnodepool) - Update a node pool
* [ListLksAvailableVersions](#listlksavailableversions) - List available Kubernetes versions
* [ListLksSites](#listlkssites) - List sites available for LKS

## ListLksClusters

Lists every LKS cluster of a project. The response is not paginated; `meta.total` is the number of clusters returned.


### Example Usage: Empty

<!-- UsageSnippet language="go" operationID="list-lks-clusters" method="get" path="/lks/clusters" example="Empty" -->
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

    res, err := s.Lks.ListLksClusters(ctx, "proj_6059EqYkOQj8p")
    if err != nil {
        log.Fatal(err)
    }
    if res.LksClusters != nil {
        // handle response
    }
}
```
### Example Usage: OneCluster

<!-- UsageSnippet language="go" operationID="list-lks-clusters" method="get" path="/lks/clusters" example="OneCluster" -->
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

    res, err := s.Lks.ListLksClusters(ctx, "proj_6059EqYkOQj8p")
    if err != nil {
        log.Fatal(err)
    }
    if res.LksClusters != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `projectID`                                                                    | `string`                                                                       | :heavy_check_mark:                                                             | Project `id_hash` or slug. Required — clusters are always scoped to a project. | proj_6059EqYkOQj8p                                                             |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.ListLksClustersResponse](../../models/operations/listlksclustersresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateLksCluster

Creates an LKS cluster. The cluster is the control plane only — worker capacity is added separately through node pools (`POST /lks/clusters/{cluster_id}/nodepools`).

`site` must be one of the slugs returned by `GET /lks/sites`, and `kubernetes_version` must be a patch listed by `GET /lks/available_versions` with `available_for_creation: true`.

The cluster is returned immediately with `status: "provisioning"`. Poll `GET /lks/clusters/{id}` until `status` is `ready`, then fetch the kubeconfig.


### Example Usage: InsufficientPermissions

<!-- UsageSnippet language="go" operationID="create-lks-cluster" method="post" path="/lks/clusters" example="InsufficientPermissions" -->
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

    res, err := s.Lks.CreateLksCluster(ctx, components.CreateLksCluster{
        Data: components.CreateLksClusterData{
            Type: components.CreateLksClusterTypeLksClusters,
            Attributes: components.CreateLksClusterAttributes{
                Name: "<value>",
                ProjectID: "<id>",
                Site: "<value>",
                KubernetesVersion: "1.36.3",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```
### Example Usage: Minimal

<!-- UsageSnippet language="go" operationID="create-lks-cluster" method="post" path="/lks/clusters" example="Minimal" -->
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

    res, err := s.Lks.CreateLksCluster(ctx, components.CreateLksCluster{
        Data: components.CreateLksClusterData{
            Type: components.CreateLksClusterTypeLksClusters,
            Attributes: components.CreateLksClusterAttributes{
                Name: "production",
                ProjectID: "proj_6059EqYkOQj8p",
                Site: "DAL2",
                KubernetesVersion: "1.36.3",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```
### Example Usage: Provisioning

<!-- UsageSnippet language="go" operationID="create-lks-cluster" method="post" path="/lks/clusters" example="Provisioning" -->
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

    res, err := s.Lks.CreateLksCluster(ctx, components.CreateLksCluster{
        Data: components.CreateLksClusterData{
            Type: components.CreateLksClusterTypeLksClusters,
            Attributes: components.CreateLksClusterAttributes{
                Name: "<value>",
                ProjectID: "<id>",
                Site: "<value>",
                KubernetesVersion: "1.36.3",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```
### Example Usage: WithNetworkOverrides

<!-- UsageSnippet language="go" operationID="create-lks-cluster" method="post" path="/lks/clusters" example="WithNetworkOverrides" -->
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

    res, err := s.Lks.CreateLksCluster(ctx, components.CreateLksCluster{
        Data: components.CreateLksClusterData{
            Type: components.CreateLksClusterTypeLksClusters,
            Attributes: components.CreateLksClusterAttributes{
                Name: "production",
                ProjectID: "proj_6059EqYkOQj8p",
                Site: "DAL2",
                KubernetesVersion: "1.36.3",
                Description: latitudeshgosdk.Pointer("Main production cluster"),
                Network: &components.CreateLksClusterNetwork{
                    PodCidrs: []string{
                        "10.70.0.0/16",
                    },
                    ServiceCidrs: []string{
                        "10.71.0.0/16",
                    },
                    NodeCidrs: []string{
                        "10.72.0.0/24",
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.CreateLksCluster](../../models/components/createlkscluster.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateLksClusterResponse](../../models/operations/createlksclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetLksCluster

Retrieves a single LKS cluster.

`status` is an open enum sourced from the platform controller — `provisioning`, `ready`, `updating`, `scaling`, `upgrading`, `paused`, `deleting` and `deleted` are the values in use today, and new ones may appear without notice. `reason` and `message` carry the machine-readable and human-readable detail behind the current `status`.


### Example Usage: Provisioning

<!-- UsageSnippet language="go" operationID="get-lks-cluster" method="get" path="/lks/clusters/{id}" example="Provisioning" -->
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

    res, err := s.Lks.GetLksCluster(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```
### Example Usage: Ready

<!-- UsageSnippet language="go" operationID="get-lks-cluster" method="get" path="/lks/clusters/{id}" example="Ready" -->
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

    res, err := s.Lks.GetLksCluster(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The cluster ID (format: `lksc_<hash>`).                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetLksClusterResponse](../../models/operations/getlksclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteLksCluster

Marks the cluster for deletion. The call returns as soon as the tombstone is written; the platform then tears the cluster and its node pools down asynchronously. A cluster that is already deleted, or that is paused, rejects the request.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-lks-cluster" method="delete" path="/lks/clusters/{id}" -->
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

    res, err := s.Lks.DeleteLksCluster(ctx, "<id>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The cluster ID (format: `lksc_<hash>`).                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteLksClusterResponse](../../models/operations/deletelksclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409            | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UpdateLksCluster

Renames a cluster, edits its description, or upgrades the control plane. At least one of `name`, `description` or `kubernetes_version` must be provided.

Setting a newer `kubernetes_version` consents to a control-plane upgrade; the value must be a patch listed by `GET /lks/available_versions` with `available_for_upgrade: true` and must not be lower than the current one. Node pools are upgraded separately and must never run a patch newer than the control plane.

The cluster must be idle: a cluster that is provisioning, updating, scaling, upgrading, paused or deleting rejects the request with 409.


### Example Usage: InsufficientPermissions

<!-- UsageSnippet language="go" operationID="update-lks-cluster" method="patch" path="/lks/clusters/{id}" example="InsufficientPermissions" -->
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

    res, err := s.Lks.UpdateLksCluster(ctx, "<id>", components.UpdateLksCluster{
        Data: components.UpdateLksClusterData{
            Type: components.UpdateLksClusterTypeLksClusters,
            Attributes: &components.UpdateLksClusterAttributes{
                KubernetesVersion: latitudeshgosdk.Pointer("1.37.2"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```
### Example Usage: Rename

<!-- UsageSnippet language="go" operationID="update-lks-cluster" method="patch" path="/lks/clusters/{id}" example="Rename" -->
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

    res, err := s.Lks.UpdateLksCluster(ctx, "<id>", components.UpdateLksCluster{
        Data: components.UpdateLksClusterData{
            Type: components.UpdateLksClusterTypeLksClusters,
            Attributes: &components.UpdateLksClusterAttributes{
                Name: latitudeshgosdk.Pointer("production-eu"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```
### Example Usage: UpgradeControlPlane

<!-- UsageSnippet language="go" operationID="update-lks-cluster" method="patch" path="/lks/clusters/{id}" example="UpgradeControlPlane" -->
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

    res, err := s.Lks.UpdateLksCluster(ctx, "<id>", components.UpdateLksCluster{
        Data: components.UpdateLksClusterData{
            Type: components.UpdateLksClusterTypeLksClusters,
            Attributes: &components.UpdateLksClusterAttributes{
                KubernetesVersion: latitudeshgosdk.Pointer("1.37.2"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```
### Example Usage: Upgrading

<!-- UsageSnippet language="go" operationID="update-lks-cluster" method="patch" path="/lks/clusters/{id}" example="Upgrading" -->
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

    res, err := s.Lks.UpdateLksCluster(ctx, "<id>", components.UpdateLksCluster{
        Data: components.UpdateLksClusterData{
            Type: components.UpdateLksClusterTypeLksClusters,
            Attributes: &components.UpdateLksClusterAttributes{
                KubernetesVersion: latitudeshgosdk.Pointer("1.37.2"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksCluster != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `id`                                                                       | `string`                                                                   | :heavy_check_mark:                                                         | The cluster ID (format: `lksc_<hash>`).                                    |
| `updateLksCluster`                                                         | [components.UpdateLksCluster](../../models/components/updatelkscluster.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.UpdateLksClusterResponse](../../models/operations/updatelksclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409, 422       | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetLksClusterKubeconfig

Returns the kubeconfig for the cluster. It only exists once the control plane is up, so this endpoint answers 409 `NOT_READY` while the cluster is still provisioning — poll `GET /lks/clusters/{id}` until `kubeconfig_url` is set.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-lks-cluster-kubeconfig" method="get" path="/lks/clusters/{id}/kubeconfig" example="Ready" -->
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

    res, err := s.Lks.GetLksClusterKubeconfig(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LksClusterKubeconfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The cluster ID (format: `lksc_<hash>`).                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetLksClusterKubeconfigResponse](../../models/operations/getlksclusterkubeconfigresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409            | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListLksNodePools

Lists every node pool of an LKS cluster. The response is not paginated; `meta.total` is the number of node pools returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-lks-node-pools" method="get" path="/lks/clusters/{cluster_id}/nodepools" example="OnePool" -->
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

    res, err := s.Lks.ListLksNodePools(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePools != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `clusterID`                                              | `string`                                                 | :heavy_check_mark:                                       | The cluster ID (format: `lksc_<hash>`).                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListLksNodePoolsResponse](../../models/operations/listlksnodepoolsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateLksNodePool

Adds a node pool to an LKS cluster. The platform provisions `count` servers of `plan` from stock, so both fields are required.

`kubernetes_version` defaults to the control-plane patch and may never be newer than it. `max_pods_per_node` is set once, here — it cannot be changed later.


### Example Usage: InsufficientPermissions

<!-- UsageSnippet language="go" operationID="create-lks-node-pool" method="post" path="/lks/clusters/{cluster_id}/nodepools" example="InsufficientPermissions" -->
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

    res, err := s.Lks.CreateLksNodePool(ctx, "<id>", components.CreateLksNodePool{
        Data: components.CreateLksNodePoolData{
            Type: components.CreateLksNodePoolTypeLksNodePools,
            Attributes: components.CreateLksNodePoolAttributes{
                Plan: "<value>",
                Count: 973849,
                KubernetesVersion: latitudeshgosdk.Pointer("1.36.3"),
                MaxPodsPerNode: latitudeshgosdk.Pointer[int64](110),
                Taints: []components.LksNodePoolTaint{
                    components.LksNodePoolTaint{
                        Key: "dedicated",
                        Value: latitudeshgosdk.Pointer("gpu"),
                        Effect: components.EffectNoSchedule,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: OnDemand

<!-- UsageSnippet language="go" operationID="create-lks-node-pool" method="post" path="/lks/clusters/{cluster_id}/nodepools" example="OnDemand" -->
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

    res, err := s.Lks.CreateLksNodePool(ctx, "<id>", components.CreateLksNodePool{
        Data: components.CreateLksNodePoolData{
            Type: components.CreateLksNodePoolTypeLksNodePools,
            Attributes: components.CreateLksNodePoolAttributes{
                Plan: "c2-small-x86",
                Count: 2,
                Name: latitudeshgosdk.Pointer("pool-a"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: Provisioning

<!-- UsageSnippet language="go" operationID="create-lks-node-pool" method="post" path="/lks/clusters/{cluster_id}/nodepools" example="Provisioning" -->
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

    res, err := s.Lks.CreateLksNodePool(ctx, "<id>", components.CreateLksNodePool{
        Data: components.CreateLksNodePoolData{
            Type: components.CreateLksNodePoolTypeLksNodePools,
            Attributes: components.CreateLksNodePoolAttributes{
                Plan: "<value>",
                Count: 973849,
                KubernetesVersion: latitudeshgosdk.Pointer("1.36.3"),
                MaxPodsPerNode: latitudeshgosdk.Pointer[int64](110),
                Taints: []components.LksNodePoolTaint{
                    components.LksNodePoolTaint{
                        Key: "dedicated",
                        Value: latitudeshgosdk.Pointer("gpu"),
                        Effect: components.EffectNoSchedule,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: WithLabelsAndTaints

<!-- UsageSnippet language="go" operationID="create-lks-node-pool" method="post" path="/lks/clusters/{cluster_id}/nodepools" example="WithLabelsAndTaints" -->
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

    res, err := s.Lks.CreateLksNodePool(ctx, "<id>", components.CreateLksNodePool{
        Data: components.CreateLksNodePoolData{
            Type: components.CreateLksNodePoolTypeLksNodePools,
            Attributes: components.CreateLksNodePoolAttributes{
                Plan: "g3-xlarge-x86",
                Count: 3,
                MaxPodsPerNode: latitudeshgosdk.Pointer[int64](250),
                Name: latitudeshgosdk.Pointer("gpu"),
                Labels: map[string]string{
                    "workload": "training",
                },
                Taints: []components.LksNodePoolTaint{
                    components.LksNodePoolTaint{
                        Key: "dedicated",
                        Value: latitudeshgosdk.Pointer("gpu"),
                        Effect: components.EffectNoSchedule,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `clusterID`                                                                  | `string`                                                                     | :heavy_check_mark:                                                           | The cluster ID (format: `lksc_<hash>`).                                      |
| `createLksNodePool`                                                          | [components.CreateLksNodePool](../../models/components/createlksnodepool.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateLksNodePoolResponse](../../models/operations/createlksnodepoolresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 422            | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## GetLksNodePool

Retrieves a single node pool of an LKS cluster. `ready_nodes` reports how many of the pool's `count` nodes have joined the cluster.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-lks-node-pool" method="get" path="/lks/clusters/{cluster_id}/nodepools/{id}" example="Ready" -->
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

    res, err := s.Lks.GetLksNodePool(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `clusterID`                                              | `string`                                                 | :heavy_check_mark:                                       | The cluster ID (format: `lksc_<hash>`).                  |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The node pool ID (format: `lksnp_<hash>`).               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetLksNodePoolResponse](../../models/operations/getlksnodepoolresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404                 | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## DeleteLksNodePool

Marks the node pool for deletion. The call returns as soon as the tombstone is written; the platform then drains and releases the nodes asynchronously.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-lks-node-pool" method="delete" path="/lks/clusters/{cluster_id}/nodepools/{id}" -->
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

    res, err := s.Lks.DeleteLksNodePool(ctx, "<id>", "<id>")
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
| `clusterID`                                              | `string`                                                 | :heavy_check_mark:                                       | The cluster ID (format: `lksc_<hash>`).                  |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The node pool ID (format: `lksnp_<hash>`).               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteLksNodePoolResponse](../../models/operations/deletelksnodepoolresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409            | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## UpdateLksNodePool

Scales, renames, upgrades or re-labels a node pool. At least one attribute must be provided.

`labels` and `taints` are declarative replacements, not merges: omit the field to leave it untouched, send the whole map/list to replace it, or send `{}` / `[]` to clear it.

A newer `kubernetes_version` rolls a node-recreating upgrade; it must not be lower than the pool's current patch nor newer than the control-plane patch. `max_pods_per_node` is immutable — a PATCH that carries it is rejected with 422 even if the value is unchanged.

The node pool must be idle: one that is provisioning, updating, scaling, upgrading, paused or deleting rejects the request with 409.


### Example Usage: ClearTaints

<!-- UsageSnippet language="go" operationID="update-lks-node-pool" method="patch" path="/lks/clusters/{cluster_id}/nodepools/{id}" example="ClearTaints" -->
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

    res, err := s.Lks.UpdateLksNodePool(ctx, "<id>", "<id>", components.UpdateLksNodePool{
        Data: components.UpdateLksNodePoolData{
            Type: components.UpdateLksNodePoolTypeLksNodePools,
            Attributes: &components.UpdateLksNodePoolAttributes{
                Taints: []components.LksNodePoolTaint{},
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: InsufficientPermissions

<!-- UsageSnippet language="go" operationID="update-lks-node-pool" method="patch" path="/lks/clusters/{cluster_id}/nodepools/{id}" example="InsufficientPermissions" -->
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

    res, err := s.Lks.UpdateLksNodePool(ctx, "<id>", "<id>", components.UpdateLksNodePool{
        Data: components.UpdateLksNodePoolData{
            Type: components.UpdateLksNodePoolTypeLksNodePools,
            Attributes: &components.UpdateLksNodePoolAttributes{
                Taints: []components.LksNodePoolTaint{
                    components.LksNodePoolTaint{
                        Key: "dedicated",
                        Value: latitudeshgosdk.Pointer("gpu"),
                        Effect: components.EffectNoSchedule,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: ReplaceLabels

<!-- UsageSnippet language="go" operationID="update-lks-node-pool" method="patch" path="/lks/clusters/{cluster_id}/nodepools/{id}" example="ReplaceLabels" -->
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

    res, err := s.Lks.UpdateLksNodePool(ctx, "<id>", "<id>", components.UpdateLksNodePool{
        Data: components.UpdateLksNodePoolData{
            Type: components.UpdateLksNodePoolTypeLksNodePools,
            Attributes: &components.UpdateLksNodePoolAttributes{
                Labels: map[string]string{
                    "env": "prod",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: Scale

<!-- UsageSnippet language="go" operationID="update-lks-node-pool" method="patch" path="/lks/clusters/{cluster_id}/nodepools/{id}" example="Scale" -->
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

    res, err := s.Lks.UpdateLksNodePool(ctx, "<id>", "<id>", components.UpdateLksNodePool{
        Data: components.UpdateLksNodePoolData{
            Type: components.UpdateLksNodePoolTypeLksNodePools,
            Attributes: &components.UpdateLksNodePoolAttributes{
                Count: latitudeshgosdk.Pointer[int64](4),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: Scaling

<!-- UsageSnippet language="go" operationID="update-lks-node-pool" method="patch" path="/lks/clusters/{cluster_id}/nodepools/{id}" example="Scaling" -->
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

    res, err := s.Lks.UpdateLksNodePool(ctx, "<id>", "<id>", components.UpdateLksNodePool{
        Data: components.UpdateLksNodePoolData{
            Type: components.UpdateLksNodePoolTypeLksNodePools,
            Attributes: &components.UpdateLksNodePoolAttributes{
                Taints: []components.LksNodePoolTaint{
                    components.LksNodePoolTaint{
                        Key: "dedicated",
                        Value: latitudeshgosdk.Pointer("gpu"),
                        Effect: components.EffectNoSchedule,
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```
### Example Usage: Upgrade

<!-- UsageSnippet language="go" operationID="update-lks-node-pool" method="patch" path="/lks/clusters/{cluster_id}/nodepools/{id}" example="Upgrade" -->
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

    res, err := s.Lks.UpdateLksNodePool(ctx, "<id>", "<id>", components.UpdateLksNodePool{
        Data: components.UpdateLksNodePoolData{
            Type: components.UpdateLksNodePoolTypeLksNodePools,
            Attributes: &components.UpdateLksNodePoolAttributes{
                KubernetesVersion: latitudeshgosdk.Pointer("1.36.3"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LksNodePool != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `clusterID`                                                                  | `string`                                                                     | :heavy_check_mark:                                                           | The cluster ID (format: `lksc_<hash>`).                                      |
| `id`                                                                         | `string`                                                                     | :heavy_check_mark:                                                           | The node pool ID (format: `lksnp_<hash>`).                                   |
| `updateLksNodePool`                                                          | [components.UpdateLksNodePool](../../models/components/updatelksnodepool.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.UpdateLksNodePoolResponse](../../models/operations/updatelksnodepoolresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 403, 404, 409, 422       | application/vnd.api+json |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListLksAvailableVersions

Lists the Kubernetes patches the platform offers, with their lifecycle flags. Use a version with `available_for_creation: true` when creating a cluster or a node pool, and one with `available_for_upgrade: true` when upgrading. Exactly one entry has `default: true`.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-lks-available-versions" method="get" path="/lks/available_versions" example="TwoVersions" -->
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

    res, err := s.Lks.ListLksAvailableVersions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.LksKubernetesVersions != nil {
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

**[*operations.ListLksAvailableVersionsResponse](../../models/operations/listlksavailableversionsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## ListLksSites

Lists the sites that can host an LKS cluster, one entry per site. Pass an entry's `id` (the site slug) as `site` when creating a cluster. `country` is null when the underlying site has no region assigned.


### Example Usage

<!-- UsageSnippet language="go" operationID="list-lks-sites" method="get" path="/lks/sites" example="TwoSites" -->
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

    res, err := s.Lks.ListLksSites(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.LksSites != nil {
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

**[*operations.ListLksSitesResponse](../../models/operations/listlkssitesresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 502                      | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |