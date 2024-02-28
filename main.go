package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	// & is used to create a pointer to a variable
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	// ListenAndServe starts an HTTP server with a given address and handler
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("An error occured: ", err)
	}
}

// First paramter is the response writer, which is used to write the response back to the client
// second is the request object which is the inboud request from the client
// we can tell this is a pointer because of the * in front
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World from Orders-API!"))

}
