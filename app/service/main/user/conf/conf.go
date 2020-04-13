package conf

import (
	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/conf"
	"angrymiao-go/punk/database/orm"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/rpc"
	"angrymiao-go/punk/net/rpc/warden"
	"angrymiao-go/punk/net/trace"
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
)

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
	BM *bm.ServerConfig

	SMSClient *warden.ClientConfig

	Db             *orm.Config
	Redis          *redis.Config
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
