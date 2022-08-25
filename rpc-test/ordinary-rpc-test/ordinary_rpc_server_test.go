package ordinary_rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func Test_RPC_Server(t *testing.T) {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	fmt.Println("RPC服务启动, 端口: 1234")
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
