protoset.bin tutorial/addressbook.pb.go: tutorial/addressbook.proto
	protoc --descriptor_set_out=protoset.bin --include_imports --go_out=paths=source_relative:. "$<"

envelopes/envs.pb.go: envelopes/envs.proto
	protoc --go_out=paths=source_relative:. "$<"

message.bin: generate/generate.go tutorial/addressbook.pb.go
	go run generate/generate.go > "$@"

example: message.bin
	go run main.go

clean:
	rm -f {protoset,message}.bin tutorial/addressbook.pb.go
