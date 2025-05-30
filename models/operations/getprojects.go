// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/internal/utils"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type GetProjectsRequest struct {
	// The project name to filter by
	FilterName *string `queryParam:"style=form,explode=true,name=filter[name]"`
	// The project slug to filter by
	FilterSlug *string `queryParam:"style=form,explode=true,name=filter[slug]"`
	// The project description to filter by
	FilterDescription *string `queryParam:"style=form,explode=true,name=filter[description]"`
	// The billing type to filter by
	FilterBillingType *string `queryParam:"style=form,explode=true,name=filter[billing_type]"`
	// The environment to filter by
	FilterEnvironment *string `queryParam:"style=form,explode=true,name=filter[environment]"`
	// The tags ids to filter by, separated by comma, e.g. `filter[tags]=tag_1,tag_2`will return projects with `tag_1` AND `tag_2`
	FilterTags *string `queryParam:"style=form,explode=true,name=filter[tags]"`
	// The `last_renewal_date` and `next_renewal_date` are provided as extra attributes that show previous and future billing cycle dates. To request it, just set `extra_fields[projects]=last_renewal_date,next_renewal_date` in the query string.
	ExtraFieldsProjects *string `queryParam:"style=form,explode=true,name=extra_fields[projects]"`
	// Number of items to return per page
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
	// Page number to return (starts at 1)
	PageNumber *int64 `default:"1" queryParam:"style=form,explode=true,name=page[number]"`
}

func (g GetProjectsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetProjectsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetProjectsRequest) GetFilterName() *string {
	if o == nil {
		return nil
	}
	return o.FilterName
}

func (o *GetProjectsRequest) GetFilterSlug() *string {
	if o == nil {
		return nil
	}
	return o.FilterSlug
}

func (o *GetProjectsRequest) GetFilterDescription() *string {
	if o == nil {
		return nil
	}
	return o.FilterDescription
}

func (o *GetProjectsRequest) GetFilterBillingType() *string {
	if o == nil {
		return nil
	}
	return o.FilterBillingType
}

func (o *GetProjectsRequest) GetFilterEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.FilterEnvironment
}

func (o *GetProjectsRequest) GetFilterTags() *string {
	if o == nil {
		return nil
	}
	return o.FilterTags
}

func (o *GetProjectsRequest) GetExtraFieldsProjects() *string {
	if o == nil {
		return nil
	}
	return o.ExtraFieldsProjects
}

func (o *GetProjectsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetProjectsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type GetProjectsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Filtered by multiple tags
	Projects *components.Projects

	Next func() (*GetProjectsResponse, error)
}

func (o *GetProjectsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetProjectsResponse) GetProjects() *components.Projects {
	if o == nil {
		return nil
	}
	return o.Projects
}
