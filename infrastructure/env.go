package infrastructure

import (
	"log"
	"os"
)

func (env *Environment) LoadConfig() {
	connection := os.Getenv("DB_URL")
	port := os.Getenv("APP_PORT")
	apps := app{Appname: "Gratitude_Journal", Port: port}
	db := database{Name: "postgres1", Connection: connection}
	env.Databases = map[string]database{"db": db}
	env.App = apps
	log.Println("Config load successfully!")
}
