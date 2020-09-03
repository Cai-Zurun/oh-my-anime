package main

import (
	"github.com/gogf/gf/frame/g"
	_ "oh-my-anime_gf/boot"
	_ "oh-my-anime_gf/router"
)

func main() {
	g.Server().Run()
}
