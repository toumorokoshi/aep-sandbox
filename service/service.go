package main

import (
	pb "github.com/toumorokoshi/aep-sandbox/service/proto"
)

type BookStoreServer struct {
	pb.UnimplementedBookStoreServer
}

func NewBookStoreServer() *BookStoreServer {
	return &BookStoreServer{}
}
