package psqldb

import (
	_ "github.com/lib/pq"

	"database/sql"
	"log"
)

const connString = "user=postgres dbname=postgres sslmode=disable"



var db *sql.DB

func Open() (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", connString)
	return db, err
}

func GetDepthCount(name string, depth int) int {
	var query string
	switch depth {
	case 1:
		query = DepthOne
		break
	case 2:
		query = DepthTwo
		break
	case 3:
		query = DepthThree
		break
	case 4:
		query = DepthFour
		break
	case 5:
		query = DepthFive
		break
	}

	rows, err := db.Query(query, name)
	if err != nil {
		log.Panic(err)
	}

	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	return count
}
