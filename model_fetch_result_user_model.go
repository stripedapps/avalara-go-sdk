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

// checks if the FetchResultUserModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchResultUserModel{}

// FetchResultUserModel Represents a fetch result
type FetchResultUserModel struct {
	// The number of rows returned by your query, prior to pagination.
	RecordsetCount *int32 `json:"@recordsetCount,omitempty"`
	// The paginated and filtered list of records matching the parameters you supplied.
	Value []UserModel `json:"value,omitempty"`
	// The link to the next page of results
	NextLink *string `json:"@nextLink,omitempty"`
}

// NewFetchResultUserModel instantiates a new FetchResultUserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchResultUserModel() *FetchResultUserModel {
	this := FetchResultUserModel{}
	return &this
}

// NewFetchResultUserModelWithDefaults instantiates a new FetchResultUserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchResultUserModelWithDefaults() *FetchResultUserModel {
	this := FetchResultUserModel{}
	return &this
}

// GetRecordsetCount returns the RecordsetCount field value if set, zero value otherwise.
func (o *FetchResultUserModel) GetRecordsetCount() int32 {
	if o == nil || IsNil(o.RecordsetCount) {
		var ret int32
		return ret
	}
	return *o.RecordsetCount
}

// GetRecordsetCountOk returns a tuple with the RecordsetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchResultUserModel) GetRecordsetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsetCount) {
		return nil, false
	}
	return o.RecordsetCount, true
}

// HasRecordsetCount returns a boolean if a field has been set.
func (o *FetchResultUserModel) HasRecordsetCount() bool {
	if o != nil && !IsNil(o.RecordsetCount) {
		return true
	}

	return false
}

// SetRecordsetCount gets a reference to the given int32 and assigns it to the RecordsetCount field.
func (o *FetchResultUserModel) SetRecordsetCount(v int32) {
	o.RecordsetCount = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FetchResultUserModel) GetValue() []UserModel {
	if o == nil || IsNil(o.Value) {
		var ret []UserModel
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchResultUserModel) GetValueOk() ([]UserModel, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FetchResultUserModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []UserModel and assigns it to the Value field.
func (o *FetchResultUserModel) SetValue(v []UserModel) {
	o.Value = v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *FetchResultUserModel) GetNextLink() string {
	if o == nil || IsNil(o.NextLink) {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchResultUserModel) GetNextLinkOk() (*string, bool) {
	if o == nil || IsNil(o.NextLink) {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *FetchResultUserModel) HasNextLink() bool {
	if o != nil && !IsNil(o.NextLink) {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *FetchResultUserModel) SetNextLink(v string) {
	o.NextLink = &v
}

func (o FetchResultUserModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchResultUserModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordsetCount) {
		toSerialize["@recordsetCount"] = o.RecordsetCount
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.NextLink) {
		toSerialize["@nextLink"] = o.NextLink
	}
	return toSerialize, nil
}

type NullableFetchResultUserModel struct {
	value *FetchResultUserModel
	isSet bool
}

func (v NullableFetchResultUserModel) Get() *FetchResultUserModel {
	return v.value
}

func (v *NullableFetchResultUserModel) Set(val *FetchResultUserModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchResultUserModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchResultUserModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchResultUserModel(val *FetchResultUserModel) *NullableFetchResultUserModel {
	return &NullableFetchResultUserModel{value: val, isSet: true}
}

func (v NullableFetchResultUserModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchResultUserModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

