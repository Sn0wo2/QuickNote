package static

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Sn0wo2/QuickNote/pkg/common"
	"github.com/Sn0wo2/QuickNote/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Setup(router fiber.Router) {
	/*router.Static("/", staticRoot, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 60 * time.Second,
		MaxAge:        60,
		Index:         index,
		Next: func(ctx *fiber.Ctx) bool {
			log.Instance.Info("Send static file",
				zap.String("ctx", common.FiberContextString(ctx)))
			return false
		},
	})*/
	router.Use("*", func(ctx *fiber.Ctx) error {
		if ctx.Method() != fiber.MethodGet {
			return ctx.Next()
		}

		if path := GetStaticFile("index.html", "./static", ctx.Path()); path != "" {
			log.Instance.Info("Send static file",
				zap.String("file", path),
				zap.String("ctx", common.FiberContextString(ctx)))

			return ctx.SendFile(path, true)
		}

		return ctx.Next()
	})
}

// GetStaticFile react router
func GetStaticFile(index, staticRoot, reqPath string) string {
	base := filepath.Join(staticRoot, strings.TrimPrefix(reqPath, "/"))
	defaultIndex := filepath.Join(staticRoot, index)

	if rel, err := filepath.Rel(staticRoot, base); err != nil || strings.HasPrefix(rel, "..") {
		return defaultIndex
	}

	if info, err := os.Stat(base); err == nil {
		switch mode := info.Mode(); {
		case mode.IsDir():
			idx := filepath.Join(base, index)
			if _, err = os.Stat(idx); err == nil {
				return idx
			}

			return defaultIndex

		case mode.IsRegular():
			return base
		}
	}

	return defaultIndex
}
