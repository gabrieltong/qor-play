package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
)

type User struct {
}

func (user *User) ConfigureApplication(app *config.App) {
	app.Admin.AddResource(&models.User{})
}
