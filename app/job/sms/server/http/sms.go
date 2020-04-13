package http

import (
	pb "angrymiao-go/app/service/main/sms/api"
	bm "angrymiao-go/punk/net/http/blademaster"
)

func send(ctx *bm.Context) {
	req := new(pb.SendReq)
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(smsSvc.Send(ctx, req))
}
