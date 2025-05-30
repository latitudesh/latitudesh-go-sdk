// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/latitudesh/latitudesh-go-sdk/internal/utils"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type CreateTagTagsType string

const (
	CreateTagTagsTypeTags CreateTagTagsType = "tags"
)

func (e CreateTagTagsType) ToPointer() *CreateTagTagsType {
	return &e
}
func (e *CreateTagTagsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tags":
		*e = CreateTagTagsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTagTagsType: %v", v)
	}
}

type CreateTagTagsAttributes struct {
	// Name of the Tag
	Name *string `json:"name,omitempty"`
	// Description of the Tag
	Description *string `json:"description,omitempty"`
	// Color of the Tag
	Color *string `default:"#ffffff" json:"color"`
}

func (c CreateTagTagsAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateTagTagsAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateTagTagsAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateTagTagsAttributes) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateTagTagsAttributes) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

type CreateTagTagsData struct {
	Type       *CreateTagTagsType       `json:"type,omitempty"`
	Attributes *CreateTagTagsAttributes `json:"attributes,omitempty"`
}

func (o *CreateTagTagsData) GetType() *CreateTagTagsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CreateTagTagsData) GetAttributes() *CreateTagTagsAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

type CreateTagTagsRequestBody struct {
	Data *CreateTagTagsData `json:"data,omitempty"`
}

func (o *CreateTagTagsRequestBody) GetData() *CreateTagTagsData {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateTagResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Created
	CustomTag *components.CustomTag
}

func (o *CreateTagResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTagResponse) GetCustomTag() *components.CustomTag {
	if o == nil {
		return nil
	}
	return o.CustomTag
}
