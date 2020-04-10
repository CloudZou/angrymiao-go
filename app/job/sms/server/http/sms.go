package http

import bm "angrymiao-go/punk/net/http/blademaster"
import pb "angrymiao-go/app/job/sms/api"


func send(ctx *bm.Context) {
	req := new(pb.SendReq)
	if err := ctx.Bind(req); err != nil {
		return
	}
	ctx.JSON(smsSvc.Send(ctx, req))
}
