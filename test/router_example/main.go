package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	languageEn "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"regexp"
	"strings"
)

var (
	Validate *validator.Validate
	Uni      *ut.UniversalTranslator
)


func hello(c *gin.Context)  {
	var login Login
	_ = c.Bind(&login)
	if err := Validate.Struct(&login); err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{"err": GetUserValidationError(err.(validator.ValidationErrors))})
		return
	}
	c.JSON(200, login)
}

type Login struct {
	UserName string	`json:"user_name" validate:"required,min=3,max=6,ValidationUserNameFormat"`
	UserPwd string `validate:"required" json:"user_pwd"`
}

// RegisterTagFunc 注册一个函数，获取struct tag里自定义的tag作为字段名
func RegisterTagFunc() {
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		return name
	})
}

// UserValidations 用户自定义验证动作
func UserValidations() {
	_ = Validate.RegisterValidation("ValidationUserNameFormat", ValidationUserNameFormat)
}

// ValidationUserNameFormat 自定义用户名字段验证
func ValidationUserNameFormat(fl validator.FieldLevel) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z_]+$", fl.Field().String()); !ok {
		return false
	}
	return true
}

// BaseFormValidationError 公共验证字段错误返回值方法
func BaseFormValidationError(err validator.ValidationErrors) []map[string]string {
	// 这步是在全部转换成英文，
	tans, _ := Uni.GetTranslator("en")
	// 注册英文uni
	_ = enTranslations.RegisterDefaultTranslations(Validate, tans)
	// 这步开始才是把所有的错误放到字典切片里
	var sliceErrs []map[string]string
	for _, e := range err {
		sliceErrs = append(sliceErrs, map[string]string{e.Field(): e.Translate(tans)})
	}
	return sliceErrs
}

// GetUserValidationError 修改验证字段错误返回值方法
func GetUserValidationError(err validator.ValidationErrors) []map[string]string {
	// 拿到字典切片
	sliceErrs := BaseFormValidationError(err)
	// 遍历字典切片
	for _, value := range sliceErrs {
		// 找到我们对应要替换的字段
		if errValue, ok := value["user_name"]; ok {
			// 找到我们自定义的方法，并修改字典的value
			if strings.Contains(errValue, "ValidationUserNameFormat") {
				value["user_name"] = "user_name can only be English letters"
			}
		}
	}
	return sliceErrs
}

func main() {
	en := languageEn.New()
	Uni = ut.New(en)
	Validate = validator.New()
	RegisterTagFunc()
	UserValidations()
	r := gin.Default()
	r.POST("/", hello)
	r.Run(":8000")
}