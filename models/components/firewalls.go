// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Meta struct {
}

type Firewalls struct {
	Data []Firewall `json:"data,omitempty"`
	Meta *Meta      `json:"meta,omitempty"`
}

func (o *Firewalls) GetData() []Firewall {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Firewalls) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}
