package rpc

import (
	"net/rpc"

)

//HelloServiceName 定义服务名称
const HelloServiceName="path/to/pkg.HelloService"

//HelloServiceInterface 服务实现的详细方法列表
type HelloServiceInterface=interface{
	Hello(request string,reply *string)error
}

//RegisterHelloService 注册服务
func RegisterHelloService(svc HelloServiceInterface)error{
	return rpc.RegisterName(HelloServiceName,svc)
}

type HelloServiceClient struct{
	*rpc.Client
}

var _HelloServiceInterface=(*HelloServiceClient)(nil)

func DialHelloService(network,address string)(*HelloServiceClient,error){
	c,err:=rpc.Dial(network,address)
	if err!=nil{
		return nil,err
	}
	return &HelloServiceClient{Client: c},nil
}
func (p *HelloServiceClient)Hello(request string,reply *string)error{
	return p.Client.Call(HelloServiceName+".Hello",request,reply)
}
