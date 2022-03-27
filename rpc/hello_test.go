package rpc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestHelloServer(t *testing.T){
   rpc.RegisterName("HelloService",new(HelloService))
   listener,err:=net.Listen("tcp",":1234")
   if err!=nil{
	   t.Fatal("ListenTCP error:",err)
   }
   conn,err:=listener.Accept()
   if err!=nil{
	   t.Fatal("Accept error:",err)
   }
   rpc.ServeConn(conn)
}
func TestHelloJsonRpcServer(t *testing.T){

}
func TestHelloJsonRpcClient(t *testing.T){
	client,err:=rpc.Dial("tcp","localhost:1234")
	if err!=nil{
		t.Fatal("dialing: ",err)
	}
	
	var reply string
	err=client.Call("HelloService.Hello","pingguodeli",&reply)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(reply)
}

func TestHelloJson(t *testing.T){
	conn,err:=net.Dial("tcp","localhost:1234")
	if err!=nil{
		t.Fatal("net.Dial:",err.Error())
	}
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err=client.Call("HelloService.Hello","haha",&reply)
	if err!=nil{
		t.Fatal(err.Error())
	}
	t.Logf("响应内容: %s",reply)
}