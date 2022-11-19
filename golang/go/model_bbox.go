/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

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

	// The registration ID of the bbox. The user needs to put this into the interface so that the bbox can then register via the /configs/register API call to retrieve the token.
	RegistrationId string `json:"registrationId,omitempty"`
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

// AssertRecurseBboxRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Bbox (e.g. [][]Bbox), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBboxRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBbox, ok := obj.(Bbox)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBboxRequired(aBbox)
	})
}
