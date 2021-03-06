/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GenericPostResponse struct {

	// HTTP response code.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`
}

// AssertGenericPostResponseRequired checks if the required fields are not zero-ed
func AssertGenericPostResponseRequired(obj GenericPostResponse) error {
	return nil
}

// AssertRecurseGenericPostResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GenericPostResponse (e.g. [][]GenericPostResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGenericPostResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGenericPostResponse, ok := obj.(GenericPostResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGenericPostResponseRequired(aGenericPostResponse)
	})
}
