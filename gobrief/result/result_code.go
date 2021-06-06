package result

const (
	// SuccessCode 成功状态码
	SuccessCode = 10000

	// ParamInvalid 参数表单错误
	ParamInvalid       = 10001
	ParamCheckError    = 10002
	ParamTypeBindError = 10003
	ParamNotComplete   = 10004

	// UserNotLogin 用户错误
	UserNotLogin         = 20001
	UserLoginError       = 20002
	UserAccountForbidden = 20003
	UserNotExist         = 20004
	UserHasExisted       = 20005

	// DataNone 数据错误
	DataNone           = 50001
	DataIsWrong        = 50002
	DataAlreadyExisted = 50003
	DataCreateWrong    = 50004

	// InterfaceInnerInvokeError 系统错误
	InterfaceInnerInvokeError = 60001

	// PermissionNoAccess 权限错误
	PermissionNoAccess = 70001

	// TokenNotExist Token错误
	TokenNotExist     = 80001
	TokenParamInvalid = 80002
	TokenError        = 80003
)

var resultText = map[int]string{
	SuccessCode:               "成功",
	ParamInvalid:              "参数无效",
	ParamCheckError:           "参数校验失败",
	ParamTypeBindError:        "参数类型错误",
	ParamNotComplete:          "参数缺失",
	UserNotLogin:              "未登录",
	UserLoginError:            "账号不存在或密码动态码错误",
	UserAccountForbidden:      "账号已经被禁用",
	UserNotExist:              "账户不存在",
	UserHasExisted:            "用户已存在",
	DataNone:                  "数据未找到",
	DataIsWrong:               "数据有误",
	DataAlreadyExisted:        "数据已存在",
	DataCreateWrong:           "数据创建错误",
	InterfaceInnerInvokeError: "内部系统接口调用异常",
	PermissionNoAccess:        "无权限访问",
	TokenNotExist:             "token不存在",
	TokenParamInvalid:         "头部请求为空",
	TokenError:                "token格式错误",
}

// ResultText returns a text for the Private status code. It returns the empty
// string if the code is unknown.
func ResultText(code int) string {
	return resultText[code]
}
