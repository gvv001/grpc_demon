package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	pb "grpc_demon/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr      = flag.String("addr", "localhost", "ip адрес")
	port      = flag.String("port", "8080", "порт")
	frequency = flag.Uint64("frequency", 1, "частота запросов (сек)") // N
	period    = flag.Uint64("period", 10, "Stat period (sec)")        // M
)

func main() {
	flag.Parse()

	address := fmt.Sprintf("%v:%v", *addr, *port)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("ошибка подключение к gRPC серверу: %v", err.Error())
	}

	defer conn.Close()

	// создаём клиента
	client := pb.NewStreamingServiceClient(conn)
	req := pb.RequestParams{Frequency: *frequency, Period: *period}
	stream, err := client.GetDataFromStream(context.Background(), &req)
	if err != nil {
		log.Printf("could not read: %v", err)
		return
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			continue
		} else if err == nil {
			// чистим окно терминала
			fmt.Print("\033[H\033[2J")
			value := fmt.Sprintf("%v", resp.Metrics)

			log.Println(value)
		}

		if err != nil {
			log.Fatalf("Error Response: %v", err)
		}

	}
}
