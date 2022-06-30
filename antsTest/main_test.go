package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/panjf2000/ants/v2"
	"xorm.io/xorm"
)

func myFunc(engine *xorm.Engine) interface{} {
	a := make(map[string]interface{}, 0)
	engine.Table("consult").Get(&a)
	log.Println(a)
	return a
}

func TestAnts(t *testing.T) {
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i.(*xorm.Engine))
	})
	for i := 0; i < 10000; i++ {
		p.Invoke(getEngine())
	}
}

func getEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("dbadmin:423B8BL4s2394T7E@(jt-dev.csy2ig5dsj91.rds.cn-north-1.amazonaws.com.cn:63306)/jt_dev?charset=utf8&interpolateParams=true&parseTime=true&loc=Local"))
	if err != nil {
		log.Println(err)
		return nil
	}
	engine.ShowSQL(true)
	engine.DatabaseTZ, _ = time.LoadLocation("Asia/Shanghai")
	return engine
}
