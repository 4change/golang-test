package secondary_rpc

import (
	"log"
	"net/rpc"
	"testing"
)

func Test_Secondary_RPC_Client_2(t *testing.T) {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Secondary_RPC_Client_1(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call(HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
}
