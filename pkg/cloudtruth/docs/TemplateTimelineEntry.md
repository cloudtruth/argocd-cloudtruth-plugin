# TemplateTimelineEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryType** | [**HistoryTypeEnum**](HistoryTypeEnum.md) |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**HistoryTemplate** | [**TemplateTimelineEntryHistoryTemplate**](TemplateTimelineEntryHistoryTemplate.md) |  | 

## Methods

### NewTemplateTimelineEntry

`func NewTemplateTimelineEntry(historyType HistoryTypeEnum, modifiedAt NullableTime, historyTemplate TemplateTimelineEntryHistoryTemplate, ) *TemplateTimelineEntry`

NewTemplateTimelineEntry instantiates a new TemplateTimelineEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateTimelineEntryWithDefaults

`func NewTemplateTimelineEntryWithDefaults() *TemplateTimelineEntry`

NewTemplateTimelineEntryWithDefaults instantiates a new TemplateTimelineEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryType

`func (o *TemplateTimelineEntry) GetHistoryType() HistoryTypeEnum`

GetHistoryType returns the HistoryType field if non-nil, zero value otherwise.

### GetHistoryTypeOk

`func (o *TemplateTimelineEntry) GetHistoryTypeOk() (*HistoryTypeEnum, bool)`

GetHistoryTypeOk returns a tuple with the HistoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryType

`func (o *TemplateTimelineEntry) SetHistoryType(v HistoryTypeEnum)`

SetHistoryType sets HistoryType field to given value.


### GetModifiedAt

`func (o *TemplateTimelineEntry) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TemplateTimelineEntry) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TemplateTimelineEntry) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *TemplateTimelineEntry) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *TemplateTimelineEntry) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetModifiedBy

`func (o *TemplateTimelineEntry) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TemplateTimelineEntry) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TemplateTimelineEntry) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *TemplateTimelineEntry) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetHistoryTemplate

`func (o *TemplateTimelineEntry) GetHistoryTemplate() TemplateTimelineEntryHistoryTemplate`

GetHistoryTemplate returns the HistoryTemplate field if non-nil, zero value otherwise.

### GetHistoryTemplateOk

`func (o *TemplateTimelineEntry) GetHistoryTemplateOk() (*TemplateTimelineEntryHistoryTemplate, bool)`

GetHistoryTemplateOk returns a tuple with the HistoryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryTemplate

`func (o *TemplateTimelineEntry) SetHistoryTemplate(v TemplateTimelineEntryHistoryTemplate)`

SetHistoryTemplate sets HistoryTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


