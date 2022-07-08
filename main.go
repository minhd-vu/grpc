package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/minhd-vu/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	addressBook = api.AddressBook{}
)

type server struct {
	api.UnimplementedServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetAddressBook(context.Context, *emptypb.Empty) (*api.AddressBook, error) {
	return &addressBook, nil
}

func (s *server) GetPerson(_ context.Context, request *api.Person) (*api.Person, error) {
	for _, person := range addressBook.People {
		if person.Id == request.Id {
			return person, nil
		}
	}
	return nil, errors.New("person not found")
}

func (s *server) AddPerson(_ context.Context, person *api.Person) (*emptypb.Empty, error) {
	addressBook.People = append(addressBook.People, person)
	return &emptypb.Empty{}, nil
}
