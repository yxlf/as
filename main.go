package main

import (
	"as/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	server := &http.Server{
		Addr:              "0.0.0.0:8080",
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	server.ListenAndServe()
}
