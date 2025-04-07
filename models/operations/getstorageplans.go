// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type GetStoragePlansResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	StoragePlans *components.StoragePlans
}

func (o *GetStoragePlansResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetStoragePlansResponse) GetStoragePlans() *components.StoragePlans {
	if o == nil {
		return nil
	}
	return o.StoragePlans
}
