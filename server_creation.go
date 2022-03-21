package server

import (
	"fmt"
	"log"
	"net/http"
)

// CreateServer will create a http server in specified port
func CreateServer(port int, handlers map[string]func(w http.ResponseWriter, r *http.Request)) {

	if len(handlers) == 0 {
		log.Fatal("No handlers assigned")
	}

	log.Printf("Starting http server at port %d \n", port)
	server := http.Server{
		Addr: fmt.Sprint(":", port),
	}

	for endpoint, handler := range handlers {
		log.Printf("Endpoint %s registered", endpoint)
		http.HandleFunc(endpoint, handler)
	}

	log.Printf("HTTP Server listening at port %d \n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
