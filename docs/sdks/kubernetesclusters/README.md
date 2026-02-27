# KubernetesClusters

## Overview

### Available Operations

* [ListKubernetesClusters](#listkubernetesclusters) - List Kubernetes Clusters
* [CreateKubernetesCluster](#createkubernetescluster) - Create a Kubernetes Cluster
* [GetKubernetesCluster](#getkubernetescluster) - Get a Kubernetes Cluster
* [DeleteKubernetesCluster](#deletekubernetescluster) - Delete a Kubernetes Cluster
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | The project ID to filter clusters by                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListKubernetesClustersResponse](../../models/operations/listkubernetesclustersresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 400, 401, 403            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |

## CreateKubernetesCluster

Creates a new managed Kubernetes cluster. The cluster will be provisioned using Cluster API (CAPI) on the Management Cluster.

**Note:** This feature requires the `managed_k8s_clusters_api` feature flag to be enabled for your team. Maximum of 1 cluster per project.

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
                Name: "my-cluster",
                ProjectID: "proj_6059EqYkOQj8p",
                Site: "SAN3",
                Plan: "c2-small-x86",
                SSHKeyID: "ssh_VkE1DwV37dnZJ",
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
                Name: "<value>",
                ProjectID: "<id>",
                Site: "<value>",
                Plan: "<value>",
                SSHKeyID: "<id>",
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

## GetKubernetesCluster

Retrieves detailed information about a Kubernetes cluster including its status, control plane, and worker node details.


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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `kubernetesClusterID`                                    | *string*                                                 | :heavy_check_mark:                                       | The cluster name                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `kubernetesClusterID`                                    | *string*                                                 | :heavy_check_mark:                                       | The cluster name                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteKubernetesClusterResponse](../../models/operations/deletekubernetesclusterresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403, 404, 422       | application/vnd.api+json |
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `kubernetesClusterID`                                    | *string*                                                 | :heavy_check_mark:                                       | The cluster name                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetKubernetesClusterKubeconfigResponse](../../models/operations/getkubernetesclusterkubeconfigresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| components.ErrorObject   | 401, 403, 404            | application/vnd.api+json |
| components.APIError      | 4XX, 5XX                 | \*/\*                    |