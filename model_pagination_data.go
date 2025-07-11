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

// checks if the PaginationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationData{}

// PaginationData struct for PaginationData
type PaginationData struct {
	CurrentPage  *int           `form:"current_page" json:"current_page,omitempty"`
	Path         *string        `form:"path" json:"path,omitempty"`
	PrevPageUrl  NullableString `form:"prev_page_url" json:"prev_page_url,omitempty"`
	FirstPageUrl *string        `form:"first_page_url" json:"first_page_url,omitempty"`
	LastPageUrl  *string        `form:"last_page_url" json:"last_page_url,omitempty"`
	NextPageUrl  NullableString `form:"next_page_url" json:"next_page_url,omitempty"`
	PerPage      *int           `form:"per_page" json:"per_page,omitempty"`
	LastPage     *int           `form:"last_page" json:"last_page,omitempty"`
	From         *int           `form:"from" json:"from,omitempty"`
	To           *int           `form:"to" json:"to,omitempty"`
	Total        *int           `form:"total" json:"total,omitempty"`
	Result       []MessageItem  `form:"result" json:"result,omitempty"`
	Links        []LinkItem     `form:"links" json:"links,omitempty"`
}

// NewPaginationData instantiates a new PaginationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationData() *PaginationData {
	this := PaginationData{}
	return &this
}

// NewPaginationDataWithDefaults instantiates a new PaginationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationDataWithDefaults() *PaginationData {
	this := PaginationData{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *PaginationData) GetCurrentPage() int {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetCurrentPageOk() (*int, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *PaginationData) HasCurrentPage() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int and assigns it to the CurrentPage field.
func (o *PaginationData) SetCurrentPage(v int) {
	o.CurrentPage = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PaginationData) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PaginationData) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PaginationData) SetPath(v string) {
	o.Path = &v
}

// GetPrevPageUrl returns the PrevPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationData) GetPrevPageUrl() string {
	if o == nil || IsNil(o.PrevPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PrevPageUrl.Get()
}

// GetPrevPageUrlOk returns a tuple with the PrevPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationData) GetPrevPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevPageUrl.Get(), o.PrevPageUrl.IsSet()
}

// HasPrevPageUrl returns a boolean if a field has been set.
func (o *PaginationData) HasPrevPageUrl() bool {
	if o != nil && o.PrevPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPrevPageUrl gets a reference to the given NullableString and assigns it to the PrevPageUrl field.
func (o *PaginationData) SetPrevPageUrl(v string) {
	o.PrevPageUrl.Set(&v)
}

// SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil
func (o *PaginationData) SetPrevPageUrlNil() {
	o.PrevPageUrl.Set(nil)
}

// UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
func (o *PaginationData) UnsetPrevPageUrl() {
	o.PrevPageUrl.Unset()
}

// GetFirstPageUrl returns the FirstPageUrl field value if set, zero value otherwise.
func (o *PaginationData) GetFirstPageUrl() string {
	if o == nil || IsNil(o.FirstPageUrl) {
		var ret string
		return ret
	}
	return *o.FirstPageUrl
}

// GetFirstPageUrlOk returns a tuple with the FirstPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetFirstPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FirstPageUrl) {
		return nil, false
	}
	return o.FirstPageUrl, true
}

// HasFirstPageUrl returns a boolean if a field has been set.
func (o *PaginationData) HasFirstPageUrl() bool {
	if o != nil && !IsNil(o.FirstPageUrl) {
		return true
	}

	return false
}

// SetFirstPageUrl gets a reference to the given string and assigns it to the FirstPageUrl field.
func (o *PaginationData) SetFirstPageUrl(v string) {
	o.FirstPageUrl = &v
}

// GetLastPageUrl returns the LastPageUrl field value if set, zero value otherwise.
func (o *PaginationData) GetLastPageUrl() string {
	if o == nil || IsNil(o.LastPageUrl) {
		var ret string
		return ret
	}
	return *o.LastPageUrl
}

// GetLastPageUrlOk returns a tuple with the LastPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetLastPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LastPageUrl) {
		return nil, false
	}
	return o.LastPageUrl, true
}

