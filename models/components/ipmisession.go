// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type IpmiSessionType string

const (
	IpmiSessionTypeIpmiSessions IpmiSessionType = "ipmi_sessions"
)

func (e IpmiSessionType) ToPointer() *IpmiSessionType {
	return &e
}
func (e *IpmiSessionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ipmi_sessions":
		*e = IpmiSessionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IpmiSessionType: %v", v)
	}
}

type IpmiSessionAttributes struct {
	// The IPMI IP Address
	IpmiAddress *string `json:"ipmi_address,omitempty"`
	// The IPMI username
	IpmiUsername *string `json:"ipmi_username,omitempty"`
	// The IPMI password
	IpmiPassword *string `json:"ipmi_password,omitempty"`
}

func (o *IpmiSessionAttributes) GetIpmiAddress() *string {
	if o == nil {
		return nil
	}
	return o.IpmiAddress
}

func (o *IpmiSessionAttributes) GetIpmiUsername() *string {
	if o == nil {
		return nil
	}
	return o.IpmiUsername
}

func (o *IpmiSessionAttributes) GetIpmiPassword() *string {
	if o == nil {
		return nil
	}
	return o.IpmiPassword
}

type IpmiSessionData struct {
	ID         *string                `json:"id,omitempty"`
	Type       *IpmiSessionType       `json:"type,omitempty"`
	Attributes *IpmiSessionAttributes `json:"attributes,omitempty"`
}

func (o *IpmiSessionData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IpmiSessionData) GetType() *IpmiSessionType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *IpmiSessionData) GetAttributes() *IpmiSessionAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

type IpmiSession struct {
	Data *IpmiSessionData `json:"data,omitempty"`
}

func (o *IpmiSession) GetData() *IpmiSessionData {
	if o == nil {
		return nil
	}
	return o.Data
}
