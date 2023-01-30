package secret

import (
	"crypto/ed25519"
	"crypto/rand"

	"github.com/RatelData/ratel-drive-core/app/models"
	"github.com/RatelData/ratel-drive-core/common/util"
	"go.uber.org/zap"
)

func IsPrivateKeyExist() bool {
	_, err := models.FindSecret()
	return err == nil
}

func GenPrivateKey() (ed25519.PublicKey, ed25519.PrivateKey) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		util.GetLogger().Error("Failed to generate private key", zap.String("error", err.Error()))
	}

	return publicKey, privateKey
}

func SavePrivatekey(key ed25519.PrivateKey) {
	secret := models.Secret{PrivateKey: key}
	if err := secret.Save(); err != nil {
		util.GetLogger().Error("Failed to save private key", zap.String("error", err.Error()))
	}
}
