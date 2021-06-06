package form_validation

import (
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

type Option func()
var options []Option


// RegisterTagFunc 注册一个函数，获取struct tag里自定义的tag作为字段名
func RegisterTagFunc() {
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		return name
	})
}

// RegisterCustomValidationFunc 注册自定义表单验证方法
func RegisterCustomValidationFunc(opts ...Option) {
	options = append(options, opts...)
}

// InitCustomValidationFunc 初始化自定义表单方法
func InitCustomValidationFunc() {
	for _, opt := range options {
		opt()
	}
}

// BaseFormValidationError 公共验证字段错误返回值方法
func BaseFormValidationError(err validator.ValidationErrors) []map[string]string {
	tans, _ := Uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(Validate, tans)
	var sliceErrs []map[string]string
	for _, e := range err {
		sliceErrs = append(sliceErrs, map[string]string{e.Field(): e.Translate(tans)})
	}
	return sliceErrs
}
