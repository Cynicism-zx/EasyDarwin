package db

import (
	"fmt"
	"log"

	"EasyDarwin/helper/jinzhu/gorm"
	_ "EasyDarwin/helper/jinzhu/gorm/dialects/sqlite"
	"EasyDarwin/helper/penggy/EasyGoLib/utils"
)

type Model struct {
	ID        string         `structs:"id" gorm:"primary_key" form:"id" json:"id"`
	CreatedAt utils.DateTime `structs:"-" json:"createdAt" gorm:"type:datetime"`
	UpdatedAt utils.DateTime `structs:"-" json:"updatedAt" gorm:"type:datetime"`
	// DeletedAt *time.Time `sql:"index" structs:"-"`
}

var SQLite *gorm.DB

func Init() (err error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTablename string) string {
		return "t_" + defaultTablename
	}
	dbFile := utils.DBFile()
	log.Println("db file -->", utils.DBFile())
	SQLite, err = gorm.Open("sqlite3", fmt.Sprintf("%s?loc=Asia/Shanghai", dbFile))
	if err != nil {
		return
	}
	// Sqlite cannot handle concurrent writes, so we limit sqlite to one connection.
	// see https://EasyDarwin/helper/mattn/go-sqlite3/issues/274
	SQLite.DB().SetMaxOpenConns(1)
	SQLite.SetLogger(DefaultGormLogger)
	SQLite.LogMode(false)
	return
}

func Close() {
	if SQLite != nil {
		SQLite.Close()
		SQLite = nil
	}
}
