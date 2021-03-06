/*
 * Devices Microservice
 *
 * Microservice for managing Giò Plants devices
 *
 * API version: 1.0.0
 * Contact: andrea.liut@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"flag"
	"fmt"
	"gio-device-ms/pkg/api"
	"gio-device-ms/pkg/model"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	checkVariables()

	port := flag.Int("port", 8080, "port to be used")

	flag.Parse()

	if err := model.Init(); err != nil {
		panic(err)
	}

	log.Printf("Server started on port %d", *port)

	router := api.NewRouter()

	p := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(p, router))
}

func checkVariables() {
	count, err := strconv.Atoi(os.Getenv("DEVICE_DRIVER_COUNT"))
	if err != nil {
		panic("invalid DEVICE_DRIVER_COUNT")
	}

	for i := 0; i < count; i++ {
		hostKey := fmt.Sprintf("DEVICE_DRIVER_%d_HOST", i)
		if host := os.Getenv(hostKey); host == "" {
			panic(fmt.Sprintf("%s not set", hostKey))
		}

		portKey := fmt.Sprintf("DEVICE_DRIVER_%d_PORT", i)
		if port := os.Getenv(portKey); port == "" {
			panic(fmt.Sprintf("%s not set", portKey))
		}
	}
}
