// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type GetFirewallAssignmentsRequest struct {
	// The Firewall ID
	FirewallID string `pathParam:"style=simple,explode=false,name=firewall_id"`
}

func (o *GetFirewallAssignmentsRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

type GetFirewallAssignmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Ok
	FirewallServer *components.FirewallServer
}

func (o *GetFirewallAssignmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetFirewallAssignmentsResponse) GetFirewallServer() *components.FirewallServer {
	if o == nil {
		return nil
	}
	return o.FirewallServer
}
