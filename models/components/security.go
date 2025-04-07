// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Security struct {
	Bearer *string `security:"scheme,type=apiKey,subtype=header,name=Authorization,env=latitudesh_bearer"`
}

func (o *Security) GetBearer() *string {
	if o == nil {
		return nil
	}
	return o.Bearer
}
