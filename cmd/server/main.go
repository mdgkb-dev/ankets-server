package main

import (
	"log"
	"mdgkb/ankets-server/loggerhelper"
	"mdgkb/ankets-server/migrations"
	"mdgkb/ankets-server/routing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/pro-assistance/pro-assister/config"
	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	helper := helperPack.NewHelper(*conf)
	logger := loggerhelper.NewLogger()

	router := gin.New()
	router.Use(gin.Recovery())
	//
	router.Use(loggerhelper.LoggingMiddleware(logger))
	routing.Init(router, helper)
	// helper.DB.DB.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{Logger: logger, ErrorLevel: logrus.ErrorLevel, QueryLevel: logrus.DebugLevel}))

	helper.Run(migrations.Init(), routing.Init)
}
