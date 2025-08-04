//go:build windows

package model

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"grpc_demon/lib"
	"grpc_demon/storage"

	"gopkg.in/yaml.v3"
)

// Структура для хранения метрик Windows
type MetricsList struct {
	os    string
	Items []MetricCollector
}

func MetricsInit(fileName string) (*MetricsList, error) {
	metricsList := &MetricsList{}
	metricsList.os = "Windows"

	metricsConfig, err := LoadConfig(fileName)
	if err != nil {
		return nil, err
	}

	for _, metric := range metricsConfig.Items {
		if metric.IsActive {
			metric.Cache = &storage.Cache{}
			metricsList.Add(&metric)
		}
	}

	return metricsList, nil
}

// Добавить метрику в список структуры Metrics
func (m *MetricsList) Add(metric MetricCollector) {
	go metric.CollectData()
	m.Items = append(m.Items, metric)
}

func (m MetricsList) GetOS() string {
	return m.os
}

// Структура метрики Windows

type Metric struct {
	Name     string         `yaml:"name"`
	Cache    *storage.Cache `yaml:"-"`
	Cmd      string         `yaml:"cmd"`
	IsActive bool           `yaml:"isActive"`
}

// Запись данных метрики в Кэш
func (m *Metric) CollectData() {
	for {
		// запись метрик в кэш 1 в скунду
		time.Sleep(time.Second * 1)

		stringValue, err := m.GetValue()
		if err != nil {
			log.Printf("Ошибка GetValue(): %v\n", err)
			continue
		}

		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			log.Printf("Ошибка ParseFloat(): %v\n", err)
			continue
		}

		m.Cache.Add(float32(floatValue))
	}
}

// Вернуть указатель на кэш метрики
func (m Metric) GetCachePointer() *storage.Cache {
	return m.Cache
}

// Вернуть название метрки
func (m Metric) GetName() string {
	return m.Name
}

// Получить значение метрки из shell
func (m Metric) GetValue() (string, error) {
	time.Sleep(time.Second * 1)

	result, err := lib.ExecShellCommand(m.Cmd)
	if err != nil {
		log.Printf("Ошибка парсинга: %v", err)
		return "", err
	}

	// удаляем символы \n \r и прочее
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "%", "")
	
	return result, nil
}

// Структура для хранения загруженных метрик из конфига
type MetricsConfig struct {
	Items []Metric `yaml:"WindowsMetrics"`
}

// Загрузка конфигурации метрик из раздела WindowsMetrics
func LoadConfig(fileName string) (*MetricsConfig, error) {
	var ml MetricsConfig

	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &ml)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	return &ml, nil
}
