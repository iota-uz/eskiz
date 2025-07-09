# UserInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**UserInfo**](UserInfo.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserInfoResponse

`func NewUserInfoResponse() *UserInfoResponse`

NewUserInfoResponse instantiates a new UserInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoResponseWithDefaults

`func NewUserInfoResponseWithDefaults() *UserInfoResponse`

NewUserInfoResponseWithDefaults instantiates a new UserInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserInfoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserInfoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserInfoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserInfoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *UserInfoResponse) GetData() UserInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserInfoResponse) GetDataOk() (*UserInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserInfoResponse) SetData(v UserInfo)`

SetData sets Data field to given value.

### HasData

`func (o *UserInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *UserInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInfoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserInfoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserInfoResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserInfoResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


