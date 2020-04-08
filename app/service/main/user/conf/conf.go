package conf

import (
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
	"angrymiao-go/punk/conf"
	"angrymiao-go/punk/gredis"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/netutil/breaker"
	"angrymiao-go/punk/net/rpc"
	"angrymiao-go/punk/net/rpc/warden"
	"angrymiao-go/punk/net/trace"
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

// global var
var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// Config config set
type Config struct {
	// rpc server
	RPCServer  *rpc.ServerConfig
	GRPCServer *warden.ServerConfig
	HTTPServer *bm.ServerConfig

	Db             DbConfig
	Redis          gredis.RedisConfig
	UdpTraceConfig *trace.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf
func Init() error {
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
