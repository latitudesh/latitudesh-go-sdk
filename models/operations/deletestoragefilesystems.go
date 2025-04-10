// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type DeleteStorageFilesystemsRequest struct {
	FilesystemID string `pathParam:"style=simple,explode=false,name=filesystem_id"`
}

func (o *DeleteStorageFilesystemsRequest) GetFilesystemID() string {
	if o == nil {
		return ""
	}
	return o.FilesystemID
}

type DeleteStorageFilesystemsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteStorageFilesystemsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
