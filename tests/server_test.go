package grpc_demon

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"grpc_demon/model"
	"grpc_demon/storage"
)

// тестируем загрузку конфига метрик
func TestLoadConfig(t *testing.T) {
	if _, err := model.LoadConfig("../config/metrics.yaml"); err != nil {
		log.Fatal(err)
	}
}

// тестируем инициализацию метрик
func TestMetricsInit(t *testing.T) {
	if _, err := model.MetricsInit("../config/metrics.yaml"); err != nil {
		log.Fatal(err)
	}
}

// тестируем работу метрики
func TestMetricsGetValue(t *testing.T) {

	metric := &model.Metric{
		Name:        "Load avg. (%)",
		Cache:       &storage.Cache{},
		Cmd:         "cat",
		CmdParams:   "/proc/loadavg", //nolint:
		ParseParams: "{print $1}",    //nolint:
		IsActive:    true,
	}

	stringValue, err := metric.GetValue()
	if err != nil {
		log.Fatal(err)
	}

	floatValue, err := strconv.ParseFloat(stringValue, 64)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(floatValue)
}
