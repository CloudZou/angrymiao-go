package conf

import (
	"angrymiao-go/punk/conf"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/http/blademaster/middleware/verify"
	"angrymiao-go/punk/net/rpc/warden"
	"angrymiao-go/punk/net/trace"
	"angrymiao-go/punk/queue/databus"
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
)

// Conf global variable.
var (
	Conf     = &Config{}
	client   *conf.Client
	confPath string
)

// Config struct of conf.
type Config struct {
	// log
	Xlog *log.Config
	//Tracer *conf.Tracer
	Tracer *trace.Config
	// VerifyConfig
	VerifyConfig *verify.Config
	// BM
	BM *bm.ServerConfig
	// http client
	HTTPClient *bm.ClientConfig
	// grpc server
	WardenServer *warden.ServerConfig
	// IdentifyConfig
	Identify *IdentifyConfig
	// DataBus config
	DataBus *DataBus
}

// IdentifyConfig identify config
type IdentifyConfig struct {
	AuthHost string
	// LoginLogConsumerSize goroutine size
	LoginLogConsumerSize int
	// LoginCacheExpires login check cache expires
	LoginCacheExpires int32
	// IntranetCIDR
	IntranetCIDR []string
}

// DataBus data bus config
type DataBus struct {
	UserLog *databus.Config
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
			log.Info("config event")
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

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init int config
func Init() error {
	if confPath != "" {
		return local()
	}
	return remote()
}
