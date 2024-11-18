package main

import (
	config2 "authSSO/internal/config"
	rlog "authSSO/internal/log"
	migrations2 "authSSO/internal/migrations"
	"flag"
)

func main() {

	config := config2.NewConfig()
	log := rlog.NewLogger("debug")

	action := flag.String("action", "up", "the action you want to execute (up/down")
	flag.Parse()

	migrations := migrations2.NewMigrations(config, log)

	switch *action {
	case "up":
		err := migrations.Up()
		if err != nil {
			log.Debug("Error while migrating database up!")
			panic(err)
		}
	case "down":
		err := migrations.Down()
		if err != nil {
			log.Debug("Error while migrating database down!")
			panic(err)
		}
	default:
		log.Info("No action provided! Exiting...")
		return
	}
}
