# SmsStatusData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int** |  | [optional] 
**UserId** | Pointer to **int** |  | [optional] 
**CountryId** | Pointer to **int** |  | [optional] 
**ConnectionId** | Pointer to **int** |  | [optional] 
**SmscId** | Pointer to **int** |  | [optional] 
**DispatchId** | Pointer to **NullableString** |  | [optional] 
**UserSmsId** | Pointer to **NullableString** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**TotalPrice** | Pointer to **float64** |  | [optional] 
**IsAd** | Pointer to **bool** |  | [optional] 
**Nick** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **int** |  | [optional] 
**PartsCount** | Pointer to **int** |  | [optional] 
**Parts** | Pointer to [**map[string]SmsStatusPart**](SmsStatusPart.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SmscData** | Pointer to **map[string][]string** |  | [optional] 
**SentAt** | Pointer to **time.Time** |  | [optional] 
**SubmitSmRespAt** | Pointer to **time.Time** |  | [optional] 
**DeliverySmAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSmsStatusData

`func NewSmsStatusData() *SmsStatusData`

NewSmsStatusData instantiates a new SmsStatusData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsStatusDataWithDefaults

`func NewSmsStatusDataWithDefaults() *SmsStatusData`

NewSmsStatusDataWithDefaults instantiates a new SmsStatusData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsStatusData) GetId() int`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsStatusData) GetIdOk() (*int, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsStatusData) SetId(v int)`

SetId sets Id field to given value.

### HasId

`func (o *SmsStatusData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *SmsStatusData) GetUserId() int`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SmsStatusData) GetUserIdOk() (*int, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SmsStatusData) SetUserId(v int)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SmsStatusData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCountryId

`func (o *SmsStatusData) GetCountryId() int`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *SmsStatusData) GetCountryIdOk() (*int, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *SmsStatusData) SetCountryId(v int)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *SmsStatusData) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetConnectionId

`func (o *SmsStatusData) GetConnectionId() int`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SmsStatusData) GetConnectionIdOk() (*int, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SmsStatusData) SetConnectionId(v int)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SmsStatusData) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetSmscId

`func (o *SmsStatusData) GetSmscId() int`

GetSmscId returns the SmscId field if non-nil, zero value otherwise.

### GetSmscIdOk

`func (o *SmsStatusData) GetSmscIdOk() (*int, bool)`

GetSmscIdOk returns a tuple with the SmscId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmscId

`func (o *SmsStatusData) SetSmscId(v int)`

SetSmscId sets SmscId field to given value.

### HasSmscId

`func (o *SmsStatusData) HasSmscId() bool`

HasSmscId returns a boolean if a field has been set.

### GetDispatchId

`func (o *SmsStatusData) GetDispatchId() string`

GetDispatchId returns the DispatchId field if non-nil, zero value otherwise.

### GetDispatchIdOk

`func (o *SmsStatusData) GetDispatchIdOk() (*string, bool)`

GetDispatchIdOk returns a tuple with the DispatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchId

`func (o *SmsStatusData) SetDispatchId(v string)`

SetDispatchId sets DispatchId field to given value.

### HasDispatchId

`func (o *SmsStatusData) HasDispatchId() bool`

HasDispatchId returns a boolean if a field has been set.

### SetDispatchIdNil

`func (o *SmsStatusData) SetDispatchIdNil(b bool)`

 SetDispatchIdNil sets the value for DispatchId to be an explicit nil

### UnsetDispatchId
`func (o *SmsStatusData) UnsetDispatchId()`

UnsetDispatchId ensures that no value is present for DispatchId, not even an explicit nil
### GetUserSmsId

`func (o *SmsStatusData) GetUserSmsId() string`

GetUserSmsId returns the UserSmsId field if non-nil, zero value otherwise.

### GetUserSmsIdOk

`func (o *SmsStatusData) GetUserSmsIdOk() (*string, bool)`

GetUserSmsIdOk returns a tuple with the UserSmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSmsId

`func (o *SmsStatusData) SetUserSmsId(v string)`

SetUserSmsId sets UserSmsId field to given value.

### HasUserSmsId

`func (o *SmsStatusData) HasUserSmsId() bool`

HasUserSmsId returns a boolean if a field has been set.

### SetUserSmsIdNil

`func (o *SmsStatusData) SetUserSmsIdNil(b bool)`

 SetUserSmsIdNil sets the value for UserSmsId to be an explicit nil

### UnsetUserSmsId
`func (o *SmsStatusData) UnsetUserSmsId()`

UnsetUserSmsId ensures that no value is present for UserSmsId, not even an explicit nil
### GetRequestId

