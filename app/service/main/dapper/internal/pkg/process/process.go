package process

import (
	"angrymiao-go/app/service/main/dapper/internal/model"
	"context"
)

// Processer .
type Processer interface {
	Process(ctx context.Context, protoSpan *model.ProtoSpan) error
}

// MockProcess MockProcess
type MockProcess func(ctx context.Context, protoSpan *model.ProtoSpan) error

// Process implement Processer
func (m MockProcess) Process(ctx context.Context, protoSpan *model.ProtoSpan) error {
	return m(ctx, protoSpan)
}
