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



type Telemetry struct {

	// The mac address of the scale
	M string `json:"m,omitempty"`

	// actual measurement of weight
	W float32 `json:"w,omitempty"`

	// timestamp of the measurement on one second precision (unix timestamp)
	E int64 `json:"e,omitempty"`

	// temperature
	T float32 `json:"t,omitempty"`

	// air pressure according to sensor
	A float32 `json:"a,omitempty"`

	// Battery voltage
	B float32 `json:"b,omitempty"`

	// signal Strength in %
	S int32 `json:"s,omitempty"`

	// longitude of the system
	Lo float32 `json:"lo,omitempty"`

	// latitude of the system
	La float32 `json:"la,omitempty"`

	// Humidity in percent
	H float32 `json:"h,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Telemetry) UnmarshalJSON(data []byte) error {

	type Alias Telemetry // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertTelemetryRequired checks if the required fields are not zero-ed
func AssertTelemetryRequired(obj Telemetry) error {
	return nil
}

// AssertTelemetryConstraints checks if the values respects the defined constraints
func AssertTelemetryConstraints(obj Telemetry) error {
	return nil
}
