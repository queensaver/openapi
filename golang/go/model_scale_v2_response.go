/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ScaleV2Response struct {

	// HTTP response code. Used for internal purposes, will be let out at the API level.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`

	// The measurement responses
	Values []ScaleV2 `json:"values,omitempty"`
}

// AssertScaleV2ResponseRequired checks if the required fields are not zero-ed
func AssertScaleV2ResponseRequired(obj ScaleV2Response) error {
	for _, el := range obj.Values {
		if err := AssertScaleV2Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseScaleV2ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ScaleV2Response (e.g. [][]ScaleV2Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseScaleV2ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aScaleV2Response, ok := obj.(ScaleV2Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertScaleV2ResponseRequired(aScaleV2Response)
	})
}
