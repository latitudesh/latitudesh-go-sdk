// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type IPAddresses struct {
	Data []IPAddress `json:"data,omitempty"`
}

func (o *IPAddresses) GetData() []IPAddress {
	if o == nil {
		return nil
	}
	return o.Data
}
