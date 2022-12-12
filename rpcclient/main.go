package main

import (
	"go-middleware/models"
	"log"
	"net/rpc"
)

const DefaultRPCPort = ":78"

func main() {
	var reply int64
	args := models.Args{}
	client, err := rpc.DialHTTP("tcp", "localhost"+DefaultRPCPort)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call("JSONServer.GiveBookDetail", args, &reply)
	if err != nil {
		log.Fatal("arithmetic error:", err)
	}
	log.Printf("reply %d", reply)

}
