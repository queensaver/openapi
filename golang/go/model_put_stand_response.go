/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PutStandResponse struct {

	// HTTP response code. Used for internal purposes, will be sent out at the API.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`

	Stand Stand `json:"stand,omitempty"`
}

// AssertPutStandResponseRequired checks if the required fields are not zero-ed
func AssertPutStandResponseRequired(obj PutStandResponse) error {
	if err := AssertStandRequired(obj.Stand); err != nil {
		return err
	}
	return nil
}

// AssertRecursePutStandResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PutStandResponse (e.g. [][]PutStandResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePutStandResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPutStandResponse, ok := obj.(PutStandResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPutStandResponseRequired(aPutStandResponse)
	})
}