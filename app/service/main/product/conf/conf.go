package conf

import (
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/CloudZou/punk/pkg/conf"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
	"github.com/CloudZou/punk/pkg/net/netutil/breaker"
	"github.com/CloudZou/punk/pkg/net/rpc"
	"github.com/CloudZou/punk/pkg/net/rpc/warden"
	"github.com/CloudZou/punk/pkg/net/trace"
	"time"
)

type DbConfig struct {
	DSN          string          // write data source name.
	ReadDSN      []string        // read data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  time.Duration   // connect max life time.
	QueryTimeout time.Duration   // query sql timeout
	ExecTimeout  time.Duration   // execute sql timeout
	TranTimeout  time.Duration   // transaction sql timeout
	Breaker      *breaker.Config // breaker
}

var (
	Conf     = &Config{}
	client   *conf.Client
	confPath string
)

type Config struct {
	// http
	BM        *bm.ServerConfig
	RPCClient *rpc.ClientConfig

	RPCServer *warden.ServerConfig

	DB DbConfig

	UserGrpcClient *warden.ClientConfig

	UdpTraceConfig *trace.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init create config instance.
func Init() (err error) {
	if confPath != "" {
		return local()
	}
	return remote()
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func remote() (err error) {
	if client, err = conf.New(); err != nil {
		return
	}
	if err = load(); err != nil {
		return
	}
	return
}

func load() (err error) {
	var (
		s       string
		ok      bool
		tmpConf *Config
	)
	if s, ok = client.Toml2(); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}
