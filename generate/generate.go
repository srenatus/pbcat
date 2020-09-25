package main

import (
	"os"

	"github.com/srenatus/pbcat/envelopes"
	pb "github.com/srenatus/pbcat/tutorial"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/anypb"
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

	anyMsg := anypb.Any{}
	err := anypb.MarshalFrom(&anyMsg, book, proto.MarshalOptions{})
	if err != nil {
		panic(err)
	}

	fileDesc := book.ProtoReflect().Descriptor().ParentFile()
	ds := []*descriptorpb.FileDescriptorProto{protodesc.ToFileDescriptorProto(fileDesc)}
	imports := fileDesc.Imports()
	for i := 0; i < imports.Len(); i++ {
		ds = append(ds, protodesc.ToFileDescriptorProto(imports.Get(i).FileDescriptor))
	}

	env := envelopes.EnvelopeWithDescriptor{
		DescriptorSet: &descriptorpb.FileDescriptorSet{
			File: ds,
		},
		Message: &anyMsg,
	}
	out, err := proto.Marshal(&env)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stdout.Write(out); err != nil {
		panic(err)
	}
}
