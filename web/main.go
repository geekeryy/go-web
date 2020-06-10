package main

import (
	"flag"

	"go-web/web/conf"
)

func main(){
	conf.Init()
	flag.Parse()
}
