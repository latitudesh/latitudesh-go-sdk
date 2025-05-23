// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type DeleteFirewallAssignmentRequest struct {
	// The Firewall ID
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
	// The Assignment ID
	AssignmentID string `pathParam:"style=simple,explode=false,name=assignment_id"`
}

func (o *DeleteFirewallAssignmentRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *DeleteFirewallAssignmentRequest) GetAssignmentID() string {
	if o == nil {
		return ""
	}
	return o.AssignmentID
}

type DeleteFirewallAssignmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteFirewallAssignmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
