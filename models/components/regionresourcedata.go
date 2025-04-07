// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Site struct {
	ID       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Slug     *string `json:"slug,omitempty"`
	Facility *string `json:"facility,omitempty"`
}

func (o *Site) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Site) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Site) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Site) GetFacility() *string {
	if o == nil {
		return nil
	}
	return o.Facility
}

type RegionResourceData struct {
	City    *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	Site    *Site   `json:"site,omitempty"`
}

func (o *RegionResourceData) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *RegionResourceData) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *RegionResourceData) GetSite() *Site {
	if o == nil {
		return nil
	}
	return o.Site
}
