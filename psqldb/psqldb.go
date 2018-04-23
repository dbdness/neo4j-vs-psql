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
	checkConn()

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

func GetRandomName(amount int) []string {
	checkConn()

	randoms := make([]string, amount)

	rows, err := db.Query(`SELECT name FROM person ORDER BY random() LIMIT $1;`, amount)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < amount; i++ {
		if rows.Next() == false {
			break
		}
		var name string
		rows.Scan(&name)
		randoms[i] = name
	}

	return randoms
}

func checkConn(){
	if db == nil{
		log.Panic("Neo4j connection null. Did you forget to call Open()?")
	}
}
