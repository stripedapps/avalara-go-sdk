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

// checks if the BatchFileModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchFileModel{}

// BatchFileModel Represents one file in a batch upload
type BatchFileModel struct {
	// Unique BatchFile identifier
	Id *int32 `json:"id,omitempty"`
	// BatchId that this BatchFile is associated with.
	BatchId *int32 `json:"batchId,omitempty"`
	// Logical Name of file (e.g. \"Input\" or \"Error\")
	Name string `json:"name"`
	// BatchFile content
	Content string `json:"content"`
	// Size of content in bytes
	ContentLength *int32 `json:"contentLength,omitempty"`
	// Content mime type (e.g. text/csv).  This is used for HTTP downloading.
	ContentType string `json:"contentType"`
	// File extension (e.g. CSV)
	FileExtension string `json:"fileExtension"`
	// Number of errors that occurred when processing this file.
	ErrorCount *int32 `json:"errorCount,omitempty"`
}

// NewBatchFileModel instantiates a new BatchFileModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchFileModel(name string, content string, contentType string, fileExtension string) *BatchFileModel {
	this := BatchFileModel{}
	this.Name = name
	this.Content = content
	this.ContentType = contentType
	this.FileExtension = fileExtension
	return &this
}

// NewBatchFileModelWithDefaults instantiates a new BatchFileModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchFileModelWithDefaults() *BatchFileModel {
	this := BatchFileModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BatchFileModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BatchFileModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BatchFileModel) SetId(v int32) {
	o.Id = &v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *BatchFileModel) GetBatchId() int32 {
	if o == nil || IsNil(o.BatchId) {
		var ret int32
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetBatchIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *BatchFileModel) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given int32 and assigns it to the BatchId field.
func (o *BatchFileModel) SetBatchId(v int32) {
	o.BatchId = &v
}

// GetName returns the Name field value
func (o *BatchFileModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BatchFileModel) SetName(v string) {
	o.Name = v
}

// GetContent returns the Content field value
func (o *BatchFileModel) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *BatchFileModel) SetContent(v string) {
	o.Content = v
}

// GetContentLength returns the ContentLength field value if set, zero value otherwise.
func (o *BatchFileModel) GetContentLength() int32 {
	if o == nil || IsNil(o.ContentLength) {
		var ret int32
		return ret
	}
	return *o.ContentLength
}

// GetContentLengthOk returns a tuple with the ContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetContentLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.ContentLength) {
		return nil, false
	}
	return o.ContentLength, true
}

// HasContentLength returns a boolean if a field has been set.
func (o *BatchFileModel) HasContentLength() bool {
	if o != nil && !IsNil(o.ContentLength) {
		return true
	}

	return false
}

// SetContentLength gets a reference to the given int32 and assigns it to the ContentLength field.
func (o *BatchFileModel) SetContentLength(v int32) {
	o.ContentLength = &v
}

// GetContentType returns the ContentType field value
func (o *BatchFileModel) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *BatchFileModel) SetContentType(v string) {
	o.ContentType = v
}

// GetFileExtension returns the FileExtension field value
func (o *BatchFileModel) GetFileExtension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetFileExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileExtension, true
}

// SetFileExtension sets field value
func (o *BatchFileModel) SetFileExtension(v string) {
	o.FileExtension = v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *BatchFileModel) GetErrorCount() int32 {
	if o == nil || IsNil(o.ErrorCount) {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchFileModel) GetErrorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *BatchFileModel) HasErrorCount() bool {
	if o != nil && !IsNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *BatchFileModel) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

func (o BatchFileModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchFileModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	toSerialize["name"] = o.Name
	toSerialize["content"] = o.Content
	if !IsNil(o.ContentLength) {
		toSerialize["contentLength"] = o.ContentLength
	}
	toSerialize["contentType"] = o.ContentType
	toSerialize["fileExtension"] = o.FileExtension
	if !IsNil(o.ErrorCount) {
		toSerialize["errorCount"] = o.ErrorCount
	}
	return toSerialize, nil
}

type NullableBatchFileModel struct {
	value *BatchFileModel
	isSet bool
}

func (v NullableBatchFileModel) Get() *BatchFileModel {
	return v.value
}

func (v *NullableBatchFileModel) Set(val *BatchFileModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchFileModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchFileModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchFileModel(val *BatchFileModel) *NullableBatchFileModel {
	return &NullableBatchFileModel{value: val, isSet: true}
}

func (v NullableBatchFileModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchFileModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


