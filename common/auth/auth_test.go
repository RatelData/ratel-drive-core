package auth

import (
	"testing"

	"github.com/RatelData/ratel-drive-core/common/util"
)

func TestLogin(t *testing.T) {
	util.SetAppConfigFilePath("../../config/app.json")
	_, err := Login("test@test.com", "test123456")
	if err != nil {
		t.Fail()
	}
}
