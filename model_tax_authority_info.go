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

// checks if the TaxAuthorityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxAuthorityInfo{}

// TaxAuthorityInfo Tax Authority Info
type TaxAuthorityInfo struct {
	// Avalara Id
	AvalaraId *string `json:"avalaraId,omitempty"`
	// Jurisdiction Name
	JurisdictionName string `json:"jurisdictionName"`
	// Jurisdiction Type
	JurisdictionType *string `json:"jurisdictionType,omitempty"`
	// Signature Code
	SignatureCode *string `json:"signatureCode,omitempty"`
}

// NewTaxAuthorityInfo instantiates a new TaxAuthorityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxAuthorityInfo(jurisdictionName string) *TaxAuthorityInfo {
	this := TaxAuthorityInfo{}
	this.JurisdictionName = jurisdictionName
	return &this
}

// NewTaxAuthorityInfoWithDefaults instantiates a new TaxAuthorityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxAuthorityInfoWithDefaults() *TaxAuthorityInfo {
	this := TaxAuthorityInfo{}
	return &this
}

// GetAvalaraId returns the AvalaraId field value if set, zero value otherwise.
func (o *TaxAuthorityInfo) GetAvalaraId() string {
	if o == nil || IsNil(o.AvalaraId) {
		var ret string
		return ret
	}
	return *o.AvalaraId
}

// GetAvalaraIdOk returns a tuple with the AvalaraId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxAuthorityInfo) GetAvalaraIdOk() (*string, bool) {
	if o == nil || IsNil(o.AvalaraId) {
		return nil, false
	}
	return o.AvalaraId, true
}

// HasAvalaraId returns a boolean if a field has been set.
func (o *TaxAuthorityInfo) HasAvalaraId() bool {
	if o != nil && !IsNil(o.AvalaraId) {
		return true
	}

	return false
}

// SetAvalaraId gets a reference to the given string and assigns it to the AvalaraId field.
func (o *TaxAuthorityInfo) SetAvalaraId(v string) {
	o.AvalaraId = &v
}

// GetJurisdictionName returns the JurisdictionName field value
func (o *TaxAuthorityInfo) GetJurisdictionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JurisdictionName
}

// GetJurisdictionNameOk returns a tuple with the JurisdictionName field value
// and a boolean to check if the value has been set.
func (o *TaxAuthorityInfo) GetJurisdictionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JurisdictionName, true
}

// SetJurisdictionName sets field value
func (o *TaxAuthorityInfo) SetJurisdictionName(v string) {
	o.JurisdictionName = v
}

// GetJurisdictionType returns the JurisdictionType field value if set, zero value otherwise.
func (o *TaxAuthorityInfo) GetJurisdictionType() string {
	if o == nil || IsNil(o.JurisdictionType) {
		var ret string
		return ret
	}
	return *o.JurisdictionType
}

// GetJurisdictionTypeOk returns a tuple with the JurisdictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxAuthorityInfo) GetJurisdictionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.JurisdictionType) {
		return nil, false
	}
	return o.JurisdictionType, true
}

// HasJurisdictionType returns a boolean if a field has been set.
func (o *TaxAuthorityInfo) HasJurisdictionType() bool {
	if o != nil && !IsNil(o.JurisdictionType) {
		return true
	}

	return false
}

// SetJurisdictionType gets a reference to the given string and assigns it to the JurisdictionType field.
func (o *TaxAuthorityInfo) SetJurisdictionType(v string) {
	o.JurisdictionType = &v
}

// GetSignatureCode returns the SignatureCode field value if set, zero value otherwise.
func (o *TaxAuthorityInfo) GetSignatureCode() string {
	if o == nil || IsNil(o.SignatureCode) {
		var ret string
		return ret
	}
	return *o.SignatureCode
}

// GetSignatureCodeOk returns a tuple with the SignatureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxAuthorityInfo) GetSignatureCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureCode) {
		return nil, false
	}
	return o.SignatureCode, true
}

// HasSignatureCode returns a boolean if a field has been set.
func (o *TaxAuthorityInfo) HasSignatureCode() bool {
	if o != nil && !IsNil(o.SignatureCode) {
		return true
	}

	return false
}

// SetSignatureCode gets a reference to the given string and assigns it to the SignatureCode field.
func (o *TaxAuthorityInfo) SetSignatureCode(v string) {
	o.SignatureCode = &v
}

func (o TaxAuthorityInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxAuthorityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvalaraId) {
		toSerialize["avalaraId"] = o.AvalaraId
	}
	toSerialize["jurisdictionName"] = o.JurisdictionName
	if !IsNil(o.JurisdictionType) {
		toSerialize["jurisdictionType"] = o.JurisdictionType
	}
	if !IsNil(o.SignatureCode) {
		toSerialize["signatureCode"] = o.SignatureCode
	}
	return toSerialize, nil
}

type NullableTaxAuthorityInfo struct {
	value *TaxAuthorityInfo
	isSet bool
}

func (v NullableTaxAuthorityInfo) Get() *TaxAuthorityInfo {
	return v.value
}

func (v *NullableTaxAuthorityInfo) Set(val *TaxAuthorityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxAuthorityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxAuthorityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxAuthorityInfo(val *TaxAuthorityInfo) *NullableTaxAuthorityInfo {
	return &NullableTaxAuthorityInfo{value: val, isSet: true}
}

func (v NullableTaxAuthorityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxAuthorityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

