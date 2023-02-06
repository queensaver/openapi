/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TelemetryInner struct {

	// The mac address of the scale
	MacAddress string `json:"macAddress,omitempty"`

	// actual measurement of weight
	Weight float32 `json:"weight,omitempty"`

	// timestamp of the measurement on one second precision
	Epoch int64 `json:"epoch,omitempty"`

	// temperature
	Temperature float32 `json:"temperature,omitempty"`

	// air pressure according to sensor
	AirPressure float32 `json:"airPressure,omitempty"`

	// Battery voltage
	BatteryVoltage float32 `json:"batteryVoltage,omitempty"`

	// signal Strength in %
	SignalStrength float32 `json:"signalStrength,omitempty"`
}

// AssertTelemetryInnerRequired checks if the required fields are not zero-ed
func AssertTelemetryInnerRequired(obj TelemetryInner) error {
	return nil
}

// AssertRecurseTelemetryInnerRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TelemetryInner (e.g. [][]TelemetryInner), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTelemetryInnerRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTelemetryInner, ok := obj.(TelemetryInner)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTelemetryInnerRequired(aTelemetryInner)
	})
}
