package psqldb

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
)

const connString = "user=postgres dbname=postgres sslmode=disable"
var db *sql.DB

func Open() (*sql.DB, error){
	var err error
	db, err = sql.Open("postgres", connString)
	return db, err
}

func GetDepthCount(id string, depth int) int {
	rows, err := db.Query("SELECT COUNT(*) FROM endorsement WHERE person_one_id=$1", id)
	if err != nil{
		log.Panic(err)
	}

	var count int
	for rows.Next(){
		rows.Scan(&count)
	}

	return count
}