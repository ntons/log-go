all: proto
	go build

proto: *.proto
	protoc --go_out=plugins=grpc:. $^


clean:
	rm -rf *.pb.go
