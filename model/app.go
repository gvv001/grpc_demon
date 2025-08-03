// общая модель данных
package model

import (
	"grpc_demon/storage"
)

type MetricCollector interface {
	CollectData()
	GetCachePointer() *storage.Cache
	GetName() string
}
