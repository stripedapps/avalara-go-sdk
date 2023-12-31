/*
AvaTax API

REST interface to Avalara's enterprise tax service, AvaTax.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the TaxCodeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxCodeModel{}

// TaxCodeModel Tax Code Model
type TaxCodeModel struct {
	// Id
	Id *int32 `json:"id,omitempty"`
	// Tax Code
	TaxCode string `json:"taxCode"`
	// Tax Code Type Id
	TaxCodeTypeId string `json:"taxCodeTypeId"`
	// Company Id
	CompanyId *int32 `json:"companyId,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Parent Tax Code
	ParentTaxCode *string `json:"parentTaxCode,omitempty"`
	// Is Physical
	IsPhysical *bool `json:"isPhysical,omitempty"`
	// Good Service Code
	GoodsServiceCode *int64 `json:"goodsServiceCode,omitempty"`
	// Entity Use Code
	EntityUseCode *string `json:"entityUseCode,omitempty"`
	// Is Active
	IsActive *bool `json:"isActive,omitempty"`
	// Is SST Certified
	IsSSTCertified *bool `json:"isSSTCertified,omitempty"`
	// The date when this record was created (read only)
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// The User ID of the user who created this record (read only)
	CreatedUserId *int32 `json:"createdUserId,omitempty"`
	// The date/time when this record was last modified (read only)
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// The user ID of the user who last modified this record (read only)
	ModifiedUserId *int32 `json:"modifiedUserId,omitempty"`
}

// NewTaxCodeModel instantiates a new TaxCodeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCodeModel(taxCode string, taxCodeTypeId string) *TaxCodeModel {
	this := TaxCodeModel{}
	this.TaxCode = taxCode
	this.TaxCodeTypeId = taxCodeTypeId
	return &this
}

// NewTaxCodeModelWithDefaults instantiates a new TaxCodeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCodeModelWithDefaults() *TaxCodeModel {
	this := TaxCodeModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaxCodeModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaxCodeModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TaxCodeModel) SetId(v int32) {
	o.Id = &v
}

// GetTaxCode returns the TaxCode field value
func (o *TaxCodeModel) GetTaxCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetTaxCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxCode, true
}

// SetTaxCode sets field value
func (o *TaxCodeModel) SetTaxCode(v string) {
	o.TaxCode = v
}

// GetTaxCodeTypeId returns the TaxCodeTypeId field value
func (o *TaxCodeModel) GetTaxCodeTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxCodeTypeId
}

// GetTaxCodeTypeIdOk returns a tuple with the TaxCodeTypeId field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetTaxCodeTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxCodeTypeId, true
}

// SetTaxCodeTypeId sets field value
func (o *TaxCodeModel) SetTaxCodeTypeId(v string) {
	o.TaxCodeTypeId = v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *TaxCodeModel) GetCompanyId() int32 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetCompanyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *TaxCodeModel) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *TaxCodeModel) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TaxCodeModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TaxCodeModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TaxCodeModel) SetDescription(v string) {
	o.Description = &v
}

// GetParentTaxCode returns the ParentTaxCode field value if set, zero value otherwise.
func (o *TaxCodeModel) GetParentTaxCode() string {
	if o == nil || IsNil(o.ParentTaxCode) {
		var ret string
		return ret
	}
	return *o.ParentTaxCode
}

// GetParentTaxCodeOk returns a tuple with the ParentTaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetParentTaxCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ParentTaxCode) {
		return nil, false
	}
	return o.ParentTaxCode, true
}

// HasParentTaxCode returns a boolean if a field has been set.
func (o *TaxCodeModel) HasParentTaxCode() bool {
	if o != nil && !IsNil(o.ParentTaxCode) {
		return true
	}

	return false
}

// SetParentTaxCode gets a reference to the given string and assigns it to the ParentTaxCode field.
func (o *TaxCodeModel) SetParentTaxCode(v string) {
	o.ParentTaxCode = &v
}

// GetIsPhysical returns the IsPhysical field value if set, zero value otherwise.
func (o *TaxCodeModel) GetIsPhysical() bool {
	if o == nil || IsNil(o.IsPhysical) {
		var ret bool
		return ret
	}
	return *o.IsPhysical
}

// GetIsPhysicalOk returns a tuple with the IsPhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetIsPhysicalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPhysical) {
		return nil, false
	}
	return o.IsPhysical, true
}

// HasIsPhysical returns a boolean if a field has been set.
func (o *TaxCodeModel) HasIsPhysical() bool {
	if o != nil && !IsNil(o.IsPhysical) {
		return true
	}

	return false
}

// SetIsPhysical gets a reference to the given bool and assigns it to the IsPhysical field.
func (o *TaxCodeModel) SetIsPhysical(v bool) {
	o.IsPhysical = &v
}

// GetGoodsServiceCode returns the GoodsServiceCode field value if set, zero value otherwise.
func (o *TaxCodeModel) GetGoodsServiceCode() int64 {
	if o == nil || IsNil(o.GoodsServiceCode) {
		var ret int64
		return ret
	}
	return *o.GoodsServiceCode
}

// GetGoodsServiceCodeOk returns a tuple with the GoodsServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetGoodsServiceCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.GoodsServiceCode) {
		return nil, false
	}
	return o.GoodsServiceCode, true
}

// HasGoodsServiceCode returns a boolean if a field has been set.
func (o *TaxCodeModel) HasGoodsServiceCode() bool {
	if o != nil && !IsNil(o.GoodsServiceCode) {
		return true
	}

	return false
}

// SetGoodsServiceCode gets a reference to the given int64 and assigns it to the GoodsServiceCode field.
func (o *TaxCodeModel) SetGoodsServiceCode(v int64) {
	o.GoodsServiceCode = &v
}

// GetEntityUseCode returns the EntityUseCode field value if set, zero value otherwise.
func (o *TaxCodeModel) GetEntityUseCode() string {
	if o == nil || IsNil(o.EntityUseCode) {
		var ret string
		return ret
	}
	return *o.EntityUseCode
}

// GetEntityUseCodeOk returns a tuple with the EntityUseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetEntityUseCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityUseCode) {
		return nil, false
	}
	return o.EntityUseCode, true
}

// HasEntityUseCode returns a boolean if a field has been set.
func (o *TaxCodeModel) HasEntityUseCode() bool {
	if o != nil && !IsNil(o.EntityUseCode) {
		return true
	}

	return false
}

// SetEntityUseCode gets a reference to the given string and assigns it to the EntityUseCode field.
func (o *TaxCodeModel) SetEntityUseCode(v string) {
	o.EntityUseCode = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *TaxCodeModel) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *TaxCodeModel) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *TaxCodeModel) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsSSTCertified returns the IsSSTCertified field value if set, zero value otherwise.
func (o *TaxCodeModel) GetIsSSTCertified() bool {
	if o == nil || IsNil(o.IsSSTCertified) {
		var ret bool
		return ret
	}
	return *o.IsSSTCertified
}

// GetIsSSTCertifiedOk returns a tuple with the IsSSTCertified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetIsSSTCertifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSSTCertified) {
		return nil, false
	}
	return o.IsSSTCertified, true
}

// HasIsSSTCertified returns a boolean if a field has been set.
func (o *TaxCodeModel) HasIsSSTCertified() bool {
	if o != nil && !IsNil(o.IsSSTCertified) {
		return true
	}

	return false
}

// SetIsSSTCertified gets a reference to the given bool and assigns it to the IsSSTCertified field.
func (o *TaxCodeModel) SetIsSSTCertified(v bool) {
	o.IsSSTCertified = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *TaxCodeModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *TaxCodeModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *TaxCodeModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetCreatedUserId returns the CreatedUserId field value if set, zero value otherwise.
func (o *TaxCodeModel) GetCreatedUserId() int32 {
	if o == nil || IsNil(o.CreatedUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedUserId
}

// GetCreatedUserIdOk returns a tuple with the CreatedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetCreatedUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedUserId) {
		return nil, false
	}
	return o.CreatedUserId, true
}

// HasCreatedUserId returns a boolean if a field has been set.
func (o *TaxCodeModel) HasCreatedUserId() bool {
	if o != nil && !IsNil(o.CreatedUserId) {
		return true
	}

	return false
}

// SetCreatedUserId gets a reference to the given int32 and assigns it to the CreatedUserId field.
func (o *TaxCodeModel) SetCreatedUserId(v int32) {
	o.CreatedUserId = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *TaxCodeModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TaxCodeModel) HasModifiedDate() bool {
	if o != nil && !IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *TaxCodeModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetModifiedUserId returns the ModifiedUserId field value if set, zero value otherwise.
func (o *TaxCodeModel) GetModifiedUserId() int32 {
	if o == nil || IsNil(o.ModifiedUserId) {
		var ret int32
		return ret
	}
	return *o.ModifiedUserId
}

// GetModifiedUserIdOk returns a tuple with the ModifiedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetModifiedUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ModifiedUserId) {
		return nil, false
	}
	return o.ModifiedUserId, true
}

// HasModifiedUserId returns a boolean if a field has been set.
func (o *TaxCodeModel) HasModifiedUserId() bool {
	if o != nil && !IsNil(o.ModifiedUserId) {
		return true
	}

	return false
}

// SetModifiedUserId gets a reference to the given int32 and assigns it to the ModifiedUserId field.
func (o *TaxCodeModel) SetModifiedUserId(v int32) {
	o.ModifiedUserId = &v
}

func (o TaxCodeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxCodeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["taxCode"] = o.TaxCode
	toSerialize["taxCodeTypeId"] = o.TaxCodeTypeId
	if !IsNil(o.CompanyId) {
		toSerialize["companyId"] = o.CompanyId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ParentTaxCode) {
		toSerialize["parentTaxCode"] = o.ParentTaxCode
	}
	if !IsNil(o.IsPhysical) {
		toSerialize["isPhysical"] = o.IsPhysical
	}
	if !IsNil(o.GoodsServiceCode) {
		toSerialize["goodsServiceCode"] = o.GoodsServiceCode
	}
	if !IsNil(o.EntityUseCode) {
		toSerialize["entityUseCode"] = o.EntityUseCode
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.IsSSTCertified) {
		toSerialize["isSSTCertified"] = o.IsSSTCertified
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.CreatedUserId) {
		toSerialize["createdUserId"] = o.CreatedUserId
	}
	if !IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !IsNil(o.ModifiedUserId) {
		toSerialize["modifiedUserId"] = o.ModifiedUserId
	}
	return toSerialize, nil
}

type NullableTaxCodeModel struct {
	value *TaxCodeModel
	isSet bool
}

func (v NullableTaxCodeModel) Get() *TaxCodeModel {
	return v.value
}

func (v *NullableTaxCodeModel) Set(val *TaxCodeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCodeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCodeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCodeModel(val *TaxCodeModel) *NullableTaxCodeModel {
	return &NullableTaxCodeModel{value: val, isSet: true}
}

func (v NullableTaxCodeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCodeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


