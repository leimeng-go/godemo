package rpc

import (
	"net"
	"net/rpc"
	"testing"
)
func TestHello1Server(t *testing.T){
	err:=RegisterHelloService(new(HelloService))
	if err!=nil{
		t.Fatal(err.Error())
	}
	listener,err:=net.Listen("tcp",":1234")
	if err!=nil{
		t.Fatal(err.Error())
	}
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			t.Fatal("Accept error:",err)
		}
		go rpc.ServeConn(conn)
	}
}
func TestHello1Client(t *testing.T) {
    client,err:=DialHelloService("tcp","localhost:1234")
	if err!=nil{
		t.Fatal("dialing:",err)
	}
	var reply string
	err=client.Hello("haha",&reply)
	if err!=nil{
		t.Fatal(err)
	}
	t.Logf("响应内容: %s",reply)
}
