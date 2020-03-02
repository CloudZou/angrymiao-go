package conf

import (
	"angrymiao-go/app/infra/canal/infoc"
	"fmt"
	pconf "github.com/CloudZou/punk/pkg/conf"
	"github.com/CloudZou/punk/pkg/log"
	"github.com/CloudZou/punk/pkg/queue/databus"
	xtime "github.com/CloudZou/punk/pkg/time"
	"io/ioutil"
	"net"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
)

var (
	// config change event
	event = make(chan *InsConf, 1)
)

// Addition addition attrbute of canal.
type Addition struct {
	PrimaryKey []string `toml:"primarykey"` // kafka msg key
	OmitField  []string `toml:"omitfield"`  // field will be ignored in table
}

// CTable canal table.
type CTable struct {
	PrimaryKey []string `toml:"primarykey"` // kafka msg key
	OmitField  []string `toml:"omitfield"`  // field will be ignored in table
	OmitAction []string `toml:"omitaction"` // action will be ignored in table
	Name       string   `toml:"name"`       // table name support regular expression
	Tables     []string
}

// Database represent mysql db
type Database struct {
	Schema   string          `toml:"schema"`
	Databus  *databus.Config `toml:"databus"`
	Infoc    *infoc.Config   `toml:"infoc"`
	CTables  []*CTable       `toml:"table"`
	TableMap map[string]*Addition
}

// CheckTable check database tables.
func (db *Database) CheckTable(addr, user, passwd string) (err error) {
	var (
		conn  *client.Conn
		res   *mysql.Result
		regex *regexp.Regexp
		table string
	)
	db.TableMap = make(map[string]*Addition)
	if conn, err = client.Connect(addr, user, passwd, db.Schema); err != nil {
		return
	}
	defer conn.Close()
	if res, err = conn.Execute(fmt.Sprintf("SHOW TABLES FROM `%s`", db.Schema)); err != nil {
		log.Error("conn.Execute() error(%v)", err)
		return
	}
	for _, ctable := range db.CTables {
		if regex, err = regexp.Compile(ctable.Name); err != nil {
			log.Error("regexp.Compile(%s) error(%v)", ctable.Name, err)
			return
		}
		for _, value := range res.Values {
			table = fmt.Sprintf("%s", value[0])
			if regex.MatchString(table) {
				db.TableMap[table] = &Addition{
					PrimaryKey: ctable.PrimaryKey,
					OmitField:  ctable.OmitField,
				}
				ctable.Tables = append(ctable.Tables, table)
			}
		}
		if len(ctable.Tables) == 0 {
			return fmt.Errorf("addr(%s) db(%s) subscribles nothing,table(%s) is empty", addr, db.Schema, ctable.Name)
		}
	}
	return
}

// InsConf instance config
type InsConf struct {
	*canal.Config

	MonitorPeriod xtime.Duration    `toml:"monitor_period"`
	MonitorOff    bool              `toml:"monitor_off"`
	Databases     []*Database       `toml:"db"`
	MasterInfo    *MasterInfoConfig `toml:"masterinfo"`
}

// CanalConfig config struct
type CanalConfig struct {
	Instances []*InsConf `toml:"instance"`
}

func newInsConf(fn, fc string) (c *InsConf, err error) {
	var ic struct {
		InsConf *InsConf `toml:"instance"`
	}
	ipPort := strings.TrimSuffix(fn, ".toml")
	if _, _, err = net.SplitHostPort(ipPort); err != nil {
		return
	}
	if _, err = toml.Decode(fc, &ic); err != nil {
		return
	}
	if ic.InsConf == nil {
		err = fmt.Errorf("file(%s) cannot decode toml", fn)
		return
	}
	if ic.InsConf.Addr != ipPort {
		err = fmt.Errorf("file(%s) name not equal addr(%s)", fn, ic.InsConf.Addr)
		return
	}
	if ic.InsConf.MasterInfo == nil {
		ic.InsConf.MasterInfo = Conf.MasterInfo
	}
	return ic.InsConf, nil
}

// LoadCanalConf load canal config.
func LoadCanalConf() (c *CanalConfig, err error) {
	var (
		result []*pconf.Value
		ok     bool
	)
	if canalPath != "" {
		result, err = localCanal()
	} else {
		result, ok = ConfClient.Configs()
		if !ok {
			panic("no canal-config")
		}
	}
	c = new(CanalConfig)
	im := map[string]struct{}{}
	for _, ns := range result {
		if ns.Name == "canal.toml" || ns.Name == "common.toml" {
			continue
		}

		var ic *InsConf
		if !strings.HasSuffix(ns.Name, ".toml") {
			err = fmt.Errorf("file(%s) name is not a toml", ns.Name)
			continue
		}
		if ic, err = newInsConf(ns.Name, ns.Config); err != nil {
			err = fmt.Errorf("file(%s) decode error(%v)", ns.Name, err)
			return
		}
		if _, ok := im[ic.Addr]; ok {
			err = fmt.Errorf("file(%s) repeat with other toml", ns.Name)
			return
		}
		im[ic.Addr] = struct{}{}
		c.Instances = append(c.Instances, ic)

	}
	if canalPath == "" {
		go func() {
			for name := range ConfClient.Event() {
				log.Info("config(%s) reload", name)
				reloadConfig(name)
			}
		}()
	}
	return
}

func localCanal() (vs []*pconf.Value, ok error) {
	fs, err := ioutil.ReadDir(canalPath)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		if !strings.HasSuffix(f.Name(), ".toml") {
			continue
		}
		ct, err := ioutil.ReadFile(canalPath + f.Name())
		if err != nil {
			continue
		}
		vs = append(vs, &pconf.Value{
			Name:   f.Name(),
			Config: string(ct),
		})
	}
	return
}

// Event returns config change event chan,
func Event() chan *InsConf {
	return event
}

func reloadConfig(name string) {
	var (
		cf string
		ok bool
	)
	if name == "canal.toml" || name == "common.toml" {
		LoadCanal()
		return
	}
	if !strings.HasSuffix(name, ".toml") {
		return
	}
	if cf, ok = ConfClient.Value(name); !ok {
		// TODO(felix): auto reload? or restart hard?
		return
	}

	ic, err := newInsConf(name, cf)
	if err != nil {
		return
	}
	event <- ic

}
