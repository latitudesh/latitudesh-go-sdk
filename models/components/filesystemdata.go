// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/latitudesh/latitudesh-go-sdk/internal/utils"
	"time"
)

type FilesystemDataType string

const (
	FilesystemDataTypeFilesystems FilesystemDataType = "filesystems"
)

func (e FilesystemDataType) ToPointer() *FilesystemDataType {
	return &e
}
func (e *FilesystemDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "filesystems":
		*e = FilesystemDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FilesystemDataType: %v", v)
	}
}

type FilesystemDataProject struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

func (o *FilesystemDataProject) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FilesystemDataProject) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FilesystemDataProject) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type FilesystemDataAttributes struct {
	Name      *string                `json:"name,omitempty"`
	SizeInGb  *int64                 `json:"size_in_gb,omitempty"`
	Project   *FilesystemDataProject `json:"project,omitempty"`
	CreatedAt *time.Time             `json:"created_at,omitempty"`
}

func (f FilesystemDataAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FilesystemDataAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FilesystemDataAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FilesystemDataAttributes) GetSizeInGb() *int64 {
	if o == nil {
		return nil
	}
	return o.SizeInGb
}

func (o *FilesystemDataAttributes) GetProject() *FilesystemDataProject {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *FilesystemDataAttributes) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

type FilesystemData struct {
	ID         *string                   `json:"id,omitempty"`
	Type       *FilesystemDataType       `json:"type,omitempty"`
	Attributes *FilesystemDataAttributes `json:"attributes,omitempty"`
}

func (o *FilesystemData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FilesystemData) GetType() *FilesystemDataType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *FilesystemData) GetAttributes() *FilesystemDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}
