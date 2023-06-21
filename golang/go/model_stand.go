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



type Stand struct {

	// A stand can have many hives. However, when sending POST or PUT requests you can only update the stand metadata, any hives attached to this array will be ignored. You will have to do the /hives API call instead to create / update hives with the correct stand UUID. This array will only be populated with GET requests.
	Hives []Hive `json:"hives,omitempty"`

	// Name of the stand. Can be chosen by the user. A stand is a collection of bee hives.
	Name string `json:"name"`

	// latitude of the stand
	Latitude float64 `json:"latitude,omitempty"`

	// longitude of the stand
	Longitude float64 `json:"longitude,omitempty"`

	// HTTP response code. Used for internal purposes, will be let out at the API level.
	HttpReponseCode int32 `json:"httpReponseCode,omitempty"`

	// Epoch when the data was last updated. This will be set internally, no need to add this with PUT or POST calls.
	Epoch int64 `json:"epoch,omitempty"`

	// Unique Identifier for this stand
	Uuid string `json:"uuid,omitempty"`

	// if set to true, the hive has been deleted at this epoch.
	Deleted bool `json:"deleted,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Stand) UnmarshalJSON(data []byte) error {

	type Alias Stand // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertStandRequired checks if the required fields are not zero-ed
func AssertStandRequired(obj Stand) error {
	elements := map[string]interface{}{
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Hives {
		if err := AssertHiveRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertStandConstraints checks if the values respects the defined constraints
func AssertStandConstraints(obj Stand) error {
	return nil
}
