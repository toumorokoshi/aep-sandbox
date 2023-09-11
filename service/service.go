package main

import (
	"context"

	bpb "github.com/toumorokoshi/aep-sandbox/service/bookstore"
)

type BookStoreServer struct {
	bpb.UnimplementedBookStoreServer
}

func NewBookStoreServer() *BookStoreServer {
	return &BookStoreServer{}
}

func (BookStoreServer) CreateBook(context.Context, *bpb.CreateBookRequest) (*bpb.Book, error) {
	return &bpb.Book{}, nil
}

func (BookStoreServer) ReadBook(c context.Context, r *bpb.ReadBookRequest) (*bpb.Book, error) {
	return &bpb.Book{
		Path: r.Path,
	}, nil
}
