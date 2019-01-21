package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type conn struct {
	psql *sql.DB
}

func ConnectPSQL(uri string) *conn {
	var db = new(conn)
	var err error
	db.psql, err = sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (db *conn) InsertLog(uid, text string) error {
	sqlStatement := `
INSERT INTO chatter_users (auto_user_id, message, time)
VALUES ($1, $2, $3)`
	_, err := db.psql.Exec(sqlStatement, uid, text, time.Now().UnixNano())
	return err
}
