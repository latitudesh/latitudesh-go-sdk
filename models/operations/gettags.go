// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type GetTagsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	CustomTag *components.CustomTag
}

func (o *GetTagsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTagsResponse) GetCustomTag() *components.CustomTag {
	if o == nil {
		return nil
	}
	return o.CustomTag
}
