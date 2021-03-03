package DB

import (
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

type dbModel struct {
	DB *gorm.DB
	tickChan chan interface{}
	tradeChan  chan interface{}
}

func GetDB() {

}


