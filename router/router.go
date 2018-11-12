package router

import (
	"github.com/gin-gonic/gin"
	"github.com/unstd/workbench/controller"
)

func Router(r *gin.Engine) {
	r.GET("/health_check", controller.HealthCheck)

	tool := r.Group("/tools")
	{
		tool.POST("/md5", controller.Md5)
	}
}
