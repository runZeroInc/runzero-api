/*
 * runZero API
 *
 * runZero Network Discovery API
 *
 * API version: 1.0.4
 * Contact: support@runzero.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// RunZeroAPIKey struct for RunZeroAPIKey
type RunZeroAPIKey struct {
	Id             string  `json:"id"`
	ClientId       *string `json:"client_id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	CreatedAt      *int64  `json:"created_at,omitempty"`
	CreatedBy      *string `json:"created_by,omitempty"`
	Comment        *string `json:"comment,omitempty"`
	LastUsedAt     *int64  `json:"last_used_at,omitempty"`
	LastUsedIp     *string `json:"last_used_ip,omitempty"`
	LastUsedUa     *string `json:"last_used_ua,omitempty"`
	Counter        *int64  `json:"counter,omitempty"`
	UsageToday     *int64  `json:"usage_today,omitempty"`
	UsageLimit     *int64  `json:"usage_limit,omitempty"`
	Token          *string `json:"token,omitempty"`
	Inactive       *bool   `json:"inactive,omitempty"`
}

// NewRunZeroAPIKey instantiates a new RunZeroAPIKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunZeroAPIKey(id string) *RunZeroAPIKey {
	this := RunZeroAPIKey{}
	this.Id = id
	return &this
}

// NewRunZeroAPIKeyWithDefaults instantiates a new RunZeroAPIKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunZeroAPIKeyWithDefaults() *RunZeroAPIKey {
	this := RunZeroAPIKey{}
	return &this
}

// GetId returns the Id field value
func (o *RunZeroAPIKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunZeroAPIKey) SetId(v string) {
	o.Id = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *RunZeroAPIKey) SetClientId(v string) {
	o.ClientId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *RunZeroAPIKey) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *RunZeroAPIKey) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *RunZeroAPIKey) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RunZeroAPIKey) SetComment(v string) {
	o.Comment = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetLastUsedAt() int64 {
	if o == nil || o.LastUsedAt == nil {
		var ret int64
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetLastUsedAtOk() (*int64, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given int64 and assigns it to the LastUsedAt field.
func (o *RunZeroAPIKey) SetLastUsedAt(v int64) {
	o.LastUsedAt = &v
}

// GetLastUsedIp returns the LastUsedIp field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetLastUsedIp() string {
	if o == nil || o.LastUsedIp == nil {
		var ret string
		return ret
	}
	return *o.LastUsedIp
}

// GetLastUsedIpOk returns a tuple with the LastUsedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetLastUsedIpOk() (*string, bool) {
	if o == nil || o.LastUsedIp == nil {
		return nil, false
	}
	return o.LastUsedIp, true
}

// HasLastUsedIp returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasLastUsedIp() bool {
	if o != nil && o.LastUsedIp != nil {
		return true
	}

	return false
}

// SetLastUsedIp gets a reference to the given string and assigns it to the LastUsedIp field.
func (o *RunZeroAPIKey) SetLastUsedIp(v string) {
	o.LastUsedIp = &v
}

// GetLastUsedUa returns the LastUsedUa field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetLastUsedUa() string {
	if o == nil || o.LastUsedUa == nil {
		var ret string
		return ret
	}
	return *o.LastUsedUa
}

// GetLastUsedUaOk returns a tuple with the LastUsedUa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetLastUsedUaOk() (*string, bool) {
	if o == nil || o.LastUsedUa == nil {
		return nil, false
	}
	return o.LastUsedUa, true
}

// HasLastUsedUa returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasLastUsedUa() bool {
	if o != nil && o.LastUsedUa != nil {
		return true
	}

	return false
}

// SetLastUsedUa gets a reference to the given string and assigns it to the LastUsedUa field.
func (o *RunZeroAPIKey) SetLastUsedUa(v string) {
	o.LastUsedUa = &v
}

// GetCounter returns the Counter field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetCounter() int64 {
	if o == nil || o.Counter == nil {
		var ret int64
		return ret
	}
	return *o.Counter
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetCounterOk() (*int64, bool) {
	if o == nil || o.Counter == nil {
		return nil, false
	}
	return o.Counter, true
}

// HasCounter returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasCounter() bool {
	if o != nil && o.Counter != nil {
		return true
	}

	return false
}

// SetCounter gets a reference to the given int64 and assigns it to the Counter field.
func (o *RunZeroAPIKey) SetCounter(v int64) {
	o.Counter = &v
}

// GetUsageToday returns the UsageToday field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetUsageToday() int64 {
	if o == nil || o.UsageToday == nil {
		var ret int64
		return ret
	}
	return *o.UsageToday
}

// GetUsageTodayOk returns a tuple with the UsageToday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetUsageTodayOk() (*int64, bool) {
	if o == nil || o.UsageToday == nil {
		return nil, false
	}
	return o.UsageToday, true
}

// HasUsageToday returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasUsageToday() bool {
	if o != nil && o.UsageToday != nil {
		return true
	}

	return false
}

// SetUsageToday gets a reference to the given int64 and assigns it to the UsageToday field.
func (o *RunZeroAPIKey) SetUsageToday(v int64) {
	o.UsageToday = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetUsageLimit() int64 {
	if o == nil || o.UsageLimit == nil {
		var ret int64
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetUsageLimitOk() (*int64, bool) {
	if o == nil || o.UsageLimit == nil {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasUsageLimit() bool {
	if o != nil && o.UsageLimit != nil {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int64 and assigns it to the UsageLimit field.
func (o *RunZeroAPIKey) SetUsageLimit(v int64) {
	o.UsageLimit = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RunZeroAPIKey) SetToken(v string) {
	o.Token = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *RunZeroAPIKey) GetInactive() bool {
	if o == nil || o.Inactive == nil {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunZeroAPIKey) GetInactiveOk() (*bool, bool) {
	if o == nil || o.Inactive == nil {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *RunZeroAPIKey) HasInactive() bool {
	if o != nil && o.Inactive != nil {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *RunZeroAPIKey) SetInactive(v bool) {
	o.Inactive = &v
}

func (o RunZeroAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.LastUsedAt != nil {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	if o.LastUsedIp != nil {
		toSerialize["last_used_ip"] = o.LastUsedIp
	}
	if o.LastUsedUa != nil {
		toSerialize["last_used_ua"] = o.LastUsedUa
	}
	if o.Counter != nil {
		toSerialize["counter"] = o.Counter
	}
	if o.UsageToday != nil {
		toSerialize["usage_today"] = o.UsageToday
	}
	if o.UsageLimit != nil {
		toSerialize["usage_limit"] = o.UsageLimit
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Inactive != nil {
		toSerialize["inactive"] = o.Inactive
	}
	return json.Marshal(toSerialize)
}

type NullableRunZeroAPIKey struct {
	value *RunZeroAPIKey
	isSet bool
}

func (v NullableRunZeroAPIKey) Get() *RunZeroAPIKey {
	return v.value
}

func (v *NullableRunZeroAPIKey) Set(val *RunZeroAPIKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRunZeroAPIKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRunZeroAPIKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunZeroAPIKey(val *RunZeroAPIKey) *NullableRunZeroAPIKey {
	return &NullableRunZeroAPIKey{value: val, isSet: true}
}

func (v NullableRunZeroAPIKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunZeroAPIKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
