package handler

import (
	"io"
	"mime/multipart"
	"strings"

	"github.com/Sn0wo2/QuickNote/internal/note"
	"github.com/Sn0wo2/QuickNote/pkg/common"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/Sn0wo2/QuickNote/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Import() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fileHeader, err := ctx.FormFile("import")
		if err != nil {
			log.Instance.Error("Failed to get file",
				zap.Error(err),
				zap.String("ctx", common.FiberContextString(ctx)),
			)

			return err
		}

		fileName := fileHeader.Filename
		if !strings.HasSuffix(fileName, ".qnote") {
			log.Instance.Warn("Invalid file format",
				zap.String("fileName", fileName),
				zap.String("ctx", common.FiberContextString(ctx)),
			)

			return ctx.Status(fiber.StatusBadRequest).JSON(response.New("invalid file format"))
		}

		file, err := fileHeader.Open()
		if err != nil {
			log.Instance.Error("Failed to open file",
				zap.Error(err),
				zap.String("fileName", fileName),
				zap.String("ctx", common.FiberContextString(ctx)),
			)

			return err
		}

		defer func(file multipart.File) {
			_ = file.Close()
		}(file)

		result, err := io.ReadAll(file)
		if err != nil {
			log.Instance.Error("Failed to read file",
				zap.Error(err),
				zap.String("fileName", fileName))

			return err
		}

		n := note.Note{NID: strings.TrimSuffix(fileName, ".qnote")}
		if err = n.Decode(result); err != nil {
			log.Instance.Error("Failed to decode note",
				zap.Error(err),
				zap.String("fileName", fileName),
				zap.String("ctx", common.FiberContextString(ctx)),
			)

			return err
		}

		if err = n.Write(); err != nil {
			log.Instance.Error("Failed to write file",
				zap.Error(err),
				zap.String("fileName", fileName),
				zap.String("ctx", common.FiberContextString(ctx)))

			return err
		}

		log.Instance.Debug("Import note",
			zap.String("fileName", fileName),
			zap.String("ctx", common.FiberContextString(ctx)),
		)

		return ctx.JSON(response.New("success"))
	}
}
