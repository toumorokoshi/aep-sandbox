package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	spb "github.com/toumorokoshi/aep-sandbox/service/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	spb.RegisterBookStoreServer(s, NewBookStoreServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
