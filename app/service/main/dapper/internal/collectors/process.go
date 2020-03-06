package collectors

import (
	"angrymiao-go/app/service/main/dapper/internal/model"
)

// Processer span processer
type Processer interface {
	Process(span *model.Span) error
}

// ProcessFunc implement Processer
type ProcessFunc func(*model.Span) error

// Process implement Processer
func (p ProcessFunc) Process(span *model.Span) error { return p(span) }
