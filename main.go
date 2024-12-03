package main

import (
	_ "program/config"
	"program/utils/jwtulits"
	_ "program/utils/logs"

	_ "github.com/gin-gonic/gin"
)

func main() {

	jwtulits.GenToken("XXIAOHZOU")

}
