package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"school-marks-app/internal/pkg/app"
	"school-marks-app/migration"
	"school-marks-app/pkg/config"
)

func Init() {
	environment := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*environment)
}

func main() {
	Init()
	migration.Migrate()
	migration.CreateCatalogs()

	r := gin.Default()
	a := app.New(r)

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
