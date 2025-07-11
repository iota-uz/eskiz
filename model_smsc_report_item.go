/*
СМС шлюз от Eskiz.uz

Отправляйте СМС по всему миру, в любом количестве!  В тестовом статусе для отправки тестовых смс сообщений, Вы можете использовать только нижеуказанные тексты:  - Это тест от Eskiz      - Bu Eskiz dan test      - This is test from Eskiz

API version: 1.0.0
Contact: danil@iota.uz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eskizapi

import (
	"encoding/json"
)

// checks if the SmscReportItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmscReportItem{}

// SmscReportItem struct for SmscReportItem
type SmscReportItem struct {
	Year       *int     `form:"year" json:"year,omitempty"`
	Month      *int     `form:"month" json:"month,omitempty"`
	AdParts    *int     `form:"ad_parts" json:"ad_parts,omitempty"`
	AdSpent    *float64 `form:"ad_spent" json:"ad_spent,omitempty"`
	Parts      *int     `form:"parts" json:"parts,omitempty"`
	Spent      *float64 `form:"spent" json:"spent,omitempty"`
	TotalParts *int     `form:"total_parts" json:"total_parts,omitempty"`
	TotalSpent *float64 `form:"total_spent" json:"total_spent,omitempty"`
	SmscId     *int     `form:"smsc_id" json:"smsc_id,omitempty"`
}

// NewSmscReportItem instantiates a new SmscReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmscReportItem() *SmscReportItem {
	this := SmscReportItem{}
	return &this
}

// NewSmscReportItemWithDefaults instantiates a new SmscReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmscReportItemWithDefaults() *SmscReportItem {
	this := SmscReportItem{}
	return &this
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *SmscReportItem) GetYear() int {
	if o == nil || IsNil(o.Year) {
		var ret int
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetYearOk() (*int, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *SmscReportItem) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int and assigns it to the Year field.
func (o *SmscReportItem) SetYear(v int) {
	o.Year = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *SmscReportItem) GetMonth() int {
	if o == nil || IsNil(o.Month) {
		var ret int
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetMonthOk() (*int, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *SmscReportItem) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int and assigns it to the Month field.
func (o *SmscReportItem) SetMonth(v int) {
	o.Month = &v
}

// GetAdParts returns the AdParts field value if set, zero value otherwise.
func (o *SmscReportItem) GetAdParts() int {
	if o == nil || IsNil(o.AdParts) {
		var ret int
		return ret
	}
	return *o.AdParts
}

// GetAdPartsOk returns a tuple with the AdParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetAdPartsOk() (*int, bool) {
	if o == nil || IsNil(o.AdParts) {
		return nil, false
	}
	return o.AdParts, true
}

// HasAdParts returns a boolean if a field has been set.
func (o *SmscReportItem) HasAdParts() bool {
	if o != nil && !IsNil(o.AdParts) {
		return true
	}

	return false
}

// SetAdParts gets a reference to the given int and assigns it to the AdParts field.
func (o *SmscReportItem) SetAdParts(v int) {
	o.AdParts = &v
}

// GetAdSpent returns the AdSpent field value if set, zero value otherwise.
func (o *SmscReportItem) GetAdSpent() float64 {
	if o == nil || IsNil(o.AdSpent) {
		var ret float64
		return ret
	}
	return *o.AdSpent
}

// GetAdSpentOk returns a tuple with the AdSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetAdSpentOk() (*float64, bool) {
	if o == nil || IsNil(o.AdSpent) {
		return nil, false
	}
	return o.AdSpent, true
}

// HasAdSpent returns a boolean if a field has been set.
func (o *SmscReportItem) HasAdSpent() bool {
	if o != nil && !IsNil(o.AdSpent) {
		return true
	}

	return false
}

// SetAdSpent gets a reference to the given float64 and assigns it to the AdSpent field.
func (o *SmscReportItem) SetAdSpent(v float64) {
	o.AdSpent = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *SmscReportItem) GetParts() int {
	if o == nil || IsNil(o.Parts) {
		var ret int
		return ret
	}
	return *o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetPartsOk() (*int, bool) {
	if o == nil || IsNil(o.Parts) {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *SmscReportItem) HasParts() bool {
	if o != nil && !IsNil(o.Parts) {
		return true
	}

	return false
}

// SetParts gets a reference to the given int and assigns it to the Parts field.
func (o *SmscReportItem) SetParts(v int) {
	o.Parts = &v
}

// GetSpent returns the Spent field value if set, zero value otherwise.
func (o *SmscReportItem) GetSpent() float64 {
	if o == nil || IsNil(o.Spent) {
		var ret float64
		return ret
	}
	return *o.Spent
}

// GetSpentOk returns a tuple with the Spent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetSpentOk() (*float64, bool) {
	if o == nil || IsNil(o.Spent) {
		return nil, false
	}
	return o.Spent, true
}

// HasSpent returns a boolean if a field has been set.
func (o *SmscReportItem) HasSpent() bool {
	if o != nil && !IsNil(o.Spent) {
		return true
	}

	return false
}

// SetSpent gets a reference to the given float64 and assigns it to the Spent field.
func (o *SmscReportItem) SetSpent(v float64) {
	o.Spent = &v
}

// GetTotalParts returns the TotalParts field value if set, zero value otherwise.
func (o *SmscReportItem) GetTotalParts() int {
	if o == nil || IsNil(o.TotalParts) {
		var ret int
		return ret
	}
	return *o.TotalParts
}

// GetTotalPartsOk returns a tuple with the TotalParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetTotalPartsOk() (*int, bool) {
	if o == nil || IsNil(o.TotalParts) {
		return nil, false
	}
	return o.TotalParts, true
}

// HasTotalParts returns a boolean if a field has been set.
func (o *SmscReportItem) HasTotalParts() bool {
	if o != nil && !IsNil(o.TotalParts) {
		return true
	}

	return false
}

// SetTotalParts gets a reference to the given int and assigns it to the TotalParts field.
func (o *SmscReportItem) SetTotalParts(v int) {
	o.TotalParts = &v
}

// GetTotalSpent returns the TotalSpent field value if set, zero value otherwise.
func (o *SmscReportItem) GetTotalSpent() float64 {
	if o == nil || IsNil(o.TotalSpent) {
		var ret float64
		return ret
	}
	return *o.TotalSpent
}

// GetTotalSpentOk returns a tuple with the TotalSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetTotalSpentOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalSpent) {
		return nil, false
	}
	return o.TotalSpent, true
}

// HasTotalSpent returns a boolean if a field has been set.
func (o *SmscReportItem) HasTotalSpent() bool {
	if o != nil && !IsNil(o.TotalSpent) {
		return true
	}

	return false
}

// SetTotalSpent gets a reference to the given float64 and assigns it to the TotalSpent field.
func (o *SmscReportItem) SetTotalSpent(v float64) {
	o.TotalSpent = &v
}

// GetSmscId returns the SmscId field value if set, zero value otherwise.
func (o *SmscReportItem) GetSmscId() int {
	if o == nil || IsNil(o.SmscId) {
		var ret int
		return ret
	}
	return *o.SmscId
}

// GetSmscIdOk returns a tuple with the SmscId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmscReportItem) GetSmscIdOk() (*int, bool) {
	if o == nil || IsNil(o.SmscId) {
		return nil, false
	}
	return o.SmscId, true
}

// HasSmscId returns a boolean if a field has been set.
func (o *SmscReportItem) HasSmscId() bool {
	if o != nil && !IsNil(o.SmscId) {
		return true
	}

	return false
}

// SetSmscId gets a reference to the given int and assigns it to the SmscId field.
func (o *SmscReportItem) SetSmscId(v int) {
	o.SmscId = &v
}

func (o SmscReportItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmscReportItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	if !IsNil(o.AdParts) {
		toSerialize["ad_parts"] = o.AdParts
	}
	if !IsNil(o.AdSpent) {
		toSerialize["ad_spent"] = o.AdSpent
	}
	if !IsNil(o.Parts) {
		toSerialize["parts"] = o.Parts
	}
	if !IsNil(o.Spent) {
		toSerialize["spent"] = o.Spent
	}
	if !IsNil(o.TotalParts) {
		toSerialize["total_parts"] = o.TotalParts
	}
	if !IsNil(o.TotalSpent) {
		toSerialize["total_spent"] = o.TotalSpent
	}
	if !IsNil(o.SmscId) {
		toSerialize["smsc_id"] = o.SmscId
	}
	return toSerialize, nil
}

type NullableSmscReportItem struct {
	value *SmscReportItem
	isSet bool
}

func (v NullableSmscReportItem) Get() *SmscReportItem {
	return v.value
}

func (v *NullableSmscReportItem) Set(val *SmscReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSmscReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSmscReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmscReportItem(val *SmscReportItem) *NullableSmscReportItem {
	return &NullableSmscReportItem{value: val, isSet: true}
}

func (v NullableSmscReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmscReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
