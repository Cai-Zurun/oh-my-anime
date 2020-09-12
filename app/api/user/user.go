package user

import (
	"github.com/gogf/gf/net/ghttp"
	"oh-my-anime_gf/app/service/user"
	"oh-my-anime_gf/library/response"
)

type Controller struct {}

type SignUpRequest struct {
	user.SignUpInput
}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   passport  formData string  true "用户账号名称"
// @param   password  formData string  true "用户密码"
// @param   password2 formData string  true "确认密码"
// @param   nickname  formData string false "用户昵称"
// @router  /user/sign-up [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (c *Controller) SignUp(r *ghttp.Request) {
	var data *SignUpRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	println(data)
	if err := user.SignUp(&data.SignUpInput); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.FAIL, "注册成功")
	}
}

// 这个是最初基于session的登录，新的基于jwt
//type SignInRequest struct {
//	Passport string `v:"required#账号不能为空"`
//	Password string `v:"required#密码不能为空"`
//}
//
//// @summary 用户登录接口
//// @tags    用户服务
//// @produce json
//// @param   passport formData string true "用户账号"
//// @param   password formData string true "用户密码"
//// @router  /user/sign-in [POST]
//// @success 200 {object} response.JsonResponse "执行结果"
//func (c *Controller) SignIn(r *ghttp.Request)  {
//	var data *SignInRequest
//	if err := r.Parse(&data); err != nil {
//		response.JsonExit(r, response.FAIL, err.Error())
//	}
//	if err := user.SignIn(data.Passport, data.Password, r.Session); err != nil {
//		response.JsonExit(r, response.FAIL, err.Error())
//	} else {
//		response.JsonExit(r, response.SUCCESS, "登陆成功", g.Map{"SessionId": r.Session.Id()} )
//	}
//}

// @summary 判断用户是否已经登录
// @tags    用户服务
// @produce json
// @router  /user/is-signed-in [GET]
// @success 200 {object} response.JsonResponse "执行结果:`true/false`"
func (c *Controller) IsSignedIn(r *ghttp.Request) {
	response.JsonExit(r, response.SUCCESS, "", user.IsSignedIn(r.Session))
}

// @summary 用户注销/退出接口
// @tags    用户服务
// @produce json
// @router  /user/sign-out [GET]
// @success 200 {object} response.JsonResponse "执行结果, 1: 未登录"
func (c *Controller) SignOut(r *ghttp.Request) {
	if err := user.SignOut(r.Session); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	response.JsonExit(r, response.SUCCESS, "退出成功")
}
