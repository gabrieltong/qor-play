package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
)

type Actor struct {
}

func (actor *Actor) ConfigureApplication(app *config.App) {
	app.Admin.AddResource(&models.Actor{})
}
