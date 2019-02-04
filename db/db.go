package db

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/config"
	"yac-go/log"
)

func Init(deps *config.AppDeps) *db.DB {
	dbDir := deps.Config.DatabaseDir

	d, err := db.OpenDB(dbDir)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to open database from '%s': %s", dbDir, err))
	}

	if d.ColExists("notes") == false {
		if err := d.Create("notes"); err != nil {
			log.Panic("Failed to create notes collection in database.")
		}
		log.Info("Created new database collection 'notes'")
	}

	if err := d.Scrub("notes"); err != nil {
		log.Panic("Error during database cleanup:", err)
	}

	return d
}
