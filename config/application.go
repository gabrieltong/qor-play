package config

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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

type Options []Option

type Option struct {
	Key string
	V   string
}

func (options *Options) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, options)
	case string:
		if v != "" {
			return options.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (options Options) Value() (driver.Value, error) {
	if len(options) == 0 {
		return nil, nil
	}
	return json.Marshal(options)
}
