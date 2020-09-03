package response

import (
	"github.com/gogf/gf/net/ghttp"
)

const (
	SUCCESS			= 200
	FAIL			= 400
	UNAUTHORIZED	= 401
)

type JsonResponse struct {
	Code    int         `json:"code"`    // 响应码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

func Json(r *ghttp.Request, code int, message string, data ...interface{})  {
	responseData := interface{}(nil)	//赋值为空接口类型的空值
	if len(data) > 0 {
		responseData = data[0]			//如果是错误信息的话，取所有错误信息的第一条即可;如果传过来的是数组的指针，就可以response数组值
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

func JsonExit(r *ghttp.Request, code int, msg string, data ...interface{}) {
	Json(r, code, msg, data...)
	r.Exit()
}
