package collector

import (
	"angrymiao-go/app/service/main/dapper/internal/pkg/process"
)

// Collecter collect span from different source
type Collector interface {
	Start() error
	RegisterProcess(p process.Processer)
	Close() error
}
