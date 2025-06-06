// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/internal/utils"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type GetTeamMembersRequest struct {
	// Number of items to return per page
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
	// Page number to return (starts at 1)
	PageNumber *int64 `default:"1" queryParam:"style=form,explode=true,name=page[number]"`
}

func (g GetTeamMembersRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTeamMembersRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTeamMembersRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetTeamMembersRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type GetTeamMembersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	TeamMembers *components.TeamMembers

	Next func() (*GetTeamMembersResponse, error)
}

func (o *GetTeamMembersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTeamMembersResponse) GetTeamMembers() *components.TeamMembers {
	if o == nil {
		return nil
	}
	return o.TeamMembers
}
