package note

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/ydb"
)

// DeleteByIdentifier tries to delete a note by its identifier
func DeleteByIdentifier(d *db.DB, identifier int) error {
	col := d.Use(ydb.ColNotes)

	_, err, id := GetByIdentifier(d, identifier)
	if err != nil {
		return err
	}

	return col.Delete(id)
}
