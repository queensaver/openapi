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
	"context"
	"net/http"
	"errors"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// BboxesGet - Get QBox metadata
func (s *DefaultApiService) BboxesGet(ctx context.Context) (ImplResponse, error) {
	// TODO - update BboxesGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Bbox{}) or use other options such as http.Ok ...
	//return Response(200, []Bbox{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("BboxesGet method not implemented")
}

// HivesDelete - Delete a Hive
func (s *DefaultApiService) HivesDelete(ctx context.Context, uuid string) (ImplResponse, error) {
	// TODO - update HivesDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("HivesDelete method not implemented")
}

// HivesGet - Get Hive metadata
func (s *DefaultApiService) HivesGet(ctx context.Context, epoch int64, secondsInThePast int64, uuid string) (ImplResponse, error) {
	// TODO - update HivesGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Hive{}) or use other options such as http.Ok ...
	//return Response(200, []Hive{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("HivesGet method not implemented")
}

// HivesPost - Create Hive metadata
func (s *DefaultApiService) HivesPost(ctx context.Context, hive Hive) (ImplResponse, error) {
	// TODO - update HivesPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Hive{}) or use other options such as http.Ok ...
	//return Response(200, Hive{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("HivesPost method not implemented")
}

// HivesPut - Update Hive metadata
func (s *DefaultApiService) HivesPut(ctx context.Context, hive Hive) (ImplResponse, error) {
	// TODO - update HivesPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Hive{}) or use other options such as http.Ok ...
	//return Response(200, Hive{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("HivesPut method not implemented")
}

// LoginPost - Authenticate a user against the system.
func (s *DefaultApiService) LoginPost(ctx context.Context, username string, password string) (ImplResponse, error) {
	// TODO - update LoginPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, string{}) or use other options such as http.Ok ...
	//return Response(400, string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LoginPost method not implemented")
}

// ScaleGet - Get Scale values
func (s *DefaultApiService) ScaleGet(ctx context.Context, qToken string, bhiveId string, epoch int64, secondsInThePast int64, token string) (ImplResponse, error) {
	// TODO - update ScaleGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Weight{}) or use other options such as http.Ok ...
	//return Response(200, []Weight{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ScaleGet method not implemented")
}

// StandsDelete - Delete a stand
func (s *DefaultApiService) StandsDelete(ctx context.Context, uuid string) (ImplResponse, error) {
	// TODO - update StandsDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("StandsDelete method not implemented")
}

// StandsGet - Get Stand metadata
func (s *DefaultApiService) StandsGet(ctx context.Context, qToken string, epoch int64, secondsInThePast int64, token string, uuid string, userId int64) (ImplResponse, error) {
	// TODO - update StandsGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetStandsResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetStandsResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("StandsGet method not implemented")
}

// StandsPost - Create stand metadata
func (s *DefaultApiService) StandsPost(ctx context.Context, qToken string, token string, userId int64, stand Stand) (ImplResponse, error) {
	// TODO - update StandsPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Stand{}) or use other options such as http.Ok ...
	//return Response(200, Stand{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("StandsPost method not implemented")
}

// StandsPut - Update stand metadata
func (s *DefaultApiService) StandsPut(ctx context.Context, stand Stand) (ImplResponse, error) {
	// TODO - update StandsPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Stand{}) or use other options such as http.Ok ...
	//return Response(200, Stand{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("StandsPut method not implemented")
}

// TemperatureGet - Get Temperature values
func (s *DefaultApiService) TemperatureGet(ctx context.Context, qToken string, bhiveId string, epoch int64, secondsInThePast int64, token string, userId int64) (ImplResponse, error) {
	// TODO - update TemperatureGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetTemperatureResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetTemperatureResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("TemperatureGet method not implemented")
}

// TemperaturePost - Save a new temperature measurement
func (s *DefaultApiService) TemperaturePost(ctx context.Context, userId int64, temperature Temperature) (ImplResponse, error) {
	// TODO - update TemperaturePost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(200, GenericPostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(400, GenericPostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(500, GenericPostResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("TemperaturePost method not implemented")
}

// UserPost - Create a user
func (s *DefaultApiService) UserPost(ctx context.Context, user User) (ImplResponse, error) {
	// TODO - update UserPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, string{}) or use other options such as http.Ok ...
	//return Response(400, string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UserPost method not implemented")
}

// VarroaScanGet - Get Varroa Scan images and metadata
func (s *DefaultApiService) VarroaScanGet(ctx context.Context, qToken string, bhiveId string, token string, epoch int64, uuid string, userId int64, secondsInThePast int64) (ImplResponse, error) {
	// TODO - update VarroaScanGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, VarroaScanResponse{}) or use other options such as http.Ok ...
	//return Response(200, VarroaScanResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("VarroaScanGet method not implemented")
}

// VarroaScanImagePost - Save Varroa Scan image
func (s *DefaultApiService) VarroaScanImagePost(ctx context.Context, qToken string, token string, userId int64, bhiveId string, epoch int64, scan string) (ImplResponse, error) {
	// TODO - update VarroaScanImagePost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("VarroaScanImagePost method not implemented")
}

// VarroaScanPost - Save Varroa Scan metadata
func (s *DefaultApiService) VarroaScanPost(ctx context.Context, userId int64, varroaScan VarroaScan) (ImplResponse, error) {
	// TODO - update VarroaScanPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("VarroaScanPost method not implemented")
}
