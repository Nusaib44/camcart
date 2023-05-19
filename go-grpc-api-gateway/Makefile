proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.
	# # protoc pkg/**/pb/*.proto --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false
	# protoc pkg/**/pb/*.proto --go-grpc_out=.
server:
	go run cmd/main.go
