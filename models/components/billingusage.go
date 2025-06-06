// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/latitudesh/latitudesh-go-sdk/internal/utils"
	"time"
)

// BillingUsageProject - The project in which the returned usage belongs to
type BillingUsageProject struct {
	ID   *string `json:"id,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Name *string `json:"name,omitempty"`
}

func (o *BillingUsageProject) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BillingUsageProject) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *BillingUsageProject) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// Period - The period from the returned billing cycle
type Period struct {
	Start *time.Time `json:"start,omitempty"`
	End   *time.Time `json:"end,omitempty"`
}

func (p Period) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Period) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Period) GetStart() *time.Time {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *Period) GetEnd() *time.Time {
	if o == nil {
		return nil
	}
	return o.End
}

// BillingUsageType - Type of discount (percentage or fixed amount)
type BillingUsageType string

const (
	BillingUsageTypePercent BillingUsageType = "percent"
	BillingUsageTypeAmount  BillingUsageType = "amount"
)

func (e BillingUsageType) ToPointer() *BillingUsageType {
	return &e
}
func (e *BillingUsageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "percent":
		fallthrough
	case "amount":
		*e = BillingUsageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillingUsageType: %v", v)
	}
}

type Discounts struct {
	// Description of the discount
	Description string `json:"description"`
	// Type of discount (percentage or fixed amount)
	Type BillingUsageType `json:"type"`
	// Value of the discount (percentage or amount)
	Value float32 `json:"value"`
}

func (o *Discounts) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Discounts) GetType() BillingUsageType {
	if o == nil {
		return BillingUsageType("")
	}
	return o.Type
}

func (o *Discounts) GetValue() float32 {
	if o == nil {
		return 0.0
	}
	return o.Value
}

type Unit string

const (
	UnitQuantity Unit = "quantity"
	UnitHour     Unit = "hour"
	UnitMinute   Unit = "minute"
)

func (e Unit) ToPointer() *Unit {
	return &e
}
func (e *Unit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "quantity":
		fallthrough
	case "hour":
		fallthrough
	case "minute":
		*e = Unit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Unit: %v", v)
	}
}

type UsageType string

const (
	UsageTypeLicensed UsageType = "licensed"
	UsageTypeMetered  UsageType = "metered"
)

func (e UsageType) ToPointer() *UsageType {
	return &e
}
func (e *UsageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "licensed":
		fallthrough
	case "metered":
		*e = UsageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UsageType: %v", v)
	}
}

type Metadata struct {
	ID       *string  `json:"id,omitempty"`
	Hostname *string  `json:"hostname,omitempty"`
	Plan     *string  `json:"plan,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

func (o *Metadata) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Metadata) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *Metadata) GetPlan() *string {
	if o == nil {
		return nil
	}
	return o.Plan
}

func (o *Metadata) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type Products struct {
	ID                    *string     `json:"id,omitempty"`
	Resource              *string     `json:"resource,omitempty"`
	Name                  *string     `json:"name,omitempty"`
	Proration             *bool       `json:"proration,omitempty"`
	Discounts             []Discounts `json:"discounts,omitempty"`
	Discountable          *bool       `json:"discountable,omitempty"`
	Description           *string     `json:"description,omitempty"`
	AmountWithoutDiscount *int64      `json:"amount_without_discount,omitempty"`
	Start                 *time.Time  `json:"start,omitempty"`
	End                   *time.Time  `json:"end,omitempty"`
	Unit                  *Unit       `json:"unit,omitempty"`
	// The unit price of the product in cents
	UnitPrice *float64   `json:"unit_price,omitempty"`
	UsageType *UsageType `json:"usage_type,omitempty"`
	Quantity  *float64   `json:"quantity,omitempty"`
	// The total usage price of the product in cents
	Price    *float64  `json:"price,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

func (p Products) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Products) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Products) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Products) GetResource() *string {
	if o == nil {
		return nil
	}
	return o.Resource
}

func (o *Products) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Products) GetProration() *bool {
	if o == nil {
		return nil
	}
	return o.Proration
}

func (o *Products) GetDiscounts() []Discounts {
	if o == nil {
		return nil
	}
	return o.Discounts
}

func (o *Products) GetDiscountable() *bool {
	if o == nil {
		return nil
	}
	return o.Discountable
}

func (o *Products) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Products) GetAmountWithoutDiscount() *int64 {
	if o == nil {
		return nil
	}
	return o.AmountWithoutDiscount
}

func (o *Products) GetStart() *time.Time {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *Products) GetEnd() *time.Time {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *Products) GetUnit() *Unit {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *Products) GetUnitPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}

func (o *Products) GetUsageType() *UsageType {
	if o == nil {
		return nil
	}
	return o.UsageType
}

func (o *Products) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *Products) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *Products) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type BillingUsageAttributes struct {
	// The project in which the returned usage belongs to
	Project *BillingUsageProject `json:"project,omitempty"`
	// The period from the returned billing cycle
	Period *Period `json:"period,omitempty"`
	// The total usage price in cents
	Price *float64 `json:"price,omitempty"`
	// The threshold which we use to charge your usage, in cents
	Threshold *float64   `json:"threshold,omitempty"`
	Products  []Products `json:"products,omitempty"`
}

func (o *BillingUsageAttributes) GetProject() *BillingUsageProject {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *BillingUsageAttributes) GetPeriod() *Period {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *BillingUsageAttributes) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *BillingUsageAttributes) GetThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.Threshold
}

func (o *BillingUsageAttributes) GetProducts() []Products {
	if o == nil {
		return nil
	}
	return o.Products
}

type BillingUsageData struct {
	ID         *string                 `json:"id,omitempty"`
	Attributes *BillingUsageAttributes `json:"attributes,omitempty"`
}

func (o *BillingUsageData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BillingUsageData) GetAttributes() *BillingUsageAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

type BillingUsage struct {
	Data *BillingUsageData `json:"data,omitempty"`
}

func (o *BillingUsage) GetData() *BillingUsageData {
	if o == nil {
		return nil
	}
	return o.Data
}
