# FilesystemAttachmentDataStatus

Whether the server is attached to the filesystem storage network. "attaching"/"detaching" while an attach or detach runs, "failed" when it gave up.

## Example Usage

```go
import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

value := components.FilesystemAttachmentDataStatusAttaching
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `FilesystemAttachmentDataStatusAttaching` | attaching                                 |
| `FilesystemAttachmentDataStatusAttached`  | attached                                  |
| `FilesystemAttachmentDataStatusDetaching` | detaching                                 |
| `FilesystemAttachmentDataStatusDetached`  | detached                                  |
| `FilesystemAttachmentDataStatusFailed`    | failed                                    |