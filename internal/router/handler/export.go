package handler

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/Sn0wo2/QuickNote/internal/note"
	"github.com/Sn0wo2/QuickNote/pkg/common"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Export() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		nid := ctx.Params("*")
		if nid == "" {
			log.Instance.Warn("Invalid nid",
				zap.String("ctx", common.FiberContextString(ctx)))

			return ctx.Status(fiber.StatusBadRequest).JSON(response.New("invalid nid"))
		}

		n, err := note.GetNote(nid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Instance.Warn("Note not found",
					zap.String("nid", nid),
					zap.String("ctx", common.FiberContextString(ctx)))

				return ctx.Status(fiber.StatusNotFound).JSON(response.New("note not found"))
			}

			log.Instance.Error("Failed to get note",
				zap.String("nid", nid),
				zap.Error(err),
				zap.String("ctx", common.FiberContextString(ctx)))

			return err
		}

		ctx.Set("Content-Type", "application/x-qnote")
		ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.qnote\"", n.NID))

		log.Instance.Info("Note exported",
			zap.String("nid", nid),
			zap.String("ctx", common.FiberContextString(ctx)))

		return ctx.SendStream(bytes.NewReader(n.Data))
	}
}
