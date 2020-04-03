package app

import (
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"github.com/astaxie/beego/validation"
)

// MarkErrors logs error logs
func MarkErrors(c *bm.Context, errors []*validation.Error) {
	for _, err := range errors {
		log.Infov(c, log.KV(err.Key, err.Message))
	}

	return
}
