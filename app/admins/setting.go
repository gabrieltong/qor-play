package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
	"github.com/qor/admin"
)

type Setting struct {
}

func (setting *Setting) ConfigureApplication(app *config.App) {
	app.Admin.AddResource(&models.Setting{}, &admin.Config{Singleton: true})
}
