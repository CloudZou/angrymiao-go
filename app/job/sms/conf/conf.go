package conf

import (
	"angrymiao-go/punk/conf"
	"angrymiao-go/punk/log"
	bm "angrymiao-go/punk/net/http/blademaster"
	"angrymiao-go/punk/net/rpc/warden"
	"angrymiao-go/punk/net/trace"
	"angrymiao-go/punk/queue/databus"
	xtime "angrymiao-go/punk/time"
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
	Log        *log.Config
	Tracer     *trace.Config
	Databus    *databus.Config
	SmsGRPC    *warden.ClientConfig
	HTTPClient *bm.ClientConfig
	HTTPServer *bm.ServerConfig
	UserReport *databus.Config
	AliyunSmsConfig *AliyunSmsConfig
	Wechat     *wechat
	Sms        *sms
}

type sms struct {
	// CallbackProc 处理回执的并发数
	CallbackProc int
	// SingleSendGoroutines 单发短信的并发数
	SingleSendProc int
	// BatchSendGoroutines 批量发送短信的并发数
	BatchSendProc int
	// MonitorProcDuration 定期监控databus有没有消费
	MonitorProcDuration xtime.Duration
	// Blacklist 黑名单手机号，用于压测
	Blacklist []string
}

type wechat struct {
	Token    string
	Secret   string
	Username string
}

type AliyunSmsConfig struct {
	RegionID string
	AccessKeyID string
	AccessKeySecret string
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
	err = load()
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
