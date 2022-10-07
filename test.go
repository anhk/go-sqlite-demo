package main

import (
	"github.com/cihub/seelog"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type seelogger struct{}

func (l *seelogger) Write(message []byte) (int, error) {
	seelog.Infof("%s", message)
	return len(message), nil
}

func main() {
	engine, err := xorm.NewEngine("sqlite3", "./example.db?cache=shared")
	if err != nil {
		panic(err)
	}
	engine.SetLogger(log.NewSimpleLogger2(&seelogger{}, "[xorm]", 0))
	seelog.Debugf("###")
	// fmt.Println(db.Ping())
	if err := engine.Ping(); err != nil {
		panic(err)
	}
}
