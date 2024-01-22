package main

import (
	rpc "github.com/supermicah/dionysus/rpc/kitex_gen/rpc/dionysus"
	"log"
)

func main() {
	svr := rpc.NewServer(new(DionysusImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
