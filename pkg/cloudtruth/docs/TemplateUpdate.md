# TemplateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | The template name. | 
**Description** | Pointer to **string** | (&#39;A description of the template.  You may find it helpful to document how this template is used to assist others when they need to maintain software that uses this content.&#39;,) | [optional] 
**Evaluated** | **bool** | If true, the &#x60;body&#x60; field has undergone evaluation. | [readonly] 
**Body** | Pointer to **string** | The content of the template.  Use mustache-style templating delimiters of &#x60;{{&#x60; and &#x60;}}&#x60; to reference parameter values by name for substitution into the template result. | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewTemplateUpdate

`func NewTemplateUpdate(id string, name string, evaluated bool, createdAt time.Time, modifiedAt NullableTime, ) *TemplateUpdate`

NewTemplateUpdate instantiates a new TemplateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateUpdateWithDefaults

`func NewTemplateUpdateWithDefaults() *TemplateUpdate`

NewTemplateUpdateWithDefaults instantiates a new TemplateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TemplateUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TemplateUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluated

`func (o *TemplateUpdate) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *TemplateUpdate) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *TemplateUpdate) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.


### GetBody

`func (o *TemplateUpdate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateUpdate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateUpdate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *TemplateUpdate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *TemplateUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TemplateUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TemplateUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *TemplateUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *TemplateUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


