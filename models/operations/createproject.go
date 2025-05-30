// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type CreateProjectProjectsType string

const (
	CreateProjectProjectsTypeProjects CreateProjectProjectsType = "projects"
)

func (e CreateProjectProjectsType) ToPointer() *CreateProjectProjectsType {
	return &e
}
func (e *CreateProjectProjectsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "projects":
		*e = CreateProjectProjectsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProjectProjectsType: %v", v)
	}
}

// CreateProjectProvisioningType - The provisioning type of the project. Default: on_demand
type CreateProjectProvisioningType string

const (
	CreateProjectProvisioningTypeReserved CreateProjectProvisioningType = "reserved"
	CreateProjectProvisioningTypeOnDemand CreateProjectProvisioningType = "on_demand"
)

func (e CreateProjectProvisioningType) ToPointer() *CreateProjectProvisioningType {
	return &e
}
func (e *CreateProjectProvisioningType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "reserved":
		fallthrough
	case "on_demand":
		*e = CreateProjectProvisioningType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProjectProvisioningType: %v", v)
	}
}

type CreateProjectEnvironment string

const (
	CreateProjectEnvironmentDevelopment CreateProjectEnvironment = "Development"
	CreateProjectEnvironmentStaging     CreateProjectEnvironment = "Staging"
	CreateProjectEnvironmentProduction  CreateProjectEnvironment = "Production"
)

func (e CreateProjectEnvironment) ToPointer() *CreateProjectEnvironment {
	return &e
}
func (e *CreateProjectEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Development":
		fallthrough
	case "Staging":
		fallthrough
	case "Production":
		*e = CreateProjectEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProjectEnvironment: %v", v)
	}
}

type CreateProjectProjectsAttributes struct {
	// The project name. Must be unique.
	Name string `json:"name"`
	// The provisioning type of the project. Default: on_demand
	ProvisioningType CreateProjectProvisioningType `json:"provisioning_type"`
	// The project description.
	Description *string                   `json:"description,omitempty"`
	Environment *CreateProjectEnvironment `json:"environment,omitempty"`
}

func (o *CreateProjectProjectsAttributes) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateProjectProjectsAttributes) GetProvisioningType() CreateProjectProvisioningType {
	if o == nil {
		return CreateProjectProvisioningType("")
	}
	return o.ProvisioningType
}

func (o *CreateProjectProjectsAttributes) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateProjectProjectsAttributes) GetEnvironment() *CreateProjectEnvironment {
	if o == nil {
		return nil
	}
	return o.Environment
}

type CreateProjectProjectsData struct {
	Type       CreateProjectProjectsType        `json:"type"`
	Attributes *CreateProjectProjectsAttributes `json:"attributes,omitempty"`
}

func (o *CreateProjectProjectsData) GetType() CreateProjectProjectsType {
	if o == nil {
		return CreateProjectProjectsType("")
	}
	return o.Type
}

func (o *CreateProjectProjectsData) GetAttributes() *CreateProjectProjectsAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

type CreateProjectProjectsRequestBody struct {
	Data *CreateProjectProjectsData `json:"data,omitempty"`
}

func (o *CreateProjectProjectsRequestBody) GetData() *CreateProjectProjectsData {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateProjectResponseBody - Success
type CreateProjectResponseBody struct {
	Data *components.Project `json:"data,omitempty"`
}

func (o *CreateProjectResponseBody) GetData() *components.Project {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateProjectResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *CreateProjectResponseBody
}

func (o *CreateProjectResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateProjectResponse) GetObject() *CreateProjectResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
