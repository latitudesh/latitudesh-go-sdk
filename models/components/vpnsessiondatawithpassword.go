// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type VpnSessionDataWithPasswordType string

const (
	VpnSessionDataWithPasswordTypeVpnSessions VpnSessionDataWithPasswordType = "vpn_sessions"
)

func (e VpnSessionDataWithPasswordType) ToPointer() *VpnSessionDataWithPasswordType {
	return &e
}
func (e *VpnSessionDataWithPasswordType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vpn_sessions":
		*e = VpnSessionDataWithPasswordType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VpnSessionDataWithPasswordType: %v", v)
	}
}

// VpnSessionDataWithPasswordStatus - from Firewall Response
type VpnSessionDataWithPasswordStatus string

const (
	VpnSessionDataWithPasswordStatusEnable  VpnSessionDataWithPasswordStatus = "enable"
	VpnSessionDataWithPasswordStatusDisable VpnSessionDataWithPasswordStatus = "disable"
)

func (e VpnSessionDataWithPasswordStatus) ToPointer() *VpnSessionDataWithPasswordStatus {
	return &e
}
func (e *VpnSessionDataWithPasswordStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enable":
		fallthrough
	case "disable":
		*e = VpnSessionDataWithPasswordStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VpnSessionDataWithPasswordStatus: %v", v)
	}
}

type VpnSessionDataWithPasswordAttributes struct {
	// VPN username
	UserName *string `json:"user_name,omitempty"`
	// VPN password
	Password *string `json:"password,omitempty"`
	// VPN port
	Port *string `json:"port,omitempty"`
	// VPN host
	Host   *string             `json:"host,omitempty"`
	Region *RegionResourceData `json:"region,omitempty"`
	// from Firewall Response
	Status *VpnSessionDataWithPasswordStatus `json:"status,omitempty"`
	// Time to expiry
	ExpiresAt *string `json:"expires_at,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o *VpnSessionDataWithPasswordAttributes) GetUserName() *string {
	if o == nil {
		return nil
	}
	return o.UserName
}

func (o *VpnSessionDataWithPasswordAttributes) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *VpnSessionDataWithPasswordAttributes) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *VpnSessionDataWithPasswordAttributes) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *VpnSessionDataWithPasswordAttributes) GetRegion() *RegionResourceData {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *VpnSessionDataWithPasswordAttributes) GetStatus() *VpnSessionDataWithPasswordStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *VpnSessionDataWithPasswordAttributes) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *VpnSessionDataWithPasswordAttributes) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *VpnSessionDataWithPasswordAttributes) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type VpnSessionDataWithPassword struct {
	ID         *string                               `json:"id,omitempty"`
	Type       *VpnSessionDataWithPasswordType       `json:"type,omitempty"`
	Attributes *VpnSessionDataWithPasswordAttributes `json:"attributes,omitempty"`
}

func (o *VpnSessionDataWithPassword) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *VpnSessionDataWithPassword) GetType() *VpnSessionDataWithPasswordType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *VpnSessionDataWithPassword) GetAttributes() *VpnSessionDataWithPasswordAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}
