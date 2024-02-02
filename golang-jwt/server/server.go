package server

import (
	"log"
	"net/http"

	"golang-jwt/middleware"
)

func StartServer(hostName, port string) error {
	host := hostName + ":" + port
	log.Printf("Listening on: %s", host)
	handler := middleware.NewHandler()

	http.Handle("/", handler)
	return http.ListenAndServe(host, nil)
}
