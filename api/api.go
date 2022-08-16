package api

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) GetStudent(ctx context.Context, in *StudentMessageRequest) (*StudentMessageResponse, error) {
	log.Printf("Got a Get Student request from client: %s", in.Name)
	return &StudentMessageResponse{Id: 1234, Name: "Aman"}, nil
}

