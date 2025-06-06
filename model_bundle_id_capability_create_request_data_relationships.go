/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BundleIdCapabilityCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdCapabilityCreateRequestDataRelationships{}

// BundleIdCapabilityCreateRequestDataRelationships struct for BundleIdCapabilityCreateRequestDataRelationships
type BundleIdCapabilityCreateRequestDataRelationships struct {
	BundleId BundleIdCapabilityCreateRequestDataRelationshipsBundleId `json:"bundleId"`
}

type _BundleIdCapabilityCreateRequestDataRelationships BundleIdCapabilityCreateRequestDataRelationships

// NewBundleIdCapabilityCreateRequestDataRelationships instantiates a new BundleIdCapabilityCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityCreateRequestDataRelationships(bundleId BundleIdCapabilityCreateRequestDataRelationshipsBundleId) *BundleIdCapabilityCreateRequestDataRelationships {
	this := BundleIdCapabilityCreateRequestDataRelationships{}
	this.BundleId = bundleId
	return &this
}

// NewBundleIdCapabilityCreateRequestDataRelationshipsWithDefaults instantiates a new BundleIdCapabilityCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityCreateRequestDataRelationshipsWithDefaults() *BundleIdCapabilityCreateRequestDataRelationships {
	this := BundleIdCapabilityCreateRequestDataRelationships{}
	return &this
}

// GetBundleId returns the BundleId field value
func (o *BundleIdCapabilityCreateRequestDataRelationships) GetBundleId() BundleIdCapabilityCreateRequestDataRelationshipsBundleId {
	if o == nil {
		var ret BundleIdCapabilityCreateRequestDataRelationshipsBundleId
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityCreateRequestDataRelationships) GetBundleIdOk() (*BundleIdCapabilityCreateRequestDataRelationshipsBundleId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *BundleIdCapabilityCreateRequestDataRelationships) SetBundleId(v BundleIdCapabilityCreateRequestDataRelationshipsBundleId) {
	o.BundleId = v
}

func (o BundleIdCapabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdCapabilityCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bundleId"] = o.BundleId
	return toSerialize, nil
}

func (o *BundleIdCapabilityCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bundleId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBundleIdCapabilityCreateRequestDataRelationships := _BundleIdCapabilityCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBundleIdCapabilityCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = BundleIdCapabilityCreateRequestDataRelationships(varBundleIdCapabilityCreateRequestDataRelationships)

	return err
}

type NullableBundleIdCapabilityCreateRequestDataRelationships struct {
	value *BundleIdCapabilityCreateRequestDataRelationships
	isSet bool
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationships) Get() *BundleIdCapabilityCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationships) Set(val *BundleIdCapabilityCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityCreateRequestDataRelationships(val *BundleIdCapabilityCreateRequestDataRelationships) *NullableBundleIdCapabilityCreateRequestDataRelationships {
	return &NullableBundleIdCapabilityCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


