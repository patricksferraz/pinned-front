package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"github.com/patricksferraz/pinned-front/app"
	"github.com/patricksferraz/pinned-front/config"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(basepath + "/.env")
		if err != nil {
			log.Printf("Error loading .env files")
		}
	}
}

func main() {
	var conf config.Config

	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatal(err)
	}

	app.StartApp(&conf)
}
