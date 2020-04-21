package api

import (
	"context"
	"fmt"

	"angrymiao-go/punk/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "user.service.v1"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (UserInfoClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewUserInfoClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
