package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

type FullConfig struct {
	Development Config
}

type Config struct {
	Adapter string `yaml:"Adapter"`
	User    string `yaml:"User"`
	Pass    string `yaml:"Pass"`
	Port    string `yaml:"Port"`
	Host    string `yaml:"Host"`
	Name    string `yaml:"Name"`
}

var DB *gorm.DB
var config *Config

func init() {

	yamlFile, err := ioutil.ReadFile("database.yml")
	if err != nil {
		fmt.Sprintln(err)
	}
	config = &Config{}
	// fmt.Sprintf(yamlFile)
	err = yaml.Unmarshal(yamlFile, config)
	// config := fullConfig.Development
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	// fmt.Printf("%+v\n", fullConfig)
	fmt.Printf("config.Name is %v", config.Name)
	fmt.Printf("%+v\n", config)
	DB, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Name))

	DB.LogMode(true)
}
