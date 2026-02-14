# UserData

## Overview

### Available Operations

* [~~GetProjectUsersData~~](#getprojectusersdata) - List Project user data :warning: **Deprecated**
* [~~GetProjectUserData~~](#getprojectuserdata) - Retrieve Project user data :warning: **Deprecated**
* [~~DeleteProjectUserData~~](#deleteprojectuserdata) - Delete Project user data :warning: **Deprecated**
* [~~Create~~](#create) - Create Project user data :warning: **Deprecated**
* [~~UpdateForProject~~](#updateforproject) - Update Project user data :warning: **Deprecated**
* [List](#list) - List user data
* [CreateNew](#createnew) - Create user data
* [Retrieve](#retrieve) - Retrieve user data
* [Update](#update) - Update user data
* [Delete](#delete) - Delete user data

## ~~GetProjectUsersData~~

List all Users Data in the project. These scripts can be used to configure servers with user data.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="get-project-users-data" method="get" path="/projects/{project_id}/user_data" example="Not Found" -->
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

    res, err := s.UserData.GetProjectUsersData(ctx, "123", latitudeshgosdk.Pointer("decoded_content"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UserData != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-project-users-data" method="get" path="/projects/{project_id}/user_data" example="Success" -->
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

    res, err := s.UserData.GetProjectUsersData(ctx, "proj_LYV8DZYQq5QoE", latitudeshgosdk.Pointer("decoded_content"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UserData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `projectID`                                                                                 | *string*                                                                                    | :heavy_check_mark:                                                                          | Project ID or Slug                                                                          |
| `extraFieldsUserData`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The `decoded_content` is provided as an extra attribute that shows content in decoded form. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetProjectUsersDataResponse](../../models/operations/getprojectusersdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~GetProjectUserData~~

Get User Data in the project. These scripts can be used to configure servers with user data.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="get-project-user-data" method="get" path="/projects/{project_id}/user_data/{user_data_id}" example="Not Found" -->
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

    res, err := s.UserData.GetProjectUserData(ctx, "123", "123", latitudeshgosdk.Pointer("decoded_content"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-project-user-data" method="get" path="/projects/{project_id}/user_data/{user_data_id}" example="Success" -->
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

    res, err := s.UserData.GetProjectUserData(ctx, "proj_e8pKq0GYdWAob", "ud_VLMmAD8EOwop2", latitudeshgosdk.Pointer("decoded_content"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `projectID`                                                                                 | *string*                                                                                    | :heavy_check_mark:                                                                          | Project ID or Slug                                                                          |
| `userDataID`                                                                                | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `extraFieldsUserData`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The `decoded_content` is provided as an extra attribute that shows content in decoded form. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetProjectUserDataResponse](../../models/operations/getprojectuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~DeleteProjectUserData~~

Allow you remove User Data in a project.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-project-user-data" method="delete" path="/projects/{project_id}/user_data/{user_data_id}" -->
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

    res, err := s.UserData.DeleteProjectUserData(ctx, "proj_x1ZJrdx5qg4LV", "123")
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project ID or Slug                                       |
| `userDataID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteProjectUserDataResponse](../../models/operations/deleteprojectuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~Create~~

Allows you to create User Data in a project, which can be used to perform custom setup on your servers after deploy and reinstall.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: Bad Request

<!-- UsageSnippet language="go" operationID="post-project-user-data" method="post" path="/projects/{project_id}/user_data" example="Bad Request" -->
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

    res, err := s.UserData.Create(ctx, "proj_v9BVDap8qRm1W", operations.PostProjectUserDataUserDataRequestBody{
        Data: operations.PostProjectUserDataUserDataData{
            Type: operations.PostProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PostProjectUserDataUserDataAttributes{
                Description: "User Data description",
                Content: "User Data base64 encoded content",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: Created

<!-- UsageSnippet language="go" operationID="post-project-user-data" method="post" path="/projects/{project_id}/user_data" example="Created" -->
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

    res, err := s.UserData.Create(ctx, "proj_8NkvdyMKdeLpx", operations.PostProjectUserDataUserDataRequestBody{
        Data: operations.PostProjectUserDataUserDataData{
            Type: operations.PostProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PostProjectUserDataUserDataAttributes{
                Description: "User Data description",
                Content: "I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="post-project-user-data" method="post" path="/projects/{project_id}/user_data" example="Not Found" -->
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

    res, err := s.UserData.Create(ctx, "123", operations.PostProjectUserDataUserDataRequestBody{
        Data: operations.PostProjectUserDataUserDataData{
            Type: operations.PostProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PostProjectUserDataUserDataAttributes{
                Description: "User Data description",
                Content: "#cloud-config\\nruncmd:\\n - [ touch,  /home/ubuntu/test ]",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when content is not base64 encoded

<!-- UsageSnippet language="go" operationID="post-project-user-data" method="post" path="/projects/{project_id}/user_data" example="when content is not base64 encoded" -->
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

    res, err := s.UserData.Create(ctx, "proj_LYV8DZK3O5QoE", operations.PostProjectUserDataUserDataRequestBody{
        Data: operations.PostProjectUserDataUserDataData{
            Type: operations.PostProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PostProjectUserDataUserDataAttributes{
                Description: "User Data description",
                Content: "I am not base64 encoded",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when content is not present

<!-- UsageSnippet language="go" operationID="post-project-user-data" method="post" path="/projects/{project_id}/user_data" example="when content is not present" -->
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

    res, err := s.UserData.Create(ctx, "proj_LA73qkVodaJ2o", operations.PostProjectUserDataUserDataRequestBody{
        Data: operations.PostProjectUserDataUserDataData{
            Type: operations.PostProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PostProjectUserDataUserDataAttributes{
                Description: "User Data description",
                Content: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when description is not present

<!-- UsageSnippet language="go" operationID="post-project-user-data" method="post" path="/projects/{project_id}/user_data" example="when description is not present" -->
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

    res, err := s.UserData.Create(ctx, "proj_pbV0Dgbzq4AWz", operations.PostProjectUserDataUserDataRequestBody{
        Data: operations.PostProjectUserDataUserDataData{
            Type: operations.PostProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PostProjectUserDataUserDataAttributes{
                Description: "duh drat ick fabricate from yum hovel",
                Content: "I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `projectID`                                                                                                            | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Project ID or Slug                                                                                                     |
| `requestBody`                                                                                                          | [operations.PostProjectUserDataUserDataRequestBody](../../models/operations/postprojectuserdatauserdatarequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.PostProjectUserDataResponse](../../models/operations/postprojectuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~UpdateForProject~~

Allow you update User Data in a project.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage: Bad Request

<!-- UsageSnippet language="go" operationID="put-project-user-data" method="patch" path="/projects/{project_id}/user_data/{user_data_id}" example="Bad Request" -->
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

    res, err := s.UserData.UpdateForProject(ctx, "proj_lxWpD62zDm6rk", "ud_pwg3ZDr3O5QlP", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "ud_pwg3ZDr3O5QlP",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Content: latitudeshgosdk.Pointer("New User Data content"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="put-project-user-data" method="patch" path="/projects/{project_id}/user_data/{user_data_id}" example="Not Found" -->
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

    res, err := s.UserData.UpdateForProject(ctx, "proj_695BdKNeOevVo", "invalid", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "invalid",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Content: latitudeshgosdk.Pointer("New User Content"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when both content and description are present

<!-- UsageSnippet language="go" operationID="put-project-user-data" method="patch" path="/projects/{project_id}/user_data/{user_data_id}" example="when both content and description are present" -->
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

    res, err := s.UserData.UpdateForProject(ctx, "proj_8NkvdyPXOeLpx", "ud_Qk0Ryqv1dW36X", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "ud_Qk0Ryqv1dW36X",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Description: latitudeshgosdk.Pointer("New Name"),
                Content: latitudeshgosdk.Pointer("I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when content is not present

<!-- UsageSnippet language="go" operationID="put-project-user-data" method="patch" path="/projects/{project_id}/user_data/{user_data_id}" example="when content is not present" -->
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

    res, err := s.UserData.UpdateForProject(ctx, "proj_k0Ryqv2RqW36X", "ud_m1R3zq2bqWxyn", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "ud_m1R3zq2bqWxyn",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Content: latitudeshgosdk.Pointer("<value>"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when description is not present

<!-- UsageSnippet language="go" operationID="put-project-user-data" method="patch" path="/projects/{project_id}/user_data/{user_data_id}" example="when description is not present" -->
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

    res, err := s.UserData.UpdateForProject(ctx, "proj_yQrJdNm5O30gv", "ud_aKXgRdR3qv9k5", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "ud_aKXgRdR3qv9k5",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Description: latitudeshgosdk.Pointer("toaster petal creative flame"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when only content is present

<!-- UsageSnippet language="go" operationID="put-project-user-data" method="patch" path="/projects/{project_id}/user_data/{user_data_id}" example="when only content is present" -->
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

    res, err := s.UserData.UpdateForProject(ctx, "proj_e8pKq0aKDWAob", "ud_2695BdKrOevVo", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "ud_2695BdKrOevVo",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Content: latitudeshgosdk.Pointer("I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when only description is present

<!-- UsageSnippet language="go" operationID="put-project-user-data" method="patch" path="/projects/{project_id}/user_data/{user_data_id}" example="when only description is present" -->
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

    res, err := s.UserData.UpdateForProject(ctx, "proj_QraYDPA3OpjwW", "ud_byQrJdNJd30gv", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "ud_byQrJdNJd30gv",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Description: latitudeshgosdk.Pointer("New Name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                             | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                 | :heavy_check_mark:                                                                                                    | The context to use for the request.                                                                                   |
| `projectID`                                                                                                           | *string*                                                                                                              | :heavy_check_mark:                                                                                                    | Project ID or Slug                                                                                                    |
| `userDataID`                                                                                                          | *string*                                                                                                              | :heavy_check_mark:                                                                                                    | N/A                                                                                                                   |
| `requestBody`                                                                                                         | [*operations.PutProjectUserDataUserDataRequestBody](../../models/operations/putprojectuserdatauserdatarequestbody.md) | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `opts`                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                              | :heavy_minus_sign:                                                                                                    | The options for this request.                                                                                         |

### Response

**[*operations.PutProjectUserDataResponse](../../models/operations/putprojectuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## List

List all Users Data in the project. These scripts can be used to configure servers with user data.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-users-data" method="get" path="/user_data" example="Success" -->
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

    res, err := s.UserData.List(ctx, nil, nil, latitudeshgosdk.Pointer("decoded_content"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UserData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `filterProject`                                                                             | **string*                                                                                   | :heavy_minus_sign:                                                                          | Project ID or slug                                                                          |
| `filterScope`                                                                               | **string*                                                                                   | :heavy_minus_sign:                                                                          | Filter by scope: `project` (has project), `team` (no project), or empty (all)               |
| `extraFieldsUserData`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The `decoded_content` is provided as an extra attribute that shows content in decoded form. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetUsersDataResponse](../../models/operations/getusersdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## CreateNew

Allows you to create User Data in a team, which can be used to perform custom setup on your servers after deploy and reinstall.


### Example Usage

<!-- UsageSnippet language="go" operationID="post-user-data" method="post" path="/user_data" example="Created" -->
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

    res, err := s.UserData.CreateNew(ctx, operations.PostUserDataUserDataRequestBody{
        Data: operations.PostUserDataUserDataData{
            Type: operations.PostUserDataUserDataTypeUserData,
            Attributes: &operations.PostUserDataUserDataAttributes{
                Description: "User Data description",
                Project: latitudeshgosdk.Pointer("proj_QraYDPW3qpjwW"),
                Content: "I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostUserDataUserDataRequestBody](../../models/operations/postuserdatauserdatarequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.PostUserDataResponse](../../models/operations/postuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Retrieve

Get User Data in the project. These scripts can be used to configure servers with user data.


### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="get-user-data" method="get" path="/user_data/{user_data_id}" example="Not Found" -->
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

    res, err := s.UserData.Retrieve(ctx, "123", latitudeshgosdk.Pointer("decoded_content"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="get-user-data" method="get" path="/user_data/{user_data_id}" example="Success" -->
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

    res, err := s.UserData.Retrieve(ctx, "ud_MLGXPdWgdnNWk", latitudeshgosdk.Pointer("decoded_content"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `userDataID`                                                                                | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `extraFieldsUserData`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The `decoded_content` is provided as an extra attribute that shows content in decoded form. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetUserDataResponse](../../models/operations/getuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Allow you update User Data in a team.


### Example Usage: when both content and description are present

<!-- UsageSnippet language="go" operationID="patch-user-data" method="patch" path="/user_data/{user_data_id}" example="when both content and description are present" -->
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

    res, err := s.UserData.Update(ctx, "ud_lpbV0DgRq4AWz", &operations.PatchUserDataUserDataRequestBody{
        Data: operations.PatchUserDataUserDataData{
            ID: "ud_lpbV0DgRq4AWz",
            Type: operations.PatchUserDataUserDataTypeUserData,
            Attributes: &operations.PatchUserDataUserDataAttributes{
                Description: latitudeshgosdk.Pointer("New Name"),
                Content: latitudeshgosdk.Pointer("I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when only content is present

<!-- UsageSnippet language="go" operationID="patch-user-data" method="patch" path="/user_data/{user_data_id}" example="when only content is present" -->
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

    res, err := s.UserData.Update(ctx, "ud_Av9BVDavORm1W", &operations.PatchUserDataUserDataRequestBody{
        Data: operations.PatchUserDataUserDataData{
            ID: "ud_Av9BVDavORm1W",
            Type: operations.PatchUserDataUserDataTypeUserData,
            Attributes: &operations.PatchUserDataUserDataAttributes{
                Content: latitudeshgosdk.Pointer("I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```
### Example Usage: when only description is present

<!-- UsageSnippet language="go" operationID="patch-user-data" method="patch" path="/user_data/{user_data_id}" example="when only description is present" -->
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

    res, err := s.UserData.Update(ctx, "ud_5LA73qkjdaJ2o", &operations.PatchUserDataUserDataRequestBody{
        Data: operations.PatchUserDataUserDataData{
            ID: "ud_5LA73qkjdaJ2o",
            Type: operations.PatchUserDataUserDataTypeUserData,
            Attributes: &operations.PatchUserDataUserDataAttributes{
                Description: latitudeshgosdk.Pointer("New Name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserDataObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `userDataID`                                                                                                | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `requestBody`                                                                                               | [*operations.PatchUserDataUserDataRequestBody](../../models/operations/patchuserdatauserdatarequestbody.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.PatchUserDataResponse](../../models/operations/patchuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Delete user data

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-user-data" method="delete" path="/user_data/{user_data_id}" -->
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

    res, err := s.UserData.Delete(ctx, "123")
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
| `userDataID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserDataResponse](../../models/operations/deleteuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |