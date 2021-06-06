package form_validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

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


// GetUserValidationError 修改验证字段错误返回值方法
func GetUserValidationError(err validator.ValidationErrors) []map[string]string {
	sliceErrs := BaseFormValidationError(err)
	for _, value := range sliceErrs {
		if errValue, ok := value["user_name"]; ok {
			if strings.Contains(errValue, "ValidationUserNameFormat") {
				value["user_name"] = "user_name can only be English letters"
			}
		}
	}
	return sliceErrs
}