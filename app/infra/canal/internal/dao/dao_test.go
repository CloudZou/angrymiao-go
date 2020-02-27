package dao

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/CloudZou/punk/pkg/conf/paladin"
)

var d *Dao
var ctx = context.Background()

func TestMain(m *testing.M) {
	flag.Set("conf", "../../test")
	flag.Set("f", "../../test/docker-compose.yaml")
	flag.Parse()

	var err error
	if err = paladin.Init(); err != nil {
		panic(err)
	}
	var cf func()
	if d, cf, err = newTestDao();err != nil {
		panic(err)
	}
	ret := m.Run()
	cf()

	os.Exit(ret)
}
