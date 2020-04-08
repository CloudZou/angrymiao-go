package collectors

import (
	"angrymiao-go/app/service/main/dapper/internal/model"
	protogen "angrymiao-go/punk/net/trace/proto"
	"testing"
)

func TestOperationNameProcess(t *testing.T) {
	p := NewOperationNameProcess()
	sp1, _ := model.FromProtoSpan(&model.ProtoSpan{
		ServiceName:   "http",
		OperationName: "http://www.example.com/echo",
		Tags:          []*protogen.Tag{&protogen.Tag{Key: "span.kind", Kind: protogen.Tag_STRING, Value: []byte("client")}},
	}, false)
	p.Process(sp1)
	if sp1.OperationName != "HTTP:UNKONWN" || sp1.ProtoSpan.OperationName != "HTTP:UNKONWN" {
		t.Errorf("expect operationName == , get %s %s", sp1.OperationName, sp1.ProtoSpan.OperationName)
	}
	if sp1.StringTag("http.url") != "http://www.example.com/echo" {
		t.Errorf("expect http.url be set")
	}
	if sp1.StringTag("_peer.sign") != "/echo" {
		t.Errorf("expect _peer.sign be set")
	}
}
