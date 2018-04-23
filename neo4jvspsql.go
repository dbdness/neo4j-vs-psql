package main

import (
	"./neo4jdb"
	"./psqldb"
	"log"
	"fmt"
)

func main() {
	neo4jQuery()
	psqlQuery()
}

func neo4jQuery(){
	conn, err := neo4jdb.Open()
	if err != nil{
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Starting NEO4J query...")
	count := neo4jdb.GetDepthCount("Genia Crist", 2)

	fmt.Println(count)
}

func psqlQuery(){
	conn, err := psqldb.Open()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Starting PSQL query...")
	count := psqldb.GetDepthCount("Genia Crist", 2)

	fmt.Println(count)
}
