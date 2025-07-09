# LinkItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewLinkItem

`func NewLinkItem() *LinkItem`

NewLinkItem instantiates a new LinkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkItemWithDefaults

`func NewLinkItemWithDefaults() *LinkItem`

NewLinkItemWithDefaults instantiates a new LinkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *LinkItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LinkItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *LinkItem) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *LinkItem) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLabel

`func (o *LinkItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LinkItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LinkItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LinkItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetActive

`func (o *LinkItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *LinkItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *LinkItem) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *LinkItem) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


