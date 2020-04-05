package db

import (

	"database/sql"
	"log"
	"sync"
)
type DBhelper struct {

Conn	*sql.DB

}
var once3 sync.Once
var dbhelper *DBhelper
func GetDb() *DBhelper  {

	once3.Do(func() {
		dbhelper = new(DBhelper)
		connStr := "postgres://postgres:123456@123.57.206.19:10001/lanren?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Panic(err.Error())
		}
		dbhelper.Conn=db
	})

	return  dbhelper
}
