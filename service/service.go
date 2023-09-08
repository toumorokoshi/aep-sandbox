package main

import (
	"context"

	spb "github.com/toumorokoshi/aep-sandbox/service/proto"
)

type BookStoreServer struct {
	spb.UnimplementedBookStoreServer
}

func NewBookStoreServer() *BookStoreServer {
	return &BookStoreServer{}
}

func (BookStoreServer) CreateBook(context.Context, *spb.CreateBookRequest) (*spb.Book, error) {
	return &spb.Book{}, nil
}

func (BookStoreServer) GetBook(context.Context, *spb.GetBookRequest) (*spb.Book, error) {
	return &spb.Book{}, nil
}
