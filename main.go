package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"grpc_demon/model"
	pb "grpc_demon/protos"
	"grpc_demon/server"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost", "ip адрес")
	port = flag.Int("port", 8080, "порт сервера")
)

func main() {
	flag.Parse()
	address := fmt.Sprintf("%v:%v", *addr, *port)

	// Загружаем спиоск метрик из конфига
	// Формируем список метрик в MetricsList
	metricsList, err := model.MetricsInit("config/metrics.yaml")
	if err != nil {
		log.Fatalf("Ошибка инициализации метрик: %v", err)
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStreamingServiceServer(s, server.NewGrpc(metricsList))

	log.Printf("Сервер запущен по даресу %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
