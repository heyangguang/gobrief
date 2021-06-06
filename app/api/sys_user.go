package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobrief/app/dao"
	"gobrief/gobrief/form_validation"
	"gobrief/gobrief/logger"
	. "gobrief/gobrief/result"
)

func UserCreate(c *gin.Context) {
	user := dao.NewUserModel()
	_= c.ShouldBind(user)
	if err := form_validation.Validate.Struct(user); err != nil {
		R(c)(ParamCheckError, ResultText(ParamCheckError), E("sys_user", err))(Error)
		return
	}
	logger.Debug(fmt.Sprintf("%v", user))
	R(c)(SuccessCode, ResultText(SuccessCode), nil)(OK)
}
