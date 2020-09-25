package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/srenatus/pbcat/envelopes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

const (
	msgFile = "message.bin"
)

func main() {
	msgBytes, err := ioutil.ReadFile(msgFile)
	if err != nil {
		panic(err)
	}

	envMsg := envelopes.EnvelopeWithDescriptor{}
	err = proto.Unmarshal(msgBytes, &envMsg)
	if err != nil {
		panic(err)
	}

	files, err := protodesc.NewFiles(envMsg.DescriptorSet)
	if err != nil {
		panic(err)
	}

	name := strings.Split(envMsg.Message.TypeUrl, "/")[1]
	msgDesc, err := files.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		panic(err)
	}
	msg := dynamicpb.NewMessage(msgDesc.(protoreflect.MessageDescriptor))
	err = proto.Unmarshal(envMsg.Message.Value, msg)
	if err != nil {
		panic(err)
	}
	msgJSON, err := protojson.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, string(msgJSON))
}
