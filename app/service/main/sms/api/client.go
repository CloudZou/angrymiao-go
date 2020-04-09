package v1

import (
	"context"
	"angrymiao-go/punk/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID discovery id.
const AppID = "sms.service"

// NewClient .
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (SmsClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "discovery://default/"+AppID)
	if err != nil {
		return nil, err
	}
	return NewSmsClient(conn), nil
}
