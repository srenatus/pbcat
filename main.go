package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

const (
	binFile = "protoset.bin"
	msgFile = "message.bin"
)

func main() {
	protoSet, err := ioutil.ReadFile(binFile)
	if err != nil {
		panic(err)
	}
	var fileSet descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(protoSet, &fileSet); err != nil {
		panic(err)
	}
	files, err := protodesc.NewFiles(&fileSet)
	if err != nil {
		panic(err)
	}

	msgBytes, err := ioutil.ReadFile(msgFile)
	if err != nil {
		panic(err)
	}

	msgDesc, err := files.FindDescriptorByName("tutorial.AddressBook")
	if err != nil {
		panic(err)
	}

	msg := dynamicpb.NewMessage(msgDesc.(protoreflect.MessageDescriptor))
	err = proto.Unmarshal(msgBytes, msg)
	if err != nil {
		panic(err)
	}
	msgJSON, err := protojson.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, string(msgJSON))
}
