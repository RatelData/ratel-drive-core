package database

import (
	"github.com/RatelData/ratel-drive-core/common/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInst *gorm.DB

// Using this function to get a connection.
func GetDB() *gorm.DB {
	if dbInst == nil {
		initDB()
	}
	return dbInst
}

// Opening a database and save the reference to `Database` struct.
func initDB() {
	conf := util.GetAppConfig()

	db, err := gorm.Open(sqlite.Open(conf.DatabasePath), &gorm.Config{})
	if err != nil {
		util.GetLogger().Error("failed to connect database")
		return
	}

	dbInst = db
}
