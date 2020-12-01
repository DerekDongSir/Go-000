package dao

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

var db *sql.DB

func init() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func FetchData() (interface{}, error) {
	rows, err := db.Query("select id, name from users where id = 1", 1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, xerrors.Wrap(err, "no rows found")
		}
		return nil, xerrors.Wrap(err, "fetch data error ")
	}
	return rows, nil
}
