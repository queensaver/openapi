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
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	BboxesGet(http.ResponseWriter, *http.Request)
	HivesDelete(http.ResponseWriter, *http.Request)
	HivesGet(http.ResponseWriter, *http.Request)
	HivesPost(http.ResponseWriter, *http.Request)
	HivesPut(http.ResponseWriter, *http.Request)
	LoginPost(http.ResponseWriter, *http.Request)
	ScaleGet(http.ResponseWriter, *http.Request)
	TemperatureGet(http.ResponseWriter, *http.Request)
	TemperaturePost(http.ResponseWriter, *http.Request)
	UserPost(http.ResponseWriter, *http.Request)
	VarroaScanGet(http.ResponseWriter, *http.Request)
	VarroaScanImagePost(http.ResponseWriter, *http.Request)
	VarroaScanPost(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	BboxesGet(context.Context) (ImplResponse, error)
	HivesDelete(context.Context, string) (ImplResponse, error)
	HivesGet(context.Context, string, string) (ImplResponse, error)
	HivesPost(context.Context, Hive) (ImplResponse, error)
	HivesPut(context.Context, Hive) (ImplResponse, error)
	LoginPost(context.Context, string, string) (ImplResponse, error)
	ScaleGet(context.Context, string, string, int64, int64, string) (ImplResponse, error)
	TemperatureGet(context.Context, string, string, int64, int64, string, int64) (ImplResponse, error)
	TemperaturePost(context.Context, NewTemperature) (ImplResponse, error)
	UserPost(context.Context, User) (ImplResponse, error)
	VarroaScanGet(context.Context, string, string, string, int64, string, int64, int64) (ImplResponse, error)
	VarroaScanImagePost(context.Context, string, string, int64, string, int64, string) (ImplResponse, error)
	VarroaScanPost(context.Context, int64, VarroaScan) (ImplResponse, error)
}
