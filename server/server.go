package server

import (
	"fmt"
	"log"
	"strings"
	"time"

	"grpc_demon/lib"
	"grpc_demon/model"
	pb "grpc_demon/protos"
)

var responseMsg string

type server struct {
	pb.UnimplementedStreamingServiceServer
	MetricsList *model.MetricsList
}

func NewGrpc(metricsList *model.MetricsList) *server {
	return &server{MetricsList: metricsList}
}

func (s server) GetDataFromStream(req *pb.RequestParams, srv pb.StreamingService_GetDataFromStreamServer) error {
	// period - M(сек) - интервал для усреднения информации
	period := req.GetPeriod()
	log.Printf("Новый gRPC клиент: частота запросов: %v сек | Среднее значение метрик за %v сек", req.GetFrequency(), period)

	// добавляем в sb названия всех метрик и значение из кэша усреднённое за М секунд
	// результат отправляем клиенту в виде одной строки
	var sb strings.Builder

	for {

		// частота ответа сервера N(сек)
		time.Sleep(time.Second * time.Duration(req.Frequency))

		sb.WriteString("\nOS - " + s.MetricsList.GetOS())

		for _, metric := range s.MetricsList.Items {
			name := metric.GetName()
			avg, err := lib.GetMetricAvg(metric.GetCachePointer(), int(period))

			if err != nil {
				responseMsg = fmt.Sprintf("\n %v\t = %v", name, err.Error())
			} else {
				responseMsg = fmt.Sprintf("\n %v\t = %v", name, avg)
			}

			sb.WriteString(responseMsg)
		}

		resp := pb.ResponseData{
			Metrics: sb.String(),
		}

		sb.Reset()

		if err := srv.Send(&resp); err != nil {
			log.Printf("Error srv.Send(): %v", err)
			return err
		}
	}
}
