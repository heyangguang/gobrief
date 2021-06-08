package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gobrief/app/dao"
	"gobrief/app/model"
	"gobrief/gobrief/form_validation"
	"gobrief/gobrief/logger"
	. "gobrief/gobrief/result"
)

func UserCreate(c *gin.Context) {
	user := dao.NewUserModel()
	_ = c.ShouldBind(user)
	if err := form_validation.Validate.Struct(user); err != nil {
		R(c)(ParamCheckError, ResultText(ParamCheckError), nil, E("sys_user", err))(Error)
		return
	}
	logger.Debug(fmt.Sprintf("%v", user))
	R(c)(SuccessCode, ResultText(SuccessCode), nil, nil)(OK)
}

func UserList(c *gin.Context) {
	user := model.UserModel{
		UserId:   0,
		UserName: "",
		UserPwd:  "",
	}
	R(c)(SuccessCode, ResultText(SuccessCode), user, nil)(OK)
}
