package conf

import (
	"angrymiao-go/punk/cache/redis"
	"angrymiao-go/punk/database/orm"
	"errors"
	"flag"
	"angrymiao-go/punk/conf"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/http/blademaster/middleware/permit"
	"angrymiao-go/punk/net/rpc/warden"
	"angrymiao-go/punk/net/trace"
	xtime "angrymiao-go/punk/time"

	"github.com/BurntSushi/toml"
)

// Config .
type Config struct {
	Cfg          *cfg
	App          *bm.App
	ORM          *orm.Config
	Log          *log.Config
	Tracer       *trace.Config
	Redis     *redis.Config
	HTTPServer   *bm.ServerConfig
	HTTPClient   *bm.ClientConfig
	DsbClient    *bm.ClientConfig
	UnameTicker  xtime.Duration
	WardenServer *warden.ServerConfig
	Permit       *permit.Config
}

type cfg struct {
	RankGroupMaxPs int
}

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

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
