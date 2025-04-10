// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Country struct {
	Slug *string `json:"slug,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *Country) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Country) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type RegionsAttributes struct {
	Slug    *string  `json:"slug,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Country *Country `json:"country,omitempty"`
}

func (o *RegionsAttributes) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *RegionsAttributes) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RegionsAttributes) GetCountry() *Country {
	if o == nil {
		return nil
	}
	return o.Country
}

type RegionsData struct {
	ID         *string            `json:"id,omitempty"`
	Attributes *RegionsAttributes `json:"attributes,omitempty"`
}

func (o *RegionsData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RegionsData) GetAttributes() *RegionsAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

type Regions struct {
	Data []RegionsData `json:"data,omitempty"`
}

func (o *Regions) GetData() []RegionsData {
	if o == nil {
		return nil
	}
	return o.Data
}
