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

// AuthenticateRegistrationIdPost - authenticate against the internal authentication service with a registrationId.
func (s *DefaultApiService) AuthenticateRegistrationIdPost(ctx context.Context, registrationId RegistrationId) (ImplResponse, error) {
	// TODO - update AuthenticateRegistrationIdPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, AuthenticatePostResponse{}) or use other options such as http.Ok ...
	//return Response(200, AuthenticatePostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AuthenticateRegistrationIdPost method not implemented")
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

// ConfigsBboxDelete - Delete a bbox
func (s *DefaultApiService) ConfigsBboxDelete(ctx context.Context, qToken string, uuid string, token string, userId int64) (ImplResponse, error) {
	// TODO - update ConfigsBboxDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsBboxDelete method not implemented")
}

// ConfigsBboxGet - Get Config metadata of bboxes
func (s *DefaultApiService) ConfigsBboxGet(ctx context.Context, qToken string, token string, uuid string, userId int64) (ImplResponse, error) {
	// TODO - update ConfigsBboxGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetBboxResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetBboxResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsBboxGet method not implemented")
}

// ConfigsBboxPost - Create bbox
func (s *DefaultApiService) ConfigsBboxPost(ctx context.Context, qToken string, token string, userId int64, bbox Bbox) (ImplResponse, error) {
	// TODO - update ConfigsBboxPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, PostBboxResponse{}) or use other options such as http.Ok ...
	//return Response(201, PostBboxResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsBboxPost method not implemented")
}

// ConfigsBboxPut - Update bbox metadata
func (s *DefaultApiService) ConfigsBboxPut(ctx context.Context, qToken string, uuid string, token string, userId int64, bbox Bbox) (ImplResponse, error) {
	// TODO - update ConfigsBboxPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, PutBboxResponse{}) or use other options such as http.Ok ...
	//return Response(201, PutBboxResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsBboxPut method not implemented")
}

// ConfigsBboxRegisterPost - register bbox
func (s *DefaultApiService) ConfigsBboxRegisterPost(ctx context.Context, userId int64, bbox Bbox) (ImplResponse, error) {
	// TODO - update ConfigsBboxRegisterPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, BboxConfigResponse{}) or use other options such as http.Ok ...
	//return Response(200, BboxConfigResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsBboxRegisterPost method not implemented")
}

// ConfigsBhiveAssociatePost - associate bhive (sensors) with a logical hive
func (s *DefaultApiService) ConfigsBhiveAssociatePost(ctx context.Context, bhiveId string, hiveUuid string, userId int64) (ImplResponse, error) {
	// TODO - update ConfigsBhiveAssociatePost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ConfigsBhiveAssociatePost method not implemented")
}

// HivesDelete - Delete a Hive
func (s *DefaultApiService) HivesDelete(ctx context.Context, qToken string, uuid string, token string, userId int64) (ImplResponse, error) {
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
func (s *DefaultApiService) HivesGet(ctx context.Context, epoch int64, secondsInThePast int64, uuid string, userId int64) (ImplResponse, error) {
	// TODO - update HivesGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetHivesResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetHivesResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("HivesGet method not implemented")
}

// HivesPost - Create Hive metadata
func (s *DefaultApiService) HivesPost(ctx context.Context, qToken string, token string, userId int64, hive Hive) (ImplResponse, error) {
	// TODO - update HivesPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PostHivesResponse{}) or use other options such as http.Ok ...
	//return Response(200, PostHivesResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("HivesPost method not implemented")
}

// HivesPut - Update Hive metadata
func (s *DefaultApiService) HivesPut(ctx context.Context, qToken string, token string, userId int64, hive Hive) (ImplResponse, error) {
	// TODO - update HivesPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PutHiveResponse{}) or use other options such as http.Ok ...
	//return Response(200, PutHiveResponse{}), nil

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

	//TODO: Uncomment the next line to return response Response(200, LoginPostResponse{}) or use other options such as http.Ok ...
	//return Response(200, LoginPostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, string{}) or use other options such as http.Ok ...
	//return Response(400, string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LoginPost method not implemented")
}

// ScaleGet - Get Scale values
func (s *DefaultApiService) ScaleGet(ctx context.Context, bhiveId string, epoch int64, secondsInThePast int64, qToken string, token string) (ImplResponse, error) {
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

// ScaleGetV2 - Get Scale values
func (s *DefaultApiService) ScaleGetV2(ctx context.Context, macAddress string, epoch int64, secondsInThePast int64, qToken string, token string, userId int64) (ImplResponse, error) {
	// TODO - update ScaleGetV2 with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []ScaleV2{}) or use other options such as http.Ok ...
	//return Response(200, []ScaleV2{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ScaleGetV2 method not implemented")
}

// ScalePost - Post Scale values
func (s *DefaultApiService) ScalePost(ctx context.Context, registrationId string, qToken string, token string, userId int64, weight Weight) (ImplResponse, error) {
	// TODO - update ScalePost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(200, GenericPostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(400, GenericPostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(500, GenericPostResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ScalePost method not implemented")
}

// ScalePostV2 - Post Scale values
func (s *DefaultApiService) ScalePostV2(ctx context.Context, registrationId string, userId int64, scaleV2 ScaleV2) (ImplResponse, error) {
	// TODO - update ScalePostV2 with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(200, GenericPostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(400, GenericPostResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, GenericPostResponse{}) or use other options such as http.Ok ...
	//return Response(500, GenericPostResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ScalePostV2 method not implemented")
}

// StandsDelete - Delete a stand
func (s *DefaultApiService) StandsDelete(ctx context.Context, qToken string, uuid string, token string, userId int64) (ImplResponse, error) {
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

	//TODO: Uncomment the next line to return response Response(201, PostStandsResponse{}) or use other options such as http.Ok ...
	//return Response(201, PostStandsResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("StandsPost method not implemented")
}

// StandsPut - Update stand metadata
func (s *DefaultApiService) StandsPut(ctx context.Context, qToken string, token string, userId int64, stand Stand) (ImplResponse, error) {
	// TODO - update StandsPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, PutStandResponse{}) or use other options such as http.Ok ...
	//return Response(201, PutStandResponse{}), nil

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
