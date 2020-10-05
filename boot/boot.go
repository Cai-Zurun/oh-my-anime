 package boot

import (
	"github.com/gogf/gf/frame/g"
	_ "oh-my-anime_gf/packed"
	"github.com/gogf/swagger"
) 

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}

