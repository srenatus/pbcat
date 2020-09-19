protoset.bin tutorial/addressbook.pb.go: tutorial/addressbook.proto
	protoc --descriptor_set_out="$@" --include_imports --go_out=paths=source_relative:. "$<" 

message.bin: generate/generate.go tutorial/addressbook.pb.go
	go run "$<" > "$@"

example: protoset.bin message.bin
	go run main.go

clean:
	rm -f {protoset,message}.bin tutorial/addressbook.pb.go