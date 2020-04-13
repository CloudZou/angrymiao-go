package service


import (
	"sync"
	"time"

	"angrymiao-go/app/infra/canal/conf"
	"angrymiao-go/punk/log"

	"github.com/juju/errors"
	"github.com/siddontang/go-mysql/client"
	"github.com/siddontang/go-mysql/mysql"
)

type dbMasterInfo struct {
	c *conf.MasterInfoConfig

	addr    string
	binName string
	binPos  uint32

	l sync.RWMutex

	lastSaveTime time.Time
}

func newDBMasterInfo(addr string, c *conf.MasterInfoConfig) (*dbMasterInfo, error) {
	m := &dbMasterInfo{c: c, addr: addr}
	conn, err := client.Connect(c.Addr, c.User, c.Password, c.DBName)
	if err != nil {
		log.Error("db master info client error(%v)", err)
		return nil, errors.Trace(err)
	}
	defer conn.Close()
	if m.c.Timeout > 0 {
		conn.SetDeadline(time.Now().Add(m.c.Timeout * time.Second))
	}
	r, err := conn.Execute("SELECT addr,bin_name,bin_pos FROM master_info WHERE addr=?", addr)
	if err != nil {
		log.Error("new db load master.info error(%v)", err)
		return nil, errors.Trace(err)
	}
	if r.RowNumber() == 0 {
		if m.c.Timeout > 0 {
			conn.SetDeadline(time.Now().Add(m.c.Timeout * time.Second))
		}
		if _, err = conn.Execute("INSERT INTO master_info (addr,bin_name,bin_pos) VALUE (?,'',0)", addr); err != nil {
			log.Error("insert master.info error(%v)", err)
			return nil, errors.Trace(err)
		}
	} else {
		m.addr, _ = r.GetStringByName(0, "addr")
		m.binName, _ = r.GetStringByName(0, "bin_name")
		bpos, _ := r.GetIntByName(0, "bin_pos")
		m.binPos = uint32(bpos)
	}
	return m, nil
}

func (m *dbMasterInfo) Save(pos mysql.Position, force bool) error {
	n := time.Now()
	if !force && n.Sub(m.lastSaveTime) < 2*time.Second {
		return nil
	}
	conn, err := client.Connect(m.c.Addr, m.c.User, m.c.Password, m.c.DBName)
	if err != nil {
		log.Error("db master info client error(%v)", err)
		return errors.Trace(err)
	}
	defer conn.Close()

	if m.c.Timeout > 0 {
		conn.SetDeadline(time.Now().Add(m.c.Timeout * time.Second))
	}
	if _, err = conn.Execute("UPDATE master_info SET bin_name=?,bin_pos=? WHERE addr=?", pos.Name, pos.Pos, m.addr); err != nil {
		log.Error("db save master info error(%v)", err)
		return errors.Trace(err)
	}
	m.lastSaveTime = n
	return nil
}

func (m *dbMasterInfo) Pos() mysql.Position {
	var pos mysql.Position
	m.l.RLock()
	pos.Name = m.binName
	pos.Pos = m.binPos
	m.l.RUnlock()
	return pos
}
