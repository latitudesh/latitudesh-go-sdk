// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type UpdateAPIKeyRequest struct {
	APIKeyID     string                  `pathParam:"style=simple,explode=false,name=api_key_id"`
	UpdateAPIKey components.UpdateAPIKey `request:"mediaType=application/json"`
}

func (o *UpdateAPIKeyRequest) GetAPIKeyID() string {
	if o == nil {
		return ""
	}
	return o.APIKeyID
}

func (o *UpdateAPIKeyRequest) GetUpdateAPIKey() components.UpdateAPIKey {
	if o == nil {
		return components.UpdateAPIKey{}
	}
	return o.UpdateAPIKey
}

// UpdateAPIKeyResponseBody - API Key Updated
type UpdateAPIKeyResponseBody struct {
	Data *components.APIKey `json:"data,omitempty"`
}

func (o *UpdateAPIKeyResponseBody) GetData() *components.APIKey {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateAPIKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// API Key Updated
	Object *UpdateAPIKeyResponseBody
}

func (o *UpdateAPIKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateAPIKeyResponse) GetObject() *UpdateAPIKeyResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
