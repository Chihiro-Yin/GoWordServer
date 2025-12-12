package main

import (
	"imxy.top/wordserver/internal/config"
	"imxy.top/wordserver/internal/routers"
)

func main() {
	config.InitDB()
	routers.InitAllRouters()
}
