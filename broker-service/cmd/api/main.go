package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

// Goals:
// (1) Take requests,
// (2) Forward them to the correct microservice,
// (3) Send response back to client
func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
