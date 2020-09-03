package middleware

import (
	"oh-my-anime_gf/app/service/user"
	"github.com/gogf/gf/net/ghttp"
	"oh-my-anime_gf/library/response"
)

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request){
	if user.IsSignedIn(r.Session) {
		r.Middleware.Next()
	} else {
		response.JsonExit(r, response.UNAUTHORIZED, "用户未登录")
	}
}

