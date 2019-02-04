package config

import (
	"github.com/HouzuoGuo/tiedot/db"
)

type AppDeps struct {
	Config Config
	Db *db.DB
}
