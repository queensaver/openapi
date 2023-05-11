/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type User struct {

	// Username, must be unique to the queensaver system. We encourage using email-addresses here.
	Username string `json:"username"`

	// Password
	Password string `json:"password"`

	// First name of the user
	FirstName string `json:"firstName"`

	// Last name of the user
	LastName string `json:"lastName"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"username": obj.Username,
		"password": obj.Password,
		"firstName": obj.FirstName,
		"lastName": obj.LastName,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}
