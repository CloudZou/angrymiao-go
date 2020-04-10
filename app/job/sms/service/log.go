package service

import (
	smsmdl "angrymiao-go/app/service/main/sms/model"
	"angrymiao-go/punk/log"
	"angrymiao-go/punk/queue/databus/report"
	"strings"
	"time"
)

const (
	_reportType = 111
)

func (s *Service) sendUserActionLog(l *smsmdl.ModelUserActionLog) {
	if l.Mobile == "" {
		log.Warn("sendUserActionLog mobile is empty, log(%+v)", l)
		return
	}
	for _, mobile := range strings.Split(l.Mobile, ",") {
		r := &report.UserInfo{
			Business: _reportType,
			Ctime:    time.Unix(l.Ts, 0),
			Index:    []interface{}{mobile},
			Content: map[string]interface{}{
				"msgid":    l.MsgID,
				"content":  l.Content,
				"status":   l.Status,
				"desc":     l.Desc,
				"type":     l.Type,
				"action":   l.Action,
			},
		}
		log.Info("sendUserActionLog(%+v)", r)
		report.User(r)
	}
}
