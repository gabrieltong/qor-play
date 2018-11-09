package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
)

type Game struct {
}

func (game *Game) ConfigureApplication(app *config.App) {
	app.Admin.AddResource(&models.Game{})
}
