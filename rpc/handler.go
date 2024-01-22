package main

import (
	"context"
	greet "github.com/supermicah/dionysus/rpc/kitex_gen/greet"
)

// DionysusImpl implements the last service interface defined in the IDL.
type DionysusImpl struct{}

// Greet implements the DionysusImpl interface.
func (s *DionysusImpl) Greet(ctx context.Context, req *greet.GreetRequest) (resp *greet.GreetResponse, err error) {
	// TODO: Your code here...
	return
}
