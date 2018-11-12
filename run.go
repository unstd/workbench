package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unstd/workbench/router"
)

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run()
}
