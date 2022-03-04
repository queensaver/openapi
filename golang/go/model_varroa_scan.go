/*
 * Queensaver API
 *
 * Queensaver API to send in sensor data and retrieve it. It's also used for user management.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VarroaScan struct {

	// bhiveId to identify the data source. Might become empty in a future iteration as it's redundant with the query parameter.
	BhiveId string `json:"bhiveId,omitempty"`

	// URL to the image of the varroa scan
	ImageUrl string `json:"imageUrl,omitempty"`

	// timestamp of the measurement on one second precision
	Epoch int64 `json:"epoch,omitempty"`

	Metadata []VarroaScanMetadata `json:"metadata,omitempty"`

	// Unique Identifier for this scan
	Uuid string `json:"uuid,omitempty"`

	// HTTP response code. Used for internal purposes, will be sent out at the API.
	HttpReponseCode int32 `json:"httpReponseCode,omitempty"`
}

// AssertVarroaScanRequired checks if the required fields are not zero-ed
func AssertVarroaScanRequired(obj VarroaScan) error {
	for _, el := range obj.Metadata {
		if err := AssertVarroaScanMetadataRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseVarroaScanRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of VarroaScan (e.g. [][]VarroaScan), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseVarroaScanRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aVarroaScan, ok := obj.(VarroaScan)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertVarroaScanRequired(aVarroaScan)
	})
}
