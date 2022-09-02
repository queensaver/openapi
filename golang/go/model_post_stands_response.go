/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostStandsResponse struct {

	// HTTP response code. Used for internal purposes, will be sent out at the API.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`

	Stand Stand `json:"stand,omitempty"`
}

// AssertPostStandsResponseRequired checks if the required fields are not zero-ed
func AssertPostStandsResponseRequired(obj PostStandsResponse) error {
	if err := AssertStandRequired(obj.Stand); err != nil {
		return err
	}
	return nil
}

// AssertRecursePostStandsResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PostStandsResponse (e.g. [][]PostStandsResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePostStandsResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPostStandsResponse, ok := obj.(PostStandsResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPostStandsResponseRequired(aPostStandsResponse)
	})
}