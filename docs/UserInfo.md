# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IsVip** | Pointer to **bool** |  | [optional] 
**Balance** | Pointer to **float64** |  | [optional] 

## Methods

### NewUserInfo

`func NewUserInfo() *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserInfo) GetId() int`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInfo) GetIdOk() (*int, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInfo) SetId(v int)`

SetId sets Id field to given value.

### HasId

`func (o *UserInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserInfo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *UserInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UserInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *UserInfo) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserInfo) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserInfo) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *UserInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsVip

`func (o *UserInfo) GetIsVip() bool`

GetIsVip returns the IsVip field if non-nil, zero value otherwise.

### GetIsVipOk

`func (o *UserInfo) GetIsVipOk() (*bool, bool)`

GetIsVipOk returns a tuple with the IsVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVip

`func (o *UserInfo) SetIsVip(v bool)`

SetIsVip sets IsVip field to given value.

### HasIsVip

`func (o *UserInfo) HasIsVip() bool`

HasIsVip returns a boolean if a field has been set.

### GetBalance

`func (o *UserInfo) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *UserInfo) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *UserInfo) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *UserInfo) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


