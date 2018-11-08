package main

import (
	"fmt"
	"net/http"

	"github.com/gabrieltong/qor-play/config"
	"github.com/jinzhu/gorm"
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

// Define a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Define another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
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
	DB.AutoMigrate(&User{}, &Product{})

	// Initalize
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	// Create resources from GORM-backend model
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	// Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount admin to the mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}
