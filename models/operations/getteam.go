// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type GetTeamResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Teams *components.Teams
}

func (o *GetTeamResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTeamResponse) GetTeams() *components.Teams {
	if o == nil {
		return nil
	}
	return o.Teams
}
