proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

run:
	go run cmd/main.go