package main

import (
	"flag"
	grpcSetup "github.com/nttungqn/m-highscore/internal/server/grpc"
	
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address where you connect with m-highscore service")
	flag.Parse()
	
}