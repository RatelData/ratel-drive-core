package device

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RatelData/ratel-drive-core/app/models"
	"github.com/RatelData/ratel-drive-core/common/util"
	"github.com/RatelData/ratel-drive-core/lib/requests"
	"github.com/RatelData/ratel-drive-core/lib/secret"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RegisterDeviceReq struct {
	Device struct {
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		ServicePort int    `json:"service_port"`
		PublicKey   string `json:"public_key"`
	} `json:"device"`
}

func RegisterDevice(token string) bool {
	if hasAlreadyRegistered() {
		util.GetLogger().Info("This device has already been registered")
		return true
	}

	endpoint := "/api/v1/devices/"

	registerDevReq := RegisterDeviceReq{}
	registerDevReq.Device.UUID = genDeviceUUID()
	registerDevReq.Device.Name = genDeviceName()
	registerDevReq.Device.ServicePort = util.GetAppConfig().ServerPort
	publicKey, privateKey := secret.GenPrivateKey()
	registerDevReq.Device.PublicKey = string(publicKey)

	body, _ := json.Marshal(registerDevReq)
	resp, err := requests.Post(endpoint, body)

	return handleRegisterDeviceResult(resp, err, privateKey)
}

// Generate a unique device id
func genDeviceUUID() string {
	return uuid.NewString()
}

// Generate a random device name
func genDeviceName() string {
	return fmt.Sprintf("device-%d", time.Now().Unix())
}

func handleRegisterDeviceResult(resp *resty.Response, err error, privateKey ed25519.PrivateKey) bool {
	logger := util.GetLogger()

	if err == nil && resp.StatusCode() == http.StatusCreated {
		logger.Info("Register device succeed!",
			zap.String("body", resp.String()),
		)

		secret.SavePrivatekey(privateKey)
		return true
	}

	var errInfo zapcore.Field
	if err != nil {
		errInfo = zap.String("error", err.Error())
	} else {
		errInfo = zap.String("status", resp.Status())
	}

	logger.Error("Register device failed!",
		errInfo,
	)

	return false
}

func hasAlreadyRegistered() bool {
	_, err := models.FindSecret()
	return err == nil
}
