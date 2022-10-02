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

	// ID of the object. The ID is usually determined by the DBMS.
	BboxId int32 `json:"bboxId,omitempty"`

	// A cron type of description of when the sensore measurements are supposed to be done.
	Schedule string `json:"schedule,omitempty"`

	Bhive []Bhive `json:"bhive,omitempty"`
}

// AssertBboxRequired checks if the required fields are not zero-ed
func AssertBboxRequired(obj Bbox) error {
	for _, el := range obj.Bhive {
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
