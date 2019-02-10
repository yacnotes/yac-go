package book

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/service/note"
	"yac-go/ydb"
)

func Delete(d *db.DB, id int) error {
	col := d.Use(ydb.ColBooks)
	if err := col.Delete(id); err != nil {
		return err
	}

	notes, err := note.GetAll(d, id)
	if err != nil {
		return err
	}

	for _, n := range notes {
		if err := note.DeleteById(d, n.Id); err != nil {
			return err
		}
	}

	return nil
}
