package api

import (
	"context"
	"fmt"

	"github.com/CloudZou/punk/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "user.service"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (GreeterClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewGreeterClient(cc), nil
}

// 生成 gRPC 代码
//go:generate punk tool protoc --grpc --bm api.proto
