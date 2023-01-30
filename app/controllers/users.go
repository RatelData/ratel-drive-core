package controllers

import (
	"net/http"

	"github.com/RatelData/ratel-drive-core/app/models/validators"
	"github.com/RatelData/ratel-drive-core/common/auth"
	"github.com/RatelData/ratel-drive-core/common/errors"
	"github.com/RatelData/ratel-drive-core/lib/device"
	"github.com/RatelData/ratel-drive-core/lib/requests"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	lv := validators.LoginValidator{}
	if err := lv.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.NewValidatorError(err))
		return
	}

	loginResult, loginErr := auth.Login(lv.User.Email, lv.User.Password)
	if loginErr != nil {
		c.JSON(http.StatusUnauthorized, errors.NewValidatorError(loginErr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"token":   loginResult.User.Token,
			"user_id": loginResult.User.UserID,
		},
	})

	requests.Init(loginResult.User.Token)

	device.RegisterDevice(loginResult.User.Token)
}
