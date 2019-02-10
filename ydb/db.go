package ydb

import (
	"encoding/json"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/config"
	"yac-go/log"
	"yac-go/model"
)

// ColNotes holds the name of the notes collection
const ColNotes = "notes"

// ColBooks holds the name of the book collection
const ColBooks = "books"

func Init(deps *config.AppDeps) *db.DB {
	dbDir := deps.Config.DatabaseDir

	d, err := db.OpenDB(dbDir)
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to open database from '%s': %s", dbDir, err))
	}

	if d.ColExists(ColNotes) == false {
		if err := d.Create(ColNotes); err != nil {
			log.Panic("Failed to create notes collection in database.")
		}
		log.Info("Created new database collection 'notes'")

		noteCol := d.Use(ColNotes)
		if err := noteCol.Index([]string{"key"}); err != nil {
			log.Info(err)
		}
		if err := noteCol.Index([]string{"book"}); err != nil {
			log.Info(err)
		}
	}

	if d.ColExists(ColBooks) == false {
		if err := d.Create(ColBooks); err != nil {
			log.Panic("Failed to create books collection in database.")
		}
		log.Info("Created new database collection 'books'")
	}

	//if err := d.Scrub("notes"); err != nil {
	//	log.Panic("Error during database cleanup:", err)
	//}

	return d
}

func ExecuteQuery(col *db.Col, queryStr []byte) (map[int]struct{}, error) {
	queryResult := make(map[int]struct{})
	var query interface{}
	if err := json.Unmarshal(queryStr, &query); err != nil {
		return nil, err
	}

	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		return nil, err
	}

	return queryResult, nil
}

func LoadById(col *db.Col, id int, v interface{}) error {
	plain, err := col.Read(id)
	if err != nil {
		return err
	}

	model.Marshal(plain, v)
	return nil
}
