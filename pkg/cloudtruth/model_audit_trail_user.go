/*
CloudTruth Management API

CloudTruth centralizes your configuration parameters and secrets making them easier to manage and use as a team.

API version: v1
Contact: support@cloudtruth.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudtruth

import (
	"encoding/json"
	"time"
)

// AuditTrailUser Details of the user associated with this change.
type AuditTrailUser struct {
	Url string `json:"url"`
	// The unique identifier of a user.
	Id string `json:"id"`
	// The type of user record.
	Type *string `json:"type,omitempty"`
	Name NullableString `json:"name"`
	// The user's organization name.
	OrganizationName NullableString `json:"organization_name"`
	// Membership identifier for user.
	MembershipId NullableString `json:"membership_id"`
	// The user's role in the current organization (defined by the request authorization header).
	Role NullableString `json:"role"`
	Email NullableString `json:"email"`
	PictureUrl NullableString `json:"picture_url"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt NullableTime `json:"modified_at"`
}

// NewAuditTrailUser instantiates a new AuditTrailUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditTrailUser(url string, id string, name NullableString, organizationName NullableString, membershipId NullableString, role NullableString, email NullableString, pictureUrl NullableString, createdAt time.Time, modifiedAt NullableTime) *AuditTrailUser {
	this := AuditTrailUser{}
	this.Url = url
	this.Id = id
	this.Name = name
	this.OrganizationName = organizationName
	this.MembershipId = membershipId
	this.Role = role
	this.Email = email
	this.PictureUrl = pictureUrl
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewAuditTrailUserWithDefaults instantiates a new AuditTrailUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditTrailUserWithDefaults() *AuditTrailUser {
	this := AuditTrailUser{}
	return &this
}

// GetUrl returns the Url field value
func (o *AuditTrailUser) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AuditTrailUser) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AuditTrailUser) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *AuditTrailUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuditTrailUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuditTrailUser) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuditTrailUser) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditTrailUser) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuditTrailUser) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuditTrailUser) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditTrailUser) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *AuditTrailUser) SetName(v string) {
	o.Name.Set(&v)
}

// GetOrganizationName returns the OrganizationName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditTrailUser) GetOrganizationName() string {
	if o == nil || o.OrganizationName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailUser) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// SetOrganizationName sets field value
func (o *AuditTrailUser) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// GetMembershipId returns the MembershipId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditTrailUser) GetMembershipId() string {
	if o == nil || o.MembershipId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MembershipId.Get()
}

// GetMembershipIdOk returns a tuple with the MembershipId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailUser) GetMembershipIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MembershipId.Get(), o.MembershipId.IsSet()
}

// SetMembershipId sets field value
func (o *AuditTrailUser) SetMembershipId(v string) {
	o.MembershipId.Set(&v)
}

// GetRole returns the Role field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditTrailUser) GetRole() string {
	if o == nil || o.Role.Get() == nil {
		var ret string
		return ret
	}

	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailUser) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// SetRole sets field value
func (o *AuditTrailUser) SetRole(v string) {
	o.Role.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditTrailUser) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *AuditTrailUser) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetPictureUrl returns the PictureUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuditTrailUser) GetPictureUrl() string {
	if o == nil || o.PictureUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PictureUrl.Get()
}

// GetPictureUrlOk returns a tuple with the PictureUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailUser) GetPictureUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PictureUrl.Get(), o.PictureUrl.IsSet()
}

// SetPictureUrl sets field value
func (o *AuditTrailUser) SetPictureUrl(v string) {
	o.PictureUrl.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *AuditTrailUser) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AuditTrailUser) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AuditTrailUser) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AuditTrailUser) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailUser) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// SetModifiedAt sets field value
func (o *AuditTrailUser) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

func (o AuditTrailUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["organization_name"] = o.OrganizationName.Get()
	}
	if true {
		toSerialize["membership_id"] = o.MembershipId.Get()
	}
	if true {
		toSerialize["role"] = o.Role.Get()
	}
	if true {
		toSerialize["email"] = o.Email.Get()
	}
	if true {
		toSerialize["picture_url"] = o.PictureUrl.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuditTrailUser struct {
	value *AuditTrailUser
	isSet bool
}

func (v NullableAuditTrailUser) Get() *AuditTrailUser {
	return v.value
}

func (v *NullableAuditTrailUser) Set(val *AuditTrailUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditTrailUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditTrailUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditTrailUser(val *AuditTrailUser) *NullableAuditTrailUser {
	return &NullableAuditTrailUser{value: val, isSet: true}
}

func (v NullableAuditTrailUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditTrailUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


