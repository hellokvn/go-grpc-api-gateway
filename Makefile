proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go