package main

import (
	"./neo4jdb"
	"./psqldb"
	"log"
	"fmt"
)

func main() {
	//neo4jQuery()
	psqlQuery()
}

func neo4jQuery(){
	conn, err := neo4jdb.Open()
	if err != nil{
		log.Panic(err)
	}
	defer conn.Close()

	count := neo4jdb.GetDepthCount("Genia Crist", 1)

	fmt.Println(count)
}

func psqlQuery(){
	conn, err := psqldb.Open()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	count := psqldb.GetDepthCount("20653", 1)

	fmt.Println(count)
}
