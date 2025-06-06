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

// checks if the AppAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAttributes{}

// AppAttributes struct for AppAttributes
type AppAttributes struct {
	Name *string `json:"name,omitempty"`
	BundleId *string `json:"bundleId,omitempty"`
	Sku *string `json:"sku,omitempty"`
	PrimaryLocale *string `json:"primaryLocale,omitempty"`
}

// NewAppAttributes instantiates a new AppAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAttributes() *AppAttributes {
	this := AppAttributes{}
	return &this
}

// NewAppAttributesWithDefaults instantiates a new AppAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAttributesWithDefaults() *AppAttributes {
	this := AppAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppAttributes) SetName(v string) {
	o.Name = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AppAttributes) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AppAttributes) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AppAttributes) SetBundleId(v string) {
	o.BundleId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *AppAttributes) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *AppAttributes) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *AppAttributes) SetSku(v string) {
	o.Sku = &v
}

// GetPrimaryLocale returns the PrimaryLocale field value if set, zero value otherwise.
func (o *AppAttributes) GetPrimaryLocale() string {
	if o == nil || IsNil(o.PrimaryLocale) {
		var ret string
		return ret
	}
	return *o.PrimaryLocale
}

// GetPrimaryLocaleOk returns a tuple with the PrimaryLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAttributes) GetPrimaryLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryLocale) {
		return nil, false
	}
	return o.PrimaryLocale, true
}

// HasPrimaryLocale returns a boolean if a field has been set.
func (o *AppAttributes) HasPrimaryLocale() bool {
	if o != nil && !IsNil(o.PrimaryLocale) {
		return true
	}

	return false
}

// SetPrimaryLocale gets a reference to the given string and assigns it to the PrimaryLocale field.
func (o *AppAttributes) SetPrimaryLocale(v string) {
	o.PrimaryLocale = &v
}

func (o AppAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.PrimaryLocale) {
		toSerialize["primaryLocale"] = o.PrimaryLocale
	}
	return toSerialize, nil
}

type NullableAppAttributes struct {
	value *AppAttributes
	isSet bool
}

func (v NullableAppAttributes) Get() *AppAttributes {
	return v.value
}

func (v *NullableAppAttributes) Set(val *AppAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAttributes(val *AppAttributes) *NullableAppAttributes {
	return &NullableAppAttributes{value: val, isSet: true}
}

func (v NullableAppAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


