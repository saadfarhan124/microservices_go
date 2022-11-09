package database

import (
	"database/sql"
	"log"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type DBInstance struct {
	Db *sql.DB
}

var Database DBInstance

func ConnectDB() {
	var dbUrl = "postgresql://postgres:postgres@db:5432/web_microservice?sslmode=disable"
	orm.RegisterDriver("postgres", orm.DRPostgres)

	if err := orm.RegisterDataBase("default", "postgres", dbUrl); err != nil {
		log.Fatal(err.Error())
	}

	if db, err := sql.Open("postgres", dbUrl); err != nil {
		log.Fatal(err.Error())
	} else {
		if err := orm.RunSyncdb("default", false, true); err != nil {
			log.Fatal(err.Error())
		} else {
			Database = DBInstance{Db: db}
			log.Print("Connected Succcessfully")
		}
	}
}
