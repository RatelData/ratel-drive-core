package requests

import (
	"errors"

	"github.com/RatelData/ratel-drive-core/common/util"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

var client *resty.Client

func Init(token string) {
	if client == nil {
		client = resty.New()
	}
	client.SetAuthToken(token)
}

func Get(endpoint string, queryParams map[string]string) (*resty.Response, error) {
	if ok, err := checkClientInitialized(); !ok {
		return nil, err
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		Get(util.CentralHost() + endpoint)

	return resp, err
}

func Post(endpoint string, body interface{}) (*resty.Response, error) {
	if ok, err := checkClientInitialized(); !ok {
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(util.CentralHost() + endpoint)

	return resp, err
}

func Put(endpoint string, body interface{}) (*resty.Response, error) {
	if ok, err := checkClientInitialized(); !ok {
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Put(util.CentralHost() + endpoint)

	return resp, err
}

func Patch(endpoint string, body interface{}) (*resty.Response, error) {
	if ok, err := checkClientInitialized(); !ok {
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Patch(util.CentralHost() + endpoint)

	return resp, err
}

func Delete(endpoint string) (*resty.Response, error) {
	if ok, err := checkClientInitialized(); !ok {
		return nil, err
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(util.CentralHost() + endpoint)

	return resp, err
}

func checkClientInitialized() (bool, error) {
	logger := util.GetLogger()
	if client == nil {
		err := errors.New("http client has not been initialized yet")
		logger.Error("http client!",
			zap.String("error", err.Error()),
		)
		return false, err
	}

	return true, nil
}
