package main

import (
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/srenatus/pbcat/tutorial"
)

func main() {
	p := &pb.Person{
		Id:    123,
		Name:  "bert",
		Email: "bert@ernie.com",
		Phones: []*pb.Person_PhoneNumber{{
			Number: "0123-222",
			Type:   pb.Person_HOME,
		}},
	}

	book := &pb.AddressBook{
		People: []*pb.Person{p},
	}

	out, err := proto.Marshal(book)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stdout.Write(out); err != nil {
		panic(err)
	}
}
