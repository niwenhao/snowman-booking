/*
 * スキー予約サービス
 *
 * このサービスはスキー場一覧、スキー場のホテルの一覧、スキー場のホテルに予約することできます。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface { 
	ResortsGet(http.ResponseWriter, *http.Request)
	ResortsResortIdHotelsGet(http.ResponseWriter, *http.Request)
	ResortsResortIdHotelsHotelIdBookingsOrderNoDelete(http.ResponseWriter, *http.Request)
	ResortsResortIdHotelsHotelIdBookingsOrderNoPut(http.ResponseWriter, *http.Request)
	ResortsResortIdHotelsHotelIdBookingsPost(http.ResponseWriter, *http.Request)
}


// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface { 
	ResortsGet(context.Context) (ImplResponse, error)
	ResortsResortIdHotelsGet(context.Context, int32) (ImplResponse, error)
	ResortsResortIdHotelsHotelIdBookingsOrderNoDelete(context.Context, int32, int32, string) (ImplResponse, error)
	ResortsResortIdHotelsHotelIdBookingsOrderNoPut(context.Context, int32, int32, string, Booking) (ImplResponse, error)
	ResortsResortIdHotelsHotelIdBookingsPost(context.Context, int32, int32, Booking) (ImplResponse, error)
}
