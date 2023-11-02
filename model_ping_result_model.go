/*
AvaTax API

REST interface to Avalara's enterprise tax service, AvaTax.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PingResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingResultModel{}

// PingResultModel Ping Result Model
type PingResultModel struct {
	// Version number
	Version *string `json:"version,omitempty"`
}

// NewPingResultModel instantiates a new PingResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingResultModel() *PingResultModel {
	this := PingResultModel{}
	return &this
}

// NewPingResultModelWithDefaults instantiates a new PingResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingResultModelWithDefaults() *PingResultModel {
	this := PingResultModel{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PingResultModel) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingResultModel) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PingResultModel) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PingResultModel) SetVersion(v string) {
	o.Version = &v
}

func (o PingResultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingResultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePingResultModel struct {
	value *PingResultModel
	isSet bool
}

func (v NullablePingResultModel) Get() *PingResultModel {
	return v.value
}

func (v *NullablePingResultModel) Set(val *PingResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePingResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePingResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingResultModel(val *PingResultModel) *NullablePingResultModel {
	return &NullablePingResultModel{value: val, isSet: true}
}

func (v NullablePingResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


