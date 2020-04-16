package service

import (
	"context"
	"flag"
	"path/filepath"
	"testing"
	"time"

	"angrymiao-go/app/service/main/sms/conf"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	srv *Service
	ctx = context.Background()
)

func init() {
	dir, _ := filepath.Abs("../cmd/manager-admin-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	srv = New(conf.Conf)
	time.Sleep(time.Second)
}

func WithService(f func(s *Service)) func() {
	return func() {
		// Reset(func() { CleanCache() })
		f(srv)
	}
}

func TestService_Close(t *testing.T) {
	Convey("TestService_Close", t, WithService(func(s *Service) {
		s.Close()
	}))
}

func TestService_Wait(t *testing.T) {
	Convey("TestService_Wait", t, WithService(func(s *Service) {
		s.Wait()
	}))
}
