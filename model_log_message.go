/*
СМС шлюз от Eskiz.uz

Отправляйте СМС по всему миру, в любом количестве!  В тестовом статусе для отправки тестовых смс сообщений, Вы можете использовать только нижеуказанные тексты:  - Это тест от Eskiz      - Bu Eskiz dan test      - This is test from Eskiz

API version: 1.0.0
Contact: danil@iota.uz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eskizapi

import (
	"encoding/json"
	"time"
)

// checks if the LogMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogMessage{}

// LogMessage struct for LogMessage
type LogMessage struct {
	Ip          *string    `form:"ip" json:"ip,omitempty"`
	From        *string    `form:"from" json:"from,omitempty"`
	To          *string    `form:"to" json:"to,omitempty"`
	DispatchId  *string    `form:"dispatch_id" json:"dispatch_id,omitempty"`
	UserSmsId   *string    `form:"user_sms_id" json:"user_sms_id,omitempty"`
	Timestamp   *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	Message     *string    `form:"message" json:"message,omitempty"`
	FullMessage *string    `form:"full_message" json:"full_message,omitempty"`
	LevelName   *string    `form:"level_name" json:"level_name,omitempty"`
}

// NewLogMessage instantiates a new LogMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogMessage() *LogMessage {
	this := LogMessage{}
	return &this
}

// NewLogMessageWithDefaults instantiates a new LogMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogMessageWithDefaults() *LogMessage {
	this := LogMessage{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *LogMessage) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *LogMessage) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *LogMessage) SetIp(v string) {
	o.Ip = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *LogMessage) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *LogMessage) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *LogMessage) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *LogMessage) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *LogMessage) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *LogMessage) SetTo(v string) {
	o.To = &v
}

// GetDispatchId returns the DispatchId field value if set, zero value otherwise.
func (o *LogMessage) GetDispatchId() string {
	if o == nil || IsNil(o.DispatchId) {
		var ret string
		return ret
	}
	return *o.DispatchId
}

// GetDispatchIdOk returns a tuple with the DispatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetDispatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.DispatchId) {
		return nil, false
	}
	return o.DispatchId, true
}

// HasDispatchId returns a boolean if a field has been set.
func (o *LogMessage) HasDispatchId() bool {
	if o != nil && !IsNil(o.DispatchId) {
		return true
	}

	return false
}

// SetDispatchId gets a reference to the given string and assigns it to the DispatchId field.
func (o *LogMessage) SetDispatchId(v string) {
	o.DispatchId = &v
}

// GetUserSmsId returns the UserSmsId field value if set, zero value otherwise.
func (o *LogMessage) GetUserSmsId() string {
	if o == nil || IsNil(o.UserSmsId) {
		var ret string
		return ret
	}
	return *o.UserSmsId
}

// GetUserSmsIdOk returns a tuple with the UserSmsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetUserSmsIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserSmsId) {
		return nil, false
	}
	return o.UserSmsId, true
}

// HasUserSmsId returns a boolean if a field has been set.
func (o *LogMessage) HasUserSmsId() bool {
	if o != nil && !IsNil(o.UserSmsId) {
		return true
	}

	return false
}

// SetUserSmsId gets a reference to the given string and assigns it to the UserSmsId field.
func (o *LogMessage) SetUserSmsId(v string) {
	o.UserSmsId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogMessage) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogMessage) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *LogMessage) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LogMessage) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LogMessage) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LogMessage) SetMessage(v string) {
	o.Message = &v
}

// GetFullMessage returns the FullMessage field value if set, zero value otherwise.
func (o *LogMessage) GetFullMessage() string {
	if o == nil || IsNil(o.FullMessage) {
		var ret string
		return ret
	}
	return *o.FullMessage
}

// GetFullMessageOk returns a tuple with the FullMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetFullMessageOk() (*string, bool) {
	if o == nil || IsNil(o.FullMessage) {
		return nil, false
	}
	return o.FullMessage, true
}

// HasFullMessage returns a boolean if a field has been set.
func (o *LogMessage) HasFullMessage() bool {
	if o != nil && !IsNil(o.FullMessage) {
		return true
	}

	return false
}

// SetFullMessage gets a reference to the given string and assigns it to the FullMessage field.
func (o *LogMessage) SetFullMessage(v string) {
	o.FullMessage = &v
}

// GetLevelName returns the LevelName field value if set, zero value otherwise.
func (o *LogMessage) GetLevelName() string {
	if o == nil || IsNil(o.LevelName) {
		var ret string
		return ret
	}
	return *o.LevelName
}

// GetLevelNameOk returns a tuple with the LevelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogMessage) GetLevelNameOk() (*string, bool) {
	if o == nil || IsNil(o.LevelName) {
		return nil, false
	}
	return o.LevelName, true
}

// HasLevelName returns a boolean if a field has been set.
func (o *LogMessage) HasLevelName() bool {
	if o != nil && !IsNil(o.LevelName) {
		return true
	}

	return false
}

// SetLevelName gets a reference to the given string and assigns it to the LevelName field.
func (o *LogMessage) SetLevelName(v string) {
	o.LevelName = &v
}

func (o LogMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.DispatchId) {
		toSerialize["dispatch_id"] = o.DispatchId
	}
	if !IsNil(o.UserSmsId) {
		toSerialize["user_sms_id"] = o.UserSmsId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.FullMessage) {
		toSerialize["full_message"] = o.FullMessage
	}
	if !IsNil(o.LevelName) {
		toSerialize["level_name"] = o.LevelName
	}
	return toSerialize, nil
}

type NullableLogMessage struct {
	value *LogMessage
	isSet bool
}

func (v NullableLogMessage) Get() *LogMessage {
	return v.value
}

func (v *NullableLogMessage) Set(val *LogMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableLogMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableLogMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogMessage(val *LogMessage) *NullableLogMessage {
	return &NullableLogMessage{value: val, isSet: true}
}

func (v NullableLogMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
