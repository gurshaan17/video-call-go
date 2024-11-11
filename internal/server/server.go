package server

import (
	"flag"
	"os"
)

var (
	address = flag.String("addr,":"", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key = flag.String("key", "", "")
)



func server(){

}