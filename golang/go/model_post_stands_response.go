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



type PostStandsResponse struct {

	// HTTP response code. Used for internal purposes, will be sent out at the API.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`

	Stand Stand `json:"stand,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *PostStandsResponse) UnmarshalJSON(data []byte) error {

	type Alias PostStandsResponse // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertPostStandsResponseRequired checks if the required fields are not zero-ed
func AssertPostStandsResponseRequired(obj PostStandsResponse) error {
	if err := AssertStandRequired(obj.Stand); err != nil {
		return err
	}
	return nil
}

// AssertPostStandsResponseConstraints checks if the values respects the defined constraints
func AssertPostStandsResponseConstraints(obj PostStandsResponse) error {
	return nil
}
