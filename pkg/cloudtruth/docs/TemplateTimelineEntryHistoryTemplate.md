# TemplateTimelineEntryHistoryTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**LedgerId** | **string** |  | [readonly] 
**Name** | **string** | The parameter name. | 
**Description** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** | The content of the template.  Use mustache-style templating delimiters of &#x60;{{&#x60; and &#x60;}}&#x60; to reference parameter values by name for substitution into the template result. | [optional] 

## Methods

### NewTemplateTimelineEntryHistoryTemplate

`func NewTemplateTimelineEntryHistoryTemplate(id string, ledgerId string, name string, ) *TemplateTimelineEntryHistoryTemplate`

NewTemplateTimelineEntryHistoryTemplate instantiates a new TemplateTimelineEntryHistoryTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateTimelineEntryHistoryTemplateWithDefaults

`func NewTemplateTimelineEntryHistoryTemplateWithDefaults() *TemplateTimelineEntryHistoryTemplate`

NewTemplateTimelineEntryHistoryTemplateWithDefaults instantiates a new TemplateTimelineEntryHistoryTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateTimelineEntryHistoryTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateTimelineEntryHistoryTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateTimelineEntryHistoryTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetLedgerId

`func (o *TemplateTimelineEntryHistoryTemplate) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *TemplateTimelineEntryHistoryTemplate) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *TemplateTimelineEntryHistoryTemplate) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.


### GetName

`func (o *TemplateTimelineEntryHistoryTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateTimelineEntryHistoryTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateTimelineEntryHistoryTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TemplateTimelineEntryHistoryTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateTimelineEntryHistoryTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateTimelineEntryHistoryTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateTimelineEntryHistoryTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBody

`func (o *TemplateTimelineEntryHistoryTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateTimelineEntryHistoryTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateTimelineEntryHistoryTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TemplateTimelineEntryHistoryTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

