package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RatelData/ratel-drive-core/common/util"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CentralToken() string {
	return ""
}

type LoginReq struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type LoginResult struct {
	User struct {
		UserID uint   `json:"user_id"`
		Token  string `json:"token"`
	} `json:"user"`
}

func Login(email string, password string) (LoginResult, error) {

	endpoint := "/api/v1/login"
	loginReq := LoginReq{}
	loginReq.User.Email = email
	loginReq.User.Password = password

	body, _ := json.Marshal(loginReq)

	loginResult := LoginResult{}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&loginResult).
		Post(util.CentralHost() + endpoint)

	if handleLoginResult(resp, err) {
		return loginResult, nil
	}

	return loginResult, errors.New("Login failed")
}

func handleLoginResult(resp *resty.Response, err error) bool {
	logger := util.GetLogger()

	if err == nil && resp.StatusCode() == http.StatusOK {
		logger.Info("Login succeed!",
			zap.String("body", resp.String()),
		)
		return true
	}

	var errInfo zapcore.Field
	if err != nil {
		errInfo = zap.String("error", err.Error())
	} else {
		errInfo = zap.String("status", resp.Status())
	}

	logger.Error("Login failed!",
		errInfo,
	)

	return false
}
