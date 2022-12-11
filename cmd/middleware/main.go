package main

import (
	"encoding/json"
	"fmt"
	"go-middleware/cmd/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// HandleFunc returns a HTTP handler
	coreLogicHandler := http.HandlerFunc(coreLogic)
	http.Handle("/city", filterContentType(setServerTimeCookie(coreLogicHandler)))
	http.ListenAndServe("localhost:8080", nil)

}

// filterContentType is a middleware that filters requests based on content
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Entered filterContentType middleware")
		//Filtering requests by MIME type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Only JSON is supported"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sets cookie to each response
		cookie := http.Cookie{Name: "Server-Time(UTC", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the setServerTimeCookie middleware")
		handler.ServeHTTP(w, r)
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
