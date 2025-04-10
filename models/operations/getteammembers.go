// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type GetTeamMembersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	TeamMembers *components.TeamMembers
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
