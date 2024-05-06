/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.7 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RackRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackRole{}

// RackRole Adds support for custom fields and tags.
type RackRole struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	Name                 string                 `json:"name"`
	Slug                 string                 `json:"slug"`
	Color                *string                `json:"color,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	RackCount            int32                  `json:"rack_count"`
	AdditionalProperties map[string]interface{}
}

type _RackRole RackRole

// NewRackRole instantiates a new RackRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackRole(id int32, url string, display string, name string, slug string, created NullableTime, lastUpdated NullableTime, rackCount int32) *RackRole {
	this := RackRole{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Created = created
	this.LastUpdated = lastUpdated
	this.RackCount = rackCount
	return &this
}

// NewRackRoleWithDefaults instantiates a new RackRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackRoleWithDefaults() *RackRole {
	this := RackRole{}
	return &this
}

// GetId returns the Id field value
func (o *RackRole) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RackRole) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *RackRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RackRole) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *RackRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *RackRole) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *RackRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RackRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *RackRole) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *RackRole) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *RackRole) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *RackRole) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *RackRole) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RackRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RackRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RackRole) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RackRole) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RackRole) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *RackRole) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RackRole) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackRole) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RackRole) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RackRole) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RackRole) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRole) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *RackRole) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RackRole) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackRole) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *RackRole) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetRackCount returns the RackCount field value
func (o *RackRole) GetRackCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value
// and a boolean to check if the value has been set.
func (o *RackRole) GetRackCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RackCount, true
}

// SetRackCount sets field value
func (o *RackRole) SetRackCount(v int32) {
	o.RackCount = v
}

func (o RackRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["rack_count"] = o.RackCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RackRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"created",
		"last_updated",
		"rack_count",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRackRole := _RackRole{}

	err = json.Unmarshal(data, &varRackRole)

	if err != nil {
		return err
	}

	*o = RackRole(varRackRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "rack_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackRole struct {
	value *RackRole
	isSet bool
}

func (v NullableRackRole) Get() *RackRole {
	return v.value
}

func (v *NullableRackRole) Set(val *RackRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRackRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRackRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackRole(val *RackRole) *NullableRackRole {
	return &NullableRackRole{value: val, isSet: true}
}

func (v NullableRackRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}