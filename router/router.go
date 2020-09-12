package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"oh-my-anime_gf/app/api/anime"
	"oh-my-anime_gf/app/api/user"
	"oh-my-anime_gf/app/service/middleware"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlUser := new(user.Controller)
		group.Middleware(middleware.CORS)
		group.ALL("/user", ctlUser)
		group.POST("/user/sign-in", user.GfJWTMiddleware.LoginHandler)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Group("/anime",func(group *ghttp.RouterGroup) {
				group.ALL("/get", anime.GetAnime)
				group.ALL("/all/get", anime.GetAllAnime)
				group.ALL("/type/get", anime.GetType)
			})
			group.Group("/anime",func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.Auth)
				group.ALL("/add", anime.AddAnime)
				group.ALL("/delete", anime.DeleteAnime)
				group.ALL("/update", anime.UpdateAnime)
				group.ALL("/type/add", anime.AddType)
				group.ALL("/type/delete", anime.DeleteType)
				group.ALL("/type/update", anime.UpdateType)
			})
		})
	})
}

