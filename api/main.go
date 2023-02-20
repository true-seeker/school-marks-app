package main

import (
	"flag"
	"fmt"
	"os"
	"school-marks-app/api/config"
	"school-marks-app/api/db"
	"school-marks-app/api/migration"
	"school-marks-app/api/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*environment)
	db.Init()

	migration.Migrate()
	migration.CreateCatalogs()
	server.Init()
}
