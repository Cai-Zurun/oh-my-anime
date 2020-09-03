package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"oh-my-anime_gf/app/api/anime"
	"oh-my-anime_gf/app/api/hello"
	"oh-my-anime_gf/app/api/user"
	"oh-my-anime_gf/app/service/middleware"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		ctlUser := new(user.Controller)
		group.Middleware(middleware.CORS)
		group.ALL("/", hello.Hello)
		group.ALL("/user", ctlUser)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth)
			group.Group("/anime/type",func(group *ghttp.RouterGroup) {
				group.ALL("/add", anime.AddType)
				group.ALL("/get", anime.GetType)
				group.ALL("/delete", anime.DeleteType)
				group.ALL("/update", anime.UpdateType)
			})
			group.Group("/anime",func(group *ghttp.RouterGroup) {
				group.ALL("/add", anime.AddAnime)
				group.ALL("/get", anime.GetAnime)
				group.ALL("/delete", anime.DeleteAnime)
				group.ALL("/update", anime.UpdateAnime)
			})
		})
	})
}

