package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-module/carbon/v2"
	"xorm.io/xorm"
)

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

// Consult 通用咨询，图文咨询
type Consult struct {
	ConsultId  []byte    `xorm:"not null comment('咨询id') BINARY(20)"`
	Genre      string    `xorm:"not null comment('咨询类型 1 通用咨询 2 图文咨询') VARCHAR(1)"`
	UserId     []byte    `xorm:"not null comment('用户id') BINARY(20)"`
	Source     string    `xorm:"not null comment('来源 1 Mac兑换 2 赠送 3 核销') VARCHAR(1)"`
	WriteId    []byte    `xorm:"comment('核销所属id') BINARY(20)"`
	Mac        string    `xorm:"comment('mac码') VARCHAR(50)"`
	Count      int       `xorm:"not null comment('次数') INT"`
	UseCount   int       `xorm:"not null comment('使用次数') INT"`
	MaturityAt time.Time `xorm:"not null comment('到期时间') TIMESTAMP"`
	Status     string    `xorm:"not null comment('状态 1 正常 2 冻结 3 过期 4 撤销') VARCHAR(1)"`
	CreatedAt  time.Time `xorm:"not null comment('创建时间') TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"not null comment('更新时间') TIMESTAMP"`
	DeletedAt  time.Time `xorm:"null comment('删除时间') TIMESTAMP"`
}

func (c *Consult) TableName() string {
	return "consult"
}

func TestOne(t *testing.T) {
	engine := getEngine()
	var c Consult
	b, err := engine.Where("created_at > ? ", time.Now().Unix()).Get(&c)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println(b, c)
}
func TestTime(t *testing.T) {
	now := carbon.Now(carbon.Shanghai)
	log.Println("now", now.String(), now.Timestamp())
	now2 := now.Carbon2Time()
	log.Println("now2", now2.String(), now2.Unix())
	now3 := carbon.Time2Carbon(now2)
	log.Println("now3", now3.String(), now3.Timestamp())
	now4 := now3.Carbon2Time()
	log.Println("now4", now4.String(), now4.Unix())
	now5 := carbon.Time2Carbon(now4)
	log.Println("now5", now5.String(), now5.Timestamp())
	now6 := now5.SetTimezone(carbon.Shanghai)
	log.Println("now6", now6.String(), now6.Timestamp())
}
