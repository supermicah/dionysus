package main

import (
	"context"

	greet "github.com/supermicah/dionysus/rpc/kitex_gen/greet"
)

// DionysusImpl implements the last service interface defined in the IDL.
type DionysusImpl struct{}

// Greet implements the DionysusImpl interface.
func (s *DionysusImpl) Greet(_ context.Context, _ *greet.GreetRequest) (resp *greet.GreetResponse, err error) {
	return
}
