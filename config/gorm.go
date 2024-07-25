package config

import (
	"github.com/thoriqaufar/liquipedia-valorant-api/entity"
	"github.com/thoriqaufar/liquipedia-valorant-api/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(localhost:3306)/liquipedia_valorant_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	helper.PanicIfError(err)

	err = db.AutoMigrate(&entity.Team{}, &entity.Player{})
	helper.PanicIfError(err)

	DB = db
}
