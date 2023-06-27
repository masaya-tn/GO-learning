package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/masaya-tn/GO-learning/config"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "user"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		email string,
		password string,
		created_at datatime
	)`, tableNameUser)

	Db.Exec(cmdU)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
