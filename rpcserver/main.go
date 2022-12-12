package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	rpcJson "github.com/gorilla/rpc/json"
	"go-middleware/models"
	"log"
	"net/http"
	"os"
)

type JSONServer struct{}

func main() {
	//Register RPC server
	s := rpc.NewServer()
	//Registers the type of data as JSON
	s.RegisterCodec(rpcJson.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	// Listen for requests on port 78
	http.ListenAndServe(":78", r)
}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *models.Args, reply *models.Book) error {
	var books []models.Book
	raw, err := os.ReadFile("./books.json")
	if err != nil {
		log.Println("error occurred reading books:", err)
		os.Exit(1)
	}

	// Unmarshalls JSON raw data into books array
	err = json.Unmarshal(raw, &books)
	if err != nil {
		log.Println("error unmarshalling JSON:", err)
		os.Exit(1)
	}

	//Iterates over each book to find the specified book
	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
			break
		}
	}
	return nil
}
