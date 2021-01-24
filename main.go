package main

import (
	"casbin_usage/handler"
	"casbin_usage/middleware"
	"casbin_usage/pkg"
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func init()  {
	adapter, err := gormadapter.NewAdapterByDB(pkg.DB)
	if err != nil {adapter, err := gormadapter.NewAdapterByDB(pkg.DB)
		if err != nil {
			panic(fmt.Sprintf("failed to initialize casbin adapter : %v", adapter))
		}
		panic(fmt.Sprintf("failed to initialize casbin adapter : %v", adapter))
	}
	router = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))
	router.POST("/user/login",handler.Login)
	resource := router.Group("/api")
	resource.Use(middleware.Authenticate())
	{
		resource.GET("/resource",middleware.Authorize("resource", "read", adapter),handler.ReadResource)
		resource.POST("/Resource",middleware.Authorize("resource", "write", adapter),handler.WriteResource)
	}
}

func main()  {
	err := router.Run(":8081")
	if err != nil {
		panic(fmt.Sprintf("failed to start gin engin: %v", err))
	}
	log.Println("application is now running...")
}
