// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type BandwidthPlanDataType string

const (
	BandwidthPlanDataTypeBandwidthPlan BandwidthPlanDataType = "bandwidth_plan"
)

func (e BandwidthPlanDataType) ToPointer() *BandwidthPlanDataType {
	return &e
}
func (e *BandwidthPlanDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bandwidth_plan":
		*e = BandwidthPlanDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BandwidthPlanDataType: %v", v)
	}
}

type Usd struct {
	Monthly *int64 `json:"monthly,omitempty"`
	Hourly  *int64 `json:"hourly,omitempty"`
}

func (o *Usd) GetMonthly() *int64 {
	if o == nil {
		return nil
	}
	return o.Monthly
}

func (o *Usd) GetHourly() *int64 {
	if o == nil {
		return nil
	}
	return o.Hourly
}

type Brl struct {
	Monthly *int64 `json:"monthly,omitempty"`
	Hourly  *int64 `json:"hourly,omitempty"`
}

func (o *Brl) GetMonthly() *int64 {
	if o == nil {
		return nil
	}
	return o.Monthly
}

func (o *Brl) GetHourly() *int64 {
	if o == nil {
		return nil
	}
	return o.Hourly
}

type Pricing struct {
	Usd *Usd `json:"usd,omitempty"`
	Brl *Brl `json:"brl,omitempty"`
}

func (o *Pricing) GetUsd() *Usd {
	if o == nil {
		return nil
	}
	return o.Usd
}

func (o *Pricing) GetBrl() *Brl {
	if o == nil {
		return nil
	}
	return o.Brl
}

type BandwidthPlanDataAttributes struct {
	Region    *string  `json:"region,omitempty"`
	Locations []string `json:"locations,omitempty"`
	Pricing   *Pricing `json:"pricing,omitempty"`
}

func (o *BandwidthPlanDataAttributes) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *BandwidthPlanDataAttributes) GetLocations() []string {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *BandwidthPlanDataAttributes) GetPricing() *Pricing {
	if o == nil {
		return nil
	}
	return o.Pricing
}

type BandwidthPlanData struct {
	ID         *string                      `json:"id,omitempty"`
	Type       *BandwidthPlanDataType       `json:"type,omitempty"`
	Attributes *BandwidthPlanDataAttributes `json:"attributes,omitempty"`
}

func (o *BandwidthPlanData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BandwidthPlanData) GetType() *BandwidthPlanDataType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *BandwidthPlanData) GetAttributes() *BandwidthPlanDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}
