package main

import (
	"context"

	basic "github.com/supermicah/dionysus/basic/kitex_gen/dionysus/basic"
)

// DionysusBasicImpl implements the last service interface defined in the IDL.
type DionysusBasicImpl struct{}

// SayHello implements the DionysusBasicImpl interface.
func (s *DionysusBasicImpl) SayHello(ctx context.Context, req *basic.SayHelloRequest) (resp *basic.SayHelloResponse, err error) {
	// TODO: Your code here...
	return
}
