generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  protos/grpc.proto
	
# исходники для палтформы win amd64
	
	GOOS=windows GOARCH=386 go build -o bin/server-win.exe main.go
	GOOS=windows GOARCH=386 go build -o ./bin/client-win.exe client/client.go
	
# исходники для палтформа win x86

	GOOS=linux GOARCH=386 go build -o bin/grpc_server_linux.exe main.go
	GOOS=linux GOARCH=386 go build -o ./bin/grpc_client_linux.exe client/client.go

	sudo docker rmi grpc_demon
	sudo docker build -t grpc_demon .
	sudo docker run -p 8080:8080 -it --rm grpc_demon
	
