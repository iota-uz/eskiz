# RefreshTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**LoginData**](LoginData.md) |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 

## Methods

### NewRefreshTokenResponse

`func NewRefreshTokenResponse() *RefreshTokenResponse`

NewRefreshTokenResponse instantiates a new RefreshTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenResponseWithDefaults

`func NewRefreshTokenResponseWithDefaults() *RefreshTokenResponse`

NewRefreshTokenResponseWithDefaults instantiates a new RefreshTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RefreshTokenResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RefreshTokenResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RefreshTokenResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RefreshTokenResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *RefreshTokenResponse) GetData() LoginData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RefreshTokenResponse) GetDataOk() (*LoginData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RefreshTokenResponse) SetData(v LoginData)`

SetData sets Data field to given value.

### HasData

`func (o *RefreshTokenResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTokenType

`func (o *RefreshTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *RefreshTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *RefreshTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *RefreshTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


