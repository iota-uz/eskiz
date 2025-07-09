# MessageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int** |  | [optional] 
**UserId** | Pointer to **int** |  | [optional] 
**CountryId** | Pointer to **NullableInt** |  | [optional] 
**ConnectionId** | Pointer to **int** |  | [optional] 
**SmscId** | Pointer to **int** |  | [optional] 
**DispatchId** | Pointer to **NullableInt** |  | [optional] 
**UserSmsId** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**IsAd** | Pointer to **bool** |  | [optional] 
**Nick** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Encoding** | Pointer to **int** |  | [optional] 
**PartsCount** | Pointer to **int** |  | [optional] 
**Parts** | Pointer to [**map[string]MessagePart**](MessagePart.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SmscData** | Pointer to **map[string][]string** |  | [optional] 
**SentAt** | Pointer to **time.Time** |  | [optional] 
**SubmitSmRespAt** | Pointer to **time.Time** |  | [optional] 
**DeliverySmAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMessageItem

`func NewMessageItem() *MessageItem`

NewMessageItem instantiates a new MessageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageItemWithDefaults

`func NewMessageItemWithDefaults() *MessageItem`

NewMessageItemWithDefaults instantiates a new MessageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageItem) GetId() int`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageItem) GetIdOk() (*int, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageItem) SetId(v int)`

SetId sets Id field to given value.

### HasId

`func (o *MessageItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *MessageItem) GetUserId() int`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MessageItem) GetUserIdOk() (*int, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MessageItem) SetUserId(v int)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MessageItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCountryId

`func (o *MessageItem) GetCountryId() int`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *MessageItem) GetCountryIdOk() (*int, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *MessageItem) SetCountryId(v int)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *MessageItem) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *MessageItem) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *MessageItem) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetConnectionId

`func (o *MessageItem) GetConnectionId() int`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *MessageItem) GetConnectionIdOk() (*int, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *MessageItem) SetConnectionId(v int)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *MessageItem) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetSmscId

`func (o *MessageItem) GetSmscId() int`

GetSmscId returns the SmscId field if non-nil, zero value otherwise.

### GetSmscIdOk

`func (o *MessageItem) GetSmscIdOk() (*int, bool)`

GetSmscIdOk returns a tuple with the SmscId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmscId

`func (o *MessageItem) SetSmscId(v int)`

SetSmscId sets SmscId field to given value.

### HasSmscId

`func (o *MessageItem) HasSmscId() bool`

HasSmscId returns a boolean if a field has been set.

### GetDispatchId

`func (o *MessageItem) GetDispatchId() int`

GetDispatchId returns the DispatchId field if non-nil, zero value otherwise.

### GetDispatchIdOk

`func (o *MessageItem) GetDispatchIdOk() (*int, bool)`

GetDispatchIdOk returns a tuple with the DispatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchId

`func (o *MessageItem) SetDispatchId(v int)`

SetDispatchId sets DispatchId field to given value.

### HasDispatchId

`func (o *MessageItem) HasDispatchId() bool`

HasDispatchId returns a boolean if a field has been set.

### SetDispatchIdNil

`func (o *MessageItem) SetDispatchIdNil(b bool)`

 SetDispatchIdNil sets the value for DispatchId to be an explicit nil

### UnsetDispatchId
`func (o *MessageItem) UnsetDispatchId()`

UnsetDispatchId ensures that no value is present for DispatchId, not even an explicit nil
### GetUserSmsId

`func (o *MessageItem) GetUserSmsId() string`

GetUserSmsId returns the UserSmsId field if non-nil, zero value otherwise.

### GetUserSmsIdOk

`func (o *MessageItem) GetUserSmsIdOk() (*string, bool)`

GetUserSmsIdOk returns a tuple with the UserSmsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSmsId

`func (o *MessageItem) SetUserSmsId(v string)`

SetUserSmsId sets UserSmsId field to given value.

### HasUserSmsId

`func (o *MessageItem) HasUserSmsId() bool`

HasUserSmsId returns a boolean if a field has been set.

### GetRequestId

