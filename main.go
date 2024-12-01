package main

import (
	_ "program/config"
	"program/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logs.Debug(nil, "成功传参")
	r.Run()

}