// HasLastPageUrl returns a boolean if a field has been set.
func (o *PaginationData) HasLastPageUrl() bool {
	if o != nil && !IsNil(o.LastPageUrl) {
		return true
	}

	return false
}

// SetLastPageUrl gets a reference to the given string and assigns it to the LastPageUrl field.
func (o *PaginationData) SetLastPageUrl(v string) {
	o.LastPageUrl = &v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationData) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationData) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *PaginationData) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *PaginationData) SetNextPageUrl(v string) {
	o.NextPageUrl.Set(&v)
}

// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *PaginationData) SetNextPageUrlNil() {
	o.NextPageUrl.Set(nil)
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *PaginationData) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *PaginationData) GetPerPage() int {
	if o == nil || IsNil(o.PerPage) {
		var ret int
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetPerPageOk() (*int, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *PaginationData) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int and assigns it to the PerPage field.
func (o *PaginationData) SetPerPage(v int) {
	o.PerPage = &v
}

// GetLastPage returns the LastPage field value if set, zero value otherwise.
func (o *PaginationData) GetLastPage() int {
	if o == nil || IsNil(o.LastPage) {
		var ret int
		return ret
	}
	return *o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetLastPageOk() (*int, bool) {
	if o == nil || IsNil(o.LastPage) {
		return nil, false
	}
	return o.LastPage, true
}

// HasLastPage returns a boolean if a field has been set.
func (o *PaginationData) HasLastPage() bool {
	if o != nil && !IsNil(o.LastPage) {
		return true
	}

	return false
}

// SetLastPage gets a reference to the given int and assigns it to the LastPage field.
func (o *PaginationData) SetLastPage(v int) {
	o.LastPage = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PaginationData) GetFrom() int {
	if o == nil || IsNil(o.From) {
		var ret int
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetFromOk() (*int, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PaginationData) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int and assigns it to the From field.
func (o *PaginationData) SetFrom(v int) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *PaginationData) GetTo() int {
	if o == nil || IsNil(o.To) {
		var ret int
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetToOk() (*int, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *PaginationData) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given int and assigns it to the To field.
func (o *PaginationData) SetTo(v int) {
	o.To = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PaginationData) GetTotal() int {
	if o == nil || IsNil(o.Total) {
		var ret int
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetTotalOk() (*int, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PaginationData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int and assigns it to the Total field.
func (o *PaginationData) SetTotal(v int) {
	o.Total = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *PaginationData) GetResult() []MessageItem {
	if o == nil || IsNil(o.Result) {
		var ret []MessageItem
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetResultOk() ([]MessageItem, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PaginationData) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []MessageItem and assigns it to the Result field.
func (o *PaginationData) SetResult(v []MessageItem) {
	o.Result = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PaginationData) GetLinks() []LinkItem {
	if o == nil || IsNil(o.Links) {
		var ret []LinkItem
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetLinksOk() ([]LinkItem, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PaginationData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkItem and assigns it to the Links field.
func (o *PaginationData) SetLinks(v []LinkItem) {
	o.Links = v
}

func (o PaginationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentPage) {
		toSerialize["current_page"] = o.CurrentPage
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if o.PrevPageUrl.IsSet() {
		toSerialize["prev_page_url"] = o.PrevPageUrl.Get()
	}
	if !IsNil(o.FirstPageUrl) {
		toSerialize["first_page_url"] = o.FirstPageUrl
	}
	if !IsNil(o.LastPageUrl) {
		toSerialize["last_page_url"] = o.LastPageUrl
	}
	if o.NextPageUrl.IsSet() {
		toSerialize["next_page_url"] = o.NextPageUrl.Get()
	}
	if !IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	if !IsNil(o.LastPage) {
		toSerialize["last_page"] = o.LastPage
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullablePaginationData struct {
	value *PaginationData
	isSet bool
}

func (v NullablePaginationData) Get() *PaginationData {
	return v.value
}

func (v *NullablePaginationData) Set(val *PaginationData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationData(val *PaginationData) *NullablePaginationData {
	return &NullablePaginationData{value: val, isSet: true}
}

func (v NullablePaginationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
