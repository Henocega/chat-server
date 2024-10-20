package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	chat "github.com/Henocega/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
)

const grpcPort = 50051

type server struct {
	chat.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, req *chat.CreateRequest) (*chat.CreateResponse, error) {
	fmt.Printf("Create request: %v", req)
	log.Printf("Context: %v", ctx)

	return &chat.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *chat.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete request: %v", req)
	log.Printf("Context: %v", ctx)
	return nil, nil
}

func (s *server) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf("Send message request: %v", req)
	log.Printf("Context: %v", ctx)
	return nil, nil
}
