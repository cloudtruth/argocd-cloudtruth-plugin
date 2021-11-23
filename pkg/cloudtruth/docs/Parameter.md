# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | A unique identifier for the parameter. | [readonly] 
**Name** | **string** | The parameter name. | 
**Description** | Pointer to **string** | A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content. | [optional] 
**Secret** | Pointer to **bool** | Indicates if this content is secret or not.  When a parameter is considered to be a secret, any internal values are stored in a dedicated vault for your organization for maximum security.  External values are inspected on-demand to ensure they align with the parameter&#39;s secret setting and if they do not, those external values are not allowed to be used. | [optional] 
**Type** | Pointer to [**ParameterTypeEnum**](ParameterTypeEnum.md) |  | [optional] 
**Rules** | [**[]ParameterRule**](ParameterRule.md) | Rules applied to this parameter. | [readonly] 
**Project** | **string** | The project that the parameter is within. | [readonly] 
**ProjectName** | **string** | The project name that the parameter is within | [readonly] 
**ReferencingTemplates** | **[]string** | Templates that reference this Parameter. | [readonly] 
**ReferencingValues** | **[]string** | Dynamic values that reference this Parameter. | [readonly] 
**Values** | [**map[string]Value**](Value.md) |              Each parameter has an effective value in every environment based on             environment inheritance and which environments have had a value set.              Environments inherit from a single parent to form a tree, as a result             a single parameter may have different values present for each environment.             When a value is not explicitly set in an environment, the parent environment             is consulted to see if it has a value defined, and so on.              The dictionary of values has an environment url as the key, and the optional             Value record that it resolves to.  If the Value.environment matches the key,             then it is an explicit value set for that environment.  If they differ, the             value was obtained from a parent environment (directly or indirectly).  If the             value is None then no value has ever been set in any environment for this             parameter.              key: Environment url             value: optional Value record          | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewParameter

`func NewParameter(url string, id string, name string, rules []ParameterRule, project string, projectName string, referencingTemplates []string, referencingValues []string, values map[string]Value, createdAt time.Time, modifiedAt time.Time, ) *Parameter`

NewParameter instantiates a new Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterWithDefaults

`func NewParameterWithDefaults() *Parameter`

NewParameterWithDefaults instantiates a new Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Parameter) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Parameter) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Parameter) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Parameter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Parameter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Parameter) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Parameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Parameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Parameter) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Parameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Parameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Parameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Parameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *Parameter) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Parameter) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Parameter) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Parameter) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *Parameter) GetType() ParameterTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Parameter) GetTypeOk() (*ParameterTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Parameter) SetType(v ParameterTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *Parameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRules

`func (o *Parameter) GetRules() []ParameterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Parameter) GetRulesOk() (*[]ParameterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Parameter) SetRules(v []ParameterRule)`

SetRules sets Rules field to given value.


### GetProject

`func (o *Parameter) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Parameter) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Parameter) SetProject(v string)`

SetProject sets Project field to given value.


### GetProjectName

`func (o *Parameter) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Parameter) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Parameter) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReferencingTemplates

`func (o *Parameter) GetReferencingTemplates() []string`

GetReferencingTemplates returns the ReferencingTemplates field if non-nil, zero value otherwise.

### GetReferencingTemplatesOk

`func (o *Parameter) GetReferencingTemplatesOk() (*[]string, bool)`

GetReferencingTemplatesOk returns a tuple with the ReferencingTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingTemplates

`func (o *Parameter) SetReferencingTemplates(v []string)`

SetReferencingTemplates sets ReferencingTemplates field to given value.


### GetReferencingValues

`func (o *Parameter) GetReferencingValues() []string`

GetReferencingValues returns the ReferencingValues field if non-nil, zero value otherwise.

### GetReferencingValuesOk

`func (o *Parameter) GetReferencingValuesOk() (*[]string, bool)`

GetReferencingValuesOk returns a tuple with the ReferencingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingValues

`func (o *Parameter) SetReferencingValues(v []string)`

SetReferencingValues sets ReferencingValues field to given value.


### GetValues

`func (o *Parameter) GetValues() map[string]Value`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Parameter) GetValuesOk() (*map[string]Value, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Parameter) SetValues(v map[string]Value)`

SetValues sets Values field to given value.


### GetCreatedAt

`func (o *Parameter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Parameter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Parameter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Parameter) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Parameter) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Parameter) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


