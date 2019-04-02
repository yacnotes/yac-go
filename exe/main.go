package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	"yac-go/config"
	"yac-go/ydb"
	"yac-go/log"
	"yac-go/middleware"
	"yac-go/routes"
)

const APP_VERSION = "v0.1.0"

func main() {
	var conf config.Config
	conf.LoadFromEnv(APP_VERSION)

	log.Level = conf.Log.Level
	log.Info("Initialized logger with level", log.Level)
	log.Info("Application version:", APP_VERSION)

	if conf.Env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	deps := &config.AppDeps{
		Config: conf,
	}
	deps.Db = ydb.Init(deps)

	r := gin.New()
	r.Use(middleware.Logger(deps))
	r.Use(middleware.Injector(deps))
	r.Use(gin.Recovery())

	routes.LoadRoutes(r, deps)

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(conf.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       log.GetLogger(),
	}

	log.Info("Running server on port", conf.Port)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic("Server listen error:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	if err := deps.Db.Close(); err != nil {
		log.Error("Error while closing database:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Panic("Server Shutdown:", err)
	}
	log.Info("Server stopped.")
}
