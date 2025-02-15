gen: 
	protoc --proto_path=./protos --go_out=./ protos/*.proto

clean:
	rm pb/*.go

run:
	go run main.go