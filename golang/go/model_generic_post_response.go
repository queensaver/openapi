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



type GenericPostResponse struct {

	// HTTP response code.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *GenericPostResponse) UnmarshalJSON(data []byte) error {

	type Alias GenericPostResponse // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertGenericPostResponseRequired checks if the required fields are not zero-ed
func AssertGenericPostResponseRequired(obj GenericPostResponse) error {
	return nil
}

// AssertGenericPostResponseConstraints checks if the values respects the defined constraints
func AssertGenericPostResponseConstraints(obj GenericPostResponse) error {
	return nil
}
