# PatchedTemplateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | The template name. | [optional] 
**Description** | Pointer to **string** | (&#39;A description of the template.  You may find it helpful to document how this template is used to assist others when they need to maintain software that uses this content.&#39;,) | [optional] 
**Evaluated** | Pointer to **bool** | If true, the &#x60;body&#x60; field has undergone evaluation. | [optional] [readonly] 
**Body** | Pointer to **string** | The content of the template.  Use mustache-style templating delimiters of &#x60;{{&#x60; and &#x60;}}&#x60; to reference parameter values by name for substitution into the template result. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedTemplateUpdate

`func NewPatchedTemplateUpdate() *PatchedTemplateUpdate`

NewPatchedTemplateUpdate instantiates a new PatchedTemplateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTemplateUpdateWithDefaults

`func NewPatchedTemplateUpdateWithDefaults() *PatchedTemplateUpdate`

NewPatchedTemplateUpdateWithDefaults instantiates a new PatchedTemplateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTemplateUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTemplateUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTemplateUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTemplateUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedTemplateUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTemplateUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTemplateUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTemplateUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTemplateUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTemplateUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTemplateUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTemplateUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluated

`func (o *PatchedTemplateUpdate) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *PatchedTemplateUpdate) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *PatchedTemplateUpdate) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.

### HasEvaluated

`func (o *PatchedTemplateUpdate) HasEvaluated() bool`

HasEvaluated returns a boolean if a field has been set.

### GetBody

`func (o *PatchedTemplateUpdate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PatchedTemplateUpdate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PatchedTemplateUpdate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PatchedTemplateUpdate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedTemplateUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedTemplateUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedTemplateUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedTemplateUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedTemplateUpdate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedTemplateUpdate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedTemplateUpdate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedTemplateUpdate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *PatchedTemplateUpdate) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *PatchedTemplateUpdate) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


