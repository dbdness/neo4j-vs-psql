package neo4jdb

import (
	driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"log"
	"strconv"
)

const connString = "bolt://neo4j:class@localhost:7687"

var conn driver.Conn

func Open() (driver.Conn, error) {
	var err error
	conn, err = driver.NewDriver().OpenNeo(connString)
	return conn, err
}

func GetDepthCount(name string, depth int) int64 {
	checkConn()

	depthStr := strconv.Itoa(depth)

	cypher := `MATCH (:Person{name:{name}})-[*` + depthStr + `]->(other) RETURN count(other)`
	//cypher := `MATCH (:Person{name:{name}})-[*` + depthStr + `]->(other) RETURN other`

	data, _, _, err := conn.QueryNeoAll(cypher, map[string]interface{}{"name": name})
	if err != nil {
		log.Panic(err)
	}

	count := data[0][0].(int64)
	return count
}

func GetRandomName(amount int) []string {
	checkConn()

	names := make([]string, amount)

	cypher := `MATCH (p:Person) WITH p, rand() as random RETURN p.name ORDER BY random LIMIT {limit}`

	data, _, _, err := conn.QueryNeoAll(cypher, map[string]interface{}{"limit": amount})
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < amount; i++ {
		name := data[i][0].(string)
		names[i] = name
	}

	return names
}

func checkConn() {
	if conn == nil {
		log.Panic("Neo4j connection null. Did you forget to call Open()?")
	}
}
