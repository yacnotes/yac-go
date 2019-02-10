package note

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/ydb"
)

// DeleteById tries to delete a note by its id
func DeleteById(d *db.DB, id int) error {
	col := d.Use(ydb.ColNotes)
	return col.Delete(id)
}
