// Code generated by Kitex v0.7.3. DO NOT EDIT.

package dionysusbasic

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	basic "github.com/supermicah/dionysus/basic/kitex_gen/dionysus/basic"
)

func serviceInfo() *kitex.ServiceInfo {
	return dionysusBasicServiceInfo
}

var dionysusBasicServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "DionysusBasic"
	handlerType := (*basic.DionysusBasic)(nil)
	methods := map[string]kitex.MethodInfo{
		"sayHello": kitex.NewMethodInfo(sayHelloHandler, newDionysusBasicSayHelloArgs, newDionysusBasicSayHelloResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "basic",
		"ServiceFilePath": `../idl/dionysus_basic.thrift`,
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

func sayHelloHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*basic.DionysusBasicSayHelloArgs)
	realResult := result.(*basic.DionysusBasicSayHelloResult)
	success, err := handler.(basic.DionysusBasic).SayHello(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDionysusBasicSayHelloArgs() interface{} {
	return basic.NewDionysusBasicSayHelloArgs()
}

func newDionysusBasicSayHelloResult() interface{} {
	return basic.NewDionysusBasicSayHelloResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SayHello(ctx context.Context, req *basic.SayHelloRequest) (r *basic.SayHelloResponse, err error) {
	var _args basic.DionysusBasicSayHelloArgs
	_args.Req = req
	var _result basic.DionysusBasicSayHelloResult
	if err = p.c.Call(ctx, "sayHello", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}