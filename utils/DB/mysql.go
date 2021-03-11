package DB

import (
	"ecust-trading/strategy"
	"github.com/jinzhu/gorm"
)

var (
	DB *dbModel
)

type dbModel struct {
	DB        *gorm.DB
	tickChan  chan interface{}
	tradeChan chan interface{}
}

func initDB() {
	DB = new(dbModel)
}

func GetDB() *dbModel {
	if DB == nil {
		initDB()
	}
	return DB
}

func (d *dbModel) SaveTradeData(msg *strategy.TradeMsg) (err error) {
	return
}
