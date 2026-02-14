# Tags

## Overview

### Available Operations

* [List](#list) - List tags
* [Create](#create) - Create tag
* [Update](#update) - Update tag
* [Delete](#delete) - Delete tag

## List

List all Tags in the team.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-tags" method="get" path="/tags" example="Success" -->
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

    res, err := s.Tags.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTags != nil {
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

**[*operations.GetTagsResponse](../../models/operations/gettagsresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Create

Create a Tag in the team.


### Example Usage: Created

<!-- UsageSnippet language="go" operationID="create-tag" method="post" path="/tags" example="Created" -->
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

    res, err := s.Tags.Create(ctx, operations.CreateTagTagsRequestBody{
        Data: &operations.CreateTagTagsData{
            Type: operations.CreateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.CreateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Tag Name"),
                Description: latitudeshgosdk.Pointer("Tag Description"),
                Color: latitudeshgosdk.Pointer("#bebebe"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when any attribute is missing or not filled

<!-- UsageSnippet language="go" operationID="create-tag" method="post" path="/tags" example="when any attribute is missing or not filled" -->
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

    res, err := s.Tags.Create(ctx, operations.CreateTagTagsRequestBody{
        Data: &operations.CreateTagTagsData{
            Type: operations.CreateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.CreateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Tag Name"),
                Description: latitudeshgosdk.Pointer(""),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when color is in wrong format

<!-- UsageSnippet language="go" operationID="create-tag" method="post" path="/tags" example="when color is in wrong format" -->
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

    res, err := s.Tags.Create(ctx, operations.CreateTagTagsRequestBody{
        Data: &operations.CreateTagTagsData{
            Type: operations.CreateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.CreateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Tag Name"),
                Description: latitudeshgosdk.Pointer("Tag Description"),
                Color: latitudeshgosdk.Pointer("bebebe"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when description is too long

<!-- UsageSnippet language="go" operationID="create-tag" method="post" path="/tags" example="when description is too long" -->
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

    res, err := s.Tags.Create(ctx, operations.CreateTagTagsRequestBody{
        Data: &operations.CreateTagTagsData{
            Type: operations.CreateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.CreateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Tag Name"),
                Description: latitudeshgosdk.Pointer("lj553vy8r5j831s7dp78fscnq4z0iv9jjwtkeudrpsdsefghd2w3cd6b5qb3knpq6bin7uurf5qj1ya4xd2yhyzv6o4krqirha1aoubs0nvc04jd21hdp9etq6g3a4o7vb8ol0avqc2j4hlbw5o6yqxa2vsm9jhyf5kt9wy78gxr7jlaol3bxe18xau5fcff3b9qjmy14b2nw5tjynjefh1kjdmfmsn0wu1tg32mr563ndefj3y24j0cyyrlbl7b"),
                Color: latitudeshgosdk.Pointer("#bebebe"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when name is already taken

<!-- UsageSnippet language="go" operationID="create-tag" method="post" path="/tags" example="when name is already taken" -->
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

    res, err := s.Tags.Create(ctx, operations.CreateTagTagsRequestBody{
        Data: &operations.CreateTagTagsData{
            Type: operations.CreateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.CreateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Argon"),
                Description: latitudeshgosdk.Pointer("Tag Description"),
                Color: latitudeshgosdk.Pointer("#bebebe"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateTagTagsRequestBody](../../models/operations/createtagtagsrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateTagResponse](../../models/operations/createtagresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Update

Update a Tag in the team.


### Example Usage: Not Found

<!-- UsageSnippet language="go" operationID="update-tag" method="patch" path="/tags/{tag_id}" example="Not Found" -->
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

    res, err := s.Tags.Update(ctx, "invalid", operations.UpdateTagTagsRequestBody{
        Data: &operations.UpdateTagTagsData{
            ID: latitudeshgosdk.Pointer("invalid"),
            Type: operations.UpdateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.UpdateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Tag Name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: Success

<!-- UsageSnippet language="go" operationID="update-tag" method="patch" path="/tags/{tag_id}" example="Success" -->
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

    res, err := s.Tags.Update(ctx, "tag_XBlke2r5RyiyVpG9LPK8tWjalLL", operations.UpdateTagTagsRequestBody{
        Data: &operations.UpdateTagTagsData{
            ID: latitudeshgosdk.Pointer("tag_XBlke2r5RyiyVpG9LPK8tWjalLL"),
            Type: operations.UpdateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.UpdateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Tag Name"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when any attribute is blank

<!-- UsageSnippet language="go" operationID="update-tag" method="patch" path="/tags/{tag_id}" example="when any attribute is blank" -->
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

    res, err := s.Tags.Update(ctx, "tag_X0oz3V4bB4Cy94Wm3G0MfWga24B", operations.UpdateTagTagsRequestBody{
        Data: &operations.UpdateTagTagsData{
            ID: latitudeshgosdk.Pointer("tag_X0oz3V4bB4Cy94Wm3G0MfWga24B"),
            Type: operations.UpdateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.UpdateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer(""),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when color is invalid

<!-- UsageSnippet language="go" operationID="update-tag" method="patch" path="/tags/{tag_id}" example="when color is invalid" -->
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

    res, err := s.Tags.Update(ctx, "tag_pppwBaKxmacw7eJmMEJ2tj8o7Lb", operations.UpdateTagTagsRequestBody{
        Data: &operations.UpdateTagTagsData{
            ID: latitudeshgosdk.Pointer("tag_pppwBaKxmacw7eJmMEJ2tj8o7Lb"),
            Type: operations.UpdateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.UpdateTagTagsAttributes{
                Color: latitudeshgosdk.Pointer("bebebe"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when description is too long

<!-- UsageSnippet language="go" operationID="update-tag" method="patch" path="/tags/{tag_id}" example="when description is too long" -->
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

    res, err := s.Tags.Update(ctx, "tag_K5G3oWwnPbfjVJ2v4V3xSbEMex0", operations.UpdateTagTagsRequestBody{
        Data: &operations.UpdateTagTagsData{
            ID: latitudeshgosdk.Pointer("tag_K5G3oWwnPbfjVJ2v4V3xSbEMex0"),
            Type: operations.UpdateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.UpdateTagTagsAttributes{
                Description: latitudeshgosdk.Pointer("sdfil1bdxt8hoe5tf0q54af15q2984xhaxmyqzkqr945acgllrmy1h7nrqy70lvz3lfiqla2on8ulx12949f6ffejxog1x5hzj5ec2eqkx1keeabd5k4b4jrfa0yzpii9a04xevll2r1530u2yzjexvqku7budmlmrp5y5o2ypxpds5wh2o69hjkpjw7fapf1lafjdibw6f5xd8n730qjh40eh9rqujovey0xovs7rn6b4w3qbjaxac48fcvr23e"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```
### Example Usage: when name is already taken

<!-- UsageSnippet language="go" operationID="update-tag" method="patch" path="/tags/{tag_id}" example="when name is already taken" -->
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

    res, err := s.Tags.Update(ctx, "tag_lPZ90rpX3PcJ8EYlyjrWCE9AyEk", operations.UpdateTagTagsRequestBody{
        Data: &operations.UpdateTagTagsData{
            ID: latitudeshgosdk.Pointer("tag_lPZ90rpX3PcJ8EYlyjrWCE9AyEk"),
            Type: operations.UpdateTagTagsTypeTags.ToPointer(),
            Attributes: &operations.UpdateTagTagsAttributes{
                Name: latitudeshgosdk.Pointer("Tag Name 2"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomTag != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `tagID`                                                                                    | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `requestBody`                                                                              | [operations.UpdateTagTagsRequestBody](../../models/operations/updatetagtagsrequestbody.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateTagResponse](../../models/operations/updatetagresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |

## Delete

Update a Tag in the team.


### Example Usage

<!-- UsageSnippet language="go" operationID="destroy-tag" method="delete" path="/tags/{tag_id}" -->
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

    res, err := s.Tags.Delete(ctx, "invalid-id")
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
| `tagID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DestroyTagResponse](../../models/operations/destroytagresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| components.APIError | 4XX, 5XX            | \*/\*               |