/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BundleIdRelationshipsProfiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdRelationshipsProfiles{}

// BundleIdRelationshipsProfiles struct for BundleIdRelationshipsProfiles
type BundleIdRelationshipsProfiles struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []BundleIdRelationshipsProfilesDataInner `json:"data,omitempty"`
}

// NewBundleIdRelationshipsProfiles instantiates a new BundleIdRelationshipsProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdRelationshipsProfiles() *BundleIdRelationshipsProfiles {
	this := BundleIdRelationshipsProfiles{}
	return &this
}

// NewBundleIdRelationshipsProfilesWithDefaults instantiates a new BundleIdRelationshipsProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdRelationshipsProfilesWithDefaults() *BundleIdRelationshipsProfiles {
	this := BundleIdRelationshipsProfiles{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BundleIdRelationshipsProfiles) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsProfiles) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BundleIdRelationshipsProfiles) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *BundleIdRelationshipsProfiles) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BundleIdRelationshipsProfiles) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsProfiles) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BundleIdRelationshipsProfiles) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BundleIdRelationshipsProfiles) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BundleIdRelationshipsProfiles) GetData() []BundleIdRelationshipsProfilesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []BundleIdRelationshipsProfilesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsProfiles) GetDataOk() ([]BundleIdRelationshipsProfilesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BundleIdRelationshipsProfiles) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BundleIdRelationshipsProfilesDataInner and assigns it to the Data field.
func (o *BundleIdRelationshipsProfiles) SetData(v []BundleIdRelationshipsProfilesDataInner) {
	o.Data = v
}

func (o BundleIdRelationshipsProfiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdRelationshipsProfiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBundleIdRelationshipsProfiles struct {
	value *BundleIdRelationshipsProfiles
	isSet bool
}

func (v NullableBundleIdRelationshipsProfiles) Get() *BundleIdRelationshipsProfiles {
	return v.value
}

func (v *NullableBundleIdRelationshipsProfiles) Set(val *BundleIdRelationshipsProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdRelationshipsProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdRelationshipsProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdRelationshipsProfiles(val *BundleIdRelationshipsProfiles) *NullableBundleIdRelationshipsProfiles {
	return &NullableBundleIdRelationshipsProfiles{value: val, isSet: true}
}

func (v NullableBundleIdRelationshipsProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdRelationshipsProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


