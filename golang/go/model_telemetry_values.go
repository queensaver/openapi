/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TelemetryValues struct {

	// HTTP response code. Used for internal purposes, will be let out at the API level.
	HttpResponseCode int32 `json:"httpResponseCode,omitempty"`

	// The measurement responses
	Values []Telemetry `json:"values,omitempty"`
}

// AssertTelemetryValuesRequired checks if the required fields are not zero-ed
func AssertTelemetryValuesRequired(obj TelemetryValues) error {
	for _, el := range obj.Values {
		if err := AssertTelemetryRequired(el); err != nil {
			return err
		}
	}
	return nil
}