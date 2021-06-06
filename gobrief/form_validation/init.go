package form_validation

import (
	languageEn "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// use a single instance , it caches struct info
var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
	ModelFuncSlice []ModelFunc
)

type ModelFunc map[string]func(err validator.ValidationErrors) []map[string]string

func InitFormValidation() {
	// 注册自定义表单方法
	RegisterCustomValidationFunc(UserValidations)

	// 初始化自定义表单方法
	InitCustomValidationFunc()
}

func init() {
	en := languageEn.New()
	Uni = ut.New(en)
	Validate = validator.New()
	RegisterTagFunc()
	ModelFuncSlice = append(ModelFuncSlice, ModelFunc{"sys_user": GetUserValidationError})
}


