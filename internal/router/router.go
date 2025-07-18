package router

import (
	"github.com/Sn0wo2/QuickNote/internal/router/handler"
	"github.com/Sn0wo2/QuickNote/internal/router/notfound"
	"github.com/Sn0wo2/QuickNote/internal/router/static"
	"github.com/Sn0wo2/QuickNote/internal/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(router fiber.Router) {
	router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	router.Use(cors.New())

	api := router.Group(setup.APIVersion)

	api.Group("health", handler.Health())
	api.Group("notes/*", handler.Note())
	api.Group("export/*", handler.Export())
	api.Group("import", handler.Import())

	notfound.Setup("api not found", api)

	static.Setup(router)

	notfound.Setup("static not found", router)
}
