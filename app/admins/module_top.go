package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
)

type ModuleTop struct {
}

func (moduleTop *ModuleTop) ConfigureApplication(app *config.App) {
	app.Admin.AddResource(&models.ModuleTop{})
}
