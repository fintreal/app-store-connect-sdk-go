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

// checks if the ProfileCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileCreateRequestDataRelationships{}

// ProfileCreateRequestDataRelationships struct for ProfileCreateRequestDataRelationships
type ProfileCreateRequestDataRelationships struct {
	BundleId BundleIdCapabilityCreateRequestDataRelationshipsBundleId `json:"bundleId"`
	Devices *ProfileRelationshipsDevices `json:"devices,omitempty"`
	Certificates ProfileCreateRequestDataRelationshipsCertificates `json:"certificates"`
}

type _ProfileCreateRequestDataRelationships ProfileCreateRequestDataRelationships

// NewProfileCreateRequestDataRelationships instantiates a new ProfileCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileCreateRequestDataRelationships(bundleId BundleIdCapabilityCreateRequestDataRelationshipsBundleId, certificates ProfileCreateRequestDataRelationshipsCertificates) *ProfileCreateRequestDataRelationships {
	this := ProfileCreateRequestDataRelationships{}
	this.BundleId = bundleId
	this.Certificates = certificates
	return &this
}

// NewProfileCreateRequestDataRelationshipsWithDefaults instantiates a new ProfileCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileCreateRequestDataRelationshipsWithDefaults() *ProfileCreateRequestDataRelationships {
	this := ProfileCreateRequestDataRelationships{}
	return &this
}

// GetBundleId returns the BundleId field value
func (o *ProfileCreateRequestDataRelationships) GetBundleId() BundleIdCapabilityCreateRequestDataRelationshipsBundleId {
	if o == nil {
		var ret BundleIdCapabilityCreateRequestDataRelationshipsBundleId
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *ProfileCreateRequestDataRelationships) GetBundleIdOk() (*BundleIdCapabilityCreateRequestDataRelationshipsBundleId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *ProfileCreateRequestDataRelationships) SetBundleId(v BundleIdCapabilityCreateRequestDataRelationshipsBundleId) {
	o.BundleId = v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *ProfileCreateRequestDataRelationships) GetDevices() ProfileRelationshipsDevices {
	if o == nil || IsNil(o.Devices) {
		var ret ProfileRelationshipsDevices
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCreateRequestDataRelationships) GetDevicesOk() (*ProfileRelationshipsDevices, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *ProfileCreateRequestDataRelationships) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given ProfileRelationshipsDevices and assigns it to the Devices field.
func (o *ProfileCreateRequestDataRelationships) SetDevices(v ProfileRelationshipsDevices) {
	o.Devices = &v
}

// GetCertificates returns the Certificates field value
func (o *ProfileCreateRequestDataRelationships) GetCertificates() ProfileCreateRequestDataRelationshipsCertificates {
	if o == nil {
		var ret ProfileCreateRequestDataRelationshipsCertificates
		return ret
	}

	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value
// and a boolean to check if the value has been set.
func (o *ProfileCreateRequestDataRelationships) GetCertificatesOk() (*ProfileCreateRequestDataRelationshipsCertificates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificates, true
}

// SetCertificates sets field value
func (o *ProfileCreateRequestDataRelationships) SetCertificates(v ProfileCreateRequestDataRelationshipsCertificates) {
	o.Certificates = v
}

func (o ProfileCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bundleId"] = o.BundleId
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	toSerialize["certificates"] = o.Certificates
	return toSerialize, nil
}

func (o *ProfileCreateRequestDataRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bundleId",
		"certificates",
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

	varProfileCreateRequestDataRelationships := _ProfileCreateRequestDataRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varProfileCreateRequestDataRelationships)

	if err != nil {
		return err
	}

	*o = ProfileCreateRequestDataRelationships(varProfileCreateRequestDataRelationships)

	return err
}

type NullableProfileCreateRequestDataRelationships struct {
	value *ProfileCreateRequestDataRelationships
	isSet bool
}

func (v NullableProfileCreateRequestDataRelationships) Get() *ProfileCreateRequestDataRelationships {
	return v.value
}

func (v *NullableProfileCreateRequestDataRelationships) Set(val *ProfileCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileCreateRequestDataRelationships(val *ProfileCreateRequestDataRelationships) *NullableProfileCreateRequestDataRelationships {
	return &NullableProfileCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableProfileCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


