package DB

import (
	"github.com/jinzhu/gorm"
)

var (
	DB *dbModel
)

type dbModel struct {
	DB        *gorm.DB
	tradeChan chan interface{}
	dayChan   chan *dayCandle
}

type dayCandle struct {
	Symbol string
	Open   float64
	Close  float64
	Low    float64
	High   float64
	TS     int64
}

func initDB() {
	conn, err := gorm.Open("mysql", buildDSN())
	if err != nil {
		panic(err)
	}
	DB = &dbModel{
		DB:      conn,
		dayChan: make(chan *dayCandle, 1024),
	}
}

func buildDSN() string {
	return "aaa"
}

func GetDB() *dbModel {
	if DB == nil {
		initDB()
	}
	return DB
}
