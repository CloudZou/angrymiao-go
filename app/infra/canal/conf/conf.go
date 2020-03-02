package conf

import (
	"errors"
	"flag"
	"github.com/BurntSushi/toml"
	pconf "github.com/CloudZou/punk/pkg/conf"
	"github.com/CloudZou/punk/pkg/log"
	"github.com/CloudZou/punk/pkg/naming/discovery"
	bm "github.com/CloudZou/punk/pkg/net/http/blademaster"
	"github.com/CloudZou/punk/pkg/net/netutil/breaker"
	"time"
)

var (
	canalPath, confPath string
	// ConfClient get config client
	ConfClient *pconf.Client
	// Conf canal config variable
	Conf = &Config{}
)

// Config canal config struct
type Config struct {
	Monitor *Monitor
	// xlog
	Log *log.Config
	// http client
	HTTPClient *bm.ClientConfig
	// http server
	BM *bm.ServerConfig
	// master info
	MasterInfo *MasterInfoConfig
	// discovery
	Discovery *discovery.Config
	// db
	DB *DbConfig
}

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

// Monitor wechat monitor
type Monitor struct {
	User   string
	Token  string
	Secret string
}

// MasterInfoConfig save pos of binlog in file or db
type MasterInfoConfig struct {
	Addr     string        `toml:"addr"`
	DBName   string        `toml:"dbName"`
	User     string        `toml:"user"`
	Password string        `toml:"password"`
	Timeout  time.Duration `toml:"timeout"`
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
	flag.StringVar(&canalPath, "canal", "", "canal instance path")
}

//Init int config
func Init() (err error) {
	if confPath != "" {
		_, err = toml.DecodeFile(confPath, &Conf)
		return
	}
	err = remote()
	return
}

func remote() (err error) {
	if ConfClient, err = pconf.New(); err != nil {
		return
	}
	ConfClient.WatchAll()
	err = LoadCanal()
	return
}

// LoadCanal canal config
func LoadCanal() (err error) {
	var (
		s       string
		ok      bool
		tmpConf *Config
	)
	if s, ok = ConfClient.Value("canal.toml"); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}
