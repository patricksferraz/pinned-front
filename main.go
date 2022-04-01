package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Netflix/go-env"
	"github.com/c-4u/pinned-front/app"
	"github.com/c-4u/pinned-front/config"
	"github.com/joho/godotenv"
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
