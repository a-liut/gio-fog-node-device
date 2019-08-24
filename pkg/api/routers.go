/*
 * Devices service
 *
 * Microservice for managing Giò Plants devices
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"encoding/json"
	"gio-device-ms/pkg/logging"
	"gio-device-ms/pkg/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func errorHandler(w http.ResponseWriter, status int32, message string) {
	r := model.ApiResponse{Code: status, Message: message}
	w.WriteHeader(int(status))
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println(err)
	}
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logging.Logger(handler, route.Name)

		router.
			Methods(route.Method...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	Route{
		"GetDeviceById",
		[]string{http.MethodGet},
		"/devices/{deviceId}",
		GetDeviceById,
	},

	Route{
		"GetDevices",
		[]string{http.MethodGet},
		"/devices",
		GetDevices,
	},

	Route{
		"CreateDevice",
		[]string{http.MethodPost},
		"/devices",
		CreateDevice,
	},

	Route{
		"GetDeviceReadings",
		[]string{http.MethodGet},
		"/devices/{deviceId}/readings",
		GetDeviceReadings,
	},

	Route{
		"CreateDeviceReadings",
		[]string{http.MethodPost},
		"/devices/{deviceId}/readings",
		CreateDeviceReadings,
	},

	Route{
		"GetRoomById",
		[]string{http.MethodGet},
		"/rooms/{roomId}",
		GetRoomById,
	},

	Route{
		"GetRooms",
		[]string{http.MethodGet},
		"/rooms",
		GetRooms,
	},

	Route{
		"CreateRoom",
		[]string{http.MethodPost},
		"/rooms",
		CreateRoom,
	},
}
