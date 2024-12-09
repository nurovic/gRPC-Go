package main

import (
	"context"
	"time"

	pb "github.com/nurovic/gRPC_Go/proto/todo/v1"
)


func (s *server) AddTask(_ context.Context, in *pb.AddTaskRequest) (*pb.AddTaskResponse, error) {
	id, _ := s.d.addTask(in.Description, in.DueDate.AsTime())

	return &pb.AddTaskResponse{Id: id}, nil
}
func (s *server) ListTasks(req *pb.ListTasksRequest, stream pb.TodoService_ListTasksServer) error {
	return s.d.getTasks(func(t interface{}) error {
	  task := t.(*pb.Task)
	  overdue := task.DueDate != nil && !task.Done &&
		task.DueDate.AsTime().Before(time.Now().UTC())
	  err := stream.Send(&pb.ListTasksResponse{
  		Task: task,
		Overdue: overdue,
	  })
  return err })
  }