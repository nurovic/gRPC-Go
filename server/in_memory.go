package main

import (
	"time"

	pb "github.com/nurovic/gRPC_Go/proto/todo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)
type inMemoryDb struct {
  tasks []*pb.Task
}
func New() db {
  return &inMemoryDb{}
}
func (d *inMemoryDb) addTask(description string, dueDate time.Time) (uint64, error) {
  nextId := uint64(len(d.tasks) + 1)
  task := &pb.Task{
    Id: nextId,
    Description: description,
    DueDate: timestamppb.New(dueDate),
}
  d.tasks = append(d.tasks, task)
  return nextId, nil
}

func (d *inMemoryDb) getTasks(f func(interface{}) error) error {
	for _, task := range d.tasks {
		if err := f(task); err != nil {
			return err
		}
	}
	return nil
}