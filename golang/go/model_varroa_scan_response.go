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



type VarroaScanResponse struct {

	// HTTP response code. Used for internal purposes, will be sent out at the API.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`

	VarroaScans []VarroaScan `json:"varroaScans,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *VarroaScanResponse) UnmarshalJSON(data []byte) error {

	type Alias VarroaScanResponse // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertVarroaScanResponseRequired checks if the required fields are not zero-ed
func AssertVarroaScanResponseRequired(obj VarroaScanResponse) error {
	for _, el := range obj.VarroaScans {
		if err := AssertVarroaScanRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertVarroaScanResponseConstraints checks if the values respects the defined constraints
func AssertVarroaScanResponseConstraints(obj VarroaScanResponse) error {
	return nil
}
