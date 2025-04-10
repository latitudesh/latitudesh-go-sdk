// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

// GetPlansOperatingSystemResponseBody - Success
type GetPlansOperatingSystemResponseBody struct {
	Data []components.OperatingSystems `json:"data,omitempty"`
}

func (o *GetPlansOperatingSystemResponseBody) GetData() []components.OperatingSystems {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetPlansOperatingSystemResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetPlansOperatingSystemResponseBody
}

func (o *GetPlansOperatingSystemResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPlansOperatingSystemResponse) GetObject() *GetPlansOperatingSystemResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
