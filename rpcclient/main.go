package main

import (
	"log"
	"net/rpc"
)

type Args struct{}

const DefaultRPCPort = ":78"

func main() {
	var reply int64
	args := Args{}
	client, err := rpc.DialHTTP("tcp", "localhost"+DefaultRPCPort)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("reply %d", reply)

}
