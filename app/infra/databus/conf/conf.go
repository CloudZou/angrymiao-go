package conf

import (
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/CloudZou/punk/pkg/log"
	pconf "github.com/CloudZou/punk/pkg/conf"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
	"github.com/CloudZou/punk/pkg/net/netutil/breaker"
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
	// Conf global config variable
	Conf     = &Config{}
	confPath string
	client   *pconf.Client
)

// Config databus config struct
type Config struct {
	// base
	Addr     string
	Clusters map[string]*Kafka
	// Log
	Log *log.Config
	// http
	HTTPServer *bm.ServerConfig
	// mysql
	MySQL *DbConfig
}

// Kafka contains cluster, brokers, sync.
type Kafka struct {
	Cluster string
	Brokers []string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

//Init int config
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
	if client, err = pconf.New(); err != nil {
		return
	}
	if err = load(); err != nil {
		return
	}
	go func() {
		for range client.Event() {
			log.Info("config reload")
			if load() != nil {
				log.Error("config reload error (%v)", err)
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
	if s, ok = client.Toml(); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}