`func (o *MessageItem) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *MessageItem) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *MessageItem) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *MessageItem) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetPrice

`func (o *MessageItem) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MessageItem) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MessageItem) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MessageItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetIsAd

`func (o *MessageItem) GetIsAd() bool`

GetIsAd returns the IsAd field if non-nil, zero value otherwise.

### GetIsAdOk

`func (o *MessageItem) GetIsAdOk() (*bool, bool)`

GetIsAdOk returns a tuple with the IsAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAd

`func (o *MessageItem) SetIsAd(v bool)`

SetIsAd sets IsAd field to given value.

### HasIsAd

`func (o *MessageItem) HasIsAd() bool`

HasIsAd returns a boolean if a field has been set.

### GetNick

`func (o *MessageItem) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *MessageItem) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *MessageItem) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *MessageItem) HasNick() bool`

HasNick returns a boolean if a field has been set.

### GetTo

`func (o *MessageItem) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MessageItem) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MessageItem) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MessageItem) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMessage

`func (o *MessageItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MessageItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEncoding

`func (o *MessageItem) GetEncoding() int`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *MessageItem) GetEncodingOk() (*int, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *MessageItem) SetEncoding(v int)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *MessageItem) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPartsCount

`func (o *MessageItem) GetPartsCount() int`

GetPartsCount returns the PartsCount field if non-nil, zero value otherwise.

### GetPartsCountOk

`func (o *MessageItem) GetPartsCountOk() (*int, bool)`

GetPartsCountOk returns a tuple with the PartsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsCount

`func (o *MessageItem) SetPartsCount(v int)`

SetPartsCount sets PartsCount field to given value.

### HasPartsCount

`func (o *MessageItem) HasPartsCount() bool`

HasPartsCount returns a boolean if a field has been set.

### GetParts

`func (o *MessageItem) GetParts() map[string]MessagePart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MessageItem) GetPartsOk() (*map[string]MessagePart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MessageItem) SetParts(v map[string]MessagePart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MessageItem) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetStatus

`func (o *MessageItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSmscData

`func (o *MessageItem) GetSmscData() map[string][]string`

GetSmscData returns the SmscData field if non-nil, zero value otherwise.

### GetSmscDataOk

`func (o *MessageItem) GetSmscDataOk() (*map[string][]string, bool)`

GetSmscDataOk returns a tuple with the SmscData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmscData

`func (o *MessageItem) SetSmscData(v map[string][]string)`

SetSmscData sets SmscData field to given value.

### HasSmscData

`func (o *MessageItem) HasSmscData() bool`

HasSmscData returns a boolean if a field has been set.

### GetSentAt

`func (o *MessageItem) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *MessageItem) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *MessageItem) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *MessageItem) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetSubmitSmRespAt

`func (o *MessageItem) GetSubmitSmRespAt() time.Time`

GetSubmitSmRespAt returns the SubmitSmRespAt field if non-nil, zero value otherwise.

### GetSubmitSmRespAtOk

`func (o *MessageItem) GetSubmitSmRespAtOk() (*time.Time, bool)`

GetSubmitSmRespAtOk returns a tuple with the SubmitSmRespAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitSmRespAt

`func (o *MessageItem) SetSubmitSmRespAt(v time.Time)`

SetSubmitSmRespAt sets SubmitSmRespAt field to given value.

### HasSubmitSmRespAt

`func (o *MessageItem) HasSubmitSmRespAt() bool`

HasSubmitSmRespAt returns a boolean if a field has been set.

### GetDeliverySmAt

`func (o *MessageItem) GetDeliverySmAt() time.Time`

GetDeliverySmAt returns the DeliverySmAt field if non-nil, zero value otherwise.

### GetDeliverySmAtOk

`func (o *MessageItem) GetDeliverySmAtOk() (*time.Time, bool)`

GetDeliverySmAtOk returns a tuple with the DeliverySmAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySmAt

`func (o *MessageItem) SetDeliverySmAt(v time.Time)`

SetDeliverySmAt sets DeliverySmAt field to given value.

### HasDeliverySmAt

`func (o *MessageItem) HasDeliverySmAt() bool`

HasDeliverySmAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MessageItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessageItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MessageItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MessageItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


