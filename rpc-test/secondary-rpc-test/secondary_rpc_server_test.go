package secondary_rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"testing"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func Test_Secondary_RPC_Server(t *testing.T) {
	RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		fmt.Println("RPC服务端已启动, 监听端口: 1234")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
