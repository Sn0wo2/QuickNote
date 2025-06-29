package handler

import (
	"github.com/Sn0wo2/QuickNote/pkg/common"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Health() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		log.Instance.Debug("Health check",
			zap.String("ctx", common.FiberContextString(ctx)))

		return ctx.JSON(response.New("success"))
	}
}
