package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	dirname, _ := os.Getwd()
	env := os.Getenv("V2FREE_ENV")
	if env == "" {
		env = "dev"
	}
	log.Println("env:", env)
	err := godotenv.Load(dirname + "/" + "config/" + env)
	if err != nil {
		log.Fatal(err)
	}
}
