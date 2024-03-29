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



type VarroaScanMetadata struct {

	// classification of the object - we currently support: bee_leg,bee,mite,bee_wing,ant,wax_moth_droppings,wax_platelets,pollen,bee_droppings,cell_cover_grist
	Class string `json:"class,omitempty"`

	// how confident the AI is regarding the result
	Confidence float64 `json:"confidence,omitempty"`

	// center of the object on the x axis
	XCenter float64 `json:"xCenter,omitempty"`

	// center of the object on the y axis
	YCenter float64 `json:"yCenter,omitempty"`

	// width of the object
	Width float64 `json:"width,omitempty"`

	// height of the object
	Height float64 `json:"height,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *VarroaScanMetadata) UnmarshalJSON(data []byte) error {

	type Alias VarroaScanMetadata // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertVarroaScanMetadataRequired checks if the required fields are not zero-ed
func AssertVarroaScanMetadataRequired(obj VarroaScanMetadata) error {
	return nil
}

// AssertVarroaScanMetadataConstraints checks if the values respects the defined constraints
func AssertVarroaScanMetadataConstraints(obj VarroaScanMetadata) error {
	return nil
}
