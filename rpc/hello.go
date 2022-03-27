package rpc
import (
	pb "godemo/protobuf"
)

type HelloService struct{}
// Hello heelo方法必须满足Go语言的RPC规则：方法只能有两个序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开方法。
func (h *HelloService)Hello(request *pb.String,reply *pb.String)error{
	reply.Value="hello:"+request.GetValue()
	return nil
}