`func (o *SmsStatusData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SmsStatusData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SmsStatusData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SmsStatusData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetPrice

`func (o *SmsStatusData) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SmsStatusData) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SmsStatusData) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SmsStatusData) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTotalPrice

`func (o *SmsStatusData) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *SmsStatusData) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *SmsStatusData) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *SmsStatusData) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetIsAd

`func (o *SmsStatusData) GetIsAd() bool`

GetIsAd returns the IsAd field if non-nil, zero value otherwise.

### GetIsAdOk

`func (o *SmsStatusData) GetIsAdOk() (*bool, bool)`

GetIsAdOk returns a tuple with the IsAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAd

`func (o *SmsStatusData) SetIsAd(v bool)`

SetIsAd sets IsAd field to given value.

### HasIsAd

`func (o *SmsStatusData) HasIsAd() bool`

HasIsAd returns a boolean if a field has been set.

### GetNick

`func (o *SmsStatusData) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *SmsStatusData) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *SmsStatusData) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *SmsStatusData) HasNick() bool`

HasNick returns a boolean if a field has been set.

### GetTo

`func (o *SmsStatusData) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SmsStatusData) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SmsStatusData) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SmsStatusData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMessage

`func (o *SmsStatusData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SmsStatusData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SmsStatusData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SmsStatusData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEncoding

`func (o *SmsStatusData) GetEncoding() int`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *SmsStatusData) GetEncodingOk() (*int, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *SmsStatusData) SetEncoding(v int)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *SmsStatusData) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPartsCount

`func (o *SmsStatusData) GetPartsCount() int`

GetPartsCount returns the PartsCount field if non-nil, zero value otherwise.

### GetPartsCountOk

`func (o *SmsStatusData) GetPartsCountOk() (*int, bool)`

GetPartsCountOk returns a tuple with the PartsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsCount

`func (o *SmsStatusData) SetPartsCount(v int)`

SetPartsCount sets PartsCount field to given value.

### HasPartsCount

`func (o *SmsStatusData) HasPartsCount() bool`

HasPartsCount returns a boolean if a field has been set.

### GetParts

`func (o *SmsStatusData) GetParts() map[string]SmsStatusPart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *SmsStatusData) GetPartsOk() (*map[string]SmsStatusPart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *SmsStatusData) SetParts(v map[string]SmsStatusPart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *SmsStatusData) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetStatus

`func (o *SmsStatusData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmsStatusData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmsStatusData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmsStatusData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSmscData

`func (o *SmsStatusData) GetSmscData() map[string][]string`

GetSmscData returns the SmscData field if non-nil, zero value otherwise.

### GetSmscDataOk

`func (o *SmsStatusData) GetSmscDataOk() (*map[string][]string, bool)`

GetSmscDataOk returns a tuple with the SmscData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmscData

`func (o *SmsStatusData) SetSmscData(v map[string][]string)`

SetSmscData sets SmscData field to given value.

### HasSmscData

`func (o *SmsStatusData) HasSmscData() bool`

HasSmscData returns a boolean if a field has been set.

### GetSentAt

`func (o *SmsStatusData) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *SmsStatusData) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *SmsStatusData) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *SmsStatusData) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetSubmitSmRespAt

`func (o *SmsStatusData) GetSubmitSmRespAt() time.Time`

GetSubmitSmRespAt returns the SubmitSmRespAt field if non-nil, zero value otherwise.

### GetSubmitSmRespAtOk

`func (o *SmsStatusData) GetSubmitSmRespAtOk() (*time.Time, bool)`

GetSubmitSmRespAtOk returns a tuple with the SubmitSmRespAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitSmRespAt

`func (o *SmsStatusData) SetSubmitSmRespAt(v time.Time)`

SetSubmitSmRespAt sets SubmitSmRespAt field to given value.

### HasSubmitSmRespAt

`func (o *SmsStatusData) HasSubmitSmRespAt() bool`

HasSubmitSmRespAt returns a boolean if a field has been set.

### GetDeliverySmAt

`func (o *SmsStatusData) GetDeliverySmAt() time.Time`

GetDeliverySmAt returns the DeliverySmAt field if non-nil, zero value otherwise.

### GetDeliverySmAtOk

`func (o *SmsStatusData) GetDeliverySmAtOk() (*time.Time, bool)`

GetDeliverySmAtOk returns a tuple with the DeliverySmAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySmAt

`func (o *SmsStatusData) SetDeliverySmAt(v time.Time)`

SetDeliverySmAt sets DeliverySmAt field to given value.

### HasDeliverySmAt

`func (o *SmsStatusData) HasDeliverySmAt() bool`

HasDeliverySmAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SmsStatusData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SmsStatusData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SmsStatusData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SmsStatusData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SmsStatusData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SmsStatusData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SmsStatusData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SmsStatusData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


