# AccessScope

`fullaccess` grants access to all of the project's buckets. `limited_access` restricts the key to the buckets listed in `bucket_permissions`.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/operations"
)

value := operations.AccessScopeFullaccess
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `AccessScopeFullaccess`    | fullaccess                 |
| `AccessScopeLimitedAccess` | limited_access             |