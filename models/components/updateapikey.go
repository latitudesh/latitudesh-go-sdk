// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/latitudesh/latitudesh-go-sdk/internal/utils"
)

type UpdateAPIKeyType string

const (
	UpdateAPIKeyTypeAPIKeys UpdateAPIKeyType = "api_keys"
)

func (e UpdateAPIKeyType) ToPointer() *UpdateAPIKeyType {
	return &e
}
func (e *UpdateAPIKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_keys":
		*e = UpdateAPIKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateAPIKeyType: %v", v)
	}
}

type UpdateAPIKeyAttributes struct {
	// Name of the API Key
	Name *string `default:"Name of the API Key" json:"name"`
	// The API version to use
	APIVersion *string `default:"2023-06-01" json:"api_version"`
}

func (u UpdateAPIKeyAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateAPIKeyAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAPIKeyAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateAPIKeyAttributes) GetAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.APIVersion
}

type UpdateAPIKeyData struct {
	ID         *string                 `json:"id,omitempty"`
	Type       UpdateAPIKeyType        `json:"type"`
	Attributes *UpdateAPIKeyAttributes `json:"attributes,omitempty"`
}

func (o *UpdateAPIKeyData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateAPIKeyData) GetType() UpdateAPIKeyType {
	if o == nil {
		return UpdateAPIKeyType("")
	}
	return o.Type
}

func (o *UpdateAPIKeyData) GetAttributes() *UpdateAPIKeyAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

type UpdateAPIKey struct {
	Data *UpdateAPIKeyData `json:"data,omitempty"`
}

func (o *UpdateAPIKey) GetData() *UpdateAPIKeyData {
	if o == nil {
		return nil
	}
	return o.Data
}
