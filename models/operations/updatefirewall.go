// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/latitudesh/latitudesh-go-sdk/models/components"
)

type UpdateFirewallFirewallsType string

const (
	UpdateFirewallFirewallsTypeFirewalls UpdateFirewallFirewallsType = "firewalls"
)

func (e UpdateFirewallFirewallsType) ToPointer() *UpdateFirewallFirewallsType {
	return &e
}
func (e *UpdateFirewallFirewallsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "firewalls":
		*e = UpdateFirewallFirewallsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateFirewallFirewallsType: %v", v)
	}
}

type UpdateFirewallFirewallsProtocol string

const (
	UpdateFirewallFirewallsProtocolTCP UpdateFirewallFirewallsProtocol = "TCP"
	UpdateFirewallFirewallsProtocolUDP UpdateFirewallFirewallsProtocol = "UDP"
)

func (e UpdateFirewallFirewallsProtocol) ToPointer() *UpdateFirewallFirewallsProtocol {
	return &e
}
func (e *UpdateFirewallFirewallsProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TCP":
		fallthrough
	case "UDP":
		*e = UpdateFirewallFirewallsProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateFirewallFirewallsProtocol: %v", v)
	}
}

type UpdateFirewallFirewallsRules struct {
	From     *string                          `json:"from,omitempty"`
	To       *string                          `json:"to,omitempty"`
	Protocol *UpdateFirewallFirewallsProtocol `json:"protocol,omitempty"`
	// Port number or range (e.g., "80", "80-443")
	Port *string `json:"port,omitempty"`
}

func (o *UpdateFirewallFirewallsRules) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *UpdateFirewallFirewallsRules) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *UpdateFirewallFirewallsRules) GetProtocol() *UpdateFirewallFirewallsProtocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *UpdateFirewallFirewallsRules) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

type UpdateFirewallFirewallsAttributes struct {
	Name  *string                        `json:"name,omitempty"`
	Rules []UpdateFirewallFirewallsRules `json:"rules,omitempty"`
}

func (o *UpdateFirewallFirewallsAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateFirewallFirewallsAttributes) GetRules() []UpdateFirewallFirewallsRules {
	if o == nil {
		return nil
	}
	return o.Rules
}

type UpdateFirewallFirewallsData struct {
	Type       UpdateFirewallFirewallsType        `json:"type"`
	Attributes *UpdateFirewallFirewallsAttributes `json:"attributes,omitempty"`
}

func (o *UpdateFirewallFirewallsData) GetType() UpdateFirewallFirewallsType {
	if o == nil {
		return UpdateFirewallFirewallsType("")
	}
	return o.Type
}

func (o *UpdateFirewallFirewallsData) GetAttributes() *UpdateFirewallFirewallsAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

type UpdateFirewallFirewallsRequestBody struct {
	Data UpdateFirewallFirewallsData `json:"data"`
}

func (o *UpdateFirewallFirewallsRequestBody) GetData() UpdateFirewallFirewallsData {
	if o == nil {
		return UpdateFirewallFirewallsData{}
	}
	return o.Data
}

type UpdateFirewallRequest struct {
	// The Firewall ID
	FirewallID  string                             `pathParam:"style=simple,explode=false,name=firewall_id"`
	RequestBody UpdateFirewallFirewallsRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateFirewallRequest) GetFirewallID() string {
	if o == nil {
		return ""
	}
	return o.FirewallID
}

func (o *UpdateFirewallRequest) GetRequestBody() UpdateFirewallFirewallsRequestBody {
	if o == nil {
		return UpdateFirewallFirewallsRequestBody{}
	}
	return o.RequestBody
}

type UpdateFirewallResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Firewall *components.Firewall
}

func (o *UpdateFirewallResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateFirewallResponse) GetFirewall() *components.Firewall {
	if o == nil {
		return nil
	}
	return o.Firewall
}
