package lib

import (
	"errors"
	"math"

	"grpc_demon/storage"
)

const (
	ErrOutOfRange = "размер кэша (Len) меньше интервала запроса (M). Ожидайте прогрева кэша"
)

// считаем информацию, усредненную за последние M (seconds) секунд.
func GetMetricAvg(c *storage.Cache, period int) (float64, error) {
	if c.Len < period {
		return 0, errors.New(ErrOutOfRange)
	}

	var sum float64

	node := c.Head
	for i := 0; i < period; i++ {
		sum += float64(node.Data)
		node = node.Prev
	}

	avg := sum / float64(period)

	return math.Round(avg*100) / 100, nil
}
