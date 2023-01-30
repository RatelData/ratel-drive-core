package device

import (
	"testing"

	"github.com/RatelData/ratel-drive-core/common/auth"
	"github.com/RatelData/ratel-drive-core/common/util"
)

func TestRegisterDevice(t *testing.T) {
	util.SetAppConfigFilePath("../../config/app.json")

	loginResult, err := auth.Login("test@rateldata.io", "test123456")
	if err != nil || !RegisterDevice(loginResult.User.Token) {
		t.Fail()
	}
}
