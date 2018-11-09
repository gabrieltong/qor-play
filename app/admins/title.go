package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
)

type Title struct {
}

func (title *Title) ConfigureApplication(app *config.App) {
	app.Admin.AddResource(&models.Title{})
}
