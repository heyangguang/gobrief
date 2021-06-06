package result

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gobrief/gobrief/form_validation"
	"sync"
)

var resultPool *sync.Pool
type resultFunc func(code int, msg interface{}, err interface{}) func(sf PutStatusCodeFunc)
type PutStatusCodeFunc func(c *gin.Context, v interface{})

type RetErrorType struct {
	ModelName string
	Err error
}

func init() {
	resultPool = &sync.Pool{
		New: func() interface{} {
			return NewResult(0, nil, nil)
		},
	}
}

type Result struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Err interface{} `json:"err"`
}

func NewResult(code int, msg interface{}, err interface{}) *Result {
	return &Result{
		Code: code,
		Msg: msg,
		Err: err,
	}
}

// E 用于构建RetErrorType方法
func E(mn string, e error) RetErrorType {
	return RetErrorType{
		ModelName: mn,
		Err:       e,
	}
}

// R 用于统一构建返回结构体
// 其中为了处理validator.ValidationErrors特殊的错误返回，使用了RetErrorType
func R(c *gin.Context) resultFunc {
	return func(code int, msg interface{}, err interface{}) func(sf PutStatusCodeFunc) {
		r := resultPool.Get().(*Result)
		defer resultPool.Put(r)
		r.Code = code
		r.Msg = msg
		r.Err = err
		if v, ok := err.(RetErrorType); ok {
			for _, mf := range form_validation.ModelFuncSlice {
				for key, value := range mf {
					if key == v.ModelName {
						r.Err = value(v.Err.(validator.ValidationErrors))
					}
				}
			}
		}
		return func(sf PutStatusCodeFunc) {
			sf(c, r)
		}
	}
}

// OK 200
func OK(c *gin.Context, v interface{}) {
	c.JSON(200, v)
}

// Error 400
func Error(c *gin.Context, v interface{}) {
	c.JSON(400, v)
}