package config

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

// // Define a GORM-backend model
// type User struct {
// 	gorm.Model
// 	Name string
// }

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Define another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}

type AdminInterface interface {
	ConfigureApplication(*App)
}

type App struct {
	Admin *admin.Admin
	DB    *gorm.DB
}

func (app *App) Use(admin AdminInterface) {
	admin.ConfigureApplication(app)
}
