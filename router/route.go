package router

import (
	"go-dev-test/utils/config"
	"go-dev-test/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.Conf.Server.AppMode)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	router := engine.Group("api/v1")
	{
		router.GET("/appmode", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  errmsg.SUCCESS,
				"message": config.Conf.Server.AppMode,
			})
		})
	}

	engine.Run(config.Conf.Server.HttpPort)
}
