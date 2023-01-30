package models

import "github.com/RatelData/ratel-drive-core/common/database"

func AutoMigrate() {
	db := database.GetDB()
	db.AutoMigrate(&Secret{})
}
