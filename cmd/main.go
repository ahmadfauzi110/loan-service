package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/ahmadfauzi110/loan-service/config"
	"github.com/ahmadfauzi110/loan-service/internal/router"
	"github.com/ahmadfauzi110/loan-service/util"
)

func main() {
	// Load .env
	cfg := config.Initialize(".")
	dbCon := config.SetupDatabase(cfg.DB)

	// Setup Echo
	e := echo.New()
	e.Static(config.CurrentConfig.STATIC_PATH, "storage/letter")
	e.Validator = util.NewValidator()

	router.SetupRoutes(e, dbCon)

	// Run server

	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Error(err)
	}

	e.Logger.Info("Server shutdown complete")
}
