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

// checks if the BundleIdRelationshipsApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdRelationshipsApp{}

// BundleIdRelationshipsApp struct for BundleIdRelationshipsApp
type BundleIdRelationshipsApp struct {
	Links *RelationshipLinks `json:"links,omitempty"`
	Data *BundleIdRelationshipsAppData `json:"data,omitempty"`
}

// NewBundleIdRelationshipsApp instantiates a new BundleIdRelationshipsApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdRelationshipsApp() *BundleIdRelationshipsApp {
	this := BundleIdRelationshipsApp{}
	return &this
}

// NewBundleIdRelationshipsAppWithDefaults instantiates a new BundleIdRelationshipsApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdRelationshipsAppWithDefaults() *BundleIdRelationshipsApp {
	this := BundleIdRelationshipsApp{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BundleIdRelationshipsApp) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsApp) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BundleIdRelationshipsApp) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *BundleIdRelationshipsApp) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BundleIdRelationshipsApp) GetData() BundleIdRelationshipsAppData {
	if o == nil || IsNil(o.Data) {
		var ret BundleIdRelationshipsAppData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationshipsApp) GetDataOk() (*BundleIdRelationshipsAppData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BundleIdRelationshipsApp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BundleIdRelationshipsAppData and assigns it to the Data field.
func (o *BundleIdRelationshipsApp) SetData(v BundleIdRelationshipsAppData) {
	o.Data = &v
}

func (o BundleIdRelationshipsApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdRelationshipsApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBundleIdRelationshipsApp struct {
	value *BundleIdRelationshipsApp
	isSet bool
}

func (v NullableBundleIdRelationshipsApp) Get() *BundleIdRelationshipsApp {
	return v.value
}

func (v *NullableBundleIdRelationshipsApp) Set(val *BundleIdRelationshipsApp) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdRelationshipsApp) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdRelationshipsApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdRelationshipsApp(val *BundleIdRelationshipsApp) *NullableBundleIdRelationshipsApp {
	return &NullableBundleIdRelationshipsApp{value: val, isSet: true}
}

func (v NullableBundleIdRelationshipsApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdRelationshipsApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


