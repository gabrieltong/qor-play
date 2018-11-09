package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
)

type Letter struct {
}

func (letter *Letter) ConfigureApplication(app *config.App) {
	app.Admin.AddResource(&models.Letter{})
}
