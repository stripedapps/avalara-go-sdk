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

// checks if the AdjustTransactionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdjustTransactionModel{}

// AdjustTransactionModel A request to adjust tax for a previously existing transaction
type AdjustTransactionModel struct {
	// A reason code indicating why this adjustment was made
	AdjustmentReason *string `json:"adjustmentReason,omitempty"`
	// If the AdjustmentReason is \"Other\", specify the reason here
	AdjustmentDescription *string `json:"adjustmentDescription,omitempty"`
	NewTransaction *CreateTransactionModel `json:"newTransaction,omitempty"`
}

// NewAdjustTransactionModel instantiates a new AdjustTransactionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustTransactionModel() *AdjustTransactionModel {
	this := AdjustTransactionModel{}
	return &this
}

// NewAdjustTransactionModelWithDefaults instantiates a new AdjustTransactionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustTransactionModelWithDefaults() *AdjustTransactionModel {
	this := AdjustTransactionModel{}
	return &this
}

// GetAdjustmentReason returns the AdjustmentReason field value if set, zero value otherwise.
func (o *AdjustTransactionModel) GetAdjustmentReason() string {
	if o == nil || IsNil(o.AdjustmentReason) {
		var ret string
		return ret
	}
	return *o.AdjustmentReason
}

// GetAdjustmentReasonOk returns a tuple with the AdjustmentReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustTransactionModel) GetAdjustmentReasonOk() (*string, bool) {
	if o == nil || IsNil(o.AdjustmentReason) {
		return nil, false
	}
	return o.AdjustmentReason, true
}

// HasAdjustmentReason returns a boolean if a field has been set.
func (o *AdjustTransactionModel) HasAdjustmentReason() bool {
	if o != nil && !IsNil(o.AdjustmentReason) {
		return true
	}

	return false
}

// SetAdjustmentReason gets a reference to the given string and assigns it to the AdjustmentReason field.
func (o *AdjustTransactionModel) SetAdjustmentReason(v string) {
	o.AdjustmentReason = &v
}

// GetAdjustmentDescription returns the AdjustmentDescription field value if set, zero value otherwise.
func (o *AdjustTransactionModel) GetAdjustmentDescription() string {
	if o == nil || IsNil(o.AdjustmentDescription) {
		var ret string
		return ret
	}
	return *o.AdjustmentDescription
}

// GetAdjustmentDescriptionOk returns a tuple with the AdjustmentDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustTransactionModel) GetAdjustmentDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AdjustmentDescription) {
		return nil, false
	}
	return o.AdjustmentDescription, true
}

// HasAdjustmentDescription returns a boolean if a field has been set.
func (o *AdjustTransactionModel) HasAdjustmentDescription() bool {
	if o != nil && !IsNil(o.AdjustmentDescription) {
		return true
	}

	return false
}

// SetAdjustmentDescription gets a reference to the given string and assigns it to the AdjustmentDescription field.
func (o *AdjustTransactionModel) SetAdjustmentDescription(v string) {
	o.AdjustmentDescription = &v
}

// GetNewTransaction returns the NewTransaction field value if set, zero value otherwise.
func (o *AdjustTransactionModel) GetNewTransaction() CreateTransactionModel {
	if o == nil || IsNil(o.NewTransaction) {
		var ret CreateTransactionModel
		return ret
	}
	return *o.NewTransaction
}

// GetNewTransactionOk returns a tuple with the NewTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustTransactionModel) GetNewTransactionOk() (*CreateTransactionModel, bool) {
	if o == nil || IsNil(o.NewTransaction) {
		return nil, false
	}
	return o.NewTransaction, true
}

// HasNewTransaction returns a boolean if a field has been set.
func (o *AdjustTransactionModel) HasNewTransaction() bool {
	if o != nil && !IsNil(o.NewTransaction) {
		return true
	}

	return false
}

// SetNewTransaction gets a reference to the given CreateTransactionModel and assigns it to the NewTransaction field.
func (o *AdjustTransactionModel) SetNewTransaction(v CreateTransactionModel) {
	o.NewTransaction = &v
}

func (o AdjustTransactionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdjustTransactionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdjustmentReason) {
		toSerialize["adjustmentReason"] = o.AdjustmentReason
	}
	if !IsNil(o.AdjustmentDescription) {
		toSerialize["adjustmentDescription"] = o.AdjustmentDescription
	}
	if !IsNil(o.NewTransaction) {
		toSerialize["newTransaction"] = o.NewTransaction
	}
	return toSerialize, nil
}

type NullableAdjustTransactionModel struct {
	value *AdjustTransactionModel
	isSet bool
}

func (v NullableAdjustTransactionModel) Get() *AdjustTransactionModel {
	return v.value
}

func (v *NullableAdjustTransactionModel) Set(val *AdjustTransactionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustTransactionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustTransactionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustTransactionModel(val *AdjustTransactionModel) *NullableAdjustTransactionModel {
	return &NullableAdjustTransactionModel{value: val, isSet: true}
}

func (v NullableAdjustTransactionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustTransactionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


