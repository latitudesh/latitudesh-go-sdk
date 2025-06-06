// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type DestroyTeamMemberRequest struct {
	// The user ID
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (o *DestroyTeamMemberRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type DestroyTeamMemberResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DestroyTeamMemberResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
