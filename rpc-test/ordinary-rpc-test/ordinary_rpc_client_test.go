package ordinary_rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func Test_RPC_Client(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "hello fly", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("拨号连接本机1234端口的RPC服务成功, 返回: ", reply)
}