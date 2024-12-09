package main

import (
	pb "github.com/nurovic/gRPC_Go/proto/todo/v1"
)
type server struct {
d db
  pb.UnimplementedTodoServiceServer
}