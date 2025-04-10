// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type DestroyTagRequest struct {
	TagID string `pathParam:"style=simple,explode=false,name=tag_id"`
}

func (o *DestroyTagRequest) GetTagID() string {
	if o == nil {
		return ""
	}
	return o.TagID
}

type DestroyTagResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DestroyTagResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
