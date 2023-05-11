/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Bhive struct {

	// The ID of a beehive. Is unique for that user.
	BhiveId string `json:"bhiveId,omitempty"`

	// Scale offset number (calibration data)
	ScaleOffset float32 `json:"scaleOffset,omitempty"`

	// Scale Reference Unit (calibration data)
	ScaleReferenceUnit float32 `json:"scaleReferenceUnit,omitempty"`
}

// AssertBhiveRequired checks if the required fields are not zero-ed
func AssertBhiveRequired(obj Bhive) error {
	return nil
}
