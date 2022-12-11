package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Args is a struct that holds information about arguments passed from the RPC client to the server
// The Args struct here has no fields because this server is not expecting the client to
// send any arguments
type Args struct{}

// TimeServer is a struct that shows the object type that the RPC server wishes to export
type TimeServer int64

func main() {
	timeserver := new(TimeServer)
	//Register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// Listen for requests on port 78
	l, err := net.Listen("tcp", ":78")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}

// GiveServerTime is the function to be called by the RPC client, and returns the current server time
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil

}
