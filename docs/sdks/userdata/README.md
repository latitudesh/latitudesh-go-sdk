# UserData
(*UserData*)

## Overview

### Available Operations

* [~~List~~](#list) - List all Project User Data :warning: **Deprecated**
* [~~PostProjectUserData~~](#postprojectuserdata) - Create a Project User Data :warning: **Deprecated**
* [~~Get~~](#get) - Retrieve a Project User Data :warning: **Deprecated**
* [~~PutProjectUserData~~](#putprojectuserdata) - Update a Project User Data :warning: **Deprecated**
* [GetUsersData](#getusersdata) - List all User Data
* [PostUserData](#postuserdata) - Create an User Data
* [GetUserData](#getuserdata) - Retrieve an User Data
* [PatchUserData](#patchuserdata) - Update an User Data
* [DeleteUserData](#deleteuserdata) - Delete an User Data

## ~~List~~

List all Users Data in the project. These scripts can be used to configure servers with user data.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.UserData.List(ctx, "proj_z2A3DV4wdnawP", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

## ~~PostProjectUserData~~

Allows you to create User Data in a project, which can be used to perform custom setup on your servers after deploy and reinstall.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.UserData.PostProjectUserData(ctx, "proj_1ZJrdxvyDg4LV", operations.PostProjectUserDataUserDataRequestBody{
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
    if res.UserData != nil {
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

## ~~Get~~

Get User Data in the project. These scripts can be used to configure servers with user data.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.UserData.Get(ctx, "proj_vYAZqG44DMQ94", "ud_lQraYDPeOpjwW", nil)
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
| `userDataID`                                                                                | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `extraFieldsUserData`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The `decoded_content` is provided as an extra attribute that shows content in decoded form. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetProjectUserDataResponse](../../models/operations/getprojectuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## ~~PutProjectUserData~~

Allow you update User Data in a project.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.UserData.PutProjectUserData(ctx, "proj_e8pKq0aKDWAob", "ud_2695BdKrOevVo", &operations.PutProjectUserDataUserDataRequestBody{
        Data: operations.PutProjectUserDataUserDataData{
            ID: "ud_2695BdKrOevVo",
            Type: operations.PutProjectUserDataUserDataTypeUserData,
            Attributes: &operations.PutProjectUserDataUserDataAttributes{
                Content: latitudeshgosdk.String("I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserData != nil {
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

## GetUsersData

List all Users Data in the project. These scripts can be used to configure servers with user data.


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

    res, err := s.UserData.GetUsersData(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `extraFieldsUserData`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The `decoded_content` is provided as an extra attribute that shows content in decoded form. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetUsersDataResponse](../../models/operations/getusersdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## PostUserData

Allows you to create User Data in a team, which can be used to perform custom setup on your servers after deploy and reinstall.


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

    res, err := s.UserData.PostUserData(ctx, operations.PostUserDataUserDataRequestBody{
        Data: operations.PostUserDataUserDataData{
            Type: operations.PostUserDataUserDataTypeUserData,
            Attributes: &operations.PostUserDataUserDataAttributes{
                Description: "User Data description",
                Content: "I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserData != nil {
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

## GetUserData

Get User Data in the project. These scripts can be used to configure servers with user data.


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

    res, err := s.UserData.GetUserData(ctx, "ud_7vYAZqGBdMQ94", nil)
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
| `userDataID`                                                                                | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `extraFieldsUserData`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The `decoded_content` is provided as an extra attribute that shows content in decoded form. |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.GetUserDataResponse](../../models/operations/getuserdataresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## PatchUserData

Allow you update User Data in a team.


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

    res, err := s.UserData.PatchUserData(ctx, "ud_Av9BVDavORm1W", &operations.PatchUserDataUserDataRequestBody{
        Data: operations.PatchUserDataUserDataData{
            ID: "ud_Av9BVDavORm1W",
            Type: operations.PatchUserDataUserDataTypeUserData,
            Attributes: &operations.PatchUserDataUserDataAttributes{
                Content: latitudeshgosdk.String("I2Nsb3VkLWNvbmZpZwpydW5jbWQ6CiAtIFsgdG91Y2gsICAvaG9tZS91YnVudHUvdGVzdCBd"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserData != nil {
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

## DeleteUserData

Delete an User Data

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

    res, err := s.UserData.DeleteUserData(ctx, "123")
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