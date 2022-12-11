package main

import (
	"fmt"
	"net/http"
)

func main() {
	// HandleFunc returns a HTTP handler
	coreLogicHandler := http.HandlerFunc(coreLogic)
	http.Handle("/", middleware(coreLogicHandler))
	http.ListenAndServe("localhost:8080", nil)

}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before forwarding request to REST API")
		// Passes back control to the handler
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware before sending REST API response to client")
	})
}

func coreLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Simulating some business logic")
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
