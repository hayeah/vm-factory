package factory

import (
	"database/sql"
	"log"

	// _ "github.com/lib/pq"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

var DB *runner.DB

func init() {
	db, err := sql.Open("postgres", "dbname=vmfactory sslmode=disable")
	if err != nil {
		log.Fatalln("DB connect failed:", err)
		return
	}
	runner.MustPing(db)

	DB = runner.NewDB(db, "postgres")
}

// var DB *xorm.Engine
// func init() {
// 	engine, err := xorm.NewEngine("postgres", "dbname=vmfactory sslmode=disable")
// 	if err != nil {
// 		log.Fatalln("DB failure:", err)
// 	}
// 	DB = engine
// }
