package conf

import (
	"angrymiao-go/punk/conf"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/http/blademaster/middleware/verify"
	"angrymiao-go/punk/net/rpc"
	"angrymiao-go/punk/net/rpc/warden"
	"angrymiao-go/punk/net/trace"
	"angrymiao-go/punk/queue/databus"
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
)

type cfg struct {
	RankGroupMaxPs int
}

var (
	Conf     = &Config{}
	client   *conf.Client
	confPath string
)

// Config struct of conf.
type Config struct {
	App        *bm.App
	Xlog       *log.Config
	Verify     *verify.Config
	HTTPServer *bm.ServerConfig
	HTTPClient *bm.ClientConfig
	RPCServer  *rpc.ServerConfig
	GRPC       *warden.ServerConfig
	Tracer     *trace.Config
	Databus    *databus.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init .
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
	go func() {
		for range client.Event() {
			log.Info("config reload")
			if load() != nil {
				log.Error("config reload err")
			}
		}
	}()
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
