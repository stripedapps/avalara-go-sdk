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

// checks if the FetchResultUPCModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchResultUPCModel{}

// FetchResultUPCModel Represents a fetch result
type FetchResultUPCModel struct {
	// The number of rows returned by your query, prior to pagination.
	RecordsetCount *int32 `json:"@recordsetCount,omitempty"`
	// The paginated and filtered list of records matching the parameters you supplied.
	Value []UPCModel `json:"value,omitempty"`
	// The link to the next page of results
	NextLink *string `json:"@nextLink,omitempty"`
}

// NewFetchResultUPCModel instantiates a new FetchResultUPCModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchResultUPCModel() *FetchResultUPCModel {
	this := FetchResultUPCModel{}
	return &this
}

// NewFetchResultUPCModelWithDefaults instantiates a new FetchResultUPCModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchResultUPCModelWithDefaults() *FetchResultUPCModel {
	this := FetchResultUPCModel{}
	return &this
}

// GetRecordsetCount returns the RecordsetCount field value if set, zero value otherwise.
func (o *FetchResultUPCModel) GetRecordsetCount() int32 {
	if o == nil || IsNil(o.RecordsetCount) {
		var ret int32
		return ret
	}
	return *o.RecordsetCount
}

// GetRecordsetCountOk returns a tuple with the RecordsetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchResultUPCModel) GetRecordsetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsetCount) {
		return nil, false
	}
	return o.RecordsetCount, true
}

// HasRecordsetCount returns a boolean if a field has been set.
func (o *FetchResultUPCModel) HasRecordsetCount() bool {
	if o != nil && !IsNil(o.RecordsetCount) {
		return true
	}

	return false
}

// SetRecordsetCount gets a reference to the given int32 and assigns it to the RecordsetCount field.
func (o *FetchResultUPCModel) SetRecordsetCount(v int32) {
	o.RecordsetCount = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FetchResultUPCModel) GetValue() []UPCModel {
	if o == nil || IsNil(o.Value) {
		var ret []UPCModel
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchResultUPCModel) GetValueOk() ([]UPCModel, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FetchResultUPCModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []UPCModel and assigns it to the Value field.
func (o *FetchResultUPCModel) SetValue(v []UPCModel) {
	o.Value = v
}

// GetNextLink returns the NextLink field value if set, zero value otherwise.
func (o *FetchResultUPCModel) GetNextLink() string {
	if o == nil || IsNil(o.NextLink) {
		var ret string
		return ret
	}
	return *o.NextLink
}

// GetNextLinkOk returns a tuple with the NextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchResultUPCModel) GetNextLinkOk() (*string, bool) {
	if o == nil || IsNil(o.NextLink) {
		return nil, false
	}
	return o.NextLink, true
}

// HasNextLink returns a boolean if a field has been set.
func (o *FetchResultUPCModel) HasNextLink() bool {
	if o != nil && !IsNil(o.NextLink) {
		return true
	}

	return false
}

// SetNextLink gets a reference to the given string and assigns it to the NextLink field.
func (o *FetchResultUPCModel) SetNextLink(v string) {
	o.NextLink = &v
}

func (o FetchResultUPCModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchResultUPCModel) ToMap() (map[string]interface{}, error) {
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

type NullableFetchResultUPCModel struct {
	value *FetchResultUPCModel
	isSet bool
}

func (v NullableFetchResultUPCModel) Get() *FetchResultUPCModel {
	return v.value
}

func (v *NullableFetchResultUPCModel) Set(val *FetchResultUPCModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchResultUPCModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchResultUPCModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchResultUPCModel(val *FetchResultUPCModel) *NullableFetchResultUPCModel {
	return &NullableFetchResultUPCModel{value: val, isSet: true}
}

func (v NullableFetchResultUPCModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchResultUPCModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


