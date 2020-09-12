package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"oh-my-anime_gf/app/api/user"
)

// 鉴权中间件，只有登录成功之后才能通过; 基于session的
//func Auth(r *ghttp.Request){
//	if user.IsSignedIn(r.Session) {
//		r.Middleware.Next()
//	} else {
//		response.JsonExit(r, response.UNAUTHORIZED, "用户未登录")
//	}
//}

func Auth(r *ghttp.Request) {
	user.GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}