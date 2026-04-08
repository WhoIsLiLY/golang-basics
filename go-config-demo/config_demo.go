package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Version string `mapstructure:"version"`
		Port    int    `mapstructure:"port"`
	} `mapstructure:"app"`
	Database struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		Name string `mapstructure:"name"`
		User string `mapstructure:"user"`
	} `mapstructure:"database"`
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Error unmarshaling config:", err)
	}

	fmt.Printf("App: %s v%s\n", config.App.Name, config.App.Version)
	fmt.Printf("Port: %d\n", config.App.Port)
	fmt.Printf("Database: %s@%s:%d/%s\n", 
		config.Database.User, 
		config.Database.Host, 
		config.Database.Port, 
		config.Database.Name)
}