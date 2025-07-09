# PaginationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PrevPageUrl** | Pointer to **NullableString** |  | [optional] 
**FirstPageUrl** | Pointer to **string** |  | [optional] 
**LastPageUrl** | Pointer to **string** |  | [optional] 
**NextPageUrl** | Pointer to **NullableString** |  | [optional] 
**PerPage** | Pointer to **int** |  | [optional] 
**LastPage** | Pointer to **int** |  | [optional] 
**From** | Pointer to **int** |  | [optional] 
**To** | Pointer to **int** |  | [optional] 
**Total** | Pointer to **int** |  | [optional] 
**Result** | Pointer to [**[]MessageItem**](MessageItem.md) |  | [optional] 
**Links** | Pointer to [**[]LinkItem**](LinkItem.md) |  | [optional] 

## Methods

### NewPaginationData

`func NewPaginationData() *PaginationData`

NewPaginationData instantiates a new PaginationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationDataWithDefaults

`func NewPaginationDataWithDefaults() *PaginationData`

NewPaginationDataWithDefaults instantiates a new PaginationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PaginationData) GetCurrentPage() int`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PaginationData) GetCurrentPageOk() (*int, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PaginationData) SetCurrentPage(v int)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *PaginationData) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetPath

`func (o *PaginationData) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PaginationData) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PaginationData) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PaginationData) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPrevPageUrl

`func (o *PaginationData) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *PaginationData) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *PaginationData) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.

### HasPrevPageUrl

`func (o *PaginationData) HasPrevPageUrl() bool`

HasPrevPageUrl returns a boolean if a field has been set.

### SetPrevPageUrlNil

`func (o *PaginationData) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *PaginationData) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetFirstPageUrl

`func (o *PaginationData) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *PaginationData) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *PaginationData) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.

### HasFirstPageUrl

`func (o *PaginationData) HasFirstPageUrl() bool`

HasFirstPageUrl returns a boolean if a field has been set.

### GetLastPageUrl

`func (o *PaginationData) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *PaginationData) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *PaginationData) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.

### HasLastPageUrl

`func (o *PaginationData) HasLastPageUrl() bool`

HasLastPageUrl returns a boolean if a field has been set.

### GetNextPageUrl

`func (o *PaginationData) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *PaginationData) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *PaginationData) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.

### HasNextPageUrl

`func (o *PaginationData) HasNextPageUrl() bool`

HasNextPageUrl returns a boolean if a field has been set.

### SetNextPageUrlNil

`func (o *PaginationData) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *PaginationData) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPerPage

`func (o *PaginationData) GetPerPage() int`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *PaginationData) GetPerPageOk() (*int, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *PaginationData) SetPerPage(v int)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *PaginationData) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetLastPage

`func (o *PaginationData) GetLastPage() int`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *PaginationData) GetLastPageOk() (*int, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *PaginationData) SetLastPage(v int)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *PaginationData) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetFrom

`func (o *PaginationData) GetFrom() int`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PaginationData) GetFromOk() (*int, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PaginationData) SetFrom(v int)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PaginationData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *PaginationData) GetTo() int`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PaginationData) GetToOk() (*int, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PaginationData) SetTo(v int)`

SetTo sets To field to given value.

### HasTo

`func (o *PaginationData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotal

`func (o *PaginationData) GetTotal() int`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationData) GetTotalOk() (*int, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationData) SetTotal(v int)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginationData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResult

`func (o *PaginationData) GetResult() []MessageItem`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PaginationData) GetResultOk() (*[]MessageItem, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PaginationData) SetResult(v []MessageItem)`

SetResult sets Result field to given value.

### HasResult

`func (o *PaginationData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetLinks

`func (o *PaginationData) GetLinks() []LinkItem`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginationData) GetLinksOk() (*[]LinkItem, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginationData) SetLinks(v []LinkItem)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginationData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


