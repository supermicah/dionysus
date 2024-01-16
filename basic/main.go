package main

import (
	basic "github.com/supermicah/dionysus/basic/kitex_gen/dionysus/basic/dionysusbasic"
	"log"
)

func main() {
	svr := basic.NewServer(new(DionysusBasicImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
