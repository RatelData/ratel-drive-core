package models

import (
	"github.com/RatelData/ratel-drive-core/common/database"
	"gorm.io/gorm"
)

type Secret struct {
	gorm.Model
	PrivateKey []byte
}

func FindSecret(conds ...interface{}) (Secret, error) {
	var secret Secret
	err := database.GetDB().First(&secret, conds...).Error
	return secret, err
}

func (secret *Secret) Save() error {
	return database.GetDB().Save(secret).Error
}
