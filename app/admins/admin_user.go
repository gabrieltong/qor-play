package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

type AdminUser struct {
}

func (admin_user *AdminUser) ConfigureApplication(app *config.App) {
	// Create resources from GORM-backend model
	AdminUser := app.Admin.AddResource(&models.AdminUser{})

	AdminUser.Scope(&admin.Scope{
		Name: "IsSuper",
		// Label: strings.Title(strings.Replace(state, "_", " ", -1)),
		Group: "Role",
		Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
			return db.Where("is_super = ?", 1)
		},
	})

	AdminUser.Scope(&admin.Scope{
		Name: "IsAuthor",
		// Label: strings.Title(strings.Replace(state, "_", " ", -1)),
		Group: "Role",
		Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
			return db.Where("is_author = ?", 1)
		},
	})
}
