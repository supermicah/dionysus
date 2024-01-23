// Code generated by Kitex v0.7.3. DO NOT EDIT.

package dionysus

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	greet "github.com/supermicah/dionysus/rpc/kitex_gen/greet"
	rpc "github.com/supermicah/dionysus/rpc/kitex_gen/rpc"
)

func serviceInfo() *kitex.ServiceInfo {
	return dionysusServiceInfo
}

var dionysusServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Dionysus"
	handlerType := (*rpc.Dionysus)(nil)
	methods := map[string]kitex.MethodInfo{
		"greet": kitex.NewMethodInfo(greetHandler, newDionysusGreetArgs, newDionysusGreetResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "rpc",
		"ServiceFilePath": `../idl/rpc.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func greetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*rpc.DionysusGreetArgs)
	realResult := result.(*rpc.DionysusGreetResult)
	success, err := handler.(rpc.Dionysus).Greet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDionysusGreetArgs() interface{} {
	return rpc.NewDionysusGreetArgs()
}

func newDionysusGreetResult() interface{} {
	return rpc.NewDionysusGreetResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Greet(ctx context.Context, req *greet.GreetRequest) (r *greet.GreetResponse, err error) {
	var _args rpc.DionysusGreetArgs
	_args.Req = req
	var _result rpc.DionysusGreetResult
	if err = p.c.Call(ctx, "greet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}