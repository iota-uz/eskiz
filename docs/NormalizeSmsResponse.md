# NormalizeSmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecialCharacters** | Pointer to [**[]SpecialCharacter**](SpecialCharacter.md) |  | [optional] 
**Recommendations** | Pointer to [**[]Recommendation**](Recommendation.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewNormalizeSmsResponse

`func NewNormalizeSmsResponse() *NormalizeSmsResponse`

NewNormalizeSmsResponse instantiates a new NormalizeSmsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNormalizeSmsResponseWithDefaults

`func NewNormalizeSmsResponseWithDefaults() *NormalizeSmsResponse`

NewNormalizeSmsResponseWithDefaults instantiates a new NormalizeSmsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecialCharacters

`func (o *NormalizeSmsResponse) GetSpecialCharacters() []SpecialCharacter`

GetSpecialCharacters returns the SpecialCharacters field if non-nil, zero value otherwise.

### GetSpecialCharactersOk

`func (o *NormalizeSmsResponse) GetSpecialCharactersOk() (*[]SpecialCharacter, bool)`

GetSpecialCharactersOk returns a tuple with the SpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharacters

`func (o *NormalizeSmsResponse) SetSpecialCharacters(v []SpecialCharacter)`

SetSpecialCharacters sets SpecialCharacters field to given value.

### HasSpecialCharacters

`func (o *NormalizeSmsResponse) HasSpecialCharacters() bool`

HasSpecialCharacters returns a boolean if a field has been set.

### GetRecommendations

`func (o *NormalizeSmsResponse) GetRecommendations() []Recommendation`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *NormalizeSmsResponse) GetRecommendationsOk() (*[]Recommendation, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *NormalizeSmsResponse) SetRecommendations(v []Recommendation)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *NormalizeSmsResponse) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetMessage

`func (o *NormalizeSmsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NormalizeSmsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NormalizeSmsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NormalizeSmsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


