package main

import (
	"encoding/json"
	"fmt"
	"go-middleware/cmd/models"
	"net/http"
)

func main() {
	// HandleFunc returns a HTTP handler
	coreLogicHandler := http.HandlerFunc(coreLogic)
	http.Handle("/city", middleware(coreLogicHandler))
	http.ListenAndServe("localhost:8080", nil)

}

// middleware accepts a handler and returns a handler
func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before forwarding request to REST API")
		// Passes back control to the handler
		// This allows a handler to execute the handler logic, that is coreLogic
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware before sending REST API response to client")
	})
}

func coreLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	// Check if HTTP method is POST
	if r.Method == "POST" {
		var city models.City
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&city)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		//ToDO: Resource creation logic goes here
		fmt.Printf("Got %s City with area of %d sq miles!\n", city.Name, city.Area)
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}
