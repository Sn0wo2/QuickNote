package listener

import (
	"github.com/Sn0wo2/QuickNote/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func Start(app *fiber.App) error {
	if config.Instance.TLS.Cert != "" && config.Instance.TLS.Key != "" {
		return app.ListenTLS(config.Instance.Listener.Address, config.Instance.TLS.Cert, config.Instance.TLS.Key)
	}

	return app.Listen(config.Instance.Listener.Address)
}
