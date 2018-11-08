package main

import (
	"fmt"
	"net/http"

	"github.com/gabrieltong/qor-play/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
)

func main() {

	// activity := &models.Activity{}

	// config.DB.First(&activity)
	mux := http.NewServeMux()
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&models.Activity{})
	admin := admin.New(&admin.AdminConfig{DB: DB})
	// admin.AddResource(&models.Activity{})
	admin.AddResource(&models.Activity{})
	admin.MountTo("/admin", mux)

	mux.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", req.URL.Path[1:])
	})

	http.ListenAndServe(":9001", mux)
	fmt.Println("started")
	// log.Fatal()
}
