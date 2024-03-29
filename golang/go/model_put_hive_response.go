/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"encoding/json"
)



type PutHiveResponse struct {

	// HTTP response code. Used for internal purposes, will be sent out at the API.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`

	Hive Hive `json:"hive,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *PutHiveResponse) UnmarshalJSON(data []byte) error {

	type Alias PutHiveResponse // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertPutHiveResponseRequired checks if the required fields are not zero-ed
func AssertPutHiveResponseRequired(obj PutHiveResponse) error {
	if err := AssertHiveRequired(obj.Hive); err != nil {
		return err
	}
	return nil
}

// AssertPutHiveResponseConstraints checks if the values respects the defined constraints
func AssertPutHiveResponseConstraints(obj PutHiveResponse) error {
	return nil
}
