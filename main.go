package main

import (
	"fmt"
	"net/http"

	"github.com/gabrieltong/qor-play/app/admins"
	"github.com/gabrieltong/qor-play/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
)

type Config struct {
	Adapter string `yaml:"Adapter"`
	User    string `yaml:"User"`
	Pass    string `yaml:"Pass"`
	Port    string `yaml:"Port"`
	Host    string `yaml:"Host"`
	Name    string `yaml:"Name"`
}

func main() {
	// config := &Config{
	// 	Name: "rails_play",
	// 	User: "root",
	// 	Pass: "1qaz2wsx",
	// 	Host: "127.0.0.1",
	// 	Port: "3306",
	// }
	// // Set up the database
	// DB, _ := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Name))
	// DB, _ := gorm.Open("sqlite3", "demo.db")
	DB := config.DB
	// DB.AutoMigrate(&User{}, &Product{})

	// Initalize
	Admin := admin.New(&admin.AdminConfig{DB: DB})
	App := &config.App{Admin, DB}

	App.Use(&admins.AdminUser{})
	App.Use(&admins.Letter{})
	App.Use(&admins.Setting{})
	App.Use(&admins.ModuleTop{})
	App.Use(&admins.Actor{})
	App.Use(&admins.User{})
	App.Use(&admins.Play{})
	App.Use(&admins.Game{})
	App.Use(&admins.Title{})

	// Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount admin to the mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9876")
	http.ListenAndServe(":9876", mux)
}
