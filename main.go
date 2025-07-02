package main

import (
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"

	"github.com/Sn0wo2/QuickNote/internal/listener"
	"github.com/Sn0wo2/QuickNote/internal/router"
	"github.com/Sn0wo2/QuickNote/internal/setup"
	"github.com/Sn0wo2/QuickNote/pkg/config"
	"github.com/Sn0wo2/QuickNote/pkg/database/orm"
	"github.com/Sn0wo2/QuickNote/pkg/database/table"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/version"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// stw
	debug.SetGCPercent(50)

	_ = godotenv.Load()

	err := config.Init()
	if err != nil {
		// log not init~
		panic(err)
	}

	log.Init()
}

func main() {
	defer func() {
		_ = log.Instance.Sync()
	}()

	if !fiber.IsChild() {
		log.Instance.Info("Starting QuickNote...", zap.String("version", version.GetFormatVersion()))
	}

	if err := orm.Init(config.Instance.Database.Type, config.Instance.Database.URL); err != nil {
		log.Instance.Fatal("Failed to initialize database",
			zap.Error(err),
		)
	}

	if err := table.Init(); err != nil {
		log.Instance.Fatal("Failed to initialize tables",
			zap.Error(err),
		)
	}

	if !fiber.IsChild() {
		log.Instance.Info("Starting server",
			zap.String("address", config.Instance.Listener.Address),
			zap.Int("pid", os.Getpid()),
		)
	}

	app := setup.Fiber()
	router.Setup(app)

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := listener.Start(app); err != nil {
			log.Instance.Fatal("Server failed to start",
				zap.Error(err),
			)
		}
	}()

	<-shutdownChan

	if err := app.Shutdown(); err != nil {
		log.Instance.Error("Server shutdown error",
			zap.Error(err),
		)
	}
}
