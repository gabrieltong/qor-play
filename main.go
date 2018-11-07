package main

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
)

func main() {
	activity := &models.Activity{}

	config.DB.First(&activity)
}
