// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type FirewallAssignmentDataType string

const (
	FirewallAssignmentDataTypeFirewallAssignments FirewallAssignmentDataType = "firewall_assignments"
)

func (e FirewallAssignmentDataType) ToPointer() *FirewallAssignmentDataType {
	return &e
}
func (e *FirewallAssignmentDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "firewall_assignments":
		*e = FirewallAssignmentDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FirewallAssignmentDataType: %v", v)
	}
}

type FirewallAssignmentDataServer struct {
	ID          *string `json:"id,omitempty"`
	PrimaryIpv4 *string `json:"primary_ipv4,omitempty"`
	Hostname    *string `json:"hostname,omitempty"`
}

func (o *FirewallAssignmentDataServer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FirewallAssignmentDataServer) GetPrimaryIpv4() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryIpv4
}

func (o *FirewallAssignmentDataServer) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

type FirewallAssignmentDataAttributes struct {
	Server     *FirewallAssignmentDataServer `json:"server,omitempty"`
	FirewallID *string                       `json:"firewall_id,omitempty"`
}

func (o *FirewallAssignmentDataAttributes) GetServer() *FirewallAssignmentDataServer {
	if o == nil {
		return nil
	}
	return o.Server
}

func (o *FirewallAssignmentDataAttributes) GetFirewallID() *string {
	if o == nil {
		return nil
	}
	return o.FirewallID
}

type FirewallAssignmentData struct {
	ID         *string                           `json:"id,omitempty"`
	Type       *FirewallAssignmentDataType       `json:"type,omitempty"`
	Attributes *FirewallAssignmentDataAttributes `json:"attributes,omitempty"`
}

func (o *FirewallAssignmentData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FirewallAssignmentData) GetType() *FirewallAssignmentDataType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *FirewallAssignmentData) GetAttributes() *FirewallAssignmentDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}
