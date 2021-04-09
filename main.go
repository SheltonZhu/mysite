package main

import (
	"flag"
	"fmt"
	"github.com/SheltonZhu/mysite/grpc"
	"github.com/SheltonZhu/mysite/routers"
	"github.com/SheltonZhu/mysite/services"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr   = flag.String("http", ":8000", "Http listen address.")
	webEnabled = flag.Bool("web", true, "Enable web server. Default enabled.")
	masterAddr = flag.String("master", "", "RPC master address.")
	rpcEnabled = flag.Bool("rpc", false, "Enable RPC server. Default disabled.")
	rpcPort    = flag.Int("rpc_port", 10000, "The RPC server port")
)

func main() {
	flag.Parse()
	if *rpcEnabled { // the master is the rpc server:
		grpc.InitServer(*rpcPort)
	}

	if *masterAddr != "" {
		services.MasterAddr = *masterAddr
	}
	if *webEnabled {
		router := routers.InitRouter()
		if err := router.Run(*httpAddr); err != nil {
			panic(fmt.Sprintf("http serve at %v start failed.", *httpAddr))
		}
	}
}
