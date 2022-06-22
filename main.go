package main

import (
	"fmt"

	"github.com/minhd-vu/grpc/api"
)

func main() {
	person := api.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*api.Person_PhoneNumber{{
			Number: "555-4321",
			Type:   api.Person_HOME,
		}},
	}
	fmt.Println(person.String())
}
