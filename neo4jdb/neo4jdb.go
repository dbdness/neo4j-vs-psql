package neo4jdb

import (
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"log"
)

const connString  = "bolt://neo4j:class@localhost:7687"
var conn driver.Conn

func Open() (driver.Conn, error) {
	var err error
	conn, err = driver.NewDriver().OpenNeo(connString)
	return conn, err
}

func GetDepthCount(name string, depth string) int64 {
	if conn == nil{
		log.Panic("Neo4j connection null. Did you forget to call Open()?")
	}

	cypher := `MATCH (:Person{name:{name}})-[*`+depth+`]->(other) RETURN count(other)`

	data, _, _, err := conn.QueryNeoAll(cypher, map[string]interface{}{"name":name})

	if err != nil {
		log.Panic(err)
	}

	count := data[0][0].(int64)
	return count
}
