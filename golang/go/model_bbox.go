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



type Bbox struct {

	// The ID of the bbox electronix (QBox). Is usually a mac address of a network interface.
	BboxId string `json:"bboxId,omitempty"`

	// A cron type of description of when the sensore measurements are supposed to be done.
	Schedule string `json:"schedule,omitempty"`

	// Unique Identifier for this bbox
	Uuid string `json:"uuid,omitempty"`

	Bhives []Bhive `json:"bhives,omitempty"`

	// If the bbox turns the power off after a successful run and wakes up later according to the schedule.
	PowerSave bool `json:"powerSave,omitempty"`

	// The registration ID of the bbox. The user needs to put this into the interface so that the bbox can then register via the /configs/bbox/register API call to retrieve the token.
	RegistrationId string `json:"registrationId,omitempty"`

	// Hardware type of the bbox - could be a varroa-scanner or a scale, etc.
	HardwareType string `json:"hardwareType,omitempty"`

	// Hardware revision - newer revisions might have different features which are important to know.
	HardwareRevision int64 `json:"hardwareRevision,omitempty"`

	// Name of the bbox - optional
	Name string `json:"name,omitempty"`

	// Associated Hive with this bbox
	HiveUuid string `json:"hiveUuid,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Bbox) UnmarshalJSON(data []byte) error {

	type Alias Bbox // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertBboxRequired checks if the required fields are not zero-ed
func AssertBboxRequired(obj Bbox) error {
	for _, el := range obj.Bhives {
		if err := AssertBhiveRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertBboxConstraints checks if the values respects the defined constraints
func AssertBboxConstraints(obj Bbox) error {
	return nil
}